package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// HxLoadingSm returns a small loading indicator.
func HxLoadingSm(id ...string) gomponents.Node {
	return hxLoading(SizeSm, id...)
}

// HxLoadingMd returns a loading indicator.
func HxLoadingMd(id ...string) gomponents.Node {
	return hxLoading(SizeMd, id...)
}

// HxLoadingLg returns a large loading indicator.
func HxLoadingLg(id ...string) gomponents.Node {
	return hxLoading(SizeLg, id...)
}

func hxLoading(size size, id ...string) gomponents.Node {
	pickedID := ""
	if len(id) > 0 {
		pickedID = id[0]
	}

	return html.Div(
		gomponents.If(
			pickedID != "",
			html.ID(pickedID),
		),
		html.Class("htmx-indicator inline-block"),
		func() gomponents.Node {
			switch size {
			case SizeSm:
				return SpinnerSm()
			case SizeMd:
				return SpinnerMd()
			case SizeLg:
				return SpinnerLg()
			default:
				return SpinnerMd()
			}
		}(),
	)
}
