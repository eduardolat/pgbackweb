package databases

import (
	"database/sql"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) editDatabaseHandler(c echo.Context) error {
	ctx := c.Request().Context()

	databaseID, err := uuid.Parse(c.Param("databaseID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	var formData createDatabaseDTO
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	_, err = h.servs.DatabasesService.UpdateDatabase(
		ctx, dbgen.DatabasesServiceUpdateDatabaseParams{
			ID:               databaseID,
			Name:             sql.NullString{String: formData.Name, Valid: true},
			PgVersion:        sql.NullString{String: formData.Version, Valid: true},
			ConnectionString: sql.NullString{String: formData.ConnectionString, Valid: true},
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondToastSuccess(c, "Database updated")
}

func editDatabaseButton(
	database dbgen.DatabasesServicePaginateDatabasesRow,
) gomponents.Node {
	idPref := "edit-database-" + database.ID.String()
	formID := idPref + "-form"
	btnClass := idPref + "-btn"
	loadingID := idPref + "-loading"

	htmxAttributes := func(url string) gomponents.Node {
		return gomponents.Group([]gomponents.Node{
			htmx.HxPost(url),
			htmx.HxInclude("#" + formID),
			htmx.HxDisabledELT("." + btnClass),
			htmx.HxIndicator("#" + loadingID),
			htmx.HxValidate("true"),
		})
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Edit database",
		Content: []gomponents.Node{
			html.Form(
				html.ID(formID),
				html.Class("space-y-2"),

				component.InputControl(component.InputControlParams{
					Name:        "name",
					Label:       "Name",
					Placeholder: "My database",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "A name to easily identify the database",
					Children: []gomponents.Node{
						html.Value(database.Name),
					},
				}),

				component.SelectControl(component.SelectControlParams{
					Name:     "version",
					Label:    "Version",
					Required: true,
					HelpText: "The version of the database",
					Children: []gomponents.Node{
						html.Option(
							gomponents.If(database.PgVersion == "13", html.Selected()),
							html.Value("13"), gomponents.Text("PostgreSQL 13"),
						),
						html.Option(
							gomponents.If(database.PgVersion == "14", html.Selected()),
							html.Value("14"), gomponents.Text("PostgreSQL 14"),
						),
						html.Option(
							gomponents.If(database.PgVersion == "15", html.Selected()),
							html.Value("15"), gomponents.Text("PostgreSQL 15"),
						),
						html.Option(
							gomponents.If(database.PgVersion == "16", html.Selected()),
							html.Value("16"), gomponents.Text("PostgreSQL 16"),
						),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "connection_string",
					Label:       "Connection string",
					Placeholder: "postgresql://user:password@localhost:5432/mydb",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "It should be a valid PostgreSQL connection string including the database name. It will be stored securely using PGP encryption.",
					Children: []gomponents.Node{
						html.Value(database.DecryptedConnectionString),
					},
				}),
			),

			html.Div(
				html.Class("flex justify-between items-center pt-4"),
				html.Div(
					html.Button(
						htmxAttributes("/dashboard/databases/test"),
						components.Classes{
							btnClass:                      true,
							"btn btn-neutral btn-outline": true,
						},
						html.Type("button"),
						component.SpanText("Test connection"),
						lucide.DatabaseZap(),
					),
				),
				html.Div(
					html.Class("flex justify-end items-center space-x-2"),
					component.HxLoadingMd(loadingID),
					html.Button(
						htmxAttributes("/dashboard/databases/"+database.ID.String()+"/edit"),
						components.Classes{
							btnClass:          true,
							"btn btn-primary": true,
						},
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
		html.Class("btn btn-neutral btn-sm btn-square btn-ghost"),
		lucide.Pencil(),
	)

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		html.Div(
			html.Class("inline-block tooltip tooltip-right"),
			html.Data("tip", "Edit database"),
			button,
		),
	)
}
