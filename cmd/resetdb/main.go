package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eduardolat/pgbackweb/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Are you sure you want to reset the database? (yes/no)")
	var answer string

	for answer != "yes" && answer != "no" {
		if _, err := fmt.Scanln(&answer); err != nil {
			panic(err)
		}
		if answer == "no" {
			log.Println("Exiting...")
			return
		}
		if answer == "yes" {
			log.Println("Resetting database...")
			break
		}
		log.Println("Please enter 'yes' or 'no'")
	}

	env, err := config.GetEnv()
	if err != nil {
		panic(err)
	}

	db := connectDB(env)

	_, err = db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	if err != nil {
		panic(fmt.Errorf("❌ Could not reset DB: %w", err))
	}

	log.Println("✅ Database reset")
}

func connectDB(env config.Env) *sql.DB {
	db, err := sql.Open("postgres", env.PBW_POSTGRES_CONN_STRING)
	if err != nil {
		panic(fmt.Errorf("❌ Could not connect to DB: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("❌ Could not ping DB: %w", err))
	}

	db.SetMaxOpenConns(1)
	log.Println("✅ Connected to DB")

	return db
}
