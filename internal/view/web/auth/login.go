package auth

import (
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) loginPageHandler(c echo.Context) error {
	ctx := c.Request().Context()

	usersQty, err := h.servs.UsersService.GetUsersQty(ctx)
	if err != nil {
		logger.Error("failed to get users qty", logger.KV{
			"ip":    c.RealIP(),
			"ua":    c.Request().UserAgent(),
			"error": err,
		})
		return c.String(http.StatusInternalServerError, "Internal server error")
	}
	if usersQty == 0 {
		return c.Redirect(http.StatusFound, "/auth/create-first-user")
	}

	return echoutil.RenderGomponent(c, http.StatusOK, loginPage())
}

func loginPage() gomponents.Node {
	content := []gomponents.Node{
		component.H1Text("Login"),

		html.Form(
			htmx.HxPost("/auth/login"),
			htmx.HxDisabledELT("find button"),
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
				html.Class("pt-2 flex justify-end items-center space-x-2"),
				component.HxLoadingMd(),
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

func (h *handlers) loginHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required,max=50"`
	}
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	session, err := h.servs.AuthService.Login(
		ctx, formData.Email, formData.Password, c.RealIP(), c.Request().UserAgent(),
	)
	if err != nil {
		logger.Error("login failed", logger.KV{
			"email": formData.Email,
			"ip":    c.RealIP(),
			"ua":    c.Request().UserAgent(),
			"err":   err,
		})
		return htmx.RespondToastError(c, "Login failed")
	}

	h.servs.AuthService.SetSessionCookie(c, session)
	return htmx.RespondRedirect(c, "/dashboard")
}
