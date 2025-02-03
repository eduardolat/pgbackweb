package component

import (
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

type EmptyResultsParams struct {
	Title    string
	Subtitle string
}

func EmptyResults(params EmptyResultsParams) nodx.Node {
	return nodx.Div(
		nodx.Class("flex flex-col justify-center items-center space-x-1"),
		lucide.FileSearch(nodx.Class("size-8")),
		nodx.If(
			params.Title != "",
			nodx.SpanEl(
				nodx.Class("text-xl"),
				nodx.Text(params.Title),
			),
		),
		nodx.If(
			params.Subtitle != "",
			nodx.SpanEl(
				nodx.Class("text-base"),
				nodx.Text(params.Subtitle),
			),
		),
	)
}

func EmptyResultsTr(params EmptyResultsParams) nodx.Node {
	return nodx.Tr(
		nodx.Td(
			nodx.Colspan("100%"),
			nodx.Class("py-10"),
			EmptyResults(params),
		),
	)
}
