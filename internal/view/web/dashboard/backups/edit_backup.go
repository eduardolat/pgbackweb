package backups

import (
	"database/sql"
	"fmt"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/staticdata"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) editBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	var formData struct {
		Name           string `form:"name" validate:"required"`
		CronExpression string `form:"cron_expression" validate:"required"`
		TimeZone       string `form:"time_zone" validate:"required"`
		IsActive       string `form:"is_active" validate:"required,oneof=true false"`
		DestDir        string `form:"dest_dir" validate:"required"`
		RetentionDays  int16  `form:"retention_days"`
		OptDataOnly    string `form:"opt_data_only" validate:"required,oneof=true false"`
		OptSchemaOnly  string `form:"opt_schema_only" validate:"required,oneof=true false"`
		OptClean       string `form:"opt_clean" validate:"required,oneof=true false"`
		OptIfExists    string `form:"opt_if_exists" validate:"required,oneof=true false"`
		OptCreate      string `form:"opt_create" validate:"required,oneof=true false"`
		OptNoComments  string `form:"opt_no_comments" validate:"required,oneof=true false"`
	}
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	_, err = h.servs.BackupsService.UpdateBackup(
		ctx, dbgen.BackupsServiceUpdateBackupParams{
			ID:             backupID,
			Name:           sql.NullString{String: formData.Name, Valid: true},
			CronExpression: sql.NullString{String: formData.CronExpression, Valid: true},
			TimeZone:       sql.NullString{String: formData.TimeZone, Valid: true},
			IsActive:       sql.NullBool{Bool: formData.IsActive == "true", Valid: true},
			DestDir:        sql.NullString{String: formData.DestDir, Valid: true},
			RetentionDays:  sql.NullInt16{Int16: formData.RetentionDays, Valid: true},
			OptDataOnly:    sql.NullBool{Bool: formData.OptDataOnly == "true", Valid: true},
			OptSchemaOnly:  sql.NullBool{Bool: formData.OptSchemaOnly == "true", Valid: true},
			OptClean:       sql.NullBool{Bool: formData.OptClean == "true", Valid: true},
			OptIfExists:    sql.NullBool{Bool: formData.OptIfExists == "true", Valid: true},
			OptCreate:      sql.NullBool{Bool: formData.OptCreate == "true", Valid: true},
			OptNoComments:  sql.NullBool{Bool: formData.OptNoComments == "true", Valid: true},
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondAlertWithRefresh(c, "Backup updated")
}

func editBackupButton(backup dbgen.BackupsServicePaginateBackupsRow) gomponents.Node {
	yesNoOptions := func(value bool) gomponents.Node {
		return gomponents.Group([]gomponents.Node{
			html.Option(
				html.Value("true"),
				gomponents.Text("Yes"),
				gomponents.If(value, html.Selected()),
			),
			html.Option(
				html.Value("false"),
				gomponents.Text("No"),
				gomponents.If(!value, html.Selected()),
			),
		})
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeLg,
		Title: "Edit backup",
		Content: []gomponents.Node{
			html.Form(
				htmx.HxPost("/dashboard/backups/"+backup.ID.String()+"/edit"),
				htmx.HxDisabledELT("find button"),
				html.Class("space-y-2 text-base"),

				component.InputControl(component.InputControlParams{
					Name:        "name",
					Label:       "Name",
					Placeholder: "My backup",
					Required:    true,
					Type:        component.InputTypeText,
					Children: []gomponents.Node{
						html.Value(backup.Name),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "cron_expression",
					Label:       "Cron expression",
					Placeholder: "* * * * *",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "The cron expression to schedule the backup",
					Pattern:     `^\S+\s+\S+\s+\S+\s+\S+\s+\S+$`,
					Children: []gomponents.Node{
						html.Value(backup.CronExpression),
					},
					HelpButtonChildren: cronExpressionHelp(),
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
								return html.Option(
									html.Value(tz.TzCode),
									gomponents.Text(tz.Label),
									gomponents.If(
										tz.TzCode == backup.TimeZone,
										html.Selected(),
									),
								)
							},
						),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:               "dest_dir",
					Label:              "Destination directory",
					Placeholder:        "/path/to/backup",
					Required:           true,
					Type:               component.InputTypeText,
					HelpText:           "Relative to the base directory of the destination",
					HelpButtonChildren: destinationDirectoryHelp(),
					Pattern:            `^\/\S*[^\/]$`,
					Children: []gomponents.Node{
						html.Value(backup.DestDir),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:               "retention_days",
					Label:              "Retention days",
					Placeholder:        "30",
					Required:           true,
					Type:               component.InputTypeNumber,
					Pattern:            "[0-9]+",
					HelpButtonChildren: retentionDaysHelp(),
					Children: []gomponents.Node{
						html.Min("0"),
						html.Max("36500"),
						html.Value(fmt.Sprintf("%d", backup.RetentionDays)),
					},
				}),

				component.SelectControl(component.SelectControlParams{
					Name:     "is_active",
					Label:    "Activate backup",
					Required: true,
					Children: []gomponents.Node{
						yesNoOptions(backup.IsActive),
					},
				}),

				html.Div(
					html.Class("pt-4"),
					html.Div(
						html.Class("flex justify-start items-center space-x-1"),
						component.H2Text("Options"),
						component.HelpButtonModal(component.HelpButtonModalParams{
							ModalTitle: "Backup options",
							Children:   pgDumpOptionsHelp(),
						}),
					),

					html.Div(
						html.Class("mt-2 grid grid-cols-2 gap-2"),
						component.SelectControl(component.SelectControlParams{
							Name:     "opt_data_only",
							Label:    "--data-only",
							Required: true,
							Children: []gomponents.Node{
								yesNoOptions(backup.OptDataOnly),
							},
						}),

						component.SelectControl(component.SelectControlParams{
							Name:     "opt_schema_only",
							Label:    "--schema-only",
							Required: true,
							Children: []gomponents.Node{
								yesNoOptions(backup.OptSchemaOnly),
							},
						}),

						component.SelectControl(component.SelectControlParams{
							Name:     "opt_clean",
							Label:    "--clean",
							Required: true,
							Children: []gomponents.Node{
								yesNoOptions(backup.OptClean),
							},
						}),

						component.SelectControl(component.SelectControlParams{
							Name:     "opt_if_exists",
							Label:    "--if-exists",
							Required: true,
							Children: []gomponents.Node{
								yesNoOptions(backup.OptIfExists),
							},
						}),

						component.SelectControl(component.SelectControlParams{
							Name:     "opt_create",
							Label:    "--create",
							Required: true,
							Children: []gomponents.Node{
								yesNoOptions(backup.OptCreate),
							},
						}),

						component.SelectControl(component.SelectControlParams{
							Name:     "opt_no_comments",
							Label:    "--no-comments",
							Required: true,
							Children: []gomponents.Node{
								yesNoOptions(backup.OptNoComments),
							},
						}),
					),
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
			html.Data("tip", "Edit backup"),
			button,
		),
	)
}
