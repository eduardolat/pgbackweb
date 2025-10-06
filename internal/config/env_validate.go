package config

import (
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/validate"
)

// validateEnv runs additional validations on the environment variables.
func validateEnv(env Env) error {
	if !validate.ListenHost(env.PBW_LISTEN_HOST) {
		return fmt.Errorf("invalid listen address %s", env.PBW_LISTEN_HOST)
	}

	if !validate.Port(env.PBW_LISTEN_PORT) {
		return fmt.Errorf("invalid listen port %s, valid values are 1-65535", env.PBW_LISTEN_PORT)
	}

	if !validate.PathPrefix(env.PBW_PATH_PREFIX) {
		return fmt.Errorf("invalid path prefix %s, must start with / and not end with / (or be empty)", env.PBW_PATH_PREFIX)
	}

	return nil
}
