package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func ChangeThemeButton(
	position dropdownPosition, alignsToEnd bool,
) gomponents.Node {
	return html.Div(
		alpine.XData(`{
			theme: "",
			
			getCurrentTheme() {
				const el = document.querySelector("html");
				const theme = el.getAttribute("data-theme");
				if (theme) {
					this.theme = theme;
					return
				}
				this.theme = "system";
			},

			init() {
				setTimeout(() => {
					this.getCurrentTheme();
				}, 200)
			}
		}`),
		alpine.XCloak(),
		alpine.XOn("click", "getCurrentTheme()"),
		alpine.XOn("click.outside", "getCurrentTheme()"),

		components.Classes{
			"dropdown":        true,
			"dropdown-end":    alignsToEnd,
			"dropdown-right":  position == DropdownPositionRight,
			"dropdown-left":   position == DropdownPositionLeft,
			"dropdown-top":    position == DropdownPositionTop,
			"dropdown-bottom": position == DropdownPositionBottom,
		},
		html.Div(
			html.TabIndex("0"),
			html.Role("button"),
			html.Class("btn btn-neutral space-x-1"),

			html.Div(
				html.Class("inline-block size-4"),
				lucide.Laptop(alpine.XShow(`theme === "system"`)),
				lucide.Sun(alpine.XShow(`theme === "light"`)),
				lucide.Moon(alpine.XShow(`theme === "dark"`)),
			),

			SpanText("Theme"),
			lucide.ChevronDown(),
		),
		html.Ul(
			html.TabIndex("0"),
			components.Classes{
				"dropdown-content":                   true,
				"bg-base-100":                        true,
				"rounded-btn shadow-md":              true,
				"z-[1] w-[150px] p-2 space-y-2 my-2": true,
			},
			html.Li(
				html.Button(
					html.Data("set-theme", ""),
					html.Class("btn btn-neutral btn-block"),
					html.Type("button"),
					lucide.Laptop(html.Class("mr-1")),
					SpanText("System"),
				),
			),
			html.Li(
				html.Button(
					html.Data("set-theme", "light"),
					html.Class("btn btn-neutral btn-block"),
					html.Type("button"),
					lucide.Sun(html.Class("mr-1")),
					SpanText("Light"),
				),
			),
			html.Li(
				html.Button(
					html.Data("set-theme", "dark"),
					html.Class("btn btn-neutral btn-block"),
					html.Type("button"),
					lucide.Moon(html.Class("mr-1")),
					SpanText("Dark"),
				),
			),
		),
	)
}
