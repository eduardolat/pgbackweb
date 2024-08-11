package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/eduardolat/pgbackweb/internal/util/strutil"
)

const (
	localBackupsDir string = "/backups"
)

// LocalUpload Creates a new file using the provided path and reader relative
// to the local backups directory.
//
// Returns the size of the file created.
func (Client) LocalUpload(relativeFilePath string, fileReader io.Reader) (int64, error) {
	fullPath := strutil.CreatePath(true, localBackupsDir, relativeFilePath)
	dir := filepath.Dir(fullPath)

	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return 0, fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return 0, fmt.Errorf("failed to create file %s: %w", fullPath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, fileReader)
	if err != nil {
		return 0, fmt.Errorf("failed to write file %s: %w", fullPath, err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("failed to get file info %s: %w", fullPath, err)
	}

	return fileInfo.Size(), nil
}

// LocalDelete Deletes a file using the provided path relative to the local
// backups directory.
func (Client) LocalDelete(relativeFilePath string) error {
	fullPath := strutil.CreatePath(true, localBackupsDir, relativeFilePath)

	err := os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", fullPath, err)
	}

	return nil
}

// LocalGetFullPath Returns the full path of a file using the provided relative
// file path to the local backups directory.
func (Client) LocalGetFullPath(relativeFilePath string) string {
	return strutil.CreatePath(true, localBackupsDir, relativeFilePath)
}
