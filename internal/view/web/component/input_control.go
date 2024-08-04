package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/google/uuid"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type InputControlParams struct {
	ID                 string
	Name               string
	Label              string
	Placeholder        string
	Required           bool
	Type               inputType
	HelpText           string
	Color              color
	AutoComplete       string
	Pattern            string
	Children           []gomponents.Node
	HelpButtonChildren []gomponents.Node
}

func InputControl(params InputControlParams) gomponents.Node {
	id := params.ID
	if id == "" {
		id = "input-control-" + uuid.NewString()
	}

	if params.Type.Value == "" {
		params.Type = InputTypeText
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
		html.Input(
			components.Classes{
				"input input-bordered w-full":    true,
				getInputColorClass(params.Color): true,
			},
			html.ID(id),
			html.Type(params.Type.Value),
			html.Name(params.Name),
			html.Placeholder(params.Placeholder),
			gomponents.If(
				params.Required,
				html.Required(),
			),
			gomponents.If(
				params.AutoComplete != "",
				html.AutoComplete(params.AutoComplete),
			),
			gomponents.If(
				params.Pattern != "",
				html.Pattern(params.Pattern),
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
	)
}

func getInputColorClass(color color) string {
	switch color {
	case ColorPrimary:
		return "input-primary"
	case ColorSecondary:
		return "input-secondary"
	case ColorAccent:
		return "input-accent"
	case ColorNeutral:
		return "input-neutral"
	case ColorInfo:
		return "input-info"
	case ColorSuccess:
		return "input-success"
	case ColorWarning:
		return "input-warning"
	case ColorError:
		return "input-error"
	default:
		return "input-neutral"
	}
}
