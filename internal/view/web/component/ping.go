package component

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	nodx "github.com/nodxdev/nodxgo"
)

func Ping(color color) nodx.Node {
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

	return nodx.SpanEl(
		nodx.Class("relative flex h-3 w-3"),
		nodx.SpanEl(
			nodx.ClassMap{
				"absolute inline-flex h-full w-full":   true,
				"animate-ping rounded-full opacity-75": true,
				bgClass:                                true,
			},
		),
		nodx.SpanEl(
			nodx.ClassMap{
				"relative inline-flex rounded-full h-3 w-3": true,
				bgClass: true,
			},
		),
	)
}

func IsActivePing(isActive bool) nodx.Node {
	pingColor := ColorSuccess
	if !isActive {
		pingColor = ColorError
	}

	return nodx.Div(
		nodx.Class("tooltip tooltip-right"),
		nodx.If(isActive, nodx.Data("tip", "Active")),
		nodx.If(!isActive, nodx.Data("tip", "Inactive")),
		Ping(pingColor),
	)
}

func HealthStatusPing(
	testOk sql.NullBool, testError sql.NullString, lastTestAt sql.NullTime,
) nodx.Node {
	pingColor := ColorWarning
	if testOk.Valid {
		if testOk.Bool {
			pingColor = ColorSuccess
		} else {
			pingColor = ColorError
		}
	}

	var moOpenerAttr, moHTML nodx.Node

	if testOk.Valid {
		statusText := "Healthy"
		if !testOk.Bool {
			statusText = "Unhealthy"
		}

		mo := Modal(ModalParams{
			Size:  SizeSm,
			Title: "Health check details",
			Content: []nodx.Node{
				nodx.Div(
					nodx.Class("overflow-x-auto"),
					nodx.Table(
						nodx.Class("table [&_th]:text-nowrap"),
						nodx.Tr(
							nodx.Th(SpanText("Status")),
							nodx.Td(SpanText(statusText)),
						),
						nodx.If(
							testError.Valid && testError.String != "",
							nodx.Tr(
								nodx.Th(SpanText("Error")),
								nodx.Td(
									nodx.Class("break-all"),
									SpanText(testError.String),
								),
							),
						),
						nodx.If(
							lastTestAt.Valid,
							nodx.Tr(
								nodx.Th(SpanText("Tested at")),
								nodx.Td(SpanText(
									lastTestAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
								)),
							),
						),
						nodx.Tr(
							nodx.Td(
								nodx.Colspan("2"),
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

	return nodx.Div(
		nodx.Class("tooltip tooltip-right"),
		nodx.Data("tip", tooltipText),
		moHTML,
		nodx.SpanEl(
			moOpenerAttr,
			nodx.Class("cursor-pointer"),
			Ping(pingColor),
		),
	)
}
