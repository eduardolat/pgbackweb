package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
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

func H1(children ...gomponents.Node) gomponents.Node {
	return html.H1(
		html.Class("text-2xl font-bold desk:text-4xl"),
		gomponents.Group(children),
	)
}

func H2(children ...gomponents.Node) gomponents.Node {
	return html.H2(
		html.Class("text-xl font-bold desk:text-2xl"),
		gomponents.Group(children),
	)
}

func H3(children ...gomponents.Node) gomponents.Node {
	return html.H3(
		html.Class("text-lg font-bold desk:text-xl"),
		gomponents.Group(children),
	)
}

func H4(children ...gomponents.Node) gomponents.Node {
	return html.H4(
		html.Class("text-base font-bold desk:text-lg"),
		gomponents.Group(children),
	)
}

func H5(children ...gomponents.Node) gomponents.Node {
	return html.H5(
		html.Class("text-sm font-bold desk:text-base"),
		gomponents.Group(children),
	)
}

func H6(children ...gomponents.Node) gomponents.Node {
	return html.H6(
		html.Class("text-xs font-bold desk:text-sm"),
		gomponents.Group(children),
	)
}

// H1Text is a convenience function to create an H1 element with a
// simple text node as its child.
func H1Text(text string) gomponents.Node {
	return H1(gomponents.Text(text))
}

// H2Text is a convenience function to create an H2 element with a
// simple text node as its child.
func H2Text(text string) gomponents.Node {
	return H2(gomponents.Text(text))
}

// H3Text is a convenience function to create an H3 element with a
// simple text node as its child.
func H3Text(text string) gomponents.Node {
	return H3(gomponents.Text(text))
}

// H4Text is a convenience function to create an H4 element with a
// simple text node as its child.
func H4Text(text string) gomponents.Node {
	return H4(gomponents.Text(text))
}

// H5Text is a convenience function to create an H5 element with a
// simple text node as its child.
func H5Text(text string) gomponents.Node {
	return H5(gomponents.Text(text))
}

// H6Text is a convenience function to create an H6 element with a
// simple text node as its child.
func H6Text(text string) gomponents.Node {
	return H6(gomponents.Text(text))
}

// PText is a convenience function to create a P element with a
// simple text node as its child.
func PText(text string) gomponents.Node {
	return html.P(gomponents.Text(text))
}

// SpanText is a convenience function to create a Span element with a
// simple text node as its child.
func SpanText(text string) gomponents.Node {
	return html.Span(gomponents.Text(text))
}

// BText is a convenience function to create a B element with a
// simple text node as its child.
func BText(text string) gomponents.Node {
	return html.B(gomponents.Text(text))
}
