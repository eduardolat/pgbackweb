package component

import (
	nodx "github.com/nodxdev/nodxgo"
)

func SkeletonTr(rows int) nodx.Node {
	rs := make([]nodx.Node, rows)
	for i := range rs {
		rs[i] = nodx.Tr(
			nodx.Td(
				nodx.Colspan("100%"),
				nodx.Div(
					nodx.Class("animate-pulse h-4 w-full bg-base-300 rounded-badge"),
				),
			),
		)
	}

	return nodx.Group(rs...)

}
