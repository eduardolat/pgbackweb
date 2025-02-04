package destinations

import (
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) editDestinationHandler(c echo.Context) error {
	ctx := c.Request().Context()

	destinationID, err := uuid.Parse(c.Param("destinationID"))
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	var formData createDestinationDTO
	if err := c.Bind(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
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
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.AlertWithRefresh(c, "Destination updated")
}

func editDestinationButton(
	destination dbgen.DestinationsServicePaginateDestinationsRow,
) nodx.Node {
	idPref := "edit-destination-" + destination.ID.String()
	formID := idPref + "-form"
	btnClass := idPref + "-btn"
	loadingID := idPref + "-loading"

	htmxAttributes := func(url string) nodx.Node {
		return nodx.Group(
			htmx.HxPost(url),
			htmx.HxInclude("#"+formID),
			htmx.HxDisabledELT("."+btnClass),
			htmx.HxIndicator("#"+loadingID),
			htmx.HxValidate("true"),
		)
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Edit destination",
		Content: []nodx.Node{
			nodx.FormEl(
				nodx.Id(formID),
				nodx.Class("space-y-2"),

				component.InputControl(component.InputControlParams{
					Name:        "name",
					Label:       "Name",
					Placeholder: "My destination",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "A name to easily identify the destination",
					Children: []nodx.Node{
						nodx.Value(destination.Name),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "bucket_name",
					Label:       "Bucket name",
					Placeholder: "my-bucket",
					Required:    true,
					Type:        component.InputTypeText,
					Children: []nodx.Node{
						nodx.Value(destination.BucketName),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "endpoint",
					Label:       "Endpoint",
					Placeholder: "s3-us-west-1.amazonaws.com",
					Required:    true,
					Type:        component.InputTypeText,
					Children: []nodx.Node{
						nodx.Value(destination.Endpoint),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "region",
					Label:       "Region",
					Placeholder: "us-west-1",
					Required:    true,
					Type:        component.InputTypeText,
					Children: []nodx.Node{
						nodx.Value(destination.Region),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "access_key",
					Label:       "Access key",
					Placeholder: "Access key",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "It will be stored securely using PGP encryption.",
					Children: []nodx.Node{
						nodx.Value(destination.DecryptedAccessKey),
					},
				}),

				component.InputControl(component.InputControlParams{
					Name:        "secret_key",
					Label:       "Secret key",
					Placeholder: "Secret key",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "It will be stored securely using PGP encryption.",
					Children: []nodx.Node{
						nodx.Value(destination.DecryptedSecretKey),
					},
				}),
			),

			nodx.Div(
				nodx.Class("flex justify-between items-center pt-4"),
				nodx.Div(
					nodx.Button(
						htmxAttributes("/dashboard/destinations/test"),
						nodx.ClassMap{
							btnClass:                      true,
							"btn btn-neutral btn-outline": true,
						},
						nodx.Type("button"),
						component.SpanText("Test connection"),
						lucide.PlugZap(),
					),
				),
				nodx.Div(
					nodx.Class("flex justify-end items-center space-x-2"),
					component.HxLoadingMd(loadingID),
					nodx.Button(
						htmxAttributes("/dashboard/destinations/"+destination.ID.String()+"/edit"),
						nodx.ClassMap{
							btnClass:          true,
							"btn btn-primary": true,
						},
						nodx.Type("button"),
						component.SpanText("Save"),
						lucide.Save(),
					),
				),
			),
		},
	})

	return nodx.Div(
		mo.HTML,
		component.OptionsDropdownButton(
			mo.OpenerAttr,
			lucide.Pencil(),
			component.SpanText("Edit destination"),
		),
	)
}
