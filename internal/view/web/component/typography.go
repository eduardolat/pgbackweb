package component

import (
	nodx "github.com/nodxdev/nodxgo"
)

// getTextColorClass returns the text color class for a text that
// is written on top of a base-100, base-200 or base-300 background.
func getTextColorClass(color color) string {
	switch color {
	case ColorPrimary:
		return "text-primary"
	case ColorSecondary:
		return "text-secondary"
	case ColorAccent:
		return "text-accent"
	case ColorNeutral:
		return "text-base-content"
	case ColorInfo:
		return "text-info"
	case ColorSuccess:
		return "text-success"
	case ColorWarning:
		return "text-warning"
	case ColorError:
		return "text-error"
	default:
		return "text-base-content"
	}
}

func H1(children ...nodx.Node) nodx.Node {
	return nodx.H1(
		nodx.Class("text-2xl font-bold desk:text-4xl"),
		nodx.Group(children...),
	)
}

func H2(children ...nodx.Node) nodx.Node {
	return nodx.H2(
		nodx.Class("text-xl font-bold desk:text-2xl"),
		nodx.Group(children...),
	)
}

func H3(children ...nodx.Node) nodx.Node {
	return nodx.H3(
		nodx.Class("text-lg font-bold desk:text-xl"),
		nodx.Group(children...),
	)
}

func H4(children ...nodx.Node) nodx.Node {
	return nodx.H4(
		nodx.Class("text-base font-bold desk:text-lg"),
		nodx.Group(children...),
	)
}

func H5(children ...nodx.Node) nodx.Node {
	return nodx.H5(
		nodx.Class("text-sm font-bold desk:text-base"),
		nodx.Group(children...),
	)
}

func H6(children ...nodx.Node) nodx.Node {
	return nodx.H6(
		nodx.Class("text-xs font-bold desk:text-sm"),
		nodx.Group(children...),
	)
}

// H1Text is a convenience function to create an H1 element with a
// simple text node as its child.
func H1Text(text string) nodx.Node {
	return H1(nodx.Text(text))
}

// H2Text is a convenience function to create an H2 element with a
// simple text node as its child.
func H2Text(text string) nodx.Node {
	return H2(nodx.Text(text))
}

// H3Text is a convenience function to create an H3 element with a
// simple text node as its child.
func H3Text(text string) nodx.Node {
	return H3(nodx.Text(text))
}

// H4Text is a convenience function to create an H4 element with a
// simple text node as its child.
func H4Text(text string) nodx.Node {
	return H4(nodx.Text(text))
}

// H5Text is a convenience function to create an H5 element with a
// simple text node as its child.
func H5Text(text string) nodx.Node {
	return H5(nodx.Text(text))
}

// H6Text is a convenience function to create an H6 element with a
// simple text node as its child.
func H6Text(text string) nodx.Node {
	return H6(nodx.Text(text))
}

// PText is a convenience function to create a P element with a
// simple text node as its child.
func PText(text string) nodx.Node {
	return nodx.P(nodx.Text(text))
}

// SpanText is a convenience function to create a Span element with a
// simple text node as its child.
func SpanText(text string) nodx.Node {
	return nodx.SpanEl(nodx.Text(text))
}

// BText is a convenience function to create a B element with a
// simple text node as its child.
func BText(text string) nodx.Node {
	return nodx.B(nodx.Text(text))
}
