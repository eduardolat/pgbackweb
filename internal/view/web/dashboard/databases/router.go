package databases

import (
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	servs *service.Service
}

func newHandlers(servs *service.Service) *handlers {
	return &handlers{servs: servs}
}

func MountRouter(
	parent *echo.Group, mids *middleware.Middleware, servs *service.Service,
) {
	h := newHandlers(servs)

	parent.GET("", h.indexPageHandler)
	parent.GET("/list", h.listDatabasesHandler)
	parent.POST("", h.createDatabaseHandler)
	parent.POST("/test", h.testDatabaseHandler)
	parent.DELETE("/:databaseID", h.deleteDatabaseHandler)
	parent.POST("/:databaseID/edit", h.editDatabaseHandler)
	parent.POST("/:databaseID/test", h.testExistingDatabaseHandler)
}
