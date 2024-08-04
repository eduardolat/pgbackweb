package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type HelpButtonModalParams struct {
	ModalTitle string
	ModalSize  size
	Children   []gomponents.Node
}

func HelpButtonModal(params HelpButtonModalParams) gomponents.Node {
	mo := Modal(ModalParams{
		Size:    params.ModalSize,
		Title:   params.ModalTitle,
		Content: params.Children,
	})

	button := html.Button(
		mo.OpenerAttr,
		html.Class("btn btn-neutral btn-ghost btn-circle btn-sm"),
		html.Type("button"),
		lucide.CircleHelp(),
	)

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		button,
	)
}
