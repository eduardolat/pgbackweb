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

func CardBox(params CardBoxParams) gomponents.Node {
	return html.Div(
		components.Classes{
			"rounded-box shadow-md bg-base-100 p-4": true,
			params.Class:                            true,
		},
		gomponents.Group(params.Children),
	)
}
