package webhooks

import (
	"database/sql"
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
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
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
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
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRedirect(c, "/dashboard/webhooks")
}

func (h *handlers) createWebhookFormHandler(c echo.Context) error {
	ctx := c.Request().Context()

	databases, err := h.servs.DatabasesService.GetAllDatabases(ctx)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	destinations, err := h.servs.DestinationsService.GetAllDestinations(ctx)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	backups, err := h.servs.BackupsService.GetAllBackups(ctx)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(c, http.StatusOK, createWebhookForm(
		databases, destinations, backups,
	))
}

func createWebhookForm(
	databases []dbgen.DatabasesServiceGetAllDatabasesRow,
	destinations []dbgen.DestinationsServiceGetAllDestinationsRow,
	backups []dbgen.Backup,
) gomponents.Node {
	return html.Form(
		htmx.HxPost("/dashboard/webhooks"),
		htmx.HxDisabledELT("find button[type='submit']"),
		html.Class("space-y-2"),

		createAndUpdateWebhookForm(databases, destinations, backups),

		html.Div(
			html.Class("flex justify-end items-center space-x-2 pt-2"),
			component.HxLoadingMd(),
			html.Button(
				html.Class("btn btn-primary"),
				html.Type("submit"),
				component.SpanText("Save"),
				lucide.Save(),
			),
		),
	)
}

func createWebhookButton() gomponents.Node {
	mo := component.Modal(component.ModalParams{
		Size:  component.SizeLg,
		Title: "Create webhook",
		Content: []gomponents.Node{
			html.Div(
				htmx.HxGet("/dashboard/webhooks/create-form"),
				htmx.HxSwap("outerHTML"),
				htmx.HxTrigger("intersect once"),
				html.Class("p-10 flex justify-center"),
				component.HxLoadingMd(),
			),
		},
	})

	button := html.Button(
		mo.OpenerAttr,
		html.Class("btn btn-primary"),
		component.SpanText("Create webhook"),
		lucide.Plus(),
	)

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		button,
	)
}
