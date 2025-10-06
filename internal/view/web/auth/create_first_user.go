package auth

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) createFirstUserPageHandler(c echo.Context) error {
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
	if usersQty > 0 {
		return c.Redirect(http.StatusFound, pathutil.BuildPath("/auth/login"))
	}

	return echoutil.RenderNodx(c, http.StatusOK, createFirstUserPage())
}

func createFirstUserPage() nodx.Node {
	content := []nodx.Node{
		component.H1Text("Create first user"),

		nodx.FormEl(
			htmx.HxPost(pathutil.BuildPath("/auth/create-first-user")),
			htmx.HxDisabledELT("find button"),
			nodx.Class("mt-4 space-y-2"),

			nodx.Div(
				component.InputControl(component.InputControlParams{
					Name:         "name",
					Label:        "Full name",
					Placeholder:  "John Doe",
					Required:     true,
					Type:         component.InputTypeText,
					AutoComplete: "name",
					Children: []nodx.Node{
						nodx.Autofocus(""),
					},
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
					Children: []nodx.Node{
						nodx.Minlength("6"),
						nodx.Maxlength("50"),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "password_confirmation",
					Label:       "Confirm password",
					Placeholder: "******",
					Required:    true,
					Type:        component.InputTypePassword,
					Children: []nodx.Node{
						nodx.Minlength("6"),
						nodx.Maxlength("50"),
					},
				}),

				nodx.Div(
					nodx.Class("pt-2 flex justify-end items-center space-x-2"),
					component.HxLoadingMd(),
					nodx.Button(
						nodx.Id("create-first-user-button"),
						nodx.Class("btn btn-primary"),
						nodx.Type("submit"),
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
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	_, err := h.servs.UsersService.CreateUser(ctx, dbgen.UsersServiceCreateUserParams{
		Name:     formData.Name,
		Email:    formData.Email,
		Password: formData.Password,
	})
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.AlertWithRedirect(
		c, "User created successfully", pathutil.BuildPath("/auth/login"),
	)
}
