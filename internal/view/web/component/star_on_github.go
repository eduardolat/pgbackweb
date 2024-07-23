package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func StarOnGithub(size size) gomponents.Node {
	small := html.IFrame(
		html.Src("https://ghbtns.com/github-btn.html?user=eduardolat&repo=pgbackweb&type=star&count=true"),
		gomponents.Attr("frameborder", "0"),
		gomponents.Attr("scrolling", "0"),
		html.Width("150"),
		html.Height("20"),
	)

	big := html.IFrame(
		html.Src("https://ghbtns.com/github-btn.html?user=eduardolat&repo=pgbackweb&type=star&count=true&size=large"),
		gomponents.Attr("frameborder", "0"),
		gomponents.Attr("scrolling", "0"),
		html.Width("170"),
		html.Height("30"),
	)

	switch size {
	case SizeSm:
		return small
	case SizeMd:
		return big
	default:
		return big
	}
}
