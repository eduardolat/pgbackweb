package component

import (
	nodx "github.com/nodxdev/nodxgo"
)

type CardBoxParams struct {
	Class    string
	BgBase   bgBase
	Children []nodx.Node
}

// CardBox renders a card box with the given children.
func CardBox(params CardBoxParams) nodx.Node {
	if params.BgBase.Value == "" {
		params.BgBase = bgBase100
	}

	return nodx.Div(
		nodx.ClassMap{
			"rounded-box shadow-md p-4": true,
			"bg-base-100":               params.BgBase == bgBase100,
			"bg-base-200":               params.BgBase == bgBase200,
			"bg-base-300":               params.BgBase == bgBase300,
			params.Class:                true,
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

// CardBoxSimpleBgBase200 is the same as CardBox, but with a less verbose
// api and default props. It renders a card box with the given children.
func CardBoxSimpleBgBase200(children ...nodx.Node) nodx.Node {
	return CardBox(CardBoxParams{
		BgBase:   bgBase200,
		Children: children,
	})
}

// CardBoxSimpleBgBase300 is the same as CardBox, but with a less verbose
// api and default props. It renders a card box with the given children.
func CardBoxSimpleBgBase300(children ...nodx.Node) nodx.Node {
	return CardBox(CardBoxParams{
		BgBase:   bgBase300,
		Children: children,
	})
}
