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

// RenderGomponent renders a gomponents component to the response of
// the echo context.
func RenderGomponent(c echo.Context, code int, component renderer) error {
	buf := bytes.Buffer{}
	err := component.Render(&buf)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(code, buf.String())
}
