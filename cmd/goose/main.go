package main

import (
	"fmt"
	"os"

	"github.com/eduardolat/pgbackweb/internal/config"
)

const migrationsFolder string = "./internal/database/migrations"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("goose command is required")
		fmt.Println("example: task goose -- up")
		return
	}

	gooseCmd := ""
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		gooseCmd += arg + " "
	}

	env := config.GetEnv()

	cmd := fmt.Sprintf(
		"goose -dir %s postgres \"%s\" %s",
		migrationsFolder,
		*env.PBW_POSTGRES_CONN_STRING,
		gooseCmd,
	)

	fmt.Println(cmd)
}
