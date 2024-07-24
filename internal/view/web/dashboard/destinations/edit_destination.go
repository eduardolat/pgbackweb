package destinations

import (
	"database/sql"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) editDestinationHandler(c echo.Context) error {
	ctx := c.Request().Context()

	destinationID, err := uuid.Parse(c.Param("destinationID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	var formData createDestinationDTO
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	_, err = h.servs.DestinationsService.UpdateDestination(
		ctx, dbgen.DestinationsServiceUpdateDestinationParams{
			ID:         destinationID,
			Name:       sql.NullString{String: formData.Name, Valid: true},
			BucketName: sql.NullString{String: formData.BucketName, Valid: true},
			Region:     sql.NullString{String: formData.Region, Valid: true},
			Endpoint:   sql.NullString{String: formData.Endpoint, Valid: true},
			AccessKey:  sql.NullString{String: formData.AccessKey, Valid: true},
			SecretKey:  sql.NullString{String: formData.SecretKey, Valid: true},
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondToastSuccess(c, "Destination updated")
}

func editDestinationButton(
	destination dbgen.DestinationsServicePaginateDestinationsRow,
) gomponents.Node {
	idPref := "edit-destination-" + destination.ID.String()
	formID := idPref + "-form"
	btnClass := idPref + "-btn"
	loadingID := idPref + "-loading"

	htmxAttributes := func(url string) gomponents.Node {
		return gomponents.Group([]gomponents.Node{
			htmx.HxPost(url),
			htmx.HxInclude("#" + formID),
			htmx.HxDisabledELT("." + btnClass),
			htmx.HxIndicator("#" + loadingID),
			htmx.HxValidate("true"),
		})
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Edit destination",
		Content: []gomponents.Node{
			html.Form(
				html.ID(formID),
				html.Class("space-y-2"),

				component.InputControl(component.InputControlParams{
					Name:        "name",
					Label:       "Name",
					Placeholder: "My destination",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "A name to easily identify the destination",
					Children: []gomponents.Node{
						html.Value(destination.Name),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "bucket_name",
					Label:       "Bucket name",
					Placeholder: "my-bucket",
					Required:    true,
					Type:        component.InputTypeText,
					Children: []gomponents.Node{
						html.Value(destination.BucketName),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "endpoint",
					Label:       "Endpoint",
					Placeholder: "s3-us-west-1.amazonaws.com",
					Required:    true,
					Type:        component.InputTypeText,
					Children: []gomponents.Node{
						html.Value(destination.Endpoint),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "region",
					Label:       "Region",
					Placeholder: "us-west-1",
					Required:    true,
					Type:        component.InputTypeText,
					Children: []gomponents.Node{
						html.Value(destination.Region),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "access_key",
					Label:       "Access key",
					Placeholder: "Access key",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "It will be stored securely using PGP encryption.",
					Children: []gomponents.Node{
						html.Value(destination.DecryptedAccessKey),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "secret_key",
					Label:       "Secret key",
					Placeholder: "Secret key",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "It will be stored securely using PGP encryption.",
					Children: []gomponents.Node{
						html.Value(destination.DecryptedSecretKey),
					},
				}),
			),

			html.Div(
				html.Class("flex justify-between items-center pt-4"),
				html.Div(
					html.Button(
						htmxAttributes("/dashboard/destinations/test"),
						components.Classes{
							btnClass:                      true,
							"btn btn-neutral btn-outline": true,
						},
						html.Type("button"),
						component.SpanText("Test connection"),
						lucide.PlugZap(),
					),
				),
				html.Div(
					html.Class("flex justify-end items-center space-x-2"),
					component.HxLoadingMd(loadingID),
					html.Button(
						htmxAttributes("/dashboard/destinations/"+destination.ID.String()+"/edit"),
						components.Classes{
							btnClass:          true,
							"btn btn-primary": true,
						},
						html.Type("button"),
						component.SpanText("Save"),
						lucide.Save(),
					),
				),
			),
		},
	})

	button := html.Button(
		mo.OpenerAttr,
		html.Class("btn btn-neutral btn-sm btn-square btn-ghost"),
		lucide.Pencil(),
	)

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		html.Div(
			html.Class("inline-block tooltip tooltip-right"),
			html.Data("tip", "Edit destination"),
			button,
		),
	)
}
