package auth

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmxs"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
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

	return echoutil.RenderNodx(c, http.StatusOK, loginPage())
}

func loginPage() nodx.Node {
	content := []nodx.Node{
		component.H1Text("Login"),

		nodx.FormEl(
			htmx.HxPost("/auth/login"),
			htmx.HxDisabledELT("find button"),
			nodx.Class("mt-4 space-y-2"),

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

			nodx.Div(
				nodx.Class("pt-2 flex justify-end items-center space-x-2"),
				component.HxLoadingMd(),
				nodx.Button(
					nodx.Class("btn btn-primary"),
					nodx.Type("submit"),
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
		return htmxs.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmxs.RespondToastError(c, err.Error())
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
		return htmxs.RespondToastError(c, "Login failed")
	}

	h.servs.AuthService.SetSessionCookie(c, session.DecryptedToken)
	return htmxs.RespondRedirect(c, "/dashboard")
}
