package middleware

import "github.com/eduardolat/pgbackweb/internal/service"

type Middleware struct {
	servs *service.Service
}

func New(servs *service.Service) *Middleware {
	return &Middleware{servs: servs}
}
