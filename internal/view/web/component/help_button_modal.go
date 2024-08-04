package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type HelpButtonModalParams struct {
	ModalTitle string
	ModalSize  size
	ButtonSize size
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
		components.Classes{
			"btn btn-neutral btn-ghost btn-circle": true,
			"btn-sm":                               params.ButtonSize == SizeSm,
			"btn-lg":                               params.ButtonSize == SizeLg,
		},
		html.Type("button"),
		lucide.CircleHelp(
			components.Classes{
				"size-4": params.ButtonSize == SizeSm,
				"size-6": params.ButtonSize == SizeMd,
				"size-8": params.ButtonSize == SizeLg,
			},
		),
	)

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		button,
	)
}
