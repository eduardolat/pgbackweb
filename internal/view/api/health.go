package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) healthHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var queryData struct {
		IncludeDatabases    bool `query:"databases"`
		IncludeDestinations bool `query:"destinations"`
	}
	if err := c.Bind(&queryData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	databasesHealthy, destinationsHealthy := true, true

	if queryData.IncludeDatabases {
		databases, err := h.servs.DatabasesService.GetAllDatabases(ctx)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		for _, db := range databases {
			if db.TestOk.Valid && !db.TestOk.Bool {
				databasesHealthy = false
				break
			}
		}

	}

	if queryData.IncludeDestinations {
		destinations, err := h.servs.DestinationsService.GetAllDestinations(ctx)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		for _, dest := range destinations {
			if dest.TestOk.Valid && !dest.TestOk.Bool {
				destinationsHealthy = false
				break
			}
		}
	}

	response := map[string]any{
		"server_healthy": true,
	}
	if queryData.IncludeDatabases {
		response["databases_healthy"] = databasesHealthy
	}
	if queryData.IncludeDestinations {
		response["destinations_healthy"] = destinationsHealthy
	}

	return c.JSON(http.StatusOK, response)
}
