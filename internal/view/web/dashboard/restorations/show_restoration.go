package restorations

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func showRestorationButton(
	restoration dbgen.RestorationsServicePaginateRestorationsRow,
) gomponents.Node {
	mo := component.Modal(component.ModalParams{
		Title: "Restoration details",
		Size:  component.SizeMd,
		Content: []gomponents.Node{
			html.Div(
				html.Class("overflow-x-auto"),
				html.Table(
					html.Class("table"),
					html.Tr(
						html.Th(component.SpanText("ID")),
						html.Td(component.SpanText(restoration.ID.String())),
					),
					html.Tr(
						html.Th(component.SpanText("Status")),
						html.Td(component.StatusBadge(restoration.Status)),
					),
					html.Tr(
						html.Th(component.SpanText("Backup")),
						html.Td(component.SpanText(restoration.BackupName)),
					),
					html.Tr(
						html.Th(component.SpanText("Database")),
						html.Td(component.SpanText(func() string {
							if restoration.DatabaseName.Valid {
								return restoration.DatabaseName.String
							}
							return "Other database"
						}())),
					),
					gomponents.If(
						restoration.Message.Valid,
						html.Tr(
							html.Th(component.SpanText("Message")),
							html.Td(
								html.Class("break-all"),
								component.SpanText(restoration.Message.String),
							),
						),
					),
					html.Tr(
						html.Th(component.SpanText("Started At")),
						html.Td(component.SpanText(
							restoration.StartedAt.Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
						)),
					),
					gomponents.If(
						restoration.FinishedAt.Valid,
						html.Tr(
							html.Th(component.SpanText("Finished At")),
							html.Td(component.SpanText(
								restoration.FinishedAt.Time.Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
							)),
						),
					),
					gomponents.If(
						restoration.FinishedAt.Valid,
						html.Tr(
							html.Th(component.SpanText("Took")),
							html.Td(component.SpanText(
								restoration.FinishedAt.Time.Sub(restoration.StartedAt).String(),
							)),
						),
					),
				),
			),
		},
	})

	button := html.Button(
		mo.OpenerAttr,
		html.Class("btn btn-square btn-sm btn-ghost"),
		lucide.Eye(),
	)

	return html.Div(
		html.Class("inline-block tooltip tooltip-right"),
		html.Data("tip", "Show details"),
		mo.HTML,
		button,
	)
}
