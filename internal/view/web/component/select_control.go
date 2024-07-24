package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type SelectControlParams struct {
	Name         string
	Label        string
	Placeholder  string
	Required     bool
	HelpText     string
	Color        color
	AutoComplete string
	Children     []gomponents.Node
}

func SelectControl(params SelectControlParams) gomponents.Node {
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
		html.Select(
			components.Classes{
				"select select-bordered w-full":   true,
				getSelectColorClass(params.Color): true,
			},
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
			html.Div(
				html.Class("label"),
				SpanText(params.HelpText),
			),
		),
	)
}

func getSelectColorClass(color color) string {
	switch color {
	case ColorPrimary:
		return "select-primary"
	case ColorSecondary:
		return "select-secondary"
	case ColorAccent:
		return "select-accent"
	case ColorNeutral:
		return "select-neutral"
	case ColorInfo:
		return "select-info"
	case ColorSuccess:
		return "select-success"
	case ColorWarning:
		return "select-warning"
	case ColorError:
		return "select-error"
	default:
		return "select-neutral"
	}
}
