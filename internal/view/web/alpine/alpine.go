package alpine

import "github.com/maragudk/gomponents"

// Template is a generic function for rendering an <template> element.
func Template(children ...gomponents.Node) gomponents.Node {
	return gomponents.El("template", children...)
}

// XData is an attribute that renders a x-data="value]" attribute.
//
// https://alpinejs.dev/directives/data
func XData(value string) gomponents.Node {
	return gomponents.Attr("x-data", value)
}

// XFor is an attribute that renders a x-for="value]" attribute.
//
// https://alpinejs.dev/directives/for
func XFor(value string) gomponents.Node {
	return gomponents.Attr("x-for", value)
}

// XInit is an attribute that renders a x-init="[value]" attribute.
//
// https://alpinejs.dev/directives/init
func XInit(value string) gomponents.Node {
	return gomponents.Attr("x-init", value)
}

// XShow is an attribute that renders a x-show="[vlue]" attribute.
//
// https://alpinejs.dev/directives/show
func XShow(value string) gomponents.Node {
	return gomponents.Attr("x-show", value)
}

// XBind is an attribute that renders a x-bind:[targetAttr]="[value]" attribute.
//
// https://alpinejs.dev/directives/bind
func XBind(targetAttr string, value string) gomponents.Node {
	return gomponents.Attr("x-bind:"+targetAttr, value)
}

// XOn is an attribute that renders a x-on:[targetEvent]="[value]" attribute.
//
// https://alpinejs.dev/directives/on
func XOn(targetEvent string, value string) gomponents.Node {
	return gomponents.Attr("x-on:"+targetEvent, value)
}

// XText is an attribute that renders a x-text="[value]" attribute.
//
// https://alpinejs.dev/directives/text
func XText(value string) gomponents.Node {
	return gomponents.Attr("x-text", value)
}

// XHTML is an attribute that renders a x-html="[value]" attribute.
//
// https://alpinejs.dev/directives/html
func XHTML(value string) gomponents.Node {
	return gomponents.Attr("x-html", value)
}

// XModel is an attribute that renders a x-model="[value]" attribute.
//
// https://alpinejs.dev/directives/model
func XModel(value string) gomponents.Node {
	return gomponents.Attr("x-model", value)
}

// XTransition is an attribute that renders a x-transition attribute.
//
// https://alpinejs.dev/directives/transition
func XTransition() gomponents.Node {
	return gomponents.Attr("x-transition")
}

// XTransitionFade is an attribute that renders a x-transition.opacity attribute.
//
// https://alpinejs.dev/directives/transition
func XTransitionFade() gomponents.Node {
	return gomponents.Attr("x-transition.opacity")
}

// XIgnore is an attribute that renders a x-ignore attribute.
//
// https://alpinejs.dev/directives/ignore
func XIgnore() gomponents.Node {
	return gomponents.Attr("x-ignore")
}

// XRef is an attribute that renders a x-ref="[value]" attribute.
//
// https://alpinejs.dev/directives/ref
func XRef(value string) gomponents.Node {
	return gomponents.Attr("x-ref", value)
}

// XCloak is an attribute that renders a x-cloak attribute.
//
// https://alpinejs.dev/directives/cloak
func XCloak() gomponents.Node {
	return gomponents.Attr("x-cloak")
}

// XTeleport is an attribute that renders a x-teleport="[value]" attribute.
//
// https://alpinejs.dev/directives/teleport
func XTeleport(value string) gomponents.Node {
	return gomponents.Attr("x-teleport", value)
}

// IF is an attribute that renders a x-if="[value]" attribute.
//
// https://alpinejs.dev/directives/if
func XIf(value string) gomponents.Node {
	return gomponents.Attr("x-if", value)
}
