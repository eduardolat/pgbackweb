package maputil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSortedStringKeys(t *testing.T) {
	// Create a map with unsorted keys
	m := map[string]any{
		"banana": 1,
		"apple":  2,
		"cherry": 3,
	}

	// Call the function
	keys := GetSortedStringKeys(m)

	// Assert that the keys are sorted
	expectedKeys := []string{"apple", "banana", "cherry"}
	assert.Equal(t, expectedKeys, keys)
}

func TestGetSortedStringKeysEmptyMap(t *testing.T) {
	// Create an empty map
	m := map[string]any{}

	// Call the function
	keys := GetSortedStringKeys(m)

	// Assert that the keys are empty
	assert.Empty(t, keys)
}
