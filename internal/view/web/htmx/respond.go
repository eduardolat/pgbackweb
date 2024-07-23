package htmx

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

// setCustomTrigger will set a custom trigger.
// Is the basic helper for all the other custom triggers.
func setCustomTrigger(c echo.Context, key string, value string) {
	value = strings.ReplaceAll(value, `"`, `\"`)
	value = url.PathEscape(value)
	s := fmt.Sprintf(`{"%s": "%s"}`, key, value)
	ServerSetTrigger(c, s)
}

// RespondAlert shows an alert.
func RespondAlert(c echo.Context, message string) error {
	ServerSetReswap(c, "none")
	setCustomTrigger(c, "ctm_alert", message)
	return c.NoContent(http.StatusOK)
}

// RespondAlertWithRefresh shows an alert and refreshes the page using HTMX.
func RespondAlertWithRefresh(c echo.Context, message string) error {
	ServerSetReswap(c, "none")
	setCustomTrigger(c, "ctm_alert_with_refresh", message)
	return c.NoContent(http.StatusOK)
}

// RespondAlertWithRedirect shows an alert and redirects the page using JS.
func RespondAlertWithRedirect(c echo.Context, message, url string) error {
	msg := fmt.Sprintf("%s-::-::-%s", message, url)

	ServerSetReswap(c, "none")
	setCustomTrigger(c, "ctm_alert_with_redirect", msg)
	return c.NoContent(http.StatusOK)
}

// RespondToastSuccess reswaps the HTMX to none and shows an success message
// inside a toast.
func RespondToastSuccess(c echo.Context, message string) error {
	ServerSetReswap(c, "none")
	setCustomTrigger(c, "ctm_toast_success", message)
	return c.NoContent(http.StatusOK)
}

// RespondToastError reswaps the HTMX to none and shows an error message inside
// a toast.
func RespondToastError(c echo.Context, message string) error {
	ServerSetReswap(c, "none")
	setCustomTrigger(c, "ctm_toast_error", message)
	return c.NoContent(http.StatusOK)
}

// RespondToastSuccessInfinite reswaps the HTMX to none and shows an success
// message inside an infinite toast.
func RespondToastSuccessInfinite(c echo.Context, message string) error {
	ServerSetReswap(c, "none")
	setCustomTrigger(c, "ctm_toast_success_infinite", message)
	return c.NoContent(http.StatusOK)
}

// RespondToastErrorInfinite reswaps the HTMX to none and shows an error message
// inside an infinite toast.
func RespondToastErrorInfinite(c echo.Context, message string) error {
	ServerSetReswap(c, "none")
	setCustomTrigger(c, "ctm_toast_error_infinite", message)
	return c.NoContent(http.StatusOK)
}

// RespondRedirect redirects the user to the given URL using HTMX.
func RespondRedirect(c echo.Context, url string) error {
	ServerSetReswap(c, "none")
	ServerSetRedirect(c, url)
	return c.NoContent(http.StatusOK)
}

// RespondRefresh refreshes the page using HTMX.
func RespondRefresh(c echo.Context) error {
	ServerSetReswap(c, "none")
	ServerSetRefresh(c)
	return c.NoContent(http.StatusOK)
}
