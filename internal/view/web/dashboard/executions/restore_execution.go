package executions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
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

func (h *handlers) restoreExecutionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData struct {
		ExecutionID uuid.UUID `form:"execution_id" validate:"required,uuid"`
		DatabaseID  uuid.UUID `form:"database_id" validate:"omitempty,uuid"`
		ConnString  string    `form:"conn_string" validate:"omitempty"`
	}
	if err := c.Bind(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	if formData.DatabaseID == uuid.Nil && formData.ConnString == "" {
		return respondhtmx.ToastError(
			c, "Database or connection string is required",
		)
	}

	if formData.DatabaseID != uuid.Nil && formData.ConnString != "" {
		return respondhtmx.ToastError(
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
			return respondhtmx.ToastError(c, err.Error())
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

	return respondhtmx.ToastSuccess(
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

	return echoutil.RenderNodx(c, http.StatusOK, restoreExecutionForm(
		execution, databases,
	))
}

func restoreExecutionForm(
	execution dbgen.ExecutionsServiceGetExecutionRow,
	databases []dbgen.DatabasesServiceGetAllDatabasesRow,
) nodx.Node {
	return nodx.FormEl(
		htmx.HxPost("/dashboard/executions/"+execution.ID.String()+"/restore"),
		htmx.HxConfirm("Are you sure you want to restore this backup?"),
		htmx.HxDisabledELT("find button"),

		alpine.XData(`{ backup_to: "database" }`),

		nodx.Input(
			nodx.Type("hidden"),
			nodx.Name("execution_id"),
			nodx.Value(execution.ID.String()),
		),

		nodx.Div(
			nodx.Class("space-y-2 text-base"),

			component.SelectControl(component.SelectControlParams{
				Name:     "backup_to",
				Label:    "Backup to",
				Required: true,
				HelpText: "You can restore the backup to an existing database or any other database using a connection string",
				Children: []nodx.Node{
					alpine.XModel("backup_to"),
					nodx.Option(
						nodx.Value("database"),
						nodx.Text("Existing database"),
						nodx.Selected(""),
					),
					nodx.Option(
						nodx.Value("conn_string"),
						nodx.Text("Other database"),
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
					Children: []nodx.Node{
						nodx.Map(
							databases,
							func(db dbgen.DatabasesServiceGetAllDatabasesRow) nodx.Node {
								return nodx.Option(
									nodx.Value(db.ID.String()),
									nodx.Text(db.Name),
									nodx.If(
										db.ID == execution.DatabaseID,
										nodx.Selected(""),
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

			nodx.Div(
				nodx.Class("pt-2"),
				nodx.Div(
					nodx.Role("alert"),
					nodx.Class("alert alert-warning"),
					lucide.TriangleAlert(),
					nodx.Div(
						nodx.P(
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

			nodx.Div(
				nodx.Class("flex justify-end items-center space-x-2 pt-2"),
				component.HxLoadingMd(),
				nodx.Button(
					nodx.Class("btn btn-primary"),
					nodx.Type("submit"),
					component.SpanText("Start restoration"),
					lucide.Zap(),
				),
			),
		),
	)
}

func restoreExecutionButton(execution dbgen.ExecutionsServicePaginateExecutionsRow) nodx.Node {
	if execution.Status != "success" || !execution.Path.Valid {
		return nil
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Restore backup execution",
		Content: []nodx.Node{
			nodx.Div(
				htmx.HxGet("/dashboard/executions/"+execution.ID.String()+"/restore-form"),
				htmx.HxSwap("outerHTML"),
				htmx.HxTrigger("intersect once"),
				nodx.Class("p-10 flex justify-center"),
				component.HxLoadingMd(),
			),
		},
	})

	return nodx.Div(
		mo.HTML,
		component.OptionsDropdownButton(
			mo.OpenerAttr,
			lucide.ArchiveRestore(),
			component.SpanText("Restore execution"),
		),
	)
}
