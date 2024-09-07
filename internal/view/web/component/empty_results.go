package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type EmptyResultsParams struct {
	Title    string
	Subtitle string
}

func EmptyResults(params EmptyResultsParams) gomponents.Node {
	return html.Div(
		html.Class("flex flex-col justify-center items-center space-x-1"),
		lucide.FileSearch(html.Class("size-8")),
		gomponents.If(
			params.Title != "",
			html.Span(
				html.Class("text-xl"),
				gomponents.Text(params.Title),
			),
		),
		gomponents.If(
			params.Subtitle != "",
			html.Span(
				html.Class("text-base"),
				gomponents.Text(params.Subtitle),
			),
		),
	)
}

func EmptyResultsTr(params EmptyResultsParams) gomponents.Node {
	return html.Tr(
		html.Td(
			html.ColSpan("100%"),
			html.Class("py-10"),
			EmptyResults(params),
		),
	)
}
