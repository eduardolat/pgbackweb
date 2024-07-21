package pgdump

import (
	"fmt"
	"os/exec"

	"github.com/eduardolat/pgbackweb/internal/util/fileutil"
	"github.com/orsinium-labs/enum"
)

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

// Ping tests the connection to the PostgreSQL database
func (Client) Ping(version PGVersion, connString string) error {
	cmd := exec.Command(version.Value.psql, connString, "-c", "SELECT 1;")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"error running psql v%s: %s",
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
// dump as a byte slice.
func (Client) Dump(
	version PGVersion, connString string, params ...DumpParams,
) ([]byte, error) {
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

	cmd := exec.Command(version.Value.pgDump, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf(
			"error running pg_dump v%s: %s",
			version.Value.version, output,
		)
	}

	return output, nil
}

// DumpZip runs the pg_dump command with the given parameters and returns the
// ZIP-compressed SQL dump as a byte slice.
func (c *Client) DumpZip(
	version PGVersion, connString string, params ...DumpParams,
) ([]byte, error) {
	dump, err := c.Dump(version, connString, params...)
	if err != nil {
		return nil, err
	}

	output, err := fileutil.CreateZip([]fileutil.ZipFile{{
		Name:  "dump.sql",
		Bytes: dump,
	}})
	if err != nil {
		return nil, fmt.Errorf("error creating zip file: %w", err)
	}

	return output, nil
}
