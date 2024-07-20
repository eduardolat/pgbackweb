package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/eduardolat/pgbackweb/internal/logger"
)

type getEnvAsStringParams struct {
	name         string
	defaultValue *string
	isRequired   bool
}

// defaultValue returns a pointer to the given value.
func newDefaultValue[T any](value T) *T {
	return &value
}

// getEnvAsString returns the value of the environment variable with the given name.
func getEnvAsString(params getEnvAsStringParams) *string { //nolint:all
	value, err := getEnvAsStringFunc(params)

	if err != nil {
		logger.FatalError(
			"error getting env variable", logger.KV{
				"name":  params.name,
				"error": err,
			},
		)
	}

	return value
}

// getEnvAsStringFunc is the outlying function for getEnvAsString.
func getEnvAsStringFunc(params getEnvAsStringParams) (*string, error) {
	if params.defaultValue != nil && params.isRequired {
		return nil, errors.New("cannot have both a default value and be required")
	}

	value, exists := os.LookupEnv(params.name)

	if !exists && params.isRequired {
		return nil, errors.New("required env variable does not exist")
	}

	if !exists {
		if params.defaultValue != nil {
			return params.defaultValue, nil
		}
		return nil, nil
	}

	return &value, nil
}

type getEnvAsIntParams struct {
	name         string
	defaultValue *int
	isRequired   bool
}

// getEnvAsInt returns the value of the environment variable with the given name.
func getEnvAsInt(params getEnvAsIntParams) *int { //nolint:all
	value, err := getEnvAsIntFunc(params)

	if err != nil {
		logger.FatalError(
			"error getting env variable", logger.KV{
				"name":  params.name,
				"error": err,
			},
		)
	}

	return value
}

// getEnvAsIntFunc is the outlying function for getEnvAsInt.
func getEnvAsIntFunc(params getEnvAsIntParams) (*int, error) {
	if params.defaultValue != nil && params.isRequired {
		return nil, errors.New("cannot have both a default value and be required")
	}

	valueStr, exists := os.LookupEnv(params.name)

	if !exists && params.isRequired {
		return nil, errors.New("required env variable does not exist")
	}

	if !exists {
		if params.defaultValue != nil {
			return params.defaultValue, nil
		}
		return nil, nil
	}

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		return nil, errors.New("env variable is not an integer")
	}

	return &value, nil
}

type getEnvAsBoolParams struct {
	name         string
	defaultValue *bool
	isRequired   bool
}

// getEnvAsBool returns the value of the environment variable with the given name.
func getEnvAsBool(params getEnvAsBoolParams) *bool { //nolint:all
	value, err := getEnvAsBoolFunc(params)

	if err != nil {
		logger.FatalError(
			"error getting env variable", logger.KV{
				"name":  params.name,
				"error": err,
			},
		)
	}

	return value
}

// getEnvAsBoolFunc is the outlying function for getEnvAsBool.
func getEnvAsBoolFunc(params getEnvAsBoolParams) (*bool, error) {
	if params.defaultValue != nil && params.isRequired {
		return nil, errors.New("cannot have both a default value and be required")
	}

	valueStr, exists := os.LookupEnv(params.name)

	if !exists && params.isRequired {
		return nil, errors.New("required env variable does not exist")
	}

	if !exists {
		if params.defaultValue != nil {
			return params.defaultValue, nil
		}
		f := false
		return &f, nil
	}

	value, err := strconv.ParseBool(valueStr)

	if err != nil {
		return nil, errors.New("env variable is not a boolean, must be true or false")
	}

	return &value, nil
}
