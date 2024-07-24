package profile

import (
	"database/sql"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) updateUserHandler(c echo.Context) error {
	reqCtx := reqctx.GetCtx(c)
	ctx := c.Request().Context()

	var formData struct {
		Name                 string `form:"name" validate:"required"`
		Email                string `form:"email" validate:"required,email"`
		Password             string `form:"password" validate:"omitempty,min=6,max=50"`
		PasswordConfirmation string `form:"password_confirmation" validate:"omitempty,eqfield=Password"`
	}
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	_, err := h.servs.UsersService.UpdateUser(ctx, dbgen.UsersServiceUpdateUserParams{
		ID:       reqCtx.User.ID,
		Name:     sql.NullString{String: formData.Name, Valid: true},
		Email:    sql.NullString{String: formData.Email, Valid: true},
		Password: sql.NullString{String: formData.Password, Valid: formData.Password != ""},
	})
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondToastSuccess(c, "Profile updated")
}

func updateUserForm(user dbgen.User) gomponents.Node {
	return component.CardBox(component.CardBoxParams{
		Children: []gomponents.Node{
			html.Form(
				htmx.HxPost("/dashboard/profile"),
				htmx.HxDisabledELT("find button"),
				html.Class("space-y-2"),

				component.H2Text("Update profile"),

				component.InputControl(component.InputControlParams{
					Name:         "name",
					Label:        "Full name",
					Placeholder:  "Your full name",
					Required:     true,
					Type:         component.InputTypeText,
					AutoComplete: "name",
					Children: []gomponents.Node{
						html.Value(user.Name),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:         "email",
					Label:        "Email",
					Placeholder:  "Your email",
					Required:     true,
					AutoComplete: "email",
					Type:         component.InputTypeEmail,
					Children: []gomponents.Node{
						html.Value(user.Email),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:         "password",
					Label:        "Change password",
					Placeholder:  "New password",
					AutoComplete: "new-password",
					Type:         component.InputTypePassword,
					HelpText:     "Leave empty to keep your current password",
				}),

				component.InputControl(component.InputControlParams{
					Name:         "password_confirmation",
					Label:        "Confirm password",
					Placeholder:  "Confirm new password",
					AutoComplete: "new-password",
					Type:         component.InputTypePassword,
				}),

				html.Div(
					html.Class("flex justify-end items-center space-x-2 pt-2"),
					component.HxLoadingMd(),
					html.Button(
						html.Class("btn btn-primary"),
						html.Type("submit"),
						component.SpanText("Save changes"),
						lucide.Save(),
					),
				),
			),
		},
	})
}
