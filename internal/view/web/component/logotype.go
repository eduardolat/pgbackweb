package component

import (
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	nodx "github.com/nodxdev/nodxgo"
)

func Logotype() nodx.Node {
	return nodx.Div(
		nodx.ClassMap{
			"inline space-x-2 select-none":    true,
			"flex justify-start items-center": true,
		},
		nodx.Img(
			nodx.Class("w-[60px] h-auto"),
			nodx.Src(pathutil.BuildPath("/images/logo.png")),
			nodx.Alt("PG Back Web"),
		),
		nodx.SpanEl(
			nodx.Class("text-2xl font-bold"),
			nodx.Text("PG Back Web"),
		),
	)
}
