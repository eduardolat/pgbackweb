package component

import (
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

type ChangeThemeButtonParams struct {
	Position    dropdownPosition
	AlignsToEnd bool
	Size        size
}

func ChangeThemeButton(params ChangeThemeButtonParams) nodx.Node {
	return nodx.Div(
		alpine.XData("alpineChangeThemeButton()"),
		alpine.XCloak(),

		nodx.ClassMap{
			"dropdown":        true,
			"dropdown-end":    params.AlignsToEnd,
			"dropdown-right":  params.Position == DropdownPositionRight,
			"dropdown-left":   params.Position == DropdownPositionLeft,
			"dropdown-top":    params.Position == DropdownPositionTop,
			"dropdown-bottom": params.Position == DropdownPositionBottom,
		},
		nodx.Div(
			nodx.Tabindex("0"),
			nodx.Role("button"),
			nodx.ClassMap{
				"btn btn-neutral space-x-1": true,
				"btn-sm":                    params.Size == SizeSm,
				"btn-lg":                    params.Size == SizeLg,
			},

			nodx.Div(
				nodx.Class("inline-block size-4"),
				lucide.Laptop(alpine.XShow(`theme === "system"`)),
				lucide.Sun(alpine.XShow(`theme === "light"`)),
				lucide.Moon(alpine.XShow(`theme === "dark"`)),
			),

			SpanText("Theme"),
			lucide.ChevronDown(),
		),
		nodx.Ul(
			nodx.Tabindex("0"),
			nodx.ClassMap{
				"dropdown-content":                   true,
				"bg-base-100":                        true,
				"rounded-btn shadow-md":              true,
				"z-[1] w-[150px] p-2 space-y-2 my-2": true,
			},
			nodx.Li(
				nodx.Button(
					alpine.XOn("click", "setTheme('')"),
					nodx.Class("btn btn-neutral btn-block"),
					nodx.Type("button"),
					lucide.Laptop(nodx.Class("mr-1")),
					SpanText("System"),
				),
			),
			nodx.Li(
				nodx.Button(
					alpine.XOn("click", "setTheme('light')"),
					nodx.Class("btn btn-neutral btn-block"),
					nodx.Type("button"),
					lucide.Sun(nodx.Class("mr-1")),
					SpanText("Light"),
				),
			),
			nodx.Li(
				nodx.Button(
					alpine.XOn("click", "setTheme('dark')"),
					nodx.Class("btn btn-neutral btn-block"),
					nodx.Type("button"),
					lucide.Moon(nodx.Class("mr-1")),
					SpanText("Dark"),
				),
			),
		),
	)
}
