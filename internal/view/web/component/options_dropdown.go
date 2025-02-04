package component

import (
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func OptionsDropdown(children ...nodx.Node) nodx.Node {
	return nodx.Div(
		nodx.Class("inline-block"),
		alpine.XData("alpineOptionsDropdown()"),
		alpine.XOn("mouseenter", "open()"),
		alpine.XOn("mouseleave", "close()"),
		nodx.Button(
			alpine.XRef("button"),
			nodx.Class("btn btn-sm btn-ghost btn-square"),
			alpine.XBind("class", "isOpen ? 'btn-active' : ''"),
			lucide.EllipsisVertical(
				nodx.Class("transition-transform"),
				alpine.XBind("class", "isOpen ? 'rotate-90' : ''"),
			),
		),
		nodx.Div(
			alpine.XRef("content"),
			nodx.ClassMap{
				"fixed hidden": true,
				"bg-base-100 rounded-box border border-base-200": true,
				"z-40 max-w-[250px] p-2 shadow-md":               true,
			},
			nodx.Group(children...),
		),
	)
}

func OptionsDropdownButton(children ...nodx.Node) nodx.Node {
	return nodx.Button(
		nodx.Class("btn btn-neutral btn-ghost btn-sm w-full flex justify-start"),
		nodx.Group(children...),
	)
}

func OptionsDropdownA(children ...nodx.Node) nodx.Node {
	return nodx.A(
		nodx.Class("btn btn-neutral btn-ghost btn-sm w-full flex justify-start"),
		nodx.Group(children...),
	)
}
