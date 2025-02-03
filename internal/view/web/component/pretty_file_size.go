package component

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	nodx "github.com/nodxdev/nodxgo"
)

// PrettyFileSize pretty prints a file size (in bytes) to a human-readable format.
// If the size is not valid, it returns an empty string.
//
// e.g. 1024 -> 1 KB
func PrettyFileSize(
	size sql.NullInt64,
) nodx.Node {
	return nodx.If(
		size.Valid,
		nodx.SpanEl(
			SpanText(strutil.FormatFileSize(size.Int64)),
		),
	)
}
