package component

import (
	"database/sql"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func PrettyDestinationName(
	isLocal bool, destinationName sql.NullString,
) gomponents.Node {
	icon := lucide.Cloud
	if !destinationName.Valid {
		destinationName = sql.NullString{
			Valid:  true,
			String: "Unknown destination",
		}
	}

	if isLocal {
		icon = lucide.HardDrive
		destinationName = sql.NullString{
			Valid:  true,
			String: "Local",
		}
	}

	return html.Span(
		html.Class("inline flex justify-start items-center space-x-1 font-mono"),
		icon(),
		SpanText(destinationName.String),
	)
}
