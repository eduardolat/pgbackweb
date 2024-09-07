package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func SkeletonTr(rows int) gomponents.Node {
	rs := make([]gomponents.Node, rows)
	for i := range rs {
		rs[i] = html.Tr(
			html.Td(
				html.ColSpan("100%"),
				html.Div(
					html.Class("animate-pulse h-4 w-full bg-base-300 rounded-badge"),
				),
			),
		)
	}

	return gomponents.Group(rs)

}
