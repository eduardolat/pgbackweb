package component

import (
	"github.com/google/uuid"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
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
	Children           []nodx.Node
	HelpButtonChildren []nodx.Node
}

func SelectControl(params SelectControlParams) nodx.Node {
	id := params.ID
	if id == "" {
		id = "select-control-" + uuid.NewString()
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
		nodx.Select(
			nodx.ClassMap{
				"w-full": true,
			},
			nodx.Id(id),
			nodx.Name(params.Name),
			nodx.If(
				params.Required,
				nodx.Required(""),
			),
			nodx.If(
				params.AutoComplete != "",
				nodx.Autocomplete(params.AutoComplete),
			),
			nodx.If(
				params.Placeholder != "",
				nodx.Option(
					nodx.Value(""),
					nodx.Disabled(""),
					nodx.Selected(""),
					nodx.Text(params.Placeholder),
				),
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
		nodx.Script(nodx.Raw(`
			new SlimSelect({select: '#`+id+`'})
		`)),
	)
}
