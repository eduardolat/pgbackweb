package webhooks

import (
	"fmt"
	"slices"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
	"github.com/eduardolat/pgbackweb/internal/util/maputil"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func createAndUpdateWebhookForm(
	databases []dbgen.DatabasesServiceGetAllDatabasesRow,
	destinations []dbgen.DestinationsServiceGetAllDestinationsRow,
	backups []dbgen.Backup,
	webhook ...dbgen.Webhook,
) gomponents.Node {
	shouldPrefill, pickedWebhook := false, dbgen.Webhook{}
	if len(webhook) > 0 {
		shouldPrefill = true
		pickedWebhook = webhook[0]
	}

	eventTypeOptions := gomponents.Group([]gomponents.Node{
		html.Option(
			gomponents.Text("Select an event type"),
			html.Disabled(),
			gomponents.If(
				!shouldPrefill,
				html.Selected(),
			),
		),

		component.GMap(
			maputil.GetSortedStringKeys(webhooks.FullEventTypes),
			func(key string) gomponents.Node {
				val := webhooks.FullEventTypes[key]
				return html.Option(
					html.Value(key),
					gomponents.Text(val),
					gomponents.If(
						shouldPrefill && key == pickedWebhook.EventType,
						html.Selected(),
					),
				)
			},
		),
	})

	databaseSelect := component.SelectControl(component.SelectControlParams{
		Name:     "target_ids",
		Label:    "Database targets",
		Required: true,
		Children: []gomponents.Node{
			alpine.XModel("targetIds"),
			html.Multiple(),
			component.GMap(
				databases,
				func(db dbgen.DatabasesServiceGetAllDatabasesRow) gomponents.Node {
					return html.Option(
						html.Value(db.ID.String()),
						gomponents.Text(db.Name),
						gomponents.If(
							shouldPrefill && slices.Contains(pickedWebhook.TargetIds, db.ID),
							html.Selected(),
						),
					)
				},
			),
		},
	})

	destinationSelect := component.SelectControl(component.SelectControlParams{
		Name:     "target_ids",
		Label:    "Destination targets",
		Required: true,
		Children: []gomponents.Node{
			alpine.XModel("targetIds"),
			html.Multiple(),
			component.GMap(
				destinations,
				func(dest dbgen.DestinationsServiceGetAllDestinationsRow) gomponents.Node {
					return html.Option(
						html.Value(dest.ID.String()),
						gomponents.Text(dest.Name),
						gomponents.If(
							shouldPrefill && slices.Contains(pickedWebhook.TargetIds, dest.ID),
							html.Selected(),
						),
					)
				},
			),
		},
	})

	backupSelect := component.SelectControl(component.SelectControlParams{
		Name:     "target_ids",
		Label:    "Backup targets",
		Required: true,
		Children: []gomponents.Node{
			alpine.XModel("targetIds"),
			html.Multiple(),
			component.GMap(
				backups,
				func(backup dbgen.Backup) gomponents.Node {
					return html.Option(
						html.Value(backup.ID.String()),
						gomponents.Text(backup.Name),
						gomponents.If(
							shouldPrefill && slices.Contains(pickedWebhook.TargetIds, backup.ID),
							html.Selected(),
						),
					)
				},
			),
		},
	})

	eventTypeSelects := map[string]gomponents.Node{
		webhooks.EventTypeDatabaseHealthy.Value.Key:      databaseSelect,
		webhooks.EventTypeDatabaseUnhealthy.Value.Key:    databaseSelect,
		webhooks.EventTypeDestinationHealthy.Value.Key:   destinationSelect,
		webhooks.EventTypeDestinationUnhealthy.Value.Key: destinationSelect,
		webhooks.EventTypeExecutionSuccess.Value.Key:     backupSelect,
		webhooks.EventTypeExecutionFailed.Value.Key:      backupSelect,
	}

	targetIdsSelect := []gomponents.Node{}
	for eventType, selectNode := range eventTypeSelects {
		targetIdsSelect = append(targetIdsSelect, alpine.Template(
			alpine.XIf(fmt.Sprintf("isEventType('%s')", eventType)),
			selectNode,
		))
	}

	return html.Div(
		html.Class("space-y-2"),

		alpine.XData(`{
			eventType: "",
			targetIds: [],

			isEventType(eventType) {
				return this.eventType === eventType
			},

			init () {
				$watch('eventType', (value, oldValue) => {
					if (value !== oldValue) {
						this.targetIds = []
					}
				})
			}
		}`),

		component.InputControl(component.InputControlParams{
			Name:        "name",
			Label:       "Name",
			Placeholder: "My webhook",
			Required:    true,
			Type:        component.InputTypeText,
			Children: []gomponents.Node{
				gomponents.If(shouldPrefill, html.Value(pickedWebhook.Name)),
			},
		}),

		component.SelectControl(component.SelectControlParams{
			Name:     "event_type",
			Label:    "Event type",
			Required: true,
			HelpButtonChildren: []gomponents.Node{
				component.H3Text("Event types"),
				component.PText(`
					These are the event types that can trigger a webhook.
				`),

				html.Div(
					html.Class("space-y-2"),

					component.CardBoxSimple(
						component.H4Text("Database healthy"),
						component.PText(`
							This event will be triggered when a database changes it's
							health status from unhealthy to healthy.
						`),
					),

					component.CardBoxSimple(
						component.H4Text("Database unhealthy"),
						component.PText(`
							This event will be triggered when a database changes it's
							health status from healthy to unhealthy.
						`),
					),

					component.CardBoxSimple(
						component.H4Text("Destination healthy"),
						component.PText(`
							This event will be triggered when a destination changes it's
							health status from unhealthy to healthy.
						`),
					),

					component.CardBoxSimple(
						component.H4Text("Destination unhealthy"),
						component.PText(`
							This event will be triggered when a destination changes it's
							health status from healthy to unhealthy.
						`),
					),

					component.CardBoxSimple(
						component.H4Text("Execution success"),
						component.PText(`
							This event will be triggered when a backup execution is
							successful.
						`),
					),

					component.CardBoxSimple(
						component.H4Text("Execution failed"),
						component.PText(`
							This event will be triggered when a backup execution fails.
						`),
					),
				),
			},
			Children: []gomponents.Node{
				alpine.XModel("eventType"),
				eventTypeOptions,
			},
		}),

		html.Div(targetIdsSelect...),

		component.SelectControl(component.SelectControlParams{
			Name:     "is_active",
			Label:    "Activate webhook",
			Required: true,
			Children: []gomponents.Node{
				html.Option(
					html.Value("true"), gomponents.Text("Yes"),
					gomponents.If(!shouldPrefill, html.Selected()),
					gomponents.If(shouldPrefill && pickedWebhook.IsActive, html.Selected()),
				),
				html.Option(
					html.Value("false"), gomponents.Text("No"),
					gomponents.If(shouldPrefill && !pickedWebhook.IsActive, html.Selected()),
				),
			},
		}),

		component.InputControl(component.InputControlParams{
			Name:        "url",
			Label:       "URL",
			Placeholder: "https://example.com/webhook",
			Required:    true,
			Type:        component.InputTypeUrl,
			Children: []gomponents.Node{
				gomponents.If(shouldPrefill, html.Value(pickedWebhook.Url)),
			},
		}),

		component.SelectControl(component.SelectControlParams{
			Name:     "method",
			Label:    "Method",
			Required: true,
			Children: []gomponents.Node{
				html.Option(
					html.Value("POST"),
					gomponents.Text("POST"),
					gomponents.If(!shouldPrefill, html.Selected()),
					gomponents.If(
						shouldPrefill && pickedWebhook.Method == "POST",
						html.Selected(),
					),
				),
				html.Option(
					html.Value("GET"),
					gomponents.Text("GET"),
					gomponents.If(
						shouldPrefill && pickedWebhook.Method == "GET",
						html.Selected(),
					),
				),
			},
		}),

		component.TextareaControl(component.TextareaControlParams{
			Name:        "headers",
			Label:       "Headers",
			Placeholder: `{ "Authorization": "Bearer my-token" }`,
			HelpText:    `By default it will send a { "Content-Type": "application/json" } header.`,
			Children: []gomponents.Node{
				gomponents.If(
					shouldPrefill, gomponents.Text(pickedWebhook.Headers.String),
				),
			},
		}),

		component.TextareaControl(component.TextareaControlParams{
			Name:        "body",
			Label:       "Body",
			Placeholder: `{ "key": "value" }`,
			HelpText:    `By default it will send an empty json object {}.`,
			Children: []gomponents.Node{
				gomponents.If(
					shouldPrefill, gomponents.Text(pickedWebhook.Body.String),
				),
			},
		}),
	)
}
