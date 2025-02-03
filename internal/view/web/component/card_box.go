package component

import (
	nodx "github.com/nodxdev/nodxgo"
)

type CardBoxParams struct {
	Class    string
	Children []nodx.Node
}

// CardBox renders a card box with the given children.
func CardBox(params CardBoxParams) nodx.Node {
	return nodx.Div(
		nodx.ClassMap{
			"rounded-box shadow-md bg-base-100 p-4": true,
			params.Class:                            true,
		},
		nodx.Group(params.Children...),
	)
}

// CardBoxSimple is the same as CardBox, but with a less verbose
// api and default props. It renders a card box with the given children.
func CardBoxSimple(children ...nodx.Node) nodx.Node {
	return CardBox(CardBoxParams{
		Children: children,
	})
}
