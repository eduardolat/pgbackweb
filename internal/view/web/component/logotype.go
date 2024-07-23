package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func Logotype() gomponents.Node {
	return html.Div(
		components.Classes{
			"inline space-x-2 select-none":    true,
			"flex justify-start items-center": true,
		},
		html.Img(
			html.Class("w-[60px] h-auto"),
			html.Src("/images/logo.png"),
			html.Alt("PG Back Web"),
		),
		html.Span(
			html.Class("text-2xl font-bold"),
			gomponents.Text("PG Back Web"),
		),
	)
}
