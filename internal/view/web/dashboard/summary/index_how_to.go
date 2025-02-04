package summary

import (
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func indexHowTo() nodx.Node {
	return nodx.Div(
		alpine.XData("alpineSummaryHowToSlider()"),
		alpine.XCloak(),
		nodx.Class("mt-6 flex flex-col justify-center items-center mx-auto"),

		component.H2Text("How to use PG Back Web"),

		component.CardBox(component.CardBoxParams{
			Class: "mt-4 space-y-4 max-w-[600px]",
			Children: []nodx.Node{
				nodx.Div(
					nodx.Class("flex justify-center"),
					nodx.Ul(
						nodx.Class("steps"),
						nodx.Li(
							nodx.Class("step"),
							alpine.XBind("class", "currentSlide >= 1 ? 'step-primary' : ''"),
							nodx.Text("Create database"),
						),
						nodx.Li(
							nodx.Class("step"),
							alpine.XBind("class", "currentSlide >= 2 ? 'step-primary' : ''"),
							nodx.Text("Create destination"),
						),
						nodx.Li(
							nodx.Class("step"),
							alpine.XBind("class", "currentSlide >= 3 ? 'step-primary' : ''"),
							nodx.Text("Create backup"),
						),
						nodx.Li(
							nodx.Class("step"),
							alpine.XBind("class", "currentSlide >= 4 ? 'step-primary' : ''"),
							nodx.Text("Wait for executions"),
						),
					),
				),

				nodx.Div(
					alpine.XShow("currentSlide === 1"),
					component.H3Text("Create database"),
					component.PText(`
						To create a database, click on the "Databases" menu item on the
						left sidebar. Then click on the "Create database" button. Fill
						in the form and click on the "Save" button. You can create as
						many databases as you want to backup.
					`),
				),

				nodx.Div(
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

				nodx.Div(
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

				nodx.Div(
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

				nodx.Div(
					nodx.Class("mt-4 flex justify-between"),
					nodx.Button(
						alpine.XBind("disabled", "!hasPrevSlide"),
						alpine.XOn("click", "prevSlide"),
						nodx.Class("btn btn-neutral btn-ghost"),
						lucide.ChevronLeft(),
						component.SpanText("Previous"),
					),
					nodx.Button(
						alpine.XBind("disabled", "!hasNextSlide"),
						alpine.XOn("click", "nextSlide"),
						nodx.Class("btn btn-neutral btn-ghost"),
						component.SpanText("Next"),
						lucide.ChevronRight(),
					),
				),
			},
		}),
	)
}
