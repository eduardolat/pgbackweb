package webhooks

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
	parent.GET("/list", h.listWebhooksHandler)
	parent.GET("/create", h.createWebhookFormHandler)
	parent.POST("/create", h.createWebhookHandler)
	parent.GET("/:webhookID/edit", h.editWebhookFormHandler)
	parent.POST("/:webhookID/edit", h.editWebhookHandler)
	parent.DELETE("/:webhookID", h.deleteWebhookHandler)
}
