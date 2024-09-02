package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type EmptyResultsParams struct {
	IsLarge  bool
	Title    string
	Subtitle string
}

func EmptyResults(params EmptyResultsParams) gomponents.Node {
	return html.Div(
		html.Class("flex flex-col justify-center items-center space-x-1"),
		lucide.FileSearch(components.Classes{
			"size-6":  !params.IsLarge,
			"size-12": params.IsLarge,
		}),
		gomponents.If(
			params.Title != "",
			html.Span(
				components.Classes{
					"text-lg":  !params.IsLarge,
					"text-2xl": params.IsLarge,
				},
				gomponents.Text(params.Title),
			),
		),
		gomponents.If(
			params.Subtitle != "",
			html.Span(
				components.Classes{
					"text-sm": !params.IsLarge,
					"text-lg": params.IsLarge,
				},
				gomponents.Text(params.Subtitle),
			),
		),
	)
}
