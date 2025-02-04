package component

import (
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

type HelpButtonModalParams struct {
	ModalTitle string
	ModalSize  size
	Children   []nodx.Node
}

func HelpButtonModal(params HelpButtonModalParams) nodx.Node {
	mo := Modal(ModalParams{
		Size:    params.ModalSize,
		Title:   params.ModalTitle,
		Content: params.Children,
	})

	button := nodx.Button(
		mo.OpenerAttr,
		nodx.Class("btn btn-neutral btn-ghost btn-circle btn-sm"),
		nodx.Type("button"),
		lucide.CircleHelp(),
	)

	return nodx.Div(
		nodx.Class("inline-block"),
		mo.HTML,
		button,
	)
}
