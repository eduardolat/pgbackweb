package component

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/integration/postgres"
	nodx "github.com/nodxdev/nodxgo"
)

func PGVersionSelectOptions(selectedVersion sql.NullString) nodx.Node {
	return nodx.Map(
		postgres.PGVersionsDesc,
		func(pgVersion postgres.PGVersion) nodx.Node {
			return nodx.Option(
				nodx.Value(pgVersion.Value.Version),
				nodx.Textf("PostgreSQL %s", pgVersion.Value.Version),
				nodx.If(
					selectedVersion.Valid && selectedVersion.String == pgVersion.Value.Version,
					nodx.Selected(""),
				),
			)
		},
	)
}
