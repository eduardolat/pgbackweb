package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	localBackupsDir string = "/backups"
)

// LocalUpload Creates a new file using the provided path and reader relative
// to the local backups directory.
func (Client) LocalUpload(relativePath string, fileReader io.Reader) error {
	fullPath := filepath.Join(localBackupsDir, relativePath)
	dir := filepath.Dir(fullPath)

	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fullPath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, fileReader)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", fullPath, err)
	}

	return nil
}

// LocalDelete Deletes a file using the provided path relative to the local
// backups directory.
func (Client) LocalDelete(relativePath string) error {
	fullPath := filepath.Join(localBackupsDir, relativePath)

	err := os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", fullPath, err)
	}

	return nil
}

// LocalReadFile Reads a file using the provided path relative to the local
// backups directory.
func (Client) LocalReadFile(relativePath string) (io.ReadCloser, error) {
	fullPath := filepath.Join(localBackupsDir, relativePath)

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", fullPath, err)
	}

	return file, nil
}
