package component

import (
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func StarOnGithub(size size) nodx.Node {
	return nodx.A(
		alpine.XData("alpineStarOnGithub()"),
		alpine.XCloak(),
		nodx.ClassMap{
			"btn btn-neutral": true,
			"btn-sm":          size == SizeSm,
			"btn-lg":          size == SizeLg,
		},
		nodx.Href("https://github.com/eduardolat/pgbackweb"),
		nodx.Target("_blank"),
		lucide.Github(),
		SpanText("Star"),
		nodx.SpanEl(
			alpine.XShow("stars"),
			alpine.XText("'( ' + stars + ' )'"),
		),
	)
}
