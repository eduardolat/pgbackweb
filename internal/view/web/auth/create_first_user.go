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

func (h *handlers) createFirstUserPageHandler(c echo.Context) error {
	return echoutil.RenderGomponent(c, http.StatusOK, createFirstUserPage())
}

func createFirstUserPage() gomponents.Node {
	content := []gomponents.Node{
		component.H1Text("Create first user"),

		html.Form(
			html.Class("mt-4 space-y-2"),

			component.InputControl(component.InputControlParams{
				Name:        "name",
				Label:       "Full name",
				Placeholder: "John Doe",
				Required:    true,
				Type:        component.InputTypeText,
			}),

			component.InputControl(component.InputControlParams{
				Name:        "email",
				Label:       "Email",
				Placeholder: "john@example.com",
				Required:    true,
				Type:        component.InputTypeEmail,
			}),

			component.InputControl(component.InputControlParams{
				Name:        "password",
				Label:       "Password",
				Placeholder: "******",
				Required:    true,
				Type:        component.InputTypePassword,
			}),

			component.InputControl(component.InputControlParams{
				Name:        "password_confirmation",
				Label:       "Confirm password",
				Placeholder: "******",
				Required:    true,
				Type:        component.InputTypePassword,
			}),

			html.Div(
				html.Class("pt-2 grid place-items-end"),
				html.Button(
					html.Class("btn btn-primary"),
					html.Type("submit"),
					component.SpanText("Create user and continue"),
					lucide.UserPlus(),
				),
			),
		),
	}

	return layout.Auth(layout.AuthParams{
		Title: "Create first user",
		Body:  content,
	})
}
