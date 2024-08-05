package summary

import (
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	ctx := c.Request().Context()

	databasesQty, err := h.servs.DatabasesService.GetDatabasesQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	destinationsQty, err := h.servs.DestinationsService.GetDestinationsQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	backupsQty, err := h.servs.BackupsService.GetBackupsQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	executionsQty, err := h.servs.ExecutionsService.GetExecutionsQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	restorationsQty, err := h.servs.RestorationsService.GetRestorationsQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK,
		indexPage(
			databasesQty, destinationsQty, backupsQty, executionsQty, restorationsQty,
		),
	)
}

func indexPage(
	databasesQty, destinationsQty, backupsQty, executionsQty, restorationsQty int64,
) gomponents.Node {
	countCard := func(title string, count int64) gomponents.Node {
		return component.CardBox(component.CardBoxParams{
			Class: "text-center",
			Children: []gomponents.Node{
				component.H2Text(title),
				html.Span(
					html.Class("text-5xl font-bold"),
					gomponents.Textf("%d", count),
				),
			},
		})
	}

	content := []gomponents.Node{
		component.H1Text("Summary"),
		html.Div(
			html.Class("mt-4 grid grid-cols-5 gap-4"),
			countCard("Databases", databasesQty),
			countCard("Destinations", destinationsQty),
			countCard("Backups", backupsQty),
			countCard("Executions", executionsQty),
			countCard("Restorations", restorationsQty),
		),
		html.Div(
			alpine.XData("genericSlider(4)"),
			alpine.XCloak(),
			html.Class("mt-6 flex flex-col justify-center items-center"),
			component.H2Text("How to use PG Back Web"),

			component.CardBox(component.CardBoxParams{
				Class: "mt-4 space-y-4 max-w-[600px]",
				Children: []gomponents.Node{
					html.Div(
						html.Class("flex justify-center"),
						html.Ul(
							html.Class("steps"),
							html.Li(
								html.Class("step"),
								alpine.XBind("class", "currentSlide === 1 ? 'step-primary' : ''"),
								gomponents.Text("Create database"),
							),
							html.Li(
								html.Class("step"),
								alpine.XBind("class", "currentSlide === 2 ? 'step-primary' : ''"),
								gomponents.Text("Create destination"),
							),
							html.Li(
								html.Class("step"),
								alpine.XBind("class", "currentSlide === 3 ? 'step-primary' : ''"),
								gomponents.Text("Create backup"),
							),
							html.Li(
								html.Class("step"),
								alpine.XBind("class", "currentSlide === 4 ? 'step-primary' : ''"),
								gomponents.Text("Wait for executions"),
							),
						),
					),

					html.Div(
						alpine.XShow("currentSlide === 1"),
						component.H3Text("Create database"),
						component.PText(`
							To create a database, click on the "Databases" menu item on the
							left sidebar. Then click on the "Create database" button. Fill
							in the form and click on the "Save" button. You can create as
							many databases as you want to backup.
						`),
					),

					html.Div(
						alpine.XShow("currentSlide === 2"),
						component.H3Text("Create destination"),
						component.PText(`
							To create a destination, click on the "Destinations" menu item on
							the left sidebar. Then click on the "Create destination" button.
							Fill in the form and click on the "Save" button. You can create
							as many destinations as you want to store the backups.
						`),
					),

					html.Div(
						alpine.XShow("currentSlide === 3"),
						component.H3Text("Create backup"),
						component.PText(`
							To create a backup you need to have at least one database and one
							destination. Click on the "Backups" menu item on the left sidebar.
							Then click on the "Create backup" button. Fill in the form and
							click on the "Save" button. You can create as many backups as you
							want including any combination of databases and destinations.
						`),
					),

					html.Div(
						alpine.XShow("currentSlide === 4"),
						component.H3Text("Wait for executions"),
						component.PText(`
							When your backup is created and active, the system will start
							creating executions based on the schedule you defined. You can
							also create executions manually by clicking the "Run backup now"
							button on the backups list. You can see the executions on the
							"Executions" menu item on the left sidebar and then click on the
							"Show details" button to see the details, logs, and download or
							delete the backup file.
						`),
					),

					html.Div(
						html.Class("mt-4 flex justify-between"),
						html.Button(
							alpine.XBind("disabled", "!hasPrevSlide"),
							alpine.XOn("click", "prevSlide"),
							html.Class("btn btn-neutral btn-ghost"),
							lucide.ChevronLeft(),
							component.SpanText("Previous"),
						),
						html.Button(
							alpine.XBind("disabled", "!hasNextSlide"),
							alpine.XOn("click", "nextSlide"),
							html.Class("btn btn-neutral btn-ghost"),
							component.SpanText("Next"),
							lucide.ChevronRight(),
						),
					),
				},
			}),
		),
	}

	return layout.Dashboard(layout.DashboardParams{
		Title: "Summary",
		Body:  content,
	})
}
