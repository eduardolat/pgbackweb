package webhooks

import (
	"database/sql"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

type createWebhookDTO struct {
	Name      string      `form:"name" validate:"required"`
	EventType string      `form:"event_type" validate:"required"`
	TargetIds []uuid.UUID `form:"target_ids" validate:"required,gt=0"`
	IsActive  string      `form:"is_active" validate:"required,oneof=true false"`
	Url       string      `form:"url" validate:"required,url"`
	Method    string      `form:"method" validate:"required,oneof=GET POST"`
	Headers   string      `form:"headers" validate:"omitempty,json"`
	Body      string      `form:"body" validate:"omitempty,json"`
}

func (h *handlers) createWebhookHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData createWebhookDTO
	if err := c.Bind(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	_, err := h.servs.WebhooksService.CreateWebhook(
		ctx, dbgen.WebhooksServiceCreateWebhookParams{
			Name:      formData.Name,
			EventType: formData.EventType,
			TargetIds: formData.TargetIds,
			IsActive:  formData.IsActive == "true",
			Url:       formData.Url,
			Method:    formData.Method,
			Headers:   sql.NullString{String: formData.Headers, Valid: true},
			Body:      sql.NullString{String: formData.Body, Valid: true},
		},
	)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.Redirect(c, pathutil.BuildPath("/dashboard/webhooks"))
}

func (h *handlers) createWebhookFormHandler(c echo.Context) error {
	ctx := c.Request().Context()

	databases, err := h.servs.DatabasesService.GetAllDatabases(ctx)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	destinations, err := h.servs.DestinationsService.GetAllDestinations(ctx)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	backups, err := h.servs.BackupsService.GetAllBackups(ctx)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return echoutil.RenderNodx(c, http.StatusOK, createWebhookForm(
		databases, destinations, backups,
	))
}

func createWebhookForm(
	databases []dbgen.DatabasesServiceGetAllDatabasesRow,
	destinations []dbgen.DestinationsServiceGetAllDestinationsRow,
	backups []dbgen.Backup,
) nodx.Node {
	return nodx.FormEl(
		htmx.HxPost("/dashboard/webhooks/create"),
		htmx.HxDisabledELT("find button[type='submit']"),
		nodx.Class("space-y-2"),

		createAndUpdateWebhookForm(databases, destinations, backups),

		nodx.Div(
			nodx.Class("flex justify-end items-center space-x-2 pt-2"),
			component.HxLoadingMd(),
			nodx.Button(
				nodx.Class("btn btn-primary"),
				nodx.Type("submit"),
				component.SpanText("Save"),
				lucide.Save(),
			),
		),
	)
}

func createWebhookButton() nodx.Node {
	mo := component.Modal(component.ModalParams{
		Size:  component.SizeLg,
		Title: "Create webhook",
		Content: []nodx.Node{
			nodx.Div(
				htmx.HxGet("/dashboard/webhooks/create"),
				htmx.HxSwap("outerHTML"),
				htmx.HxTrigger("intersect once"),
				nodx.Class("p-10 flex justify-center"),
				component.HxLoadingMd(),
			),
		},
	})

	button := nodx.Button(
		mo.OpenerAttr,
		nodx.Class("btn btn-primary"),
		component.SpanText("Create webhook"),
		lucide.Plus(),
	)

	return nodx.Div(
		nodx.Class("inline-block"),
		mo.HTML,
		button,
	)
}
