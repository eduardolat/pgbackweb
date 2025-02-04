package component

import (
	"database/sql"

	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func PrettyDestinationName(
	isLocal bool, destinationName sql.NullString,
) nodx.Node {
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

	return nodx.SpanEl(
		nodx.Class("inline flex justify-start items-center space-x-1 font-mono"),
		icon(),
		SpanText(destinationName.String),
	)
}
