package config

import (
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/joho/godotenv"
)

type Env struct {
	PBW_ENCRYPTION_KEY       *string
	PBW_POSTGRES_CONN_STRING *string
}

// GetEnv returns the environment variables.
//
// If there is an error, it will log it and exit the program.
func GetEnv() *Env {
	err := godotenv.Load()
	if err == nil {
		logger.Info("using .env file")
	}

	env := &Env{
		PBW_ENCRYPTION_KEY: getEnvAsString(getEnvAsStringParams{
			name:       "PBW_ENCRYPTION_KEY",
			isRequired: true,
		}),
		PBW_POSTGRES_CONN_STRING: getEnvAsString(getEnvAsStringParams{
			name:       "PBW_POSTGRES_CONN_STRING",
			isRequired: true,
		}),
	}

	validateEnv(env)
	return env
}
