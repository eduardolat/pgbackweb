package config

import (
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Env struct {
	PBW_ENCRYPTION_KEY       string `env:"PBW_ENCRYPTION_KEY,required"`
	PBW_POSTGRES_CONN_STRING string `env:"PBW_POSTGRES_CONN_STRING,required"`
	PBW_LISTEN_HOST          string `env:"PBW_LISTEN_HOST" envDefault:"0.0.0.0"`
	PBW_LISTEN_PORT          string `env:"PBW_LISTEN_PORT" envDefault:"8085"`

	// OIDC Configuration
	PBW_OIDC_ENABLED        bool   `env:"PBW_OIDC_ENABLED" envDefault:"false"`
	PBW_OIDC_ISSUER_URL     string `env:"PBW_OIDC_ISSUER_URL"`
	PBW_OIDC_CLIENT_ID      string `env:"PBW_OIDC_CLIENT_ID"`
	PBW_OIDC_CLIENT_SECRET  string `env:"PBW_OIDC_CLIENT_SECRET"`
	PBW_OIDC_REDIRECT_URL   string `env:"PBW_OIDC_REDIRECT_URL"`
	PBW_OIDC_SCOPES         string `env:"PBW_OIDC_SCOPES" envDefault:"openid profile email"`
	PBW_OIDC_USERNAME_CLAIM string `env:"PBW_OIDC_USERNAME_CLAIM" envDefault:"preferred_username"`
	PBW_OIDC_EMAIL_CLAIM    string `env:"PBW_OIDC_EMAIL_CLAIM" envDefault:"email"`
	PBW_OIDC_NAME_CLAIM     string `env:"PBW_OIDC_NAME_CLAIM" envDefault:"name"`
}

var (
	getEnvRes  Env
	getEnvErr  error
	getEnvOnce sync.Once
)

// GetEnv returns the environment variables.
//
// If there is an error, it will log it and exit the program.
func GetEnv(disableLogs ...bool) (Env, error) {
	getEnvOnce.Do(func() {
		_ = godotenv.Load()

		parsedEnv, err := env.ParseAs[Env]()
		if err != nil {
			getEnvErr = err
			return
		}

		if err := validateEnv(parsedEnv); err != nil {
			getEnvErr = err
			return
		}

		getEnvRes = parsedEnv
	})

	return getEnvRes, getEnvErr
}
