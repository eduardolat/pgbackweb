package echoutil

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type renderer interface {
	Render(w io.Writer) error
}

// RenderNodx renders a NodX component to the response of
// the echo context.
func RenderNodx(c echo.Context, code int, component renderer) error {
	if component == nil {
		return c.NoContent(code)
	}

	buf := bytes.Buffer{}
	err := component.Render(&buf)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(code, buf.String())
}
