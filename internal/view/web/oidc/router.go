package oidc

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/service/oidc"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	servs *service.Service
}

// MountRouter registers OIDC authentication routes on the provided Echo group if OIDC is enabled.
// It sets up login and callback endpoints under a middleware group that requires no prior authentication.
func MountRouter(
	parent *echo.Group, mids *middleware.Middleware, servs *service.Service,
) {
	if !servs.OIDCService.IsEnabled() {
		return
	}

	h := handlers{servs: servs}

	requireNoAuth := parent.Group("", mids.RequireNoAuth)

	requireNoAuth.GET("/oidc/login", h.oidcLoginHandler)
	requireNoAuth.GET("/oidc/callback", h.oidcCallbackHandler)
}

func (h *handlers) oidcLoginHandler(c echo.Context) error {
	state, err := h.servs.OIDCService.GenerateState()
	if err != nil {
		logger.Error("OIDC: failed to generate state", logger.KV{
			"ip":    c.RealIP(),
			"ua":    c.Request().UserAgent(),
			"error": err,
		})
		return handleOIDCError(c, "OIDC: Unable to initiate login")
	}

	// Store state in session/cookie for verification
	c.SetCookie(&http.Cookie{
		Name:     "oidc_state",
		Value:    state,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   300, // 5 minutes
		Path:     "/",
	})

	authURL := h.servs.OIDCService.GetAuthURL(state)
	return c.Redirect(http.StatusFound, authURL)
}

func (h *handlers) oidcCallbackHandler(c echo.Context) error {
	ctx := c.Request().Context()

	// Verify state parameter
	state := c.QueryParam("state")
	stateCookie, err := c.Cookie("oidc_state")
	if err != nil || stateCookie == nil || stateCookie.Value != state {
		expectedValue := ""
		if stateCookie != nil {
			expectedValue = stateCookie.Value
		}
		logger.Error("OIDC: state mismatch", logger.KV{
			"ip":       c.RealIP(),
			"ua":       c.Request().UserAgent(),
			"state":    state,
			"expected": expectedValue,
		})
		return handleOIDCError(c, "OIDC: Invalid state parameter")
	}

	// Clear the state cookie
	c.SetCookie(&http.Cookie{
		Name:     "oidc_state",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
		Path:     "/",
	})

	// Check for error from OIDC provider
	if errorParam := c.QueryParam("error"); errorParam != "" {
		errorDesc := c.QueryParam("error_description")
		errorMsg := "OIDC: Login failed"
		if errorDesc != "" {
			errorMsg = "OIDC: " + errorDesc
		} else {
			errorMsg = "OIDC: " + errorParam
		}
		logger.Error("OIDC provider returned error", logger.KV{
			"ip":                c.RealIP(),
			"ua":                c.Request().UserAgent(),
			"error":             errorParam,
			"error_description": errorDesc,
		})
		return handleOIDCError(c, errorMsg)
	}

	code := c.QueryParam("code")
	if code == "" {
		return handleOIDCError(c, "OIDC: Missing authorization code")
	}

	// Exchange code for user info
	userInfo, err := h.servs.OIDCService.ExchangeCode(ctx, code)
	if err != nil {
		logger.Error("failed to exchange OIDC code", logger.KV{
			"ip":    c.RealIP(),
			"ua":    c.Request().UserAgent(),
			"error": err,
		})
		return handleOIDCError(c, "OIDC: Unable to authenticate with provider")
	}

	// Create or update user
	user, err := h.servs.OIDCService.CreateOrUpdateUser(ctx, userInfo)
	if err != nil {
		errorMsg := "OIDC: Unable to create user account"
		if errors.Is(err, oidc.ErrEmailAlreadyExists) {
			errorMsg = "OIDC: Email already exists. Use regular login."
		}
		logger.Error("failed to create/update OIDC user", logger.KV{
			"ip":    c.RealIP(),
			"ua":    c.Request().UserAgent(),
			"email": userInfo.Email,
			"error": err,
		})
		return handleOIDCError(c, errorMsg)
	}

	logger.Info("OIDC: authentication successful", logger.KV{
		"email":   userInfo.Email,
		"name":    userInfo.Name,
		"subject": userInfo.Subject,
		"user_id": user.ID,
	})

	// Create session for the user
	session, err := h.servs.AuthService.LoginOIDC(
		ctx, user.ID, c.RealIP(), c.Request().UserAgent(),
	)
	if err != nil {
		logger.Error("OIDC: failed to create session for user", logger.KV{
			"ip":      c.RealIP(),
			"ua":      c.Request().UserAgent(),
			"user_id": user.ID,
			"error":   err,
		})
		return handleOIDCError(c, "OIDC: Unable to create session")
	}

	// Set session cookie and redirect to dashboard
	h.servs.AuthService.SetSessionCookie(c, session.DecryptedToken)
	return c.Redirect(http.StatusSeeOther, "/dashboard")
}

// handleOIDCError sends an OIDC error response as an HTMX toast if the request is from HTMX, or redirects to the login page with an error message for regular browser requests.
func handleOIDCError(c echo.Context, message string) error {
	// Check if it's an HTMX request
	if c.Request().Header.Get("HX-Request") != "" {
		return respondhtmx.ToastError(c, message)
	}

	// For regular browser requests, redirect to login with error parameter
	return c.Redirect(http.StatusSeeOther, "/auth/login?error="+url.QueryEscape(message))
}
