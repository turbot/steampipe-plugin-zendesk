package zendesk

import (
	"context"

	"github.com/nukosuke/go-zendesk/zendesk"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableZendeskTicketAudit() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_ticket_audit",
		Description: "Audits are a read-only history of all updates to a ticket. When a ticket is updated in Zendesk Support, an audit is stored. Each audit represents a single update to the ticket. An update can consist of one or more events.",
		List: &plugin.ListConfig{
			Hydrate: listTicketAudit,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "ticket_id"}),
			Hydrate:    getTicketAudit,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the ticket update."},
			{Name: "ticket_id", Type: proto.ColumnType_INT, Description: "The ID of the associated ticket."},
			// Other columns
			{Name: "author_id", Type: proto.ColumnType_INT, Description: "The user who created the audit."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time the audit was created."},
			{Name: "events", Type: proto.ColumnType_JSON, Description: "An array of the events that happened in this audit."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Metadata for the audit, custom and system data."},
			{Name: "via_channel", Type: proto.ColumnType_STRING, Transform: transform.FromField("Via.Channel"), Description: "How the ticket or event was created. Examples: \"web\", \"mobile\", \"rule\", \"system\"."},
			{Name: "via_followup_source_id", Type: proto.ColumnType_STRING, Description: "The id of a closed ticket when creating a follow-up ticket."},
			{Name: "via_source_from", Type: proto.ColumnType_JSON, Transform: transform.FromField("Via.Source.From"), Description: "Source the ticket was sent to."},
			{Name: "via_source_ref", Type: proto.ColumnType_STRING, Transform: transform.FromField("Via.Source.Ref"), Description: "Medium used to raise the ticket."},
			{Name: "via_source_to", Type: proto.ColumnType_JSON, Transform: transform.FromField("Via.Source.From"), Description: "Target that received the ticket."},
		},
	}
}

func listTicketAudit(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	opts := zendesk.CursorOption{}
	for {
		ticketAudits, cursor, err := conn.GetAllTicketAudits(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, t := range ticketAudits {
			d.StreamListItem(ctx, t)
		}
		opts.Cursor = cursor.AfterCursor
		if cursor.AfterCursor == "" {
			break
		}
	}
	return nil, nil
}

func getTicketAudit(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	ticketID := quals["ticket_id"].GetInt64Value()
	result, err := conn.GetTicketAudit(ctx, id, ticketID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
