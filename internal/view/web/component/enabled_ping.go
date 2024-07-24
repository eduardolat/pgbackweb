package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func EnabledPing(enabled bool) gomponents.Node {
	return html.Span(
		html.Class("relative flex h-3 w-3"),
		html.Span(
			components.Classes{
				"absolute inline-flex h-full w-full":   true,
				"animate-ping rounded-full opacity-75": true,
				"bg-success":                           enabled,
				"bg-error":                             !enabled,
			},
		),
		html.Span(
			components.Classes{
				"relative inline-flex rounded-full h-3 w-3": true,
				"bg-success": enabled,
				"bg-error":   !enabled,
			},
		),
	)
}
