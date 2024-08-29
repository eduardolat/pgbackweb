package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type CardBoxParams struct {
	Class    string
	Children []gomponents.Node
}

// CardBox renders a card box with the given children.
func CardBox(params CardBoxParams) gomponents.Node {
	return html.Div(
		components.Classes{
			"rounded-box shadow-md bg-base-100 p-4": true,
			params.Class:                            true,
		},
		gomponents.Group(params.Children),
	)
}

// CardBoxSimple is the same as CardBox, but with a less verbose
// api and default props. It renders a card box with the given children.
func CardBoxSimple(children ...gomponents.Node) gomponents.Node {
	return CardBox(CardBoxParams{
		Children: children,
	})
}
