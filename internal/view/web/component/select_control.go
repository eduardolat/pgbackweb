package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/google/uuid"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type SelectControlParams struct {
	ID                 string
	Name               string
	Label              string
	Placeholder        string
	Required           bool
	HelpText           string
	Color              color
	AutoComplete       string
	Children           []gomponents.Node
	HelpButtonChildren []gomponents.Node
}

func SelectControl(params SelectControlParams) gomponents.Node {
	id := params.ID
	if id == "" {
		id = "select-control-" + uuid.NewString()
	}

	return html.Div(
		components.Classes{
			"form-control w-full":           true,
			getTextColorClass(params.Color): true,
		},
		html.Div(
			html.Class("label flex justify-start"),
			html.Label(
				html.For(id),
				html.Class("flex justify-start items-center space-x-1"),
				SpanText(params.Label),
				gomponents.If(
					params.Required,
					lucide.Asterisk(html.Class("text-error")),
				),
			),
			gomponents.If(
				len(params.HelpButtonChildren) > 0,
				HelpButtonModal(HelpButtonModalParams{
					ModalTitle: params.Label,
					Children:   params.HelpButtonChildren,
				}),
			),
		),
		html.Select(
			components.Classes{
				"w-full": true,
			},
			html.ID(id),
			html.Name(params.Name),
			gomponents.If(
				params.Required,
				html.Required(),
			),
			gomponents.If(
				params.AutoComplete != "",
				html.AutoComplete(params.AutoComplete),
			),
			gomponents.If(
				params.Placeholder != "",
				html.Option(
					html.Value(""),
					html.Disabled(),
					html.Selected(),
					gomponents.Text(params.Placeholder),
				),
			),
			gomponents.Group(params.Children),
		),
		gomponents.If(
			params.HelpText != "",
			html.Label(
				html.Class("label"),
				html.For(id),
				SpanText(params.HelpText),
			),
		),
		html.Script(gomponents.Raw(`
			new SlimSelect({select: '#`+id+`'})
		`)),
	)
}
