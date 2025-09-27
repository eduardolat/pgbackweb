package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/database"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/cryptoutil"
	"github.com/google/uuid"
)

func main() {
	env, err := config.GetEnv()
	if err != nil {
		panic(err)
	}

	db := database.Connect(env)
	defer db.Close()
	dbg := dbgen.New(db)

	fmt.Println()
	fmt.Println()
	fmt.Println("PG Back Web - Password Reset")
	fmt.Println("---")
	fmt.Print("User email: ")
	var userID uuid.UUID

	for {
		var email string
		if _, err := fmt.Scanln(&email); err != nil {
			panic(err)
		}

		user, err := dbg.UsersServiceGetUserByEmail(
			context.Background(), email,
		)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			fmt.Print("User not found. Enter new email: ")
			continue
		}
		if err != nil {
			panic(err)
		}

		userID = user.ID
		break
	}

	newPassword := uuid.NewString()
	hashedPassword, err := cryptoutil.CreateBcryptHash(newPassword)
	if err != nil {
		panic(err)
	}

	err = dbg.UsersServiceChangePassword(
		context.Background(), dbgen.UsersServiceChangePasswordParams{
			ID:       userID,
			Password: sql.NullString{String: hashedPassword, Valid: true},
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Password reset successfully")
	fmt.Println("New password: ", newPassword)
	fmt.Println()
	fmt.Println("You can change your password after login")
	fmt.Println()
}
