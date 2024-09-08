package executions

import (
	"context"
	"fmt"
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) restoreExecutionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData struct {
		ExecutionID uuid.UUID `form:"execution_id" validate:"required,uuid"`
		DatabaseID  uuid.UUID `form:"database_id" validate:"omitempty,uuid"`
		ConnString  string    `form:"conn_string" validate:"omitempty"`
	}
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	if formData.DatabaseID == uuid.Nil && formData.ConnString == "" {
		return htmx.RespondToastError(
			c, "Database or connection string is required",
		)
	}

	if formData.DatabaseID != uuid.Nil && formData.ConnString != "" {
		return htmx.RespondToastError(
			c, "Database and connection string cannot be both set",
		)
	}

	execution, err := h.servs.ExecutionsService.GetExecution(
		ctx, formData.ExecutionID,
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if formData.ConnString != "" {
		err := h.servs.DatabasesService.TestDatabase(
			ctx, execution.DatabasePgVersion, formData.ConnString,
		)
		if err != nil {
			return htmx.RespondToastError(c, err.Error())
		}
	}

	go func() {
		ctx := context.Background()
		_ = h.servs.RestorationsService.RunRestoration(
			ctx,
			formData.ExecutionID,
			uuid.NullUUID{
				Valid: formData.DatabaseID != uuid.Nil,
				UUID:  formData.DatabaseID,
			},
			formData.ConnString,
		)
	}()

	return htmx.RespondToastSuccess(
		c, "Process started, check the restorations page for more details",
	)
}

func (h *handlers) restoreExecutionFormHandler(c echo.Context) error {
	ctx := c.Request().Context()

	executionID, err := uuid.Parse(c.Param("executionID"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	execution, err := h.servs.ExecutionsService.GetExecution(ctx, executionID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	databases, err := h.servs.DatabasesService.GetAllDatabases(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return echoutil.RenderGomponent(c, http.StatusOK, restoreExecutionForm(
		execution, databases,
	))
}

func restoreExecutionForm(
	execution dbgen.ExecutionsServiceGetExecutionRow,
	databases []dbgen.DatabasesServiceGetAllDatabasesRow,
) gomponents.Node {
	return html.Form(
		htmx.HxPost("/dashboard/executions/"+execution.ID.String()+"/restore"),
		htmx.HxConfirm("Are you sure you want to restore this backup?"),
		htmx.HxDisabledELT("find button"),

		alpine.XData(`{ backup_to: "database" }`),

		html.Input(
			html.Type("hidden"),
			html.Name("execution_id"),
			html.Value(execution.ID.String()),
		),

		html.Div(
			html.Class("space-y-2 text-base"),

			component.SelectControl(component.SelectControlParams{
				Name:     "backup_to",
				Label:    "Backup to",
				Required: true,
				HelpText: "You can restore the backup to an existing database or any other database using a connection string",
				Children: []gomponents.Node{
					alpine.XModel("backup_to"),
					html.Option(
						html.Value("database"),
						gomponents.Text("Existing database"),
						html.Selected(),
					),
					html.Option(
						html.Value("conn_string"),
						gomponents.Text("Other database"),
					),
				},
			}),

			alpine.Template(
				alpine.XIf("backup_to === 'database'"),
				component.SelectControl(component.SelectControlParams{
					Name:        "database_id",
					Label:       "Database",
					Placeholder: "Select a database",
					Required:    true,
					Children: []gomponents.Node{
						component.GMap(
							databases,
							func(db dbgen.DatabasesServiceGetAllDatabasesRow) gomponents.Node {
								return html.Option(
									html.Value(db.ID.String()),
									gomponents.Text(db.Name),
									gomponents.If(
										db.ID == execution.DatabaseID,
										html.Selected(),
									),
								)
							},
						),
					},
				}),
			),

			alpine.Template(
				alpine.XIf("backup_to === 'conn_string'"),
				component.InputControl(component.InputControlParams{
					Name:        "conn_string",
					Label:       "Connection string",
					Placeholder: "postgresql://user:password@localhost:5432/mydb",
					Type:        component.InputTypeText,
					Required:    true,
				}),
			),

			html.Div(
				html.Class("pt-2"),
				html.Div(
					html.Role("alert"),
					html.Class("alert alert-warning"),
					lucide.TriangleAlert(),
					html.Div(
						html.P(
							component.BText(fmt.Sprintf(
								"This restoration uses psql v%s", execution.DatabasePgVersion,
							)),
						),
						component.PText(`
							Please make sure the database you are restoring to is compatible
							with this version of psql and double-check that the picked
							database is the one you want to restore to.
						`),
					),
				),
			),

			html.Div(
				html.Class("flex justify-end items-center space-x-2 pt-2"),
				component.HxLoadingMd(),
				html.Button(
					html.Class("btn btn-primary"),
					html.Type("submit"),
					component.SpanText("Start restoration"),
					lucide.Zap(),
				),
			),
		),
	)
}

func restoreExecutionButton(execution dbgen.ExecutionsServicePaginateExecutionsRow) gomponents.Node {
	if execution.Status != "success" || !execution.Path.Valid {
		return nil
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Restore backup execution",
		Content: []gomponents.Node{
			html.Div(
				htmx.HxGet("/dashboard/executions/"+execution.ID.String()+"/restore-form"),
				htmx.HxSwap("outerHTML"),
				htmx.HxTrigger("intersect once"),
				html.Class("p-10 flex justify-center"),
				component.HxLoadingMd(),
			),
		},
	})

	return html.Div(
		mo.HTML,
		component.OptionsDropdownButton(
			mo.OpenerAttr,
			lucide.ArchiveRestore(),
			component.SpanText("Restore execution"),
		),
	)
}
