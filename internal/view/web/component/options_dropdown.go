package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func OptionsDropdown(children ...gomponents.Node) gomponents.Node {
	return html.Div(
		html.Class("inline-block"),
		alpine.XData("alpineOptionsDropdown()"),
		alpine.XOn("mouseenter", "open()"),
		alpine.XOn("mouseleave", "close()"),
		html.Button(
			alpine.XRef("button"),
			html.Class("btn btn-sm btn-ghost btn-square"),
			alpine.XBind("class", "isOpen ? 'btn-active' : ''"),
			lucide.EllipsisVertical(
				html.Class("transition-transform"),
				alpine.XBind("class", "isOpen ? 'rotate-90' : ''"),
			),
		),
		html.Div(
			alpine.XRef("content"),
			components.Classes{
				"fixed hidden": true,
				"bg-base-100 rounded-box border border-base-200": true,
				"z-40 max-w-[250px] p-2 shadow-md":               true,
			},
			gomponents.Group(children),
		),
	)
}

func OptionsDropdownButton(children ...gomponents.Node) gomponents.Node {
	return html.Button(
		html.Class("btn btn-neutral btn-ghost btn-sm w-full flex justify-start"),
		gomponents.Group(children),
	)
}

func OptionsDropdownA(children ...gomponents.Node) gomponents.Node {
	return html.A(
		html.Class("btn btn-neutral btn-ghost btn-sm w-full flex justify-start"),
		gomponents.Group(children),
	)
}
