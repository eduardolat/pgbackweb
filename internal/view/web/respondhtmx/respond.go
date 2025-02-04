package respondhtmx

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	htmx "github.com/nodxdev/nodxgo-htmx"
)

// setCustomTrigger will set a custom trigger.
// Is the basic helper for all the other custom triggers.
func setCustomTrigger(c echo.Context, key string, value string) {
	value = strings.ReplaceAll(value, `"`, `\"`)
	value = url.PathEscape(value)
	s := fmt.Sprintf(`{"%s": "%s"}`, key, value)
	htmx.ServerSetTrigger(c.Response().Header(), s)
}

// Alert shows an alert.
func Alert(c echo.Context, message string) error {
	htmx.ServerSetReswap(c.Response().Header(), "none")
	setCustomTrigger(c, "ctm_alert", message)
	return c.NoContent(http.StatusOK)
}

// AlertWithRefresh shows an alert and refreshes the page using HTMX.
func AlertWithRefresh(c echo.Context, message string) error {
	htmx.ServerSetReswap(c.Response().Header(), "none")
	setCustomTrigger(c, "ctm_alert_with_refresh", message)
	return c.NoContent(http.StatusOK)
}

// AlertWithRedirect shows an alert and redirects the page using JS.
func AlertWithRedirect(c echo.Context, message, url string) error {
	msg := fmt.Sprintf("%s-::-::-%s", message, url)

	htmx.ServerSetReswap(c.Response().Header(), "none")
	setCustomTrigger(c, "ctm_alert_with_redirect", msg)
	return c.NoContent(http.StatusOK)
}

// ToastSuccess reswaps the HTMX to none and shows an success message
// inside a toast.
func ToastSuccess(c echo.Context, message string) error {
	htmx.ServerSetReswap(c.Response().Header(), "none")
	setCustomTrigger(c, "ctm_toast_success", message)
	return c.NoContent(http.StatusOK)
}

// ToastError reswaps the HTMX to none and shows an error message inside
// a toast.
func ToastError(c echo.Context, message string) error {
	htmx.ServerSetReswap(c.Response().Header(), "none")
	setCustomTrigger(c, "ctm_toast_error", message)
	return c.NoContent(http.StatusOK)
}

// ToastSuccessInfinite reswaps the HTMX to none and shows an success
// message inside an infinite toast.
func ToastSuccessInfinite(c echo.Context, message string) error {
	htmx.ServerSetReswap(c.Response().Header(), "none")
	setCustomTrigger(c, "ctm_toast_success_infinite", message)
	return c.NoContent(http.StatusOK)
}

// ToastErrorInfinite reswaps the HTMX to none and shows an error message
// inside an infinite toast.
func ToastErrorInfinite(c echo.Context, message string) error {
	htmx.ServerSetReswap(c.Response().Header(), "none")
	setCustomTrigger(c, "ctm_toast_error_infinite", message)
	return c.NoContent(http.StatusOK)
}

// Redirect redirects the user to the given URL using HTMX.
func Redirect(c echo.Context, url string) error {
	htmx.ServerSetReswap(c.Response().Header(), "none")
	htmx.ServerSetRedirect(c.Response().Header(), url)
	return c.NoContent(http.StatusOK)
}

// Refresh refreshes the page using HTMX.
func Refresh(c echo.Context) error {
	htmx.ServerSetReswap(c.Response().Header(), "none")
	htmx.ServerSetRefresh(c.Response().Header(), "true")
	return c.NoContent(http.StatusOK)
}
