package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/google/uuid"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
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
	Children           []gomponents.Node
	HelpButtonChildren []gomponents.Node
}

func TextareaControl(params TextareaControlParams) gomponents.Node {
	id := params.ID
	if id == "" {
		id = "textarea-control-" + uuid.NewString()
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
		html.Textarea(
			components.Classes{
				"textarea textarea-bordered w-full": true,
				getTextareaColorClass(params.Color): true,
			},
			html.ID(id),
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
