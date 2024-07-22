package component

import (
	"github.com/maragudk/gomponents"
	gcomponents "github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// HxLoadingSm returns a small loading indicator.
func HxLoadingSm(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, SizeSm, id...)
}

// HxLoadingMd returns a loading indicator.
func HxLoadingMd(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, SizeMd, id...)
}

// HxLoadingLg returns a large loading indicator.
func HxLoadingLg(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, SizeLg, id...)
}

func hxLoading(centered bool, size size, id ...string) gomponents.Node {
	pickedID := ""
	if len(id) > 0 {
		pickedID = id[0]
	}

	return html.Div(
		gomponents.If(
			pickedID != "",
			html.ID(pickedID),
		),
		html.Class("htmx-indicator"),
		html.Div(
			gcomponents.Classes{
				"flex justify-center items-center": centered,
				"w-full h-full":                    true,
			},

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
		),
	)
}
