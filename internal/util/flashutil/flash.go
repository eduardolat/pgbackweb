package flashutil

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	flashCookieName = "pbw_flash"
)

// FlashType represents the type of flash message
type FlashType string

const (
	FlashTypeError   FlashType = "error"
	FlashTypeSuccess FlashType = "success"
	FlashTypeInfo    FlashType = "info"
)

// FlashMessage represents a flash message
type FlashMessage struct {
	Type    FlashType `json:"type"`
	Message string    `json:"message"`
}

// SetFlash sets a flash message in a cookie
func SetFlash(c echo.Context, flashType FlashType, message string) {
	// Simple format: type:message
	value := string(flashType) + ":" + message

	cookie := &http.Cookie{
		Name:     flashCookieName,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   300, // 5 minutes
		Path:     "/",
	}
	c.SetCookie(cookie)
}

// GetFlash retrieves and clears a flash message from cookies
func GetFlash(c echo.Context) (*FlashMessage, error) {
	cookie, err := c.Cookie(flashCookieName)
	if err != nil && err == http.ErrNoCookie {
		return nil, nil // No flash message
	}
	if err != nil {
		return nil, err
	}

	// Clear the flash cookie immediately
	clearCookie := &http.Cookie{
		Name:     flashCookieName,
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
		Path:     "/",
	}
	c.SetCookie(clearCookie)

	// Parse the flash message
	value := cookie.Value
	if value == "" {
		return nil, nil
	}

	// Simple parsing: find first colon
	colonIndex := -1
	for i, char := range value {
		if char == ':' {
			colonIndex = i
			break
		}
	}

	if colonIndex == -1 {
		return nil, nil // Invalid format
	}

	flashType := FlashType(value[:colonIndex])
	message := value[colonIndex+1:]

	return &FlashMessage{
		Type:    flashType,
		Message: message,
	}, nil
}
