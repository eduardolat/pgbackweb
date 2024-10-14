package summary

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func indexHowTo() gomponents.Node {
	return html.Div(
		alpine.XData("alpineSummaryHowToSlider()"),
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
					component.H3Text("Create S3 destination (optional)"),
					component.PText(`
						To create a destination, click on the "Destinations" menu item on
						the left sidebar. Then click on the "Create destination" button.
						Fill in the form and click on the "Save" button. You can create
						as many destinations as you want to store the backups. If you
						don't want to use S3 destinations and store the backups locally,
						you can skip this step.
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
	)
}
