package executions

import (
	"net/http"
	"path/filepath"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) downloadExecutionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	executionID, err := uuid.Parse(c.Param("executionID"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	isLocal, link, err := h.servs.ExecutionsService.GetExecutionDownloadLinkOrPath(
		ctx, executionID,
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if isLocal {
		return c.Attachment(link, filepath.Base(link))
	}

	return c.Redirect(http.StatusFound, link)
}

func showExecutionButton(
	execution dbgen.ExecutionsServicePaginateExecutionsRow,
) gomponents.Node {
	mo := component.Modal(component.ModalParams{
		Title: "Execution details",
		Size:  component.SizeMd,
		Content: []gomponents.Node{
			html.Div(
				html.Class("overflow-x-auto"),
				html.Table(
					html.Class("table"),
					html.Tr(
						html.Th(component.SpanText("ID")),
						html.Td(component.SpanText(execution.ID.String())),
					),
					html.Tr(
						html.Th(component.SpanText("Status")),
						html.Td(component.StatusBadge(execution.Status)),
					),
					html.Tr(
						html.Th(component.SpanText("Database")),
						html.Td(component.SpanText(execution.DatabaseName)),
					),
					html.Tr(
						html.Th(component.SpanText("Destination")),
						html.Td(component.PrettyDestinationName(
							execution.BackupIsLocal, execution.DestinationName,
						)),
					),
					gomponents.If(
						execution.Message.Valid,
						html.Tr(
							html.Th(component.SpanText("Message")),
							html.Td(
								html.Class("break-all"),
								component.SpanText(execution.Message.String),
							),
						),
					),
					html.Tr(
						html.Th(component.SpanText("Started at")),
						html.Td(component.SpanText(
							execution.StartedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
						)),
					),
					gomponents.If(
						execution.FinishedAt.Valid,
						html.Tr(
							html.Th(component.SpanText("Finished at")),
							html.Td(component.SpanText(
								execution.FinishedAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
							)),
						),
					),
					gomponents.If(
						execution.FinishedAt.Valid,
						html.Tr(
							html.Th(component.SpanText("Took")),
							html.Td(component.SpanText(
								execution.FinishedAt.Time.Sub(execution.StartedAt).String(),
							)),
						),
					),
					gomponents.If(
						execution.DeletedAt.Valid,
						html.Tr(
							html.Th(component.SpanText("Deleted at")),
							html.Td(component.SpanText(
								execution.DeletedAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
							)),
						),
					),
					gomponents.If(
						execution.FileSize.Valid,
						html.Tr(
							html.Th(component.SpanText("File size")),
							html.Td(component.PrettyFileSize(execution.FileSize)),
						),
					),
				),
				gomponents.If(
					execution.Status == "success",
					html.Div(
						html.Class("flex justify-end items-center space-x-2"),
						deleteExecutionButton(execution.ID),
						html.A(
							html.Href("/dashboard/executions/"+execution.ID.String()+"/download"),
							html.Target("_blank"),
							html.Class("btn btn-primary"),
							component.SpanText("Download"),
							lucide.Download(),
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
