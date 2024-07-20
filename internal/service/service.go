package service

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/users"
)

type Service struct {
	UsersService *users.Service
}

func New(dbgen *dbgen.Queries) *Service {
	usersService := users.New(dbgen)

	return &Service{
		UsersService: usersService,
	}
}
