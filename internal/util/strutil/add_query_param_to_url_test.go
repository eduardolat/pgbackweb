package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddQueryParamToUrl(t *testing.T) {
	assert := assert.New(t)

	// Test case when url, key or value is empty
	assert.Equal("", AddQueryParamToUrl("", "key", "value"))
	assert.Equal("http://example.com", AddQueryParamToUrl("http://example.com", "", "value"))
	assert.Equal("http://example.com", AddQueryParamToUrl("http://example.com", "key", ""))

	// Test case when url does not contain "?"
	assert.Equal("http://example.com?key=value", AddQueryParamToUrl("http://example.com", "key", "value"))

	// Test case when url ends with "?"
	assert.Equal("http://example.com?key=value", AddQueryParamToUrl("http://example.com?", "key", "value"))

	// Test case when url ends with "&"
	assert.Equal("http://example.com&?key=value", AddQueryParamToUrl("http://example.com&", "key", "value"))

	// Test case when url contains "?" but does not end with "?" or "&"
	assert.Equal("http://example.com/path?key=value", AddQueryParamToUrl("http://example.com/path?", "key", "value"))

	// Test case when multiple query params are added
	expected := "http://example.com/path?key1=value1&key2=value2"
	got := AddQueryParamToUrl("http://example.com/path", "key1", "value1")
	got = AddQueryParamToUrl(got, "key2", "value2")
	assert.Equal(expected, got)

	// Test case when value contains special characters
	expected = "http://example.com/path?key1=value%20with%20spaces"
	got = AddQueryParamToUrl("http://example.com/path", "key1", "value with spaces")
	assert.Equal(expected, got)

	// Test case when the key is empty
	expected = "http://example.com/path"
	got = AddQueryParamToUrl("http://example.com/path", "", "value")
	assert.Equal(expected, got)

	// Test case when the value is empty
	expected = "http://example.com/path?key2=val2"
	got = AddQueryParamToUrl("http://example.com/path", "key1", "")
	got = AddQueryParamToUrl(got, "key2", "val2")
	assert.Equal(expected, got)
}
