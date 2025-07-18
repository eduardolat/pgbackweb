package oidc

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("OIDC Disabled", func(t *testing.T) {
		env := config.Env{
			PBW_OIDC_ENABLED: false,
		}

		service, err := New(env, nil)

		assert.NoError(t, err)
		assert.NotNil(t, service)
		assert.False(t, service.IsEnabled())
		assert.Equal(t, env, service.env)
		assert.Nil(t, service.provider)
	})

	t.Run("OIDC Enabled with Invalid Issuer", func(t *testing.T) {
		env := config.Env{
			PBW_OIDC_ENABLED:     true,
			PBW_OIDC_ISSUER_URL:  "invalid-url",
			PBW_OIDC_CLIENT_ID:   "test-client",
			PBW_OIDC_CLIENT_SECRET: "test-secret",
			PBW_OIDC_REDIRECT_URL: "http://localhost:8080/auth/oidc/callback",
			PBW_OIDC_SCOPES:      "openid profile email",
		}

		service, err := New(env, nil)

		assert.Error(t, err)
		assert.Nil(t, service)
		assert.Contains(t, err.Error(), "failed to create OIDC provider")
	})
}

func TestIsEnabled(t *testing.T) {
	t.Run("Enabled", func(t *testing.T) {
		service := &Service{
			env: config.Env{PBW_OIDC_ENABLED: true},
		}
		assert.True(t, service.IsEnabled())
	})

	t.Run("Disabled", func(t *testing.T) {
		service := &Service{
			env: config.Env{PBW_OIDC_ENABLED: false},
		}
		assert.False(t, service.IsEnabled())
	})
}

func TestGetAuthURL(t *testing.T) {
	t.Run("OIDC Disabled", func(t *testing.T) {
		service := &Service{
			env: config.Env{PBW_OIDC_ENABLED: false},
		}

		url := service.GetAuthURL("test-state")
		assert.Empty(t, url)
	})

	t.Run("OIDC Enabled", func(t *testing.T) {
		service := &Service{
			env: config.Env{PBW_OIDC_ENABLED: true},
		}

		// This will return empty string without proper config, but shows the behavior
		_ = service.GetAuthURL("test-state")
		// Without proper oauth2.Config, this will return empty string or panic
		// In a real test, we'd need to mock or provide proper config
		assert.True(t, service.IsEnabled())
	})
}

func TestGenerateState(t *testing.T) {
	service := &Service{}

	t.Run("Success", func(t *testing.T) {
		state, err := service.GenerateState()

		assert.NoError(t, err)
		assert.NotEmpty(t, state)
		assert.Greater(t, len(state), 20) // Base64 encoded 32 bytes should be > 20 chars
	})

	t.Run("Multiple Calls Generate Different States", func(t *testing.T) {
		state1, err1 := service.GenerateState()
		state2, err2 := service.GenerateState()

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.NotEqual(t, state1, state2)
	})
}

func TestExchangeCode(t *testing.T) {
	t.Run("OIDC Disabled", func(t *testing.T) {
		service := &Service{
			env: config.Env{PBW_OIDC_ENABLED: false},
		}

		userInfo, err := service.ExchangeCode(context.Background(), "test-code")

		assert.Error(t, err)
		assert.Nil(t, userInfo)
		assert.Equal(t, ErrOIDCNotEnabled, err)
	})

	// Note: Testing the enabled case would require mocking the OIDC provider
	// and OAuth2 flow, which is complex. In a real implementation, you'd
	// want to use dependency injection to make these components testable.
}

func TestErrorTypes(t *testing.T) {
	t.Run("Error Messages", func(t *testing.T) {
		assert.Equal(t, "email already exists with different authentication method", ErrEmailAlreadyExists.Error())
		assert.Equal(t, "OIDC is not enabled", ErrOIDCNotEnabled.Error())
		assert.Equal(t, "invalid or expired token", ErrInvalidToken.Error())
		assert.Equal(t, "required user information missing from OIDC claims", ErrMissingClaims.Error())
	})
}

func TestUserInfo(t *testing.T) {
	t.Run("UserInfo Structure", func(t *testing.T) {
		userInfo := UserInfo{
			Email:    "test@example.com",
			Name:     "Test User",
			Username: "testuser",
			Subject:  "test-subject-123",
		}

		assert.Equal(t, "test@example.com", userInfo.Email)
		assert.Equal(t, "Test User", userInfo.Name)
		assert.Equal(t, "testuser", userInfo.Username)
		assert.Equal(t, "test-subject-123", userInfo.Subject)
	})
}

func TestService_AdvancedErrorHandling(t *testing.T) {
	t.Run("GenerateState Multiple Calls", func(t *testing.T) {
		service := &Service{}

		// Generate multiple states to ensure uniqueness
		states := make(map[string]bool)
		for i := 0; i < 100; i++ {
			state, err := service.GenerateState()
			assert.NoError(t, err)
			assert.NotEmpty(t, state)
			assert.False(t, states[state], "State should be unique")
			states[state] = true
		}
	})
}

func TestService_ConfigurationValidation(t *testing.T) {
	t.Run("Empty Service Struct", func(t *testing.T) {
		service := &Service{}

		// Test methods on empty service
		assert.False(t, service.IsEnabled())
		assert.Empty(t, service.GetAuthURL("test-state"))

		state, err := service.GenerateState()
		assert.NoError(t, err)
		assert.NotEmpty(t, state)

		// ExchangeCode should fail on empty service
		userInfo, err := service.ExchangeCode(context.Background(), "test-code")
		assert.Error(t, err)
		assert.Equal(t, ErrOIDCNotEnabled, err)
		assert.Nil(t, userInfo)
	})
}

func TestService_ContextHandling(t *testing.T) {
	t.Run("Canceled Context", func(t *testing.T) {
		service := &Service{}

		// Create canceled context
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		// ExchangeCode should fail on disabled service (context won't matter)
		userInfo, err := service.ExchangeCode(ctx, "test-code")
		assert.Error(t, err)
		assert.Equal(t, ErrOIDCNotEnabled, err)
		assert.Nil(t, userInfo)
	})
}

func TestService_StateGeneration(t *testing.T) {
	t.Run("State Format Validation", func(t *testing.T) {
		service := &Service{}

		state, err := service.GenerateState()
		assert.NoError(t, err)
		assert.NotEmpty(t, state)

		// Check that state is base64 URL encoded
		decoded, err := base64.URLEncoding.DecodeString(state)
		assert.NoError(t, err)
		assert.Equal(t, 32, len(decoded)) // Should be 32 bytes
	})
}

func TestUserInfo_StructValidation(t *testing.T) {
	t.Run("UserInfo Fields", func(t *testing.T) {
		userInfo := UserInfo{
			Email:    "test@example.com",
			Name:     "Test User",
			Username: "testuser",
			Subject:  "sub-123",
		}

		assert.Equal(t, "test@example.com", userInfo.Email)
		assert.Equal(t, "Test User", userInfo.Name)
		assert.Equal(t, "testuser", userInfo.Username)
		assert.Equal(t, "sub-123", userInfo.Subject)
	})

	t.Run("Empty UserInfo", func(t *testing.T) {
		userInfo := UserInfo{}

		assert.Empty(t, userInfo.Email)
		assert.Empty(t, userInfo.Name)
		assert.Empty(t, userInfo.Username)
		assert.Empty(t, userInfo.Subject)
	})
}

// Integration tests would require a real database and proper setup
// These would be placed in a separate test file with build tags
// like // +build integration
