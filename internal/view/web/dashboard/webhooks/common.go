package webhooks

import (
	"fmt"
	"slices"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
	"github.com/eduardolat/pgbackweb/internal/util/maputil"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
)

func createAndUpdateWebhookForm(
	databases []dbgen.DatabasesServiceGetAllDatabasesRow,
	destinations []dbgen.DestinationsServiceGetAllDestinationsRow,
	backups []dbgen.Backup,
	webhook ...dbgen.Webhook,
) nodx.Node {
	shouldPrefill, pickedWebhook := false, dbgen.Webhook{}
	if len(webhook) > 0 {
		shouldPrefill = true
		pickedWebhook = webhook[0]
	}

	eventTypeOptions := nodx.Group(
		nodx.Option(
			nodx.Text("Select an event type"),
			nodx.Disabled(""),
			nodx.If(
				!shouldPrefill,
				nodx.Selected(""),
			),
		),

		nodx.Map(
			maputil.GetSortedStringKeys(webhooks.FullEventTypes),
			func(key string) nodx.Node {
				val := webhooks.FullEventTypes[key]
				return nodx.Option(
					nodx.Value(key),
					nodx.Text(val),
					nodx.If(
						shouldPrefill && key == pickedWebhook.EventType,
						nodx.Selected(""),
					),
				)
			},
		),
	)

	databaseSelect := component.SelectControl(component.SelectControlParams{
		Name:     "target_ids",
		Label:    "Database targets",
		Required: true,
		Children: []nodx.Node{
			alpine.XModel("targetIds"),
			nodx.Multiple(""),
			nodx.Map(
				databases,
				func(db dbgen.DatabasesServiceGetAllDatabasesRow) nodx.Node {
					return nodx.Option(
						nodx.Value(db.ID.String()),
						nodx.Text(db.Name),
						nodx.If(
							shouldPrefill && slices.Contains(pickedWebhook.TargetIds, db.ID),
							nodx.Selected(""),
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
		Children: []nodx.Node{
			alpine.XModel("targetIds"),
			nodx.Multiple(""),
			nodx.Map(
				destinations,
				func(dest dbgen.DestinationsServiceGetAllDestinationsRow) nodx.Node {
					return nodx.Option(
						nodx.Value(dest.ID.String()),
						nodx.Text(dest.Name),
						nodx.If(
							shouldPrefill && slices.Contains(pickedWebhook.TargetIds, dest.ID),
							nodx.Selected(""),
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
		Children: []nodx.Node{
			alpine.XModel("targetIds"),
			nodx.Multiple(""),
			nodx.Map(
				backups,
				func(backup dbgen.Backup) nodx.Node {
					return nodx.Option(
						nodx.Value(backup.ID.String()),
						nodx.Text(backup.Name),
						nodx.If(
							shouldPrefill && slices.Contains(pickedWebhook.TargetIds, backup.ID),
							nodx.Selected(""),
						),
					)
				},
			),
		},
	})

	eventTypeSelects := map[string]nodx.Node{
		webhooks.EventTypeDatabaseHealthy.Value.Key:      databaseSelect,
		webhooks.EventTypeDatabaseUnhealthy.Value.Key:    databaseSelect,
		webhooks.EventTypeDestinationHealthy.Value.Key:   destinationSelect,
		webhooks.EventTypeDestinationUnhealthy.Value.Key: destinationSelect,
		webhooks.EventTypeExecutionSuccess.Value.Key:     backupSelect,
		webhooks.EventTypeExecutionFailed.Value.Key:      backupSelect,
	}

	targetIdsSelect := []nodx.Node{}
	for eventType, selectNode := range eventTypeSelects {
		targetIdsSelect = append(targetIdsSelect, alpine.Template(
			alpine.XIf(fmt.Sprintf("isEventType('%s')", eventType)),
			selectNode,
		))
	}

	pickedTargetIds := ""
	if len(pickedWebhook.TargetIds) > 0 {
		for _, tid := range pickedWebhook.TargetIds {
			pickedTargetIds += fmt.Sprintf("'%s',", tid.String())
		}
	}

	return nodx.Div(
		nodx.Class("space-y-2"),

		alpine.XData(`{
			eventType: "`+pickedWebhook.EventType+`",
			targetIds: [`+pickedTargetIds+`],

			isEventType(eventType) {
				return this.eventType === eventType
			},

			autoGrowHeadersTextarea() {
				textareaAutoGrow($refs.headersTextarea)
			},

			autoGrowBodyTextarea() {
				textareaAutoGrow($refs.bodyTextarea)
			},

			formatHeadersTextarea() {
				const el = $refs.headersTextarea
				el.value = formatJson(el.value)
				this.autoGrowHeadersTextarea()
			},

			formatBodyTextarea() {
				const el = $refs.bodyTextarea
				el.value = formatJson(el.value)
				this.autoGrowBodyTextarea()
			},

			init() {
				$watch('eventType', (value, oldValue) => {
					if (value !== oldValue) {
						this.targetIds = []
					}
				})

				this.formatHeadersTextarea()
				this.formatBodyTextarea()
			}
		}`),

		component.InputControl(component.InputControlParams{
			Name:        "name",
			Label:       "Name",
			Placeholder: "My webhook",
			Required:    true,
			Type:        component.InputTypeText,
			Children: []nodx.Node{
				nodx.If(shouldPrefill, nodx.Value(pickedWebhook.Name)),
			},
		}),

		component.SelectControl(component.SelectControlParams{
			Name:     "event_type",
			Label:    "Event type",
			Required: true,
			HelpButtonChildren: []nodx.Node{
				component.H3Text("Event types"),
				component.PText(`
					These are the event types that can trigger a webhook.
				`),

				nodx.Div(
					nodx.Class("space-y-2"),

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
			Children: []nodx.Node{
				alpine.XModel("eventType"),
				eventTypeOptions,
			},
		}),

		nodx.Div(targetIdsSelect...),

		component.SelectControl(component.SelectControlParams{
			Name:     "is_active",
			Label:    "Activate webhook",
			Required: true,
			Children: []nodx.Node{
				nodx.Option(
					nodx.Value("true"), nodx.Text("Yes"),
					nodx.If(!shouldPrefill, nodx.Selected("")),
					nodx.If(shouldPrefill && pickedWebhook.IsActive, nodx.Selected("")),
				),
				nodx.Option(
					nodx.Value("false"), nodx.Text("No"),
					nodx.If(shouldPrefill && !pickedWebhook.IsActive, nodx.Selected("")),
				),
			},
		}),

		component.InputControl(component.InputControlParams{
			Name:        "url",
			Label:       "URL",
			Placeholder: "https://example.com/webhook",
			Required:    true,
			Type:        component.InputTypeUrl,
			Children: []nodx.Node{
				nodx.If(shouldPrefill, nodx.Value(pickedWebhook.Url)),
			},
		}),

		component.SelectControl(component.SelectControlParams{
			Name:     "method",
			Label:    "Method",
			Required: true,
			Children: []nodx.Node{
				nodx.Option(
					nodx.Value("POST"),
					nodx.Text("POST"),
					nodx.If(!shouldPrefill, nodx.Selected("")),
					nodx.If(
						shouldPrefill && pickedWebhook.Method == "POST",
						nodx.Selected(""),
					),
				),
				nodx.Option(
					nodx.Value("GET"),
					nodx.Text("GET"),
					nodx.If(
						shouldPrefill && pickedWebhook.Method == "GET",
						nodx.Selected(""),
					),
				),
			},
		}),

		component.TextareaControl(component.TextareaControlParams{
			Name:        "headers",
			Label:       "Headers",
			Placeholder: `{ "Authorization": "Bearer my-token" }`,
			HelpText:    `By default it will send a { "Content-Type": "application/json" } header.`,
			Children: []nodx.Node{
				alpine.XRef("headersTextarea"),
				alpine.XOn("click.outside", "formatHeadersTextarea()"),
				alpine.XOn("input", "autoGrowHeadersTextarea()"),
				nodx.If(
					shouldPrefill, nodx.Text(pickedWebhook.Headers.String),
				),
			},
		}),

		component.TextareaControl(component.TextareaControlParams{
			Name:        "body",
			Label:       "Body",
			Placeholder: `{ "key": "value" }`,
			HelpText:    `By default it will send an empty json object {}.`,
			Children: []nodx.Node{
				alpine.XRef("bodyTextarea"),
				alpine.XOn("click.outside", "formatBodyTextarea()"),
				alpine.XOn("input", "autoGrowBodyTextarea()"),
				nodx.If(
					shouldPrefill, nodx.Text(pickedWebhook.Body.String),
				),
			},
		}),
	)
}
