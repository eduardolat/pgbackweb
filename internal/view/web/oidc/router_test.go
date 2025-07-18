package oidc

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/service/oidc"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleOIDCError(t *testing.T) {
	t.Run("HTMX Request", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("HX-Request", "true")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		err := handleOIDCError(c, "Test error message")
		
		// Should handle HTMX request (returns no error, sets headers for toast)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		// Check HTMX headers are set for toast error
		assert.Contains(t, rec.Header().Get("HX-Reswap"), "none")
		assert.Contains(t, rec.Header().Get("HX-Trigger"), "ctm_toast_error")
	})

	t.Run("Regular Browser Request", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		err := handleOIDCError(c, "Test error message")
		
		assert.NoError(t, err)
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		
		location := rec.Header().Get("Location")
		assert.Contains(t, location, "/auth/login")
		assert.Contains(t, location, url.QueryEscape("Test error message"))
	})
}

func TestOIDCLoginHandler_StateGeneration(t *testing.T) {
	t.Run("State Generation and Cookie Setting", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		// Create a mock OIDC service that's enabled
		mockOIDCService := &oidc.Service{}
		// Note: In a real test, you'd need to properly mock this
		// For now, we'll just test the structure
		
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		// This will succeed because GenerateState() doesn't depend on config
		// and GetAuthURL() will return empty string but won't error
		err := h.oidcLoginHandler(c)
		
		// The function will succeed and redirect (even to empty URL)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusFound, rec.Code)
		
		// Check that the state cookie was set
		cookies := rec.Result().Cookies()
		foundStateCookie := false
		for _, cookie := range cookies {
			if cookie.Name == "oidc_state" {
				foundStateCookie = true
				assert.NotEmpty(t, cookie.Value)
				assert.True(t, cookie.HttpOnly)
				assert.Equal(t, 300, cookie.MaxAge)
				break
			}
		}
		assert.True(t, foundStateCookie, "oidc_state cookie should be set")
	})
}

func TestOIDCCallbackHandler_StateValidation(t *testing.T) {
	t.Run("State Mismatch", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/callback?code=test-code&state=wrong-state", nil)
		req.AddCookie(&http.Cookie{
			Name:  "oidc_state",
			Value: "correct-state",
		})
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		err := h.oidcCallbackHandler(c)
		
		// Should handle state mismatch by redirecting (returning no error)
		assert.NoError(t, err)
		// Check that it redirected to login with error
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		assert.Contains(t, rec.Header().Get("Location"), "/auth/login")
	})

	t.Run("Missing State Cookie", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/callback?code=test-code&state=test-state", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		err := h.oidcCallbackHandler(c)
		
		// Should handle missing state cookie by redirecting (returns no error)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		assert.Contains(t, rec.Header().Get("Location"), "/auth/login")
	})

	t.Run("OIDC Provider Error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/callback?error=access_denied&error_description=User+denied+access", nil)
		req.AddCookie(&http.Cookie{
			Name:  "oidc_state",
			Value: "test-state",
		})
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		err := h.oidcCallbackHandler(c)
		
		// Should handle provider error by redirecting (returns no error)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		assert.Contains(t, rec.Header().Get("Location"), "/auth/login")
	})

	t.Run("Missing Authorization Code", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/callback?state=test-state", nil)
		req.AddCookie(&http.Cookie{
			Name:  "oidc_state",
			Value: "test-state",
		})
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		err := h.oidcCallbackHandler(c)
		
		// Should handle missing code by redirecting (returns no error)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		assert.Contains(t, rec.Header().Get("Location"), "/auth/login")
	})
}

func TestOIDCCallbackHandler_AdvancedScenarios(t *testing.T) {
	t.Run("State Parameter Empty", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/callback?code=test-code&state=", nil)
		req.AddCookie(&http.Cookie{
			Name:  "oidc_state",
			Value: "valid-state",
		})
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		err := h.oidcCallbackHandler(c)
		
		// Should handle empty state parameter by redirecting
		assert.NoError(t, err)
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		assert.Contains(t, rec.Header().Get("Location"), "/auth/login")
	})
	
	t.Run("State Parameter Missing", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/callback?code=test-code", nil)
		req.AddCookie(&http.Cookie{
			Name:  "oidc_state",
			Value: "valid-state",
		})
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		err := h.oidcCallbackHandler(c)
		
		// Should handle missing state parameter by redirecting
		assert.NoError(t, err)
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		assert.Contains(t, rec.Header().Get("Location"), "/auth/login")
	})
	
	t.Run("HTMX Callback Request", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/callback?code=test-code&state=wrong-state", nil)
		req.Header.Set("HX-Request", "true")
		req.AddCookie(&http.Cookie{
			Name:  "oidc_state",
			Value: "correct-state",
		})
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		err := h.oidcCallbackHandler(c)
		
		// Should handle HTMX request with toast error
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Header().Get("HX-Reswap"), "none")
		assert.Contains(t, rec.Header().Get("HX-Trigger"), "ctm_toast_error")
	})
}

func TestOIDCLoginHandler_EdgeCases(t *testing.T) {
	t.Run("HTMX Login Request", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/auth/oidc/login", nil)
		req.Header.Set("HX-Request", "true")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		err := h.oidcLoginHandler(c)
		
		// Should redirect normally even for HTMX requests
		assert.NoError(t, err)
		assert.Equal(t, http.StatusFound, rec.Code)
	})
}

func TestHandlers_ServiceIntegration(t *testing.T) {
	t.Run("Handler Creation", func(t *testing.T) {
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		h := handlers{servs: mockServices}
		
		// Verify handlers struct is created correctly
		assert.NotNil(t, h.servs)
		assert.NotNil(t, h.servs.OIDCService)
	})
}

func TestMountRouter_EnabledCheck(t *testing.T) {
	t.Run("OIDC Disabled - No Panic", func(t *testing.T) {
		e := echo.New()
		group := e.Group("/auth")
		
		// Create a disabled OIDC service
		mockOIDCService := &oidc.Service{}
		mockServices := &service.Service{
			OIDCService: mockOIDCService,
		}
		
		mockMiddleware := &middleware.Middleware{}
		
		// This should not panic even if OIDC is disabled
		assert.NotPanics(t, func() {
			MountRouter(group, mockMiddleware, mockServices)
		})
	})
}
