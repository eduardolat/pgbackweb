package main

import (
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/database"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service"
)

func main() {
	env := config.GetEnv()

	db := database.Connect(env)
	defer db.Close()
	dbgen := dbgen.New(db)

	_ = service.New(dbgen)
}
