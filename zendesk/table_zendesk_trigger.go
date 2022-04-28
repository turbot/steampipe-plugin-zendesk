package zendesk

import (
	"context"

	"github.com/nukosuke/go-zendesk/zendesk"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableZendeskTrigger() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_trigger",
		Description: "A trigger consists of one or more actions performed when a ticket is created or updated. The actions are performed only if certain conditions are met. For example, a trigger can notify the customer when an agent changes the status of a ticket to Solved.",
		List: &plugin.ListConfig{
			Hydrate: listTrigger,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTrigger,
		},
		Columns: []*plugin.Column{

			// NOTE:
			// The actions and conditions_* fields are arrays of {field,value} objects.
			// It's tempting to convert them into objects / maps, but that doesn't work
			// because they often have multiple items with the same field.

			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Automatically assigned when the trigger is created"},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "The title of the trigger"},
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "Whether the trigger is active"},
			// Other columns
			{Name: "actions", Type: proto.ColumnType_JSON, Description: "An array of actions describing what the trigger will do."},
			{Name: "conditions_all", Type: proto.ColumnType_JSON, Transform: transform.FromField("Conditions.All"), Description: "Trigger if all conditions are met."},
			{Name: "conditions_any", Type: proto.ColumnType_JSON, Transform: transform.FromField("Conditions.Any"), Description: "Trigger if any condition is met."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time the trigger was created"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the trigger"},
			{Name: "position", Type: proto.ColumnType_INT, Description: "Position of the trigger, determines the order they will execute in."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time of the last update of the trigger"},
		},
	}
}

func listTrigger(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	opts := &zendesk.TriggerListOptions{
		PageOptions: zendesk.PageOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	for true {
		triggers, page, err := conn.GetTriggers(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, t := range triggers {
			d.StreamListItem(ctx, t)
		}
		if !page.HasNext() {
			break
		}
		opts.Page++
	}
	return nil, nil
}

func getTrigger(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	result, err := conn.GetTrigger(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
