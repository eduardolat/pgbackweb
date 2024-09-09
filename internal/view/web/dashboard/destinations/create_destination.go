package destinations

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type createDestinationDTO struct {
	Name       string `form:"name" validate:"required"`
	BucketName string `form:"bucket_name" validate:"required"`
	AccessKey  string `form:"access_key" validate:"required"`
	SecretKey  string `form:"secret_key" validate:"required"`
	Region     string `form:"region" validate:"required"`
	Endpoint   string `form:"endpoint" validate:"required"`
}

func (h *handlers) createDestinationHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData createDestinationDTO
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	_, err := h.servs.DestinationsService.CreateDestination(
		ctx, dbgen.DestinationsServiceCreateDestinationParams{
			Name:       formData.Name,
			AccessKey:  formData.AccessKey,
			SecretKey:  formData.SecretKey,
			Region:     formData.Region,
			Endpoint:   formData.Endpoint,
			BucketName: formData.BucketName,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRedirect(c, "/dashboard/destinations")
}

func createDestinationButton() gomponents.Node {
	htmxAttributes := func(url string) gomponents.Node {
		return gomponents.Group([]gomponents.Node{
			htmx.HxPost(url),
			htmx.HxInclude("#add-destination-form"),
			htmx.HxDisabledELT(".add-destination-btn"),
			htmx.HxIndicator("#add-destination-loading"),
			htmx.HxValidate("true"),
		})
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Add destination",
		Content: []gomponents.Node{
			html.Form(
				html.ID("add-destination-form"),
				html.Class("space-y-2"),

				component.InputControl(component.InputControlParams{
					Name:        "name",
					Label:       "Name",
					Placeholder: "My destination",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "A name to easily identify the destination",
				}),

				component.InputControl(component.InputControlParams{
					Name:        "bucket_name",
					Label:       "Bucket name",
					Placeholder: "my-bucket",
					Required:    true,
					Type:        component.InputTypeText,
				}),

				component.InputControl(component.InputControlParams{
					Name:        "endpoint",
					Label:       "Endpoint",
					Placeholder: "s3-us-west-1.amazonaws.com",
					Required:    true,
					Type:        component.InputTypeText,
				}),

				component.InputControl(component.InputControlParams{
					Name:        "region",
					Label:       "Region",
					Placeholder: "us-west-1",
					Required:    true,
					Type:        component.InputTypeText,
				}),

				component.InputControl(component.InputControlParams{
					Name:        "access_key",
					Label:       "Access key",
					Placeholder: "Access key",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "It will be stored securely using PGP encryption.",
				}),

				component.InputControl(component.InputControlParams{
					Name:        "secret_key",
					Label:       "Secret key",
					Placeholder: "Secret key",
					Required:    true,
					Type:        component.InputTypeText,
					HelpText:    "It will be stored securely using PGP encryption.",
				}),
			),

			html.Div(
				html.Class("flex justify-between items-center pt-4"),
				html.Div(
					html.Button(
						htmxAttributes("/dashboard/destinations/test"),
						html.Class("add-destination-btn btn btn-neutral btn-outline"),
						html.Type("button"),
						component.SpanText("Test connection"),
						lucide.PlugZap(),
					),
				),
				html.Div(
					html.Class("flex justify-end items-center space-x-2"),
					component.HxLoadingMd("add-destination-loading"),
					html.Button(
						htmxAttributes("/dashboard/destinations"),
						html.Class("add-destination-btn btn btn-primary"),
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
		html.Class("btn btn-primary"),
		component.SpanText("Add destination"),
		lucide.Plus(),
	)

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		button,
	)
}
