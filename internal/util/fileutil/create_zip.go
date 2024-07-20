package fileutil

import (
	"archive/zip"
	"bytes"
)

// ZipFile represents a file to be added to a zip archive.
type ZipFile struct {
	Name  string
	Bytes []byte
}

// CreateZip creates a zip file with the given files and
// returns the zip file as a byte slice or an error if
// something went wrong.
func CreateZip(files []ZipFile) ([]byte, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			w.Close()
			return nil, err
		}
		if _, err := f.Write(file.Bytes); err != nil {
			w.Close()
			return nil, err
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
