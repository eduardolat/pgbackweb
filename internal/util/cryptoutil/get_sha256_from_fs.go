package cryptoutil

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/fs"
)

// GetSHA256FromFS takes a fs.FS and returns a SHA256 hash of the combined contents
// of all files in the filesystem.
//
// If there is an error, it returns an empty string.
func GetSHA256FromFS(fsys fs.FS) string {
	hash := sha256.New()

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		file, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := io.Copy(hash, file); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}
