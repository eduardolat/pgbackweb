package summary

import (
	"fmt"
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

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
			reqCtx, databasesQty, destinationsQty, backupsQty, executionsQty,
			restorationsQty,
		),
	)
}

func indexPage(
	reqCtx reqctx.Ctx,
	databasesQty dbgen.DatabasesServiceGetDatabasesQtyRow,
	destinationsQty dbgen.DestinationsServiceGetDestinationsQtyRow,
	backupsQty dbgen.BackupsServiceGetBackupsQtyRow,
	executionsQty dbgen.ExecutionsServiceGetExecutionsQtyRow,
	restorationsQty dbgen.RestorationsServiceGetRestorationsQtyRow,
) gomponents.Node {
	type ChartData struct {
		Label    string
		Labels   []string
		Data     []int32
		BgColors []string
	}

	countCard := func(
		title string,
		count int64,
		chartData ChartData,
	) gomponents.Node {
		chart := func() gomponents.Node {
			notAvailable := html.Div(
				html.Class("size-[218px] flex flex-col justify-center items-center"),
				html.Span(
					html.Class("text-sm text-base-content pb-[32px]"),
					gomponents.Text("Chart waiting for data"),
				),
			)

			if len(chartData.Data) < 1 {
				return notAvailable
			}

			areAllZero := true
			for _, d := range chartData.Data {
				if d != 0 {
					areAllZero = false
					break
				}
			}

			if areAllZero {
				return notAvailable
			}

			chartID := "chart-" + uuid.NewString()
			labels := ""
			for _, label := range chartData.Labels {
				labels += fmt.Sprintf("'%s',", label)
			}

			data := ""
			for _, d := range chartData.Data {
				data += fmt.Sprintf("%d,", d)
			}

			bgColors := ""
			for _, color := range chartData.BgColors {
				bgColors += fmt.Sprintf("'%s',", color)
			}

			return html.Div(
				html.Class("mt-2"),
				html.Div(html.Canvas(html.ID(chartID))),
				html.Script(gomponents.Raw(`
					new Chart(document.getElementById('`+chartID+`'), {
						type: 'doughnut',
						data: {
							labels: [`+labels+`],
							datasets: [{
								label: '`+chartData.Label+`',
								data: [`+data+`],
								backgroundColor: [`+bgColors+`],
								borderColor: 'rgba(0, 0, 0, 0)',
								borderWidth: 0
							}]
						},
						options: {
							plugins: {
								legend: {
									position: 'bottom'
								}
							}
						}
					});
				`)),
			)
		}

		return component.CardBox(component.CardBoxParams{
			Class: "flex-none text-center w-[250px]",
			Children: []gomponents.Node{
				component.H2Text(fmt.Sprintf("%d %s", count, title)),
				chart(),
			},
		})
	}

	const (
		greenColor  = "#00a96e"
		redColor    = "#ff5861"
		yellowColor = "#ffbe00"
		blueColor   = "#00b6ff"
	)

	content := []gomponents.Node{
		html.Div(
			html.Class("flex justify-center"),
			component.H1Text("Summary"),
		),
		html.Div(
			html.Class("mt-4 flex justify-center flex-wrap gap-4"),

			countCard("Databases", databasesQty.All, ChartData{
				Label:    "Quantity",
				Labels:   []string{"Healthy", "Unhealthy"},
				Data:     []int32{databasesQty.Healthy, databasesQty.Unhealthy},
				BgColors: []string{greenColor, redColor},
			}),
			countCard("Destinations", destinationsQty.All, ChartData{
				Label:    "Quantity",
				Labels:   []string{"Healthy", "Unhealthy"},
				Data:     []int32{destinationsQty.Healthy, destinationsQty.Unhealthy},
				BgColors: []string{greenColor, redColor},
			}),
			countCard("Backups", backupsQty.All, ChartData{
				Label:    "Quantity",
				Labels:   []string{"Active", "Inactive"},
				Data:     []int32{backupsQty.Active, backupsQty.Inactive},
				BgColors: []string{greenColor, redColor},
			}),
			countCard("Executions", executionsQty.All, ChartData{
				Label:  "Status",
				Labels: []string{"Running", "Success", "Failed", "Deleted"},
				Data: []int32{
					executionsQty.Running, executionsQty.Success, executionsQty.Failed,
					executionsQty.Deleted,
				},
				BgColors: []string{blueColor, greenColor, redColor, yellowColor},
			}),
			countCard("Restorations", restorationsQty.All, ChartData{
				Label:  "Status",
				Labels: []string{"Running", "Success", "Failed"},
				Data: []int32{
					restorationsQty.Running, restorationsQty.Success,
					restorationsQty.Failed,
				},
				BgColors: []string{blueColor, greenColor, redColor},
			}),
		),
		html.Div(
			alpine.XData("genericSlider(4)"),
			alpine.XCloak(),
			html.Class("mt-6 flex flex-col justify-center items-center mx-auto"),
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
								alpine.XBind("class", "currentSlide >= 1 ? 'step-primary' : ''"),
								gomponents.Text("Create database"),
							),
							html.Li(
								html.Class("step"),
								alpine.XBind("class", "currentSlide >= 2 ? 'step-primary' : ''"),
								gomponents.Text("Create destination"),
							),
							html.Li(
								html.Class("step"),
								alpine.XBind("class", "currentSlide >= 3 ? 'step-primary' : ''"),
								gomponents.Text("Create backup"),
							),
							html.Li(
								html.Class("step"),
								alpine.XBind("class", "currentSlide >= 4 ? 'step-primary' : ''"),
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

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: "Summary",
		Body:  content,
	})
}
