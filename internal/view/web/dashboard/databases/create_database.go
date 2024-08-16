package databases

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type createDatabaseDTO struct {
	Name             string `form:"name" validate:"required"`
	Version          string `form:"version" validate:"required"`
	ConnectionString string `form:"connection_string" validate:"required"`
}

func (h *handlers) createDatabaseHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData createDatabaseDTO
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	_, err := h.servs.DatabasesService.CreateDatabase(
		ctx, dbgen.DatabasesServiceCreateDatabaseParams{
			Name:             formData.Name,
			PgVersion:        formData.Version,
			ConnectionString: formData.ConnectionString,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRedirect(c, "/dashboard/databases")
}

func createDatabaseButton() gomponents.Node {
	htmxAttributes := func(url string) gomponents.Node {
		return gomponents.Group([]gomponents.Node{
			htmx.HxPost(url),
			htmx.HxInclude("#create-database-form"),
			htmx.HxDisabledELT(".create-database-btn"),
			htmx.HxIndicator("#create-database-loading"),
			htmx.HxValidate("true"),
		})
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Create database",
		Content: []gomponents.Node{
			html.Form(
				html.ID("create-database-form"),
				html.Class("space-y-2"),

				component.InputControl(component.InputControlParams{
					Name:        "name",
					Label:       "Name",
					Placeholder: "My database",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "A name to easily identify the database",
				}),

				component.SelectControl(component.SelectControlParams{
					Name:        "version",
					Label:       "Version",
					Placeholder: "Select a version",
					Required:    true,
					HelpText:    "The version of the database",
					Children: []gomponents.Node{
						html.Option(html.Value("13"), gomponents.Text("PostgreSQL 13")),
						html.Option(html.Value("14"), gomponents.Text("PostgreSQL 14")),
						html.Option(html.Value("15"), gomponents.Text("PostgreSQL 15")),
						html.Option(html.Value("16"), gomponents.Text("PostgreSQL 16")),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "connection_string",
					Label:       "Connection string",
					Placeholder: "postgresql://user:password@localhost:5432/mydb",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "It should be a valid PostgreSQL connection string including the database name. It will be stored securely using PGP encryption.",
				}),
			),

			html.Div(
				html.Class("flex justify-between items-center pt-4"),
				html.Div(
					html.Button(
						htmxAttributes("/dashboard/databases/test"),
						html.Class("create-database-btn btn btn-neutral btn-outline"),
						html.Type("button"),
						component.SpanText("Test connection"),
						lucide.DatabaseZap(),
					),
				),
				html.Div(
					html.Class("flex justify-end items-center space-x-2"),
					component.HxLoadingMd("create-database-loading"),
					html.Button(
						htmxAttributes("/dashboard/databases"),
						html.Class("create-database-btn btn btn-primary"),
						html.Type("button"),
						component.SpanText("Save"),
						lucide.Save(),
					),
				),
			),
		},
	})

	button := html.Button(
		mo.OpenerAttr,
		html.Class("btn btn-primary"),
		component.SpanText("Create database"),
		lucide.Plus(),
	)

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		button,
	)
}
