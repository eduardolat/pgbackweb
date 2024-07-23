package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type InputControlParams struct {
	Name         string
	Label        string
	Placeholder  string
	Required     bool
	Type         inputType
	HelpText     string
	Color        color
	AutoComplete string
	Children     []gomponents.Node
}

func InputControl(params InputControlParams) gomponents.Node {
	if params.Type.Value == "" {
		params.Type = InputTypeText
	}

	return html.Label(
		components.Classes{
			"form-control w-full":           true,
			getTextColorClass(params.Color): true,
		},
		html.Div(
			html.Class("label"),
			html.Div(
				html.Class("flex justify-start items-center space-x-1"),
				SpanText(params.Label),
				gomponents.If(
					params.Required,
					lucide.Asterisk(html.Class("text-error")),
				),
			),
		),
		html.Input(
			components.Classes{
				"input input-bordered w-full":    true,
				getInputColorClass(params.Color): true,
			},
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
			gomponents.Group(params.Children),
		),
		gomponents.If(
			params.HelpText != "",
			html.Div(
				html.Class("label"),
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
