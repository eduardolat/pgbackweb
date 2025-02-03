package component

import (
	"github.com/google/uuid"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

type TextareaControlParams struct {
	ID                 string
	Name               string
	Label              string
	Placeholder        string
	Required           bool
	HelpText           string
	Color              color
	AutoComplete       string
	Pattern            string
	Children           []nodx.Node
	HelpButtonChildren []nodx.Node
}

func TextareaControl(params TextareaControlParams) nodx.Node {
	id := params.ID
	if id == "" {
		id = "textarea-control-" + uuid.NewString()
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
		nodx.Textarea(
			nodx.ClassMap{
				"textarea textarea-bordered w-full": true,
				getTextareaColorClass(params.Color): true,
			},
			nodx.Id(id),
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

func getTextareaColorClass(color color) string {
	switch color {
	case ColorPrimary:
		return "textarea-primary"
	case ColorSecondary:
		return "textarea-secondary"
	case ColorAccent:
		return "textarea-accent"
	case ColorNeutral:
		return "textarea-neutral"
	case ColorInfo:
		return "textarea-info"
	case ColorSuccess:
		return "textarea-success"
	case ColorWarning:
		return "textarea-warning"
	case ColorError:
		return "textarea-error"
	default:
		return "textarea-neutral"
	}
}
