package backups

import (
	"net/http"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/staticdata"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) createBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData struct {
		DatabaseID     uuid.UUID `form:"database_id" validate:"required,uuid"`
		DestinationID  uuid.UUID `form:"destination_id" validate:"omitempty,uuid"`
		IsLocal        string    `form:"is_local" validate:"required,oneof=true false"`
		Name           string    `form:"name" validate:"required"`
		CronExpression string    `form:"cron_expression" validate:"required"`
		TimeZone       string    `form:"time_zone" validate:"required"`
		IsActive       string    `form:"is_active" validate:"required,oneof=true false"`
		DestDir        string    `form:"dest_dir" validate:"required"`
		RetentionDays  int16     `form:"retention_days"`
		OptDataOnly    string    `form:"opt_data_only" validate:"required,oneof=true false"`
		OptSchemaOnly  string    `form:"opt_schema_only" validate:"required,oneof=true false"`
		OptClean       string    `form:"opt_clean" validate:"required,oneof=true false"`
		OptIfExists    string    `form:"opt_if_exists" validate:"required,oneof=true false"`
		OptCreate      string    `form:"opt_create" validate:"required,oneof=true false"`
		OptNoComments  string    `form:"opt_no_comments" validate:"required,oneof=true false"`
	}
	if err := c.Bind(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	_, err := h.servs.BackupsService.CreateBackup(
		ctx, dbgen.BackupsServiceCreateBackupParams{
			DatabaseID: formData.DatabaseID,
			DestinationID: uuid.NullUUID{
				Valid: formData.IsLocal == "false", UUID: formData.DestinationID,
			},
			IsLocal:        formData.IsLocal == "true",
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
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.Redirect(c, pathutil.BuildPath("/dashboard/backups"))
}

func (h *handlers) createBackupFormHandler(c echo.Context) error {
	ctx := c.Request().Context()

	databases, err := h.servs.DatabasesService.GetAllDatabases(ctx)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	destinations, err := h.servs.DestinationsService.GetAllDestinations(ctx)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return echoutil.RenderNodx(
		c, http.StatusOK, createBackupForm(databases, destinations),
	)
}

func createBackupForm(
	databases []dbgen.DatabasesServiceGetAllDatabasesRow,
	destinations []dbgen.DestinationsServiceGetAllDestinationsRow,
) nodx.Node {
	yesNoOptions := func() nodx.Node {
		return nodx.Group(
			nodx.Option(nodx.Value("true"), nodx.Text("Yes")),
			nodx.Option(nodx.Value("false"), nodx.Text("No"), nodx.Selected("")),
		)
	}

	serverTZ := time.Now().Location().String()

	return nodx.FormEl(
		htmx.HxPost(pathutil.BuildPath("/dashboard/backups")),
		htmx.HxDisabledELT("find button"),
		nodx.Class("space-y-2 text-base"),

		alpine.XData(`{
			is_local: "false",
		}`),

		component.InputControl(component.InputControlParams{
			Name:        "name",
			Label:       "Name",
			Placeholder: "My backup",
			Required:    true,
			Type:        component.InputTypeText,
		}),

		component.SelectControl(component.SelectControlParams{
			Name:        "database_id",
			Label:       "Database",
			Required:    true,
			Placeholder: "Select a database",
			Children: []nodx.Node{
				nodx.Map(
					databases,
					func(db dbgen.DatabasesServiceGetAllDatabasesRow) nodx.Node {
						return nodx.Option(nodx.Value(db.ID.String()), nodx.Text(db.Name))
					},
				),
			},
		}),

		component.SelectControl(component.SelectControlParams{
			Name:     "is_local",
			Label:    "Local backup",
			Required: true,
			Children: []nodx.Node{
				alpine.XModel("is_local"),
				nodx.Option(nodx.Value("true"), nodx.Text("Yes")),
				nodx.Option(nodx.Value("false"), nodx.Text("No"), nodx.Selected("")),
			},
			HelpButtonChildren: localBackupsHelp(),
		}),

		alpine.Template(
			alpine.XIf("is_local == 'false'"),
			component.SelectControl(component.SelectControlParams{
				Name:        "destination_id",
				Label:       "Destination",
				Required:    true,
				Placeholder: "Select a destination",
				Children: []nodx.Node{
					nodx.Map(
						destinations,
						func(dest dbgen.DestinationsServiceGetAllDestinationsRow) nodx.Node {
							return nodx.Option(nodx.Value(dest.ID.String()), nodx.Text(dest.Name))
						},
					),
				},
			}),
		),

		component.InputControl(component.InputControlParams{
			Name:               "cron_expression",
			Label:              "Cron expression",
			Placeholder:        "* * * * *",
			Required:           true,
			Type:               component.InputTypeText,
			HelpText:           "The cron expression to schedule the backup",
			Pattern:            `^\S+\s+\S+\s+\S+\s+\S+\s+\S+$`,
			HelpButtonChildren: cronExpressionHelp(),
		}),

		component.SelectControl(component.SelectControlParams{
			Name:        "time_zone",
			Label:       "Time zone",
			Required:    true,
			Placeholder: "Select a time zone",
			Children: []nodx.Node{
				nodx.Map(
					staticdata.Timezones,
					func(tz staticdata.Timezone) nodx.Node {
						var selected nodx.Node
						if tz.TzCode == serverTZ {
							selected = nodx.Selected("")
						}

						return nodx.Option(nodx.Value(tz.TzCode), nodx.Text(tz.Label), selected)
					},
				),
			},
			HelpButtonChildren: timezoneFilenamesHelp(),
		}),

		component.InputControl(component.InputControlParams{
			Name:               "dest_dir",
			Label:              "Destination directory",
			Placeholder:        "/path/to/backup",
			Required:           true,
			Type:               component.InputTypeText,
			HelpText:           "Relative to the base directory of the destination",
			Pattern:            `^\/\S*[^\/]$`,
			HelpButtonChildren: destinationDirectoryHelp(),
		}),

		component.InputControl(component.InputControlParams{
			Name:               "retention_days",
			Label:              "Retention days",
			Placeholder:        "30",
			Required:           true,
			Type:               component.InputTypeNumber,
			Pattern:            "[0-9]+",
			HelpButtonChildren: retentionDaysHelp(),
			Children: []nodx.Node{
				nodx.Min("0"),
				nodx.Max("36500"),
			},
		}),

		component.SelectControl(component.SelectControlParams{
			Name:     "is_active",
			Label:    "Activate backup",
			Required: true,
			Children: []nodx.Node{
				nodx.Option(nodx.Value("true"), nodx.Text("Yes")),
				nodx.Option(nodx.Value("false"), nodx.Text("No")),
			},
		}),

		nodx.Div(
			nodx.Class("pt-4"),
			nodx.Div(
				nodx.Class("flex justify-start items-center space-x-1"),
				component.H2Text("Options"),
				component.HelpButtonModal(component.HelpButtonModalParams{
					ModalTitle: "Backup options",
					Children:   pgDumpOptionsHelp(),
				}),
			),

			nodx.Div(
				nodx.Class("mt-2 grid grid-cols-2 gap-2"),

				component.SelectControl(component.SelectControlParams{
					Name:     "opt_data_only",
					Label:    "--data-only",
					Required: true,
					Children: []nodx.Node{
						yesNoOptions(),
					},
				}),

				component.SelectControl(component.SelectControlParams{
					Name:     "opt_schema_only",
					Label:    "--schema-only",
					Required: true,
					Children: []nodx.Node{
						yesNoOptions(),
					},
				}),

				component.SelectControl(component.SelectControlParams{
					Name:     "opt_clean",
					Label:    "--clean",
					Required: true,
					Children: []nodx.Node{
						yesNoOptions(),
					},
				}),

				component.SelectControl(component.SelectControlParams{
					Name:     "opt_if_exists",
					Label:    "--if-exists",
					Required: true,
					Children: []nodx.Node{
						yesNoOptions(),
					},
				}),

				component.SelectControl(component.SelectControlParams{
					Name:     "opt_create",
					Label:    "--create",
					Required: true,
					Children: []nodx.Node{
						yesNoOptions(),
					},
				}),

				component.SelectControl(component.SelectControlParams{
					Name:     "opt_no_comments",
					Label:    "--no-comments",
					Required: true,
					Children: []nodx.Node{
						yesNoOptions(),
					},
				}),
			),
		),

		nodx.Div(
			nodx.Class("flex justify-end items-center space-x-2 pt-2"),
			component.HxLoadingMd(),
			nodx.Button(
				nodx.Class("btn btn-primary"),
				nodx.Type("submit"),
				component.SpanText("Create backup task"),
				lucide.Save(),
			),
		),
	)
}

func createBackupButton() nodx.Node {
	mo := component.Modal(component.ModalParams{
		Size:  component.SizeLg,
		Title: "Create backup task",
		Content: []nodx.Node{
			nodx.Div(
				htmx.HxGet(pathutil.BuildPath("/dashboard/backups/create-form")),
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
		component.SpanText("Create backup task"),
		lucide.Plus(),
	)

	return nodx.Div(
		nodx.Class("inline-block"),
		mo.HTML,
		button,
	)
}
