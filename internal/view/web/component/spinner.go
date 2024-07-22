package component

import (
	"fmt"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func spinner(size size) gomponents.Node {
	return lucide.LoaderCircle(components.Classes{
		"animate-spin inline-block": true,
		"size-5":                    size == SizeSm,
		"size-8":                    size == SizeMd,
		"size-12":                   size == SizeLg,
	})
}

func SpinnerSm() gomponents.Node {
	return spinner(SizeSm)
}

func SpinnerMd() gomponents.Node {
	return spinner(SizeMd)
}

func SpinnerLg() gomponents.Node {
	return spinner(SizeLg)
}

func spinnerContainer(size size, height string) gomponents.Node {
	return html.Div(
		components.Classes{
			"flex justify-center": true,
			"items-center w-full": true,
		},
		html.Style(fmt.Sprintf("height: %s;", height)),
		spinner(size),
	)
}

func SpinnerContainerSm(height ...string) gomponents.Node {
	pickedHeight := "300px"
	if len(height) > 0 {
		pickedHeight = height[0]
	}
	return spinnerContainer(SizeSm, pickedHeight)
}

func SpinnerContainerMd(height ...string) gomponents.Node {
	pickedHeight := "300px"
	if len(height) > 0 {
		pickedHeight = height[0]
	}
	return spinnerContainer(SizeMd, pickedHeight)
}

func SpinnerContainerLg(height ...string) gomponents.Node {
	pickedHeight := "300px"
	if len(height) > 0 {
		pickedHeight = height[0]
	}
	return spinnerContainer(SizeLg, pickedHeight)
}
