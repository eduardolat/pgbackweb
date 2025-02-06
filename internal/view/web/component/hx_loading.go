package component

import (
	nodx "github.com/nodxdev/nodxgo"
)

// HxLoadingSm returns a small loading indicator.
func HxLoadingSm(id ...string) nodx.Node {
	return hxLoading(SizeSm, id...)
}

// HxLoadingMd returns a loading indicator.
func HxLoadingMd(id ...string) nodx.Node {
	return hxLoading(SizeMd, id...)
}

// HxLoadingLg returns a large loading indicator.
func HxLoadingLg(id ...string) nodx.Node {
	return hxLoading(SizeLg, id...)
}

func hxLoading(size size, id ...string) nodx.Node {
	pickedID := ""
	if len(id) > 0 {
		pickedID = id[0]
	}

	return nodx.Div(
		nodx.If(
			pickedID != "",
			nodx.Id(pickedID),
		),
		nodx.Class("htmx-indicator inline-block"),
		func() nodx.Node {
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
