package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvAsStringFunc(t *testing.T) {
	// Test when environment variable exists
	os.Setenv("TEST_ENV", "test_value")
	value, err := getEnvAsStringFunc(getEnvAsStringParams{
		name:       "TEST_ENV",
		isRequired: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, "test_value", *value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable does not exist, default value is provided, and is not required
	value, err = getEnvAsStringFunc(getEnvAsStringParams{
		name:         "NON_EXISTENT_ENV",
		defaultValue: newDefaultValue("default_value"),
		isRequired:   false,
	})
	assert.NoError(t, err)
	assert.Equal(t, "default_value", *value)

	// Test when environment variable does not exist, no default value is provided, and is required
	// This should return an error
	value, err = getEnvAsStringFunc(getEnvAsStringParams{
		name:       "NON_EXISTENT_ENV",
		isRequired: true,
	})
	assert.Error(t, err)
	assert.Nil(t, value)

	// Test when environment variable exists, default value is provided, and is required
	os.Setenv("TEST_ENV", "test_value")
	value, err = getEnvAsStringFunc(getEnvAsStringParams{
		name:         "TEST_ENV",
		defaultValue: newDefaultValue("default_value"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "test_value", *value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable exists, is not required, and no default value is provided
	os.Setenv("TEST_ENV", "test_value")
	value, err = getEnvAsStringFunc(getEnvAsStringParams{
		name:       "TEST_ENV",
		isRequired: false,
	})
	assert.NoError(t, err)
	assert.Equal(t, "test_value", *value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable does not exist, is not required, and no default value is provided
	value, err = getEnvAsStringFunc(getEnvAsStringParams{
		name:       "NON_EXISTENT_ENV",
		isRequired: false,
	})
	assert.NoError(t, err)
	assert.Nil(t, value)

	// Test when default value and required are both present
	// This should return an error
	_, err = getEnvAsStringFunc(getEnvAsStringParams{
		name:         "NON_EXISTENT_ENV",
		defaultValue: newDefaultValue("default_value"),
		isRequired:   true,
	})
	assert.Error(t, err)
}

func TestGetEnvAsIntFunc(t *testing.T) {
	// Test when environment variable exists and is an integer
	os.Setenv("TEST_ENV", "123")
	value, err := getEnvAsIntFunc(getEnvAsIntParams{
		name:       "TEST_ENV",
		isRequired: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, 123, *value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable does not exist, default value is provided, and is not required
	value, err = getEnvAsIntFunc(getEnvAsIntParams{
		name:         "NON_EXISTENT_ENV",
		defaultValue: newDefaultValue(456),
	})
	assert.NoError(t, err)
	assert.Equal(t, 456, *value)

	// Test when environment variable does not exist, no default value is provided, and is required
	// This should return an error
	value, err = getEnvAsIntFunc(getEnvAsIntParams{
		name:       "NON_EXISTENT_ENV",
		isRequired: true,
	})
	assert.Error(t, err)
	assert.Nil(t, value)

	// Test when environment variable exists, is not an integer, no default value is provided, and is required
	// This should return an error
	os.Setenv("TEST_ENV", "not_an_integer")
	value, err = getEnvAsIntFunc(getEnvAsIntParams{
		name:       "TEST_ENV",
		isRequired: true,
	})
	assert.Error(t, err)
	assert.Nil(t, value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable exists, is not required, and no default value is provided
	os.Setenv("TEST_ENV", "123")
	value, err = getEnvAsIntFunc(getEnvAsIntParams{
		name:       "TEST_ENV",
		isRequired: false,
	})
	assert.NoError(t, err)
	assert.Equal(t, 123, *value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable does not exist, is not required, and no default value is provided
	value, err = getEnvAsIntFunc(getEnvAsIntParams{
		name:       "NON_EXISTENT_ENV",
		isRequired: false,
	})
	assert.NoError(t, err)
	assert.Nil(t, value)

	// Test when default value and required are both present
	// This should return an error
	_, err = getEnvAsIntFunc(getEnvAsIntParams{
		name:         "NON_EXISTENT_ENV",
		defaultValue: newDefaultValue(1),
		isRequired:   true,
	})
	assert.Error(t, err)
}

func TestGetEnvAsBoolFunc(t *testing.T) {
	// Test when environment variable exists and is a boolean
	os.Setenv("TEST_ENV", "true")
	value, err := getEnvAsBoolFunc(getEnvAsBoolParams{
		name:       "TEST_ENV",
		isRequired: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, true, *value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable exists, is not a boolean, and is required
	os.Setenv("TEST_ENV", "not_a_boolean")
	_, err = getEnvAsBoolFunc(getEnvAsBoolParams{
		name:       "TEST_ENV",
		isRequired: true,
	})
	assert.Error(t, err)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable exists, is not required, and no default value is provided
	os.Setenv("TEST_ENV", "true")
	value, err = getEnvAsBoolFunc(getEnvAsBoolParams{
		name:       "TEST_ENV",
		isRequired: false,
	})
	assert.NoError(t, err)
	assert.Equal(t, true, *value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable does not exist, is not required, and no default value is provided
	value, err = getEnvAsBoolFunc(getEnvAsBoolParams{
		name:       "NON_EXISTENT_ENV",
		isRequired: false,
	})
	assert.NoError(t, err)
	assert.Equal(t, false, *value)

	// Test when default value and required are both present
	// This should return an error
	_, err = getEnvAsBoolFunc(getEnvAsBoolParams{
		name:         "NON_EXISTENT_ENV",
		defaultValue: newDefaultValue(true),
		isRequired:   true,
	})
	assert.Error(t, err)
}
