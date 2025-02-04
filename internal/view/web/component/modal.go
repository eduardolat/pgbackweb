package component

import (
	"fmt"

	"github.com/google/uuid"
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

// ModalParams are the props for the Modal component.
type ModalParams struct {
	// ID is the ID of the modal dialog. If empty, a random ID will be generated.
	ID string
	// Content is the content of the modal dialog.
	Content []nodx.Node
	// Size is the size of the modal dialog.
	// Can be "sm", "md", and "lg".
	// The default is "md".
	Size size
	// Title is the title of the modal dialog.
	// If you need more than a string, use TitleNode instead.
	Title string
	// TitleNode is the title of the modal dialog.
	// If you need only a string, use Title instead.
	TitleNode nodx.Node
	// HTMXIndicator is an optional ID of an HTMX indicator that
	// should be inserted in the modal header.
	HTMXIndicator string
}

// ModalResult is the result of creating a modal dialog.
type ModalResult struct {
	// HTML is the modal dialog nodx.
	HTML nodx.Node
	// OpenerAttr is the attribute to add to the element that opens the modal dialog.
	OpenerAttr nodx.Node
}

// Modal renders a modal dialog.
func Modal(params ModalParams) ModalResult {
	id := params.ID
	if id == "" {
		id = "mo-" + uuid.NewString()
	}

	openEventName := fmt.Sprintf("%s_open", id)
	closeEventName := fmt.Sprintf("%s_close", id)
	openerAttr := nodx.Attr(
		"onClick",
		"event.preventDefault(); window.dispatchEvent(new Event('"+openEventName+"'));",
	)
	closerAttr := nodx.Attr(
		"onClick",
		"event.preventDefault(); window.dispatchEvent(new Event('"+closeEventName+"'));",
	)

	openCode := `document.getElementById("` + id + `").classList.remove("hidden");`
	closeCode := `document.getElementById("` + id + `").classList.add("hidden");`

	size := SizeMd
	if params.Size.Value != "" {
		size = params.Size
	}

	hasHTMXIndicator := params.HTMXIndicator != ""

	content := nodx.Div(
		alpine.XData(`{}`),
		alpine.XOn(fmt.Sprintf("%s.window", openEventName), openCode),
		alpine.XOn(fmt.Sprintf("%s.window", closeEventName), closeCode),
		alpine.XOn("keyup.escape.window", closeCode),

		nodx.Id(id),
		nodx.ClassMap{
			"hidden":                          true,
			"!p-0 !m-0 w-[100dvw] h-[100dvh]": true,
			"fixed left-0 top-0 z-[1000]":     true,
		},

		// Backdrop
		nodx.Div(
			closerAttr,
			nodx.ClassMap{
				"bg-black opacity-25": true,
				"!w-full !h-full":     true,
				"z-[1001]":            true,
			},
		),

		// Dialog
		nodx.Div(
			nodx.ClassMap{
				"absolute z-[1002] top-[50%] left-[50%]":      true,
				"translate-y-[-50%] translate-x-[-50%]":       true,
				"max-w-[calc(100dvw-30px)] max-h-[85dvh]":     true,
				"bg-base-100 rounded-box overflow-y-auto p-0": true,
				"overflow-x-hidden whitespace-normal":         true,

				"w-[400px]": size == SizeSm,
				"w-[600px]": size == SizeMd,
				"w-[800px]": size == SizeLg,
			},

			nodx.Div(
				nodx.ClassMap{
					"w-full sticky top-0 right-0 z-[1003] bg-base-100": true,
					"flex items-center justify-between":                true,
					"border-b border-base-300 px-4 py-3":               true,
				},

				nodx.Div(
					nodx.If(
						params.TitleNode != nil,
						params.TitleNode,
					),

					nodx.If(
						params.Title != "",
						nodx.SpanEl(
							nodx.Class("text-xl font-bold desk:text-2xl"),
							nodx.Text(params.Title),
						),
					),

					nodx.If(
						hasHTMXIndicator,
						nodx.Div(
							nodx.Class("inline-flex h-full items-center pl-2"),
							HxLoadingSm(params.HTMXIndicator),
						),
					),
				),

				nodx.Button(
					nodx.Class("btn btn-circle btn-ghost btn-sm"),
					lucide.X(nodx.Class("size-6")),
					closerAttr,
				),
			),

			nodx.Div(
				nodx.Class("p-4"),
				nodx.Group(params.Content...),
			),
		),
	)

	content = alpine.Template(
		alpine.XData(""),
		alpine.XTeleport("body"),
		nodx.Div(content),
	)

	return ModalResult{
		OpenerAttr: openerAttr,
		HTML:       content,
	}
}
