package static

import (
	"embed"
	"sync"

	"github.com/eduardolat/pgbackweb/internal/util/cryptoutil"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
)

//go:embed *
var StaticFs embed.FS

var (
	staticSHA256     string
	staticSHA256Once sync.Once
)

// GetStaticSHA256 returns the SHA256 hash of all the files combined in the
// static directory.
func GetStaticSHA256() string {
	staticSHA256Once.Do(func() {
		staticSHA256 = cryptoutil.GetSHA256FromFS(StaticFs)
	})
	return staticSHA256
}

// GetVersionedFilePath returns a versioned file path by appending a shortened
// SHA256 hash of the static filesystem to the query parameter.
//
// The hash is truncated to the first 8 characters for brevity.
//
// This function also prepends the configured path prefix so that static files
// are correctly referenced in HTML when the application is served under a subpath.
func GetVersionedFilePath(filePath string) string {
	hash := GetStaticSHA256()

	if len(hash) > 8 {
		hash = hash[:8]
	}

	// Prepend the path prefix to ensure static files are found
	// when the app is served under a subpath like /pgbackweb
	fullPath := pathutil.BuildPath(filePath)

	return fullPath + "?v=" + hash
}
