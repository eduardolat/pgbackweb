package main

import (
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/database"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func main() {
	env := config.GetEnv()
	db := database.Connect(env)
	defer db.Close()
	_ = dbgen.New(db)
}
