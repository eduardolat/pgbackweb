package component

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func Ping(color color) gomponents.Node {
	if color.Value == "" {
		color = ColorNeutral
	}

	bgClass := ""
	switch color {
	case ColorPrimary:
		bgClass = "bg-primary"
	case ColorSecondary:
		bgClass = "bg-secondary"
	case ColorAccent:
		bgClass = "bg-accent"
	case ColorNeutral:
		bgClass = "bg-neutral"
	case ColorInfo:
		bgClass = "bg-info"
	case ColorSuccess:
		bgClass = "bg-success"
	case ColorWarning:
		bgClass = "bg-warning"
	case ColorError:
		bgClass = "bg-error"
	}

	return html.Span(
		html.Class("relative flex h-3 w-3"),
		html.Span(
			components.Classes{
				"absolute inline-flex h-full w-full":   true,
				"animate-ping rounded-full opacity-75": true,
				bgClass:                                true,
			},
		),
		html.Span(
			components.Classes{
				"relative inline-flex rounded-full h-3 w-3": true,
				bgClass: true,
			},
		),
	)
}

func IsActivePing(isActive bool) gomponents.Node {
	pingColor := ColorSuccess
	if !isActive {
		pingColor = ColorError
	}

	return html.Div(
		html.Class("tooltip tooltip-right"),
		gomponents.If(isActive, html.Data("tip", "Active")),
		gomponents.If(!isActive, html.Data("tip", "Inactive")),
		Ping(pingColor),
	)
}

func HealthStatusPing(
	testOk sql.NullBool, testError sql.NullString, lastTestAt sql.NullTime,
) gomponents.Node {
	pingColor := ColorWarning
	if testOk.Valid {
		if testOk.Bool {
			pingColor = ColorSuccess
		} else {
			pingColor = ColorError
		}
	}

	var moOpenerAttr, moHTML gomponents.Node

	if testOk.Valid {
		statusText := "Healthy"
		if !testOk.Bool {
			statusText = "Unhealthy"
		}

		mo := Modal(ModalParams{
			Size:  SizeSm,
			Title: "Health check details",
			Content: []gomponents.Node{
				html.Div(
					html.Class("overflow-x-auto"),
					html.Table(
						html.Class("table [&_th]:text-nowrap"),
						html.Tr(
							html.Th(SpanText("Status")),
							html.Td(SpanText(statusText)),
						),
						gomponents.If(
							testError.Valid && testError.String != "",
							html.Tr(
								html.Th(SpanText("Error")),
								html.Td(
									html.Class("break-all"),
									SpanText(testError.String),
								),
							),
						),
						gomponents.If(
							lastTestAt.Valid,
							html.Tr(
								html.Th(SpanText("Tested at")),
								html.Td(SpanText(
									lastTestAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
								)),
							),
						),
						html.Tr(
							html.Td(
								html.ColSpan("2"),
								PText(`
									The health check runs automatically every 10 minutes, when
									PG Back Web starts, and when you click the "Test connection"
									button.
								`),
							),
						),
					),
				),
			},
		})

		moOpenerAttr = mo.OpenerAttr
		moHTML = mo.HTML
	}

	tooltipText := func() string {
		if testOk.Valid {
			if testOk.Bool {
				return "Healthy (click for details)"
			}
			return "Unhealthy (click for details)"
		}
		return "Waiting for next test"
	}()

	return html.Div(
		html.Class("tooltip tooltip-right"),
		html.Data("tip", tooltipText),
		moHTML,
		html.Span(
			moOpenerAttr,
			html.Class("cursor-pointer"),
			Ping(pingColor),
		),
	)
}
