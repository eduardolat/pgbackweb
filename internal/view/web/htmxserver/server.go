package htmxserver

import (
	"github.com/labstack/echo/v4"
)

// ServerIsBoosted indicates that the request is via an element using hx-boost.
func ServerIsBoosted(c echo.Context) bool {
	return c.Request().Header.Get("HX-Boosted") != ""
}

// ServerGetCurrentURL of the browser.
func ServerGetCurrentURL(c echo.Context) string {
	return c.Request().Header.Get("HX-Current-URL")
}

// ServerIsHistoryRestoreRequest returns whether the request is for history restoration
// after a miss in the local history cache.
func ServerIsHistoryRestoreRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-History-Restore-Request") == "true"
}

// ServerGetPrompt from the user response to an hx-prompt.
// See https://htmx.org/attributes/hx-prompt
func ServerGetPrompt(c echo.Context) string {
	return c.Request().Header.Get("HX-Prompt")
}

// ServerIsRequest returns whether this is a HTMX request.
func ServerIsRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

// ServerGetTarget returns the id of the target element if it exists.
func ServerGetTarget(c echo.Context) string {
	return c.Request().Header.Get("HX-Target")
}

// ServerGetTriggerName returns the name of the triggered element if it exists.
func ServerGetTriggerName(c echo.Context) string {
	return c.Request().Header.Get("HX-Trigger-Name")
}

// ServerGetTrigger returns the id of the triggered element if it exists.
func ServerGetTrigger(c echo.Context) string {
	return c.Request().Header.Get("HX-Trigger")
}

// ServerSetLocation allows you to do a client-side redirect that does
// not do a full page reload.
// See https://htmx.org/headers/hx-location
func ServerSetLocation(c echo.Context, v string) {
	c.Response().Header().Set("HX-Location", v)
}

// ServerSetPushURL pushes a new URL into the history stack.
// See https://htmx.org/headers/hx-push-url
func ServerSetPushURL(c echo.Context, v string) {
	c.Response().Header().Set("HX-Push-Url", v)
}

// ServerSetRedirect can be used to do a client-side redirect to a new location.
func ServerSetRedirect(c echo.Context, v string) {
	c.Response().Header().Set("HX-Redirect", v)
}

// ServerSetRefresh will make the client side do a a full refresh of the page.
func ServerSetRefresh(c echo.Context) {
	c.Response().Header().Set("HX-Refresh", "true")
}

// ServerSetReplaceURL replaces the current URL in the location bar.
// See https://htmx.org/headers/hx-replace-url
func ServerSetReplaceURL(c echo.Context, v string) {
	c.Response().Header().Set("HX-Replace-Url", v)
}

// ServerSetReswap allows you to specify how the response will be swapped.
// See https://htmx.org/attributes/hx-swap
func ServerSetReswap(c echo.Context, v string) {
	c.Response().Header().Set("HX-Reswap", v)
}

// ServerSetRetarget sets a CSS selector that updates the target of the
// content update to a different element on the page.
func ServerSetRetarget(c echo.Context, v string) {
	c.Response().Header().Set("HX-Retarget", v)
}

// ServerSetTrigger allows you to trigger client side events.
// See https://htmx.org/headers/hx-trigger
func ServerSetTrigger(c echo.Context, v string) {
	c.Response().Header().Set("HX-Trigger", v)
}

// ServerSetTriggerAfterSettle allows you to trigger client side events.
// See https://htmx.org/headers/hx-trigger
func ServerSetTriggerAfterSettle(c echo.Context, v string) {
	c.Response().Header().Set("HX-Trigger-After-Settle", v)
}

// ServerSetTriggerAfterSwap allows you to trigger client side events.
// See https://htmx.org/headers/hx-trigger
func ServerSetTriggerAfterSwap(c echo.Context, v string) {
	c.Response().Header().Set("HX-Trigger-After-Swap", v)
}
