package echoutil

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type okRenderer struct{}

func (m *okRenderer) Render(w io.Writer) error {
	_, err := w.Write([]byte("render ok"))
	return err
}

type nokRenderer struct{}

func (m *nokRenderer) Render(w io.Writer) error {
	return errors.New("render nok")
}

func TestRenderGomponent(t *testing.T) {
	// Setup renderer mocks
	okRendererIns := new(okRenderer)
	nokRendererIns := new(nokRenderer)

	t.Run("Successful Render", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := RenderGomponent(c, http.StatusOK, okRendererIns)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, c.Response().Status)
		assert.Equal(t, "render ok", rec.Body.String())
	})

	t.Run("Unsuccessful Render", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := RenderGomponent(c, http.StatusOK, nokRendererIns)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "render nok", rec.Body.String())
	})
}
