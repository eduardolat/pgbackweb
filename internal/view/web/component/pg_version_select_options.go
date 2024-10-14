package component

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/integration/postgres"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func PGVersionSelectOptions(selectedVersion sql.NullString) gomponents.Node {
	return GMap(
		postgres.PGVersions,
		func(pgVersion postgres.PGVersion) gomponents.Node {
			return html.Option(
				html.Value(pgVersion.Value.Version),
				gomponents.Textf("PostgreSQL %s", pgVersion.Value.Version),
				gomponents.If(
					selectedVersion.Valid && selectedVersion.String == pgVersion.Value.Version,
					html.Selected(),
				),
			)
		},
	)
}
