package component

import (
	"bytes"

	nodx "github.com/nodxdev/nodxgo"
)

// RenderableGroup renders a group of nodes without a parent element.
//
// This is because nodx.Group() cannot be directly rendered and
// needs to be wrapped in a parent element.
func RenderableGroup(children []nodx.Node) nodx.Node {
	buf := bytes.Buffer{}
	for _, child := range children {
		err := child.Render(&buf)
		if err != nil {
			return nodx.Raw("Error rendering group")
		}
	}
	return nodx.Raw(buf.String())
}
