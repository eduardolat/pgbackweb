package databases

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmxs"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
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
		return htmxs.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	_, err := h.servs.DatabasesService.CreateDatabase(
		ctx, dbgen.DatabasesServiceCreateDatabaseParams{
			Name:             formData.Name,
			PgVersion:        formData.Version,
			ConnectionString: formData.ConnectionString,
		},
	)
	if err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	return htmxs.RespondRedirect(c, "/dashboard/databases")
}

func createDatabaseButton() nodx.Node {
	htmxAttributes := func(url string) nodx.Node {
		return nodx.Group(
			htmx.HxPost(url),
			htmx.HxInclude("#create-database-form"),
			htmx.HxDisabledELT(".create-database-btn"),
			htmx.HxIndicator("#create-database-loading"),
			htmx.HxValidate("true"),
		)
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Create database",
		Content: []nodx.Node{
			nodx.FormEl(
				nodx.Id("create-database-form"),
				nodx.Class("space-y-2"),

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
					Children: []nodx.Node{
						component.PGVersionSelectOptions(sql.NullString{}),
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

			nodx.Div(
				nodx.Class("flex justify-between items-center pt-4"),
				nodx.Div(
					nodx.Button(
						htmxAttributes("/dashboard/databases/test"),
						nodx.Class("create-database-btn btn btn-neutral btn-outline"),
						nodx.Type("button"),
						component.SpanText("Test connection"),
						lucide.DatabaseZap(),
					),
				),
				nodx.Div(
					nodx.Class("flex justify-end items-center space-x-2"),
					component.HxLoadingMd("create-database-loading"),
					nodx.Button(
						htmxAttributes("/dashboard/databases"),
						nodx.Class("create-database-btn btn btn-primary"),
						nodx.Type("button"),
						component.SpanText("Save"),
						lucide.Save(),
					),
				),
			),
		},
	})

	button := nodx.Button(
		mo.OpenerAttr,
		nodx.Class("btn btn-primary"),
		component.SpanText("Create database"),
		lucide.Plus(),
	)

	return nodx.Div(
		nodx.Class("inline-block"),
		mo.HTML,
		button,
	)
}
