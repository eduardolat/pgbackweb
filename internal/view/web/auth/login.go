package auth

import (
	"net/http"
	"net/url"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
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

	// Check for error message in URL parameters
	errorMsg := c.QueryParam("error")
	if errorMsg != "" {
		// URL decode the error message to handle encoded characters
		if decodedMsg, err := url.QueryUnescape(errorMsg); err == nil {
			errorMsg = decodedMsg
		}
	}

	return echoutil.RenderNodx(c, http.StatusOK, loginPage(h.servs.OIDCService.IsEnabled(), errorMsg))
}

// loginPage constructs the login page UI as a NodX node tree, optionally displaying an error message and an OIDC login option.
// 
// If an error message is provided, a toast notification is triggered on page load. If OIDC login is enabled, a button for SSO login and a divider are included before the traditional email/password login form.
// 
// Returns the complete login page node wrapped in the authentication layout.
func loginPage(oidcEnabled bool, errorMsg string) nodx.Node {
	content := []nodx.Node{
		component.H1Text("Login"),
	}

	// Add JavaScript to show toast notification if error message is present
	if errorMsg != "" {
		// Use a data attribute to safely pass the error message to JavaScript
		content = append(content,
			nodx.Script(
				nodx.Attr("data-error-message", errorMsg),
				nodx.Text(`
					(function() {
						const errorMsg = document.currentScript.dataset.errorMessage;
						if (errorMsg) {
							window.toaster.error(errorMsg);
						}
					})();
				`),
			),
		)
	}

	// Add OIDC login option if enabled
	if oidcEnabled {
		content = append(content,
			nodx.Div(
				nodx.Class("mt-4"),
				nodx.A(
					nodx.Href("/auth/oidc/login"),
					nodx.Class("btn btn-outline btn-block"),
					component.SpanText("Login with SSO"),
					lucide.ExternalLink(),
				),
			),
			nodx.Div(
				nodx.Class("divider"),
				nodx.Text("OR"),
			),
		)
	}

	// Traditional login form
	content = append(content,
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
	)

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
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
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
		return respondhtmx.ToastError(c, "Login failed")
	}

	h.servs.AuthService.SetSessionCookie(c, session.DecryptedToken)
	return respondhtmx.Redirect(c, "/dashboard")
}
