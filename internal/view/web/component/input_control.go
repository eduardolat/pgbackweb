package component

import (
	"github.com/google/uuid"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
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
	Children           []nodx.Node
	HelpButtonChildren []nodx.Node
}

func InputControl(params InputControlParams) nodx.Node {
	id := params.ID
	if id == "" {
		id = "input-control-" + uuid.NewString()
	}

	if params.Type.Value == "" {
		params.Type = InputTypeText
	}

	return nodx.Div(
		nodx.ClassMap{
			"form-control w-full":           true,
			getTextColorClass(params.Color): true,
		},
		nodx.Div(
			nodx.Class("label flex justify-start"),
			nodx.LabelEl(
				nodx.For(id),
				nodx.Class("flex justify-start items-center space-x-1"),
				SpanText(params.Label),
				nodx.If(
					params.Required,
					lucide.Asterisk(nodx.Class("text-error")),
				),
			),
			nodx.If(
				len(params.HelpButtonChildren) > 0,
				HelpButtonModal(HelpButtonModalParams{
					ModalTitle: params.Label,
					Children:   params.HelpButtonChildren,
				}),
			),
		),
		nodx.Input(
			nodx.ClassMap{
				"input input-bordered w-full":    true,
				getInputColorClass(params.Color): true,
			},
			nodx.Id(id),
			nodx.Type(params.Type.Value),
			nodx.Name(params.Name),
			nodx.Placeholder(params.Placeholder),
			nodx.If(
				params.Required,
				nodx.Required(""),
			),
			nodx.If(
				params.AutoComplete != "",
				nodx.Autocomplete(params.AutoComplete),
			),
			nodx.If(
				params.Pattern != "",
				nodx.Pattern(params.Pattern),
			),
			nodx.Group(params.Children...),
		),
		nodx.If(
			params.HelpText != "",
			nodx.LabelEl(
				nodx.Class("label"),
				nodx.For(id),
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
