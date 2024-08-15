package component

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// PrettyFileSize pretty prints a file size (in bytes) to a human-readable format.
// If the size is not valid, it returns an empty string.
//
// e.g. 1024 -> 1 KB
func PrettyFileSize(
	size sql.NullInt64,
) gomponents.Node {
	return gomponents.If(
		size.Valid,
		html.Span(
			SpanText(strutil.FormatFileSize(size.Int64)),
		),
	)
}
