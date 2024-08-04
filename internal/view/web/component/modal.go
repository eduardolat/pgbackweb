package component

import (
	"fmt"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/google/uuid"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// ModalParams are the props for the Modal component.
type ModalParams struct {
	// ID is the ID of the modal dialog. If empty, a random ID will be generated.
	ID string
	// Content is the content of the modal dialog.
	Content []gomponents.Node
	// Size is the size of the modal dialog.
	// Can be "sm", "md", and "lg".
	// The default is "md".
	Size size
	// Title is the title of the modal dialog.
	// If you need more than a string, use TitleNode instead.
	Title string
	// TitleNode is the title of the modal dialog.
	// If you need only a string, use Title instead.
	TitleNode gomponents.Node
	// HTMXIndicator is an optional ID of an HTMX indicator that
	// should be inserted in the modal header.
	HTMXIndicator string
}

// ModalResult is the result of creating a modal dialog.
type ModalResult struct {
	// HTML is the modal dialog HTML.
	HTML gomponents.Node
	// OpenerAttr is the attribute to add to the element that opens the modal dialog.
	OpenerAttr gomponents.Node
}

// Modal renders a modal dialog.
func Modal(params ModalParams) ModalResult {
	id := params.ID
	if id == "" {
		id = "mo-" + uuid.NewString()
	}

	openEventName := fmt.Sprintf("%s_open", id)
	closeEventName := fmt.Sprintf("%s_close", id)
	openerAttr := gomponents.Attr(
		"onClick",
		"event.preventDefault(); window.dispatchEvent(new Event('"+openEventName+"'));",
	)
	closerAttr := gomponents.Attr(
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

	content := html.Div(
		alpine.XData(`{}`),
		alpine.XOn(fmt.Sprintf("%s.window", openEventName), openCode),
		alpine.XOn(fmt.Sprintf("%s.window", closeEventName), closeCode),
		alpine.XOn("keyup.escape.window", closeCode),

		html.ID(id),
		components.Classes{
			"hidden":                          true,
			"!p-0 !m-0 w-[100dvw] h-[100dvh]": true,
			"fixed left-0 top-0 z-[1000]":     true,
		},

		// Backdrop
		html.Div(
			closerAttr,
			components.Classes{
				"bg-black opacity-25": true,
				"!w-full !h-full":     true,
				"z-[1001]":            true,
			},
		),

		// Dialog
		html.Div(
			components.Classes{
				"absolute z-[1002] top-[50%] left-[50%]":      true,
				"translate-y-[-50%] translate-x-[-50%]":       true,
				"max-w-[calc(100dvw-30px)] max-h-[85dvh]":     true,
				"bg-base-100 rounded-box overflow-y-auto p-0": true,
				"overflow-x-hidden whitespace-normal":         true,

				"w-[400px]": size == SizeSm,
				"w-[600px]": size == SizeMd,
				"w-[800px]": size == SizeLg,
			},

			html.Div(
				components.Classes{
					"w-full sticky top-0 right-0 bg-base-100": true,
					"flex items-center justify-between":       true,
					"border-b border-base-300 px-4 py-3":      true,
				},

				html.Div(
					gomponents.If(
						params.TitleNode != nil,
						params.TitleNode,
					),

					gomponents.If(
						params.Title != "",
						html.Span(
							html.Class("text-xl font-bold desk:text-2xl"),
							gomponents.Text(params.Title),
						),
					),

					gomponents.If(
						hasHTMXIndicator,
						html.Div(
							html.Class("inline-flex h-full items-center pl-2"),
							HxLoadingSm(params.HTMXIndicator),
						),
					),
				),

				html.Button(
					html.Class("btn btn-circle btn-ghost btn-sm"),
					lucide.X(html.Class("size-6")),
					closerAttr,
				),
			),

			html.Div(
				html.Class("p-4"),
				gomponents.Group(params.Content),
			),
		),
	)

	content = alpine.Template(
		alpine.XData(""),
		alpine.XTeleport("body"),
		html.Div(content),
	)

	return ModalResult{
		OpenerAttr: openerAttr,
		HTML:       content,
	}
}
