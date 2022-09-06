package zendesk

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableZendeskGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_group",
		Description: "When support requests arrive in Zendesk Support, they can be assigned to a Group. Groups serve as the core element of ticket workflow; support agents are organized into Groups and tickets can be assigned to a Group only, or to an assigned agent within a Group. A ticket can never be assigned to an agent without also being assigned to a Group.",
		List: &plugin.ListConfig{
			Hydrate: listGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getGroup,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the group"},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "API url of the group"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the group"},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "True if the group has been deleted"},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time the group was created"},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time of the last update of the group"},
		},
	}
}

func listGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	// TODO - The library doesn't support paging for input?
	groups, _, err := conn.GetGroups(ctx)
	if err != nil {
		return nil, err
	}
	for _, t := range groups {
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}

func getGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	result, err := conn.GetGroup(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
