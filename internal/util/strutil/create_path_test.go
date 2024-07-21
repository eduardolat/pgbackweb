package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePath(t *testing.T) {
	assert := assert.New(t)

	// Test with addLeadingSlash as true and non-empty parts
	result := CreatePath(true, "part1", "part2", "part3")
	assert.Equal("/part1/part2/part3", result, "They should be equal")

	// Test with addLeadingSlash as false and non-empty parts
	result = CreatePath(false, "part1", "part2", "part3")
	assert.Equal("part1/part2/part3", result, "They should be equal")

	// Test with addLeadingSlash as true and some empty parts
	result = CreatePath(true, "part1", "", "part3")
	assert.Equal("/part1/part3", result, "They should be equal")

	// Test with addLeadingSlash as false and some empty parts
	result = CreatePath(false, "part1", "", "part3")
	assert.Equal("part1/part3", result, "They should be equal")

	// Test with addLeadingSlash as true and all empty parts
	result = CreatePath(true, "", "", "")
	assert.Equal("", result, "They should be equal")

	// Test with addLeadingSlash as false and all empty parts
	result = CreatePath(false, "", "", "")
	assert.Equal("", result, "They should be equal")

	// Test with addLeadingSlash as true and parts with leading slashes
	result = CreatePath(true, "/part1", "/part2", "/part3")
	assert.Equal("/part1/part2/part3", result, "They should be equal")

	// Test with addLeadingSlash as false and parts with leading slashes
	result = CreatePath(false, "/part1", "/part2", "/part3")
	assert.Equal("part1/part2/part3", result, "They should be equal")

	// Test with addLeadingSlash as false and parts with leading and trailing slashes
	result = CreatePath(false, "/part1/", "/part2/", "/part3/")
	assert.Equal("part1/part2/part3/", result, "They should be equal")

	// Test with file name and query params
	result = CreatePath(false, "/part1/", "/part2/", "/part3/", "file.pdf?param1=1&param2=2")
	assert.Equal("part1/part2/part3/file.pdf?param1=1&param2=2", result, "They should be equal")

	// Test create url
	result = CreatePath(false, "https://www.google.com/", "/search/", "golang.pdf?param1=1&param2=2")
	assert.Equal("https://www.google.com/search/golang.pdf?param1=1&param2=2", result, "They should be equal")
}
