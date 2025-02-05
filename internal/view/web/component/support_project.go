package component

import (
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func SupportProjectModal() ModalResult {
	mo := Modal(ModalParams{
		Size:  SizeMd,
		Title: "Help support PG Back Web",
		Content: []nodx.Node{
			nodx.Div(
				alpine.XData("alpineSupportProjectData()"),
			),
		},
	})

	return mo
}

func SupportProjectButton(size size) nodx.Node {
	mo := SupportProjectModal()

	return nodx.Group(
		mo.HTML,
		nodx.Button(
			mo.OpenerAttr,
			nodx.ClassMap{
				"btn btn-success": true,
				"btn-sm":          size == SizeSm,
				"btn-lg":          size == SizeLg,
			},
			lucide.HeartHandshake(),
			SpanText("Support the project"),
		),
	)
}
