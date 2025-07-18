package profile

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) updateUserHandler(c echo.Context) error {
	reqCtx := reqctx.GetCtx(c)
	ctx := c.Request().Context()

	// Check if user is OIDC user
	isOIDCUser := reqCtx.User.OidcProvider.Valid && reqCtx.User.OidcSubject.Valid

	// Block profile updates for OIDC users
	if isOIDCUser {
		return respondhtmx.ToastError(c, "Profile updates are not allowed for SSO users. Your profile is managed by your identity provider.")
	}

	var formData struct {
		Name                 string `form:"name" validate:"required"`
		Email                string `form:"email" validate:"required,email"`
		Password             string `form:"password" validate:"omitempty,min=6,max=50"`
		PasswordConfirmation string `form:"password_confirmation" validate:"omitempty,eqfield=Password"`
	}
	if err := c.Bind(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	passwordUpdate := sql.NullString{String: formData.Password, Valid: formData.Password != ""}

	_, err := h.servs.UsersService.UpdateUser(ctx, dbgen.UsersServiceUpdateUserParams{
		ID:       reqCtx.User.ID,
		Name:     sql.NullString{String: formData.Name, Valid: true},
		Email:    sql.NullString{String: formData.Email, Valid: true},
		Password: passwordUpdate,
	})
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.ToastSuccess(c, "Profile updated")
}

func updateUserForm(user dbgen.User) nodx.Node {
	// Check if user is OIDC user
	isOIDCUser := user.OidcProvider.Valid && user.OidcSubject.Valid

	// Build form fields
	formFields := []nodx.Node{
		component.H2Text("Update profile"),

		// Show different message for OIDC users
		nodx.If(isOIDCUser,
			nodx.Div(
				nodx.Class("alert alert-info mb-4"),
				nodx.Div(
					nodx.Class("flex items-center space-x-2"),
					lucide.Info(),
					nodx.Div(
						nodx.Class("text-sm"),
						nodx.Text("You are logged in via SSO. Your profile information is managed by your identity provider and cannot be changed here."),
					),
				),
			),
		),

		component.InputControl(component.InputControlParams{
			Name:         "name",
			Label:        "Full name",
			Placeholder:  "Your full name",
			Required:     !isOIDCUser, // Don't require if disabled
			Type:         component.InputTypeText,
			AutoComplete: "name",
			Children: []nodx.Node{
				nodx.Value(user.Name),
				nodx.If(isOIDCUser, nodx.Disabled("")),
				nodx.If(isOIDCUser, nodx.Readonly("")),
			},
		}),

		component.InputControl(component.InputControlParams{
			Name:         "email",
			Label:        "Email",
			Placeholder:  "Your email",
			Required:     !isOIDCUser, // Don't require if disabled
			AutoComplete: "email",
			Type:         component.InputTypeEmail,
			Children: []nodx.Node{
				nodx.Value(user.Email),
				nodx.If(isOIDCUser, nodx.Disabled("")),
				nodx.If(isOIDCUser, nodx.Readonly("")),
			},
		}),
	}

	// Add password fields only for non-OIDC users
	if !isOIDCUser {
		formFields = append(formFields,
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
		)
	}

	// Add submit button (disabled for OIDC users)
	formFields = append(formFields,
		nodx.Div(
			nodx.Class("flex justify-end items-center space-x-2 pt-2"),
			component.HxLoadingMd(),
			nodx.Button(
				nodx.ClassMap{
					"btn btn-primary":  !isOIDCUser,
					"btn btn-disabled": isOIDCUser,
				},
				nodx.Type("submit"),
				nodx.If(isOIDCUser, nodx.Disabled("")),
				component.SpanText("Save changes"),
				lucide.Save(),
			),
		),
	)

	return component.CardBox(component.CardBoxParams{
		Children: []nodx.Node{
			nodx.FormEl(
				append([]nodx.Node{
					nodx.If(!isOIDCUser, htmx.HxPost("/dashboard/profile")),
					nodx.If(!isOIDCUser, htmx.HxDisabledELT("find button")),
					nodx.Class("space-y-2"),
				}, formFields...)...,
			),
		},
	})
}
