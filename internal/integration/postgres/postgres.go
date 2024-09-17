package postgres

import (
	"bytes"
	"fmt"
	"github.com/alexmullins/zip"
	"io"
	"os"
	"os/exec"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/orsinium-labs/enum"
)

/*
	Important:
	Versions supported by PG Back Web must be supported in PostgreSQL Version Policy
	https://www.postgresql.org/support/versioning/

	Backing up a database from an old unsupported version should not be allowed.
*/

type version struct {
	version string
	pgDump  string
	psql    string
}

type PGVersion enum.Member[version]

var (
	PG13 = PGVersion{version{
		version: "13",
		pgDump:  "/usr/lib/postgresql/13/bin/pg_dump",
		psql:    "/usr/lib/postgresql/13/bin/psql",
	}}
	PG14 = PGVersion{version{
		version: "14",
		pgDump:  "/usr/lib/postgresql/14/bin/pg_dump",
		psql:    "/usr/lib/postgresql/14/bin/psql",
	}}
	PG15 = PGVersion{version{
		version: "15",
		pgDump:  "/usr/lib/postgresql/15/bin/pg_dump",
		psql:    "/usr/lib/postgresql/15/bin/psql",
	}}
	PG16 = PGVersion{version{
		version: "16",
		pgDump:  "/usr/lib/postgresql/16/bin/pg_dump",
		psql:    "/usr/lib/postgresql/16/bin/psql",
	}}
)

type Client struct{}

func New() *Client {
	return &Client{}
}

// ParseVersion returns the PGVersion enum member for the given PostgreSQL
// version as a string.
func (Client) ParseVersion(version string) (PGVersion, error) {
	switch version {
	case "13":
		return PG13, nil
	case "14":
		return PG14, nil
	case "15":
		return PG15, nil
	case "16":
		return PG16, nil
	default:
		return PGVersion{}, fmt.Errorf("pg version not allowed: %s", version)
	}
}

// Test tests the connection to the PostgreSQL database
func (Client) Test(version PGVersion, connString string) error {
	cmd := exec.Command(version.Value.psql, connString, "-c", "SELECT 1;")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"error running psql test v%s: %s",
			version.Value.version, output,
		)
	}

	return nil
}

// DumpParams contains the parameters for the pg_dump command
type DumpParams struct {
	// DataOnly (--data-only): Dump only the data, not the schema (data definitions).
	// Table data, large objects, and sequence values are dumped.
	DataOnly bool

	// SchemaOnly (--schema-only): Dump only the object definitions (schema), not data.
	SchemaOnly bool

	// Clean (--clean): Output commands to DROP all the dumped database objects
	// prior to outputting the commands for creating them. This option is useful
	// when the restore is to overwrite an existing database. If any of the
	// objects do not exist in the destination database, ignorable error messages
	// will be reported during restore, unless --if-exists is also specified.
	Clean bool

	// IfExists (--if-exists): Use DROP ... IF EXISTS commands to drop objects in
	// --clean mode. This suppresses “does not exist” errors that might otherwise
	// be reported. This option is not valid unless --clean is also specified.
	IfExists bool

	// Create (--create): Begin the output with a command to create the database
	// itself and reconnect to the created database. (With a script of this form,
	// it doesn't matter which database in the destination installation you
	// connect to before running the script.) If --clean is also specified, the
	// script drops and recreates the target database before reconnecting to it.
	Create bool

	// NoComments (--no-comments): Do not dump comments.
	NoComments bool
}

// Dump runs the pg_dump command with the given parameters. It returns the SQL
// dump as an io.Reader.
func (Client) Dump(
	version PGVersion, connString string, params ...DumpParams,
) io.Reader {
	pickedParams := DumpParams{}
	if len(params) > 0 {
		pickedParams = params[0]
	}

	args := []string{connString}
	if pickedParams.DataOnly {
		args = append(args, "--data-only")
	}
	if pickedParams.SchemaOnly {
		args = append(args, "--schema-only")
	}
	if pickedParams.Clean {
		args = append(args, "--clean")
	}
	if pickedParams.IfExists {
		args = append(args, "--if-exists")
	}
	if pickedParams.Create {
		args = append(args, "--create")
	}
	if pickedParams.NoComments {
		args = append(args, "--no-comments")
	}

	errorBuffer := &bytes.Buffer{}
	reader, writer := io.Pipe()
	cmd := exec.Command(version.Value.pgDump, args...)
	cmd.Stdout = writer
	cmd.Stderr = errorBuffer

	go func() {
		defer writer.Close()
		if err := cmd.Run(); err != nil {
			writer.CloseWithError(fmt.Errorf(
				"error running pg_dump v%s: %s",
				version.Value.version, errorBuffer.String(),
			))
		}
	}()

	return reader
}

// DumpZip runs the pg_dump command with the given parameters and returns the
// ZIP-compressed SQL dump as an io.Reader.
func (c *Client) DumpZip(
	version PGVersion, connString string, params ...DumpParams,
) io.Reader {
	dumpReader := c.Dump(version, connString, params...)
	reader, writer := io.Pipe()
	env := config.GetEnv()

	go func() {
		defer writer.Close()

		zipWriter := zip.NewWriter(writer)
		defer zipWriter.Close()

		var fileWriter io.Writer
		var err error

		if env.PBW_BACKUP_PASSWORD != nil {
			fileWriter, err = zipWriter.Encrypt("dump.sql", *env.PBW_BACKUP_PASSWORD)
		} else {
			fileWriter, err = zipWriter.Create("dump.sql")
		}

		if err != nil {
			writer.CloseWithError(fmt.Errorf("error creating zip file: %w", err))
			return
		}

		if _, err := io.Copy(fileWriter, dumpReader); err != nil {
			writer.CloseWithError(fmt.Errorf("error writing to zip file: %w", err))
			return
		}
	}()

	return reader
}

// RestoreZip downloads or copies the ZIP from the given url or path, unzips it,
// and runs the psql command to restore the database.
//
// The ZIP file must contain a dump.sql file with the SQL dump to restore.
//
//   - version: PostgreSQL version to use for the restore
//   - connString: connection string to the database
//   - isLocal: whether the ZIP file is local or a URL
//   - zipURLOrPath: URL or path to the ZIP file
func (Client) RestoreZip(
	version PGVersion, connString string, isLocal bool, zipURLOrPath string,
) error {
	env := config.GetEnv()
	workDir, err := os.MkdirTemp("", "pbw-restore-*")
	if err != nil {
		return fmt.Errorf("error creating temp dir: %w", err)
	}
	defer os.RemoveAll(workDir)
	zipPath := strutil.CreatePath(true, workDir, "dump.zip")
	dumpPath := strutil.CreatePath(true, workDir, "dump.sql")

	if isLocal {
		cmd := exec.Command("cp", zipURLOrPath, zipPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error copying ZIP file to temp dir: %s", output)
		}
	}

	if !isLocal {
		cmd := exec.Command("wget", "--no-verbose", "-O", zipPath, zipURLOrPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error downloading ZIP file: %s", output)
		}
	}

	if _, err := os.Stat(zipPath); os.IsNotExist(err) {
		return fmt.Errorf("zip file not found: %s", zipPath)
	}

	var cmd *exec.Cmd
	if env.PBW_BACKUP_PASSWORD != nil {
		cmd = exec.Command("7z", "x", "-p"+*env.PBW_BACKUP_PASSWORD, "-o"+workDir, zipPath)
	} else {
		cmd = exec.Command("7z", "x", "-o"+workDir, zipPath)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error unzipping ZIP file: %s", output)
	}

	if _, err := os.Stat(dumpPath); os.IsNotExist(err) {
		return fmt.Errorf("dump.sql file not found in ZIP file: %s", zipPath)
	}

	cmd = exec.Command(version.Value.psql, connString, "-f", dumpPath)
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"error running psql v%s command: %s",
			version.Value.version, output,
		)
	}

	return nil
}
