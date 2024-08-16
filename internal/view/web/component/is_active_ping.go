package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func IsActivePing(isActive bool) gomponents.Node {
	return html.Div(
		html.Class("tooltip tooltip-right"),
		gomponents.If(isActive, html.Data("tip", "Active")),
		gomponents.If(!isActive, html.Data("tip", "Inactive")),
		html.Span(
			html.Class("relative flex h-3 w-3"),
			html.Span(
				components.Classes{
					"absolute inline-flex h-full w-full":   true,
					"animate-ping rounded-full opacity-75": true,
					"bg-success":                           isActive,
					"bg-error":                             !isActive,
				},
			),
			html.Span(
				components.Classes{
					"relative inline-flex rounded-full h-3 w-3": true,
					"bg-success": isActive,
					"bg-error":   !isActive,
				},
			),
		),
	)
}
