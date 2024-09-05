package cryptoutil

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed get_sha256_from_fs_test_data/*
var testFS embed.FS

func TestGetSHA256FromFS_ValidFS(t *testing.T) {
	hash := GetSHA256FromFS(testFS)
	assert.NotEmpty(t, hash)

	// To generate the expected hash, you must combine the contents of all the
	// files in the test_data directory and calculate the SHA256 hash of the
	// resulting string.
	expectedHash := "d2c58b6783050a95542286a58250d4dc872877a6cf28610669516dcfacf954af"
	assert.Equal(t, expectedHash, hash)
}

func TestGetSHA256FromFS_EmptyFS(t *testing.T) {
	var emptyFS embed.FS
	hash := GetSHA256FromFS(emptyFS)

	expectedHash := sha256.New().Sum(nil)
	assert.Equal(t, hex.EncodeToString(expectedHash), hash)
}
