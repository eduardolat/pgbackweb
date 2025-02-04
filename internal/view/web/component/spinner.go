package component

import (
	"fmt"

	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func spinner(size size) nodx.Node {
	return lucide.LoaderCircle(nodx.ClassMap{
		"animate-spin inline-block": true,
		"size-5":                    size == SizeSm,
		"size-8":                    size == SizeMd,
		"size-12":                   size == SizeLg,
	})
}

func SpinnerSm() nodx.Node {
	return spinner(SizeSm)
}

func SpinnerMd() nodx.Node {
	return spinner(SizeMd)
}

func SpinnerLg() nodx.Node {
	return spinner(SizeLg)
}

func spinnerContainer(size size, height string) nodx.Node {
	return nodx.Div(
		nodx.ClassMap{
			"flex justify-center": true,
			"items-center w-full": true,
		},
		nodx.StyleAttr(fmt.Sprintf("height: %s;", height)),
		spinner(size),
	)
}

func SpinnerContainerSm(height ...string) nodx.Node {
	pickedHeight := "300px"
	if len(height) > 0 {
		pickedHeight = height[0]
	}
	return spinnerContainer(SizeSm, pickedHeight)
}

func SpinnerContainerMd(height ...string) nodx.Node {
	pickedHeight := "300px"
	if len(height) > 0 {
		pickedHeight = height[0]
	}
	return spinnerContainer(SizeMd, pickedHeight)
}

func SpinnerContainerLg(height ...string) nodx.Node {
	pickedHeight := "300px"
	if len(height) > 0 {
		pickedHeight = height[0]
	}
	return spinnerContainer(SizeLg, pickedHeight)
}
