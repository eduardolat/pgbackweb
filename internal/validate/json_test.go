package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	validJSON := []string{
		`{"name": "John", "age": 30}`,
		`{"name": "John", "friends": ["Alice", "Bob"]}`,
		`{"items": [{"id": 1, "name": "Item1"}, {"id": 2, "name": "Item2"}]}`,
		`{"emptyArray": [], "emptyObject": {}}`,
		`{"nullValue": null}`,
		`{"booleanTrue": true, "booleanFalse": false}`,
		`{"numberInt": 123, "numberFloat": 123.456}`,
		`{"string": "Hello, World!"}`,
		`{"nested": {"level1": {"level2": {"level3": "value"}}}}`,
		`{"escapedString": "This is a quote: \"}"}`,
	}

	for _, jsonStr := range validJSON {
		assert.True(t, JSON(jsonStr), "Expected JSON to be valid: %s", jsonStr)
	}

	invalidJSON := []string{
		`{"name": "John", "age": 30`,
		`{"name": "John", "friends": ["Alice", "Bob"}`,
		`{"items": [{"id": 1, "name": "Item1"}, {"id": 2, "name": "Item2"]]}`,
		`{"unclosedString": "Hello}`,
		`{"unexpectedToken": tru}`,
		`{"incompleteObject": {}`,
		`{name: "John", "age": 30}`,
		`{"missingComma": "value1" "value2"}`,
		`["mismatch": "value"]`,
		`some other thing`,
		``,
	}

	for _, jsonStr := range invalidJSON {
		assert.False(t, JSON(jsonStr), "Expected JSON to be invalid: %s", jsonStr)
	}
}
