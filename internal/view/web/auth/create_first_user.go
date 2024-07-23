package auth

import (
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
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
			htmx.HxPost("/auth/create-first-user"),
			htmx.HxDisabledELT("find button"),
			html.Class("mt-4 space-y-2"),

			html.Div(
				component.InputControl(component.InputControlParams{
					Name:         "name",
					Label:        "Full name",
					Placeholder:  "John Doe",
					Required:     true,
					Type:         component.InputTypeText,
					AutoComplete: "name",
				}),

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
					AutoComplete: "new-password",
					Children: []gomponents.Node{
						html.MinLength("6"),
						html.MaxLength("50"),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "password_confirmation",
					Label:       "Confirm password",
					Placeholder: "******",
					Required:    true,
					Type:        component.InputTypePassword,
					Children: []gomponents.Node{
						html.MinLength("6"),
						html.MaxLength("50"),
					},
				}),

				html.Div(
					html.Class("pt-2 flex justify-end items-center space-x-2"),
					component.HxLoadingMd(),
					html.Button(
						html.ID("create-first-user-button"),
						html.Class("btn btn-primary"),
						html.Type("submit"),
						component.SpanText("Create user and continue"),
						lucide.UserPlus(),
					),
				),
			),
		),
	}

	return layout.Auth(layout.AuthParams{
		Title: "Create first user",
		Body:  content,
	})
}

func (h *handlers) createFirstUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData struct {
		Name                 string `form:"name" validate:"required"`
		Email                string `form:"email" validate:"required,email"`
		Password             string `form:"password" validate:"required,min=6,max=50"`
		PasswordConfirmation string `form:"password_confirmation" validate:"required,eqfield=Password"`
	}
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	_, err := h.servs.UsersService.CreateUser(ctx, dbgen.UsersServiceCreateUserParams{
		Name:     formData.Name,
		Email:    formData.Email,
		Password: formData.Password,
	})
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondAlertWithRedirect(
		c, "User created successfully", "/auth/login",
	)
}
