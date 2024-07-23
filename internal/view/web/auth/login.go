package auth

import (
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) loginPageHandler(c echo.Context) error {
	return echoutil.RenderGomponent(c, http.StatusOK, loginPage())
}

func loginPage() gomponents.Node {
	content := []gomponents.Node{
		component.H1Text("Login"),

		html.Form(
			html.Class("mt-4 space-y-2"),

			component.InputControl(component.InputControlParams{
				Name:         "email",
				Label:        "Email",
				Placeholder:  "john@example.com",
				Required:     true,
				Type:         component.InputTypeEmail,
				AutoComplete: "email",
			}),

			component.InputControl(component.InputControlParams{
				Name:         "password",
				Label:        "Password",
				Placeholder:  "******",
				Required:     true,
				Type:         component.InputTypePassword,
				AutoComplete: "current-password",
			}),

			html.Div(
				html.Class("pt-2 grid place-items-end"),
				html.Button(
					html.Class("btn btn-primary"),
					html.Type("submit"),
					component.SpanText("Login"),
					lucide.LogIn(),
				),
			),
		),
	}

	return layout.Auth(layout.AuthParams{
		Title: "Login",
		Body:  content,
	})
}
