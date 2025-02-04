package destinations

import (
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handlers) testDestinationHandler(c echo.Context) error {
	var formData createDestinationDTO
	if err := c.Bind(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	err := h.servs.DestinationsService.TestDestination(
		formData.AccessKey, formData.SecretKey, formData.Region, formData.Endpoint,
		formData.BucketName,
	)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.ToastSuccess(c, "Connection successful")
}

func (h *handlers) testExistingDestinationHandler(c echo.Context) error {
	ctx := c.Request().Context()
	destinationID, err := uuid.Parse(c.Param("destinationID"))
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	err = h.servs.DestinationsService.TestDestinationAndStoreResult(ctx, destinationID)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.ToastSuccess(c, "Connection successful")
}
