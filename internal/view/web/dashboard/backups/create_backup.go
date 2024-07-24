package backups

import (
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/staticdata"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) createBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData struct {
		DatabaseID     uuid.UUID `form:"database_id" validate:"required,uuid"`
		DestinationID  uuid.UUID `form:"destination_id" validate:"required,uuid"`
		Name           string    `form:"name" validate:"required"`
		CronExpression string    `form:"cron_expression" validate:"required"`
		TimeZone       string    `form:"time_zone" validate:"required"`
		IsActive       string    `form:"is_active" validate:"required,oneof=true false"`
		DestDir        string    `form:"dest_dir" validate:"required"`
		RetentionDays  int16     `form:"retention_days" validate:"required"`
		OptDataOnly    string    `form:"opt_data_only" validate:"required,oneof=true false"`
		OptSchemaOnly  string    `form:"opt_schema_only" validate:"required,oneof=true false"`
		OptClean       string    `form:"opt_clean" validate:"required,oneof=true false"`
		OptIfExists    string    `form:"opt_if_exists" validate:"required,oneof=true false"`
		OptCreate      string    `form:"opt_create" validate:"required,oneof=true false"`
		OptNoComments  string    `form:"opt_no_comments" validate:"required,oneof=true false"`
	}
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	_, err := h.servs.BackupsService.CreateBackup(
		ctx, dbgen.BackupsServiceCreateBackupParams{
			DatabaseID:     formData.DatabaseID,
			DestinationID:  formData.DestinationID,
			Name:           formData.Name,
			CronExpression: formData.CronExpression,
			TimeZone:       formData.TimeZone,
			IsActive:       formData.IsActive == "true",
			DestDir:        formData.DestDir,
			RetentionDays:  formData.RetentionDays,
			OptDataOnly:    formData.OptDataOnly == "true",
			OptSchemaOnly:  formData.OptSchemaOnly == "true",
			OptClean:       formData.OptClean == "true",
			OptIfExists:    formData.OptIfExists == "true",
			OptCreate:      formData.OptCreate == "true",
			OptNoComments:  formData.OptNoComments == "true",
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRedirect(c, "/dashboard/backups")
}

func (h *handlers) createBackupFormHandler(c echo.Context) error {
	ctx := c.Request().Context()

	databases, err := h.servs.DatabasesService.GetAllDatabases(ctx)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	destinations, err := h.servs.DestinationsService.GetAllDestinations(ctx)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, createBackupForm(databases, destinations),
	)
}

func createBackupForm(
	databases []dbgen.DatabasesServiceGetAllDatabasesRow,
	destinations []dbgen.DestinationsServiceGetAllDestinationsRow,
) gomponents.Node {
	yesNoOptions := func() gomponents.Node {
		return gomponents.Group([]gomponents.Node{
			html.Option(html.Value("true"), gomponents.Text("Yes")),
			html.Option(html.Value("false"), gomponents.Text("No"), html.Selected()),
		})
	}

	return html.Form(
		htmx.HxPost("/dashboard/backups"),
		htmx.HxDisabledELT("find button"),
		html.Class("space-y-2"),

		component.SelectControl(component.SelectControlParams{
			Name:        "database_id",
			Label:       "Database",
			Required:    true,
			Placeholder: "Select a database",
			Children: []gomponents.Node{
				component.GMap(
					databases,
					func(db dbgen.DatabasesServiceGetAllDatabasesRow) gomponents.Node {
						return html.Option(html.Value(db.ID.String()), gomponents.Text(db.Name))
					},
				),
			},
		}),

		component.SelectControl(component.SelectControlParams{
			Name:        "destination_id",
			Label:       "Destination",
			Required:    true,
			Placeholder: "Select a destination",
			Children: []gomponents.Node{
				component.GMap(
					destinations,
					func(dest dbgen.DestinationsServiceGetAllDestinationsRow) gomponents.Node {
						return html.Option(html.Value(dest.ID.String()), gomponents.Text(dest.Name))
					},
				),
			},
		}),

		component.InputControl(component.InputControlParams{
			Name:        "name",
			Label:       "Name",
			Placeholder: "My backup",
			Required:    true,
			Type:        component.InputTypeText,
			HelpText:    "A name to easily identify the backup",
		}),

		component.InputControl(component.InputControlParams{
			Name:        "cron_expression",
			Label:       "Cron expression",
			Placeholder: "* * * * *",
			Required:    true,
			Type:        component.InputTypeText,
			HelpText:    "The cron expression to schedule the backup",
		}),

		component.SelectControl(component.SelectControlParams{
			Name:        "time_zone",
			Label:       "Time zone",
			Required:    true,
			Placeholder: "Select a time zone",
			HelpText:    "The time zone in which the cron expression will be evaluated",
			Children: []gomponents.Node{
				component.GMap(
					staticdata.Timezones,
					func(tz staticdata.Timezone) gomponents.Node {
						return html.Option(html.Value(tz.TzCode), gomponents.Text(tz.Label))
					},
				),
			},
		}),

		component.InputControl(component.InputControlParams{
			Name:        "dest_dir",
			Label:       "Destination directory",
			Placeholder: "/path/to/backup",
			Required:    true,
			Type:        component.InputTypeText,
		}),

		component.InputControl(component.InputControlParams{
			Name:        "retention_days",
			Label:       "Retention days",
			Placeholder: "30",
			Required:    true,
			Type:        component.InputTypeNumber,
			HelpText:    "The number of days to keep the backup. All backups older than this will be deleted",
			Children: []gomponents.Node{
				html.Min("0"),
				html.Max("36500"),
				html.Pattern("[0-9]+"),
			},
		}),

		component.SelectControl(component.SelectControlParams{
			Name:     "is_active",
			Label:    "Activate backup",
			Required: true,
			Children: []gomponents.Node{
				html.Option(html.Value("true"), gomponents.Text("Yes")),
				html.Option(html.Value("false"), gomponents.Text("No")),
			},
		}),

		html.Div(
			html.Class("grid grid-cols-2 gap-2"),

			component.SelectControl(component.SelectControlParams{
				Name:     "opt_data_only",
				Label:    "--data-only",
				Required: true,
				Children: []gomponents.Node{
					yesNoOptions(),
				},
			}),

			component.SelectControl(component.SelectControlParams{
				Name:     "opt_schema_only",
				Label:    "--schema-only",
				Required: true,
				Children: []gomponents.Node{
					yesNoOptions(),
				},
			}),

			component.SelectControl(component.SelectControlParams{
				Name:     "opt_clean",
				Label:    "--clean",
				Required: true,
				Children: []gomponents.Node{
					yesNoOptions(),
				},
			}),

			component.SelectControl(component.SelectControlParams{
				Name:     "opt_if_exists",
				Label:    "--if-exists",
				Required: true,
				Children: []gomponents.Node{
					yesNoOptions(),
				},
			}),

			component.SelectControl(component.SelectControlParams{
				Name:     "opt_create",
				Label:    "--create",
				Required: true,
				Children: []gomponents.Node{
					yesNoOptions(),
				},
			}),

			component.SelectControl(component.SelectControlParams{
				Name:     "opt_no_comments",
				Label:    "--no-comments",
				Required: true,
				Children: []gomponents.Node{
					yesNoOptions(),
				},
			}),
		),

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

func createBackupButton() gomponents.Node {
	mo := component.Modal(component.ModalParams{
		Size:  component.SizeLg,
		Title: "Create backup",
		Content: []gomponents.Node{
			html.Div(
				htmx.HxGet("/dashboard/backups/create-form"),
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
		component.SpanText("Create backup"),
		lucide.Plus(),
	)

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		button,
	)
}
