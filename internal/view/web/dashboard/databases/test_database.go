package databases

import (
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmxserver"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handlers) testDatabaseHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData createDatabaseDTO
	if err := c.Bind(&formData); err != nil {
		return htmxserver.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmxserver.RespondToastError(c, err.Error())
	}

	err := h.servs.DatabasesService.TestDatabase(
		ctx, formData.Version, formData.ConnectionString,
	)
	if err != nil {
		return htmxserver.RespondToastError(c, err.Error())
	}

	return htmxserver.RespondToastSuccess(c, "Connection successful")
}

func (h *handlers) testExistingDatabaseHandler(c echo.Context) error {
	ctx := c.Request().Context()
	databaseID, err := uuid.Parse(c.Param("databaseID"))
	if err != nil {
		return htmxserver.RespondToastError(c, err.Error())
	}

	err = h.servs.DatabasesService.TestDatabaseAndStoreResult(ctx, databaseID)
	if err != nil {
		return htmxserver.RespondToastError(c, err.Error())
	}

	return htmxserver.RespondToastSuccess(c, "Connection successful")
}
