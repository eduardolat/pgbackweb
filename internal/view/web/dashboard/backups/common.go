package backups

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
	"time"
)

func localBackupsHelp() []gomponents.Node {
	return []gomponents.Node{
		component.H3Text("Local backups"),
		component.PText(`
			Local backups are stored in the server where PG Back Web is running.
			They are stored under /backups directory so you can mount a docker
			volume to this directory to persist the backups in any way you want.
		`),

		html.Div(
			html.Class("mt-2"),
			component.H3Text("Remote backups"),
			component.PText(`
				Remote backups are stored in a destination. A destination is a remote
				S3 compatible storage. With this option you don't need to worry about
				creating and managing docker volumes.
			`),
		),
	}
}

func cronExpressionHelp() []gomponents.Node {
	return []gomponents.Node{
		component.PText(`
			A cron expression is a string used to define a schedule for running tasks
			in Unix-like operating systems. It consists of five fields representing
			the minute, hour, day of the month, month, and day of the week.
			Cron expressions enable precise scheduling of periodic tasks.
		`),

		html.Div(
			html.Class("mt-4 flex justify-end items-center space-x-1"),
			html.A(
				html.Href("https://en.wikipedia.org/wiki/Cron"),
				html.Target("_blank"),
				html.Class("btn btn-ghost"),
				component.SpanText("Learn more"),
				lucide.ExternalLink(),
			),
			html.A(
				html.Href("https://crontab.guru/examples.html"),
				html.Target("_blank"),
				html.Class("btn btn-ghost"),
				component.SpanText("Examples & common expressions"),
				lucide.ExternalLink(),
			),
		),
	}
}

func timezoneFilenamesHelp() []gomponents.Node {
	serverTimezone := time.Now().Location().String()

	return []gomponents.Node{
		html.P(
			component.SpanText(`
				Only for cron evaluation.
                Backup filenames will always use the server timezone (currently 
			`),
			component.BText(serverTimezone),
			component.SpanText(")."),
		),

		html.Div(
			html.Class("mt-4 flex justify-end items-center"),
			html.A(
				html.Href("https://github.com/eduardolat/pgbackweb?tab=readme-ov-file#configuration"),
				html.Target("_blank"),
				html.Class("btn btn-ghost"),
				component.SpanText("Learn more in project README"),
				lucide.ExternalLink(),
			),
		),
	}
}

func destinationDirectoryHelp() []gomponents.Node {
	return []gomponents.Node{
		component.PText(`
			The destination directory is the directory where the backups will be
			stored. This directory is relative to the base directory of the
			destination. It should start with a slash, should not contain any
			spaces, and should not end with a slash.
		`),

		html.Div(
			html.Class("mt-2"),
			component.H3Text("Local backups"),
			component.PText(`
				For local backups, the base directory is /backups. So, the backup files
				will be stored in:
			`),
			html.Div(
				components.Classes{
					"whitespace-nowrap p-1": true,
					"overflow-x-scroll":     true,
					"font-mono":             true,
				},
				component.BText(
					"/backups/<destination-directory>/<YYYY>/<MM>/<DD>/dump-<random-suffix>.zip",
				),
			),
		),

		html.Div(
			html.Class("mt-2"),
			component.H3Text("Remote backups"),
			component.PText(`
				For remote backups, the base directory is the root of the bucket. So,
				the backup files will be stored in:
			`),
			html.Div(
				components.Classes{
					"whitespace-nowrap p-1": true,
					"overflow-x-scroll":     true,
					"font-mono":             true,
				},
				component.BText(
					"s3://<bucket>/<destination-directory>/<YYYY>/<MM>/<DD>/dump-<random-suffix>.zip",
				),
			),
		),
	}
}

func retentionDaysHelp() []gomponents.Node {
	return []gomponents.Node{
		html.Div(
			html.Class("space-y-2"),

			component.PText(`
				Retention days specifies the number of days to keep backup files before
				they are automatically deleted. This ensures that old backups are removed
				to save storage space. The retention period is evaluated by execution.
			`),

			component.PText(`
				If you set the retention days to 0, the backups will never be deleted.
			`),
		),
	}
}

func pgDumpOptionsHelp() []gomponents.Node {
	return []gomponents.Node{
		html.Div(
			html.Class("space-y-2"),

			component.PText(`
				This software uses the battle tested pg_dump utility to create backups. It
				makes consistent backups even if the database is being used concurrently.
			`),

			component.PText(`
				These are options that will be passed to the pg_dump utility. By default,
				PG Back Web does not pass any options so the backups are full backups.
			`),

			html.Div(
				html.Class("flex justify-end"),
				html.A(
					html.Class("btn btn-ghost"),
					html.Href("https://www.postgresql.org/docs/current/app-pgdump.html"),
					html.Target("_blank"),
					component.SpanText("Learn more in pg_dump documentation"),
					lucide.ExternalLink(html.Class("ml-1")),
				),
			),
		),
	}
}
