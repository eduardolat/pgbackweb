package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func StatusBadge(status string) gomponents.Node {
	class := ""
	switch status {
	case "running":
		class = "badge-info"
	case "success":
		class = "badge-success"
	case "failed":
		class = "badge-error"
	case "deleted":
		class = "badge-warning"
	default:
		class = "badge-neutral"
	}

	return html.Span(
		components.Classes{
			"badge": true,
			class:   true,
		},
		gomponents.Text(status),
	)
}
