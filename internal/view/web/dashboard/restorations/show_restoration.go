package restorations

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func showRestorationButton(
	restoration dbgen.RestorationsServicePaginateRestorationsRow,
) nodx.Node {
	mo := component.Modal(component.ModalParams{
		Title: "Restoration details",
		Size:  component.SizeMd,
		Content: []nodx.Node{
			nodx.Div(
				nodx.Class("overflow-x-auto"),
				nodx.Table(
					nodx.Class("table [&_th]:text-nowrap"),
					nodx.Tr(
						nodx.Th(component.SpanText("ID")),
						nodx.Td(component.SpanText(restoration.ID.String())),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Status")),
						nodx.Td(component.StatusBadge(restoration.Status)),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Backup")),
						nodx.Td(component.SpanText(restoration.BackupName)),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Database")),
						nodx.Td(component.SpanText(func() string {
							if restoration.DatabaseName.Valid {
								return restoration.DatabaseName.String
							}
							return "Other database"
						}())),
					),
					nodx.If(
						restoration.Message.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText("Message")),
							nodx.Td(
								nodx.Class("break-all"),
								component.SpanText(restoration.Message.String),
							),
						),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Started At")),
						nodx.Td(component.SpanText(
							restoration.StartedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
						)),
					),
					nodx.If(
						restoration.FinishedAt.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText("Finished At")),
							nodx.Td(component.SpanText(
								restoration.FinishedAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
							)),
						),
					),
					nodx.If(
						restoration.FinishedAt.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText("Took")),
							nodx.Td(component.SpanText(
								restoration.FinishedAt.Time.Sub(restoration.StartedAt).String(),
							)),
						),
					),
				),
			),
		},
	})

	button := nodx.Button(
		mo.OpenerAttr,
		nodx.Class("btn btn-square btn-sm btn-ghost"),
		lucide.Eye(),
	)

	return nodx.Div(
		nodx.Class("inline-block tooltip tooltip-right"),
		nodx.Data("tip", "Show details"),
		mo.HTML,
		button,
	)
}
