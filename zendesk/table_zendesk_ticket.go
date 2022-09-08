package zendesk

import (
	"context"

	"github.com/nukosuke/go-zendesk/zendesk"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableZendeskTicket() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_ticket",
		Description: "Tickets are the means through which your end users (customers) communicate with agents in Zendesk Support. Tickets can originate from a number of channels, including email, Help Center, chat, phone call, Twitter, Facebook, or the API. All tickets have a core set of properties.",
		List: &plugin.ListConfig{
			Hydrate: listTicket,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTicket,
		},
		Columns: []*plugin.Column{
			{Name: "allow_attachments", Type: proto.ColumnType_BOOL, Description: "Permission for agents to add add attachments to a comment. Defaults to true"},
			{Name: "allow_channelback", Type: proto.ColumnType_BOOL, Description: "Is false if channelback is disabled, true otherwise. Only applicable for channels framework ticket"},
			{Name: "assignee_id", Type: proto.ColumnType_INT, Description: "The agent currently assigned to the ticket"},
			{Name: "brand_id", Type: proto.ColumnType_INT, Description: "Enterprise only. The id of the brand this ticket is associated with"},
			{Name: "collaborator_ids", Type: proto.ColumnType_JSON, Description: "The ids of users currently CC'ed on the ticket"},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When this record was created"},
			{Name: "custom_fields", Type: proto.ColumnType_JSON, Description: "Custom fields for the ticket."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Read-only first comment on the ticket. When creating a ticket, use comment to set the description."},
			{Name: "due_at", Type: proto.ColumnType_TIMESTAMP, Description: "If this is a ticket of type \"task\" it has a due date. Due date format uses ISO 8601 format."},
			{Name: "email_cc_ids", Type: proto.ColumnType_JSON, Description: "The ids of agents or end users currently CC'ed on the ticket."},
			{Name: "external_id", Type: proto.ColumnType_INT, Description: "An id you can use to link Zendesk Support tickets to local records"},
			{Name: "follower_ids", Type: proto.ColumnType_JSON, Description: "The ids of agents currently following the ticket."},
			{Name: "followup_ids", Type: proto.ColumnType_JSON, Description: "The ids of the followups created from this ticket. Ids are only visible once the ticket is closed"},
			{Name: "forum_topic_id", Type: proto.ColumnType_INT, Description: "The topic in the Zendesk Web portal this ticket originated from, if any. The Web portal is deprecated"},
			{Name: "group_id", Type: proto.ColumnType_INT, Description: "The group this ticket is assigned to"},
			{Name: "has_incidents", Type: proto.ColumnType_BOOL, Description: "Is true if a ticket is a problem type and has one or more incidents linked to it. Otherwise, the value is false."},
			{Name: "id", Type: proto.ColumnType_INT, Description: "Automatically assigned when the ticket is created"},
			{Name: "is_public", Type: proto.ColumnType_BOOL, Description: "Is true if any comments are public, false otherwise"},
			{Name: "macro_ids", Type: proto.ColumnType_JSON, Description: "POST requests only. List of macro IDs to be recorded in the ticket audit"},
			{Name: "organization_id", Type: proto.ColumnType_INT, Description: "The organization of the requester. You can only specify the ID of an organization associated with the requester."},
			{Name: "priority", Type: proto.ColumnType_STRING, Description: "The urgency with which the ticket should be addressed. Allowed values are \"urgent\", \"high\", \"normal\", or \"low\"."},
			{Name: "problem_id", Type: proto.ColumnType_INT, Description: "For tickets of type \"incident\", the ID of the problem the incident is linked to"},
			{Name: "raw_subject", Type: proto.ColumnType_STRING, Description: "The dynamic content placeholder, if present, or the \"subject\" value, if not."},
			{Name: "recipient", Type: proto.ColumnType_STRING, Description: "The original recipient e-mail address of the ticket"},
			{Name: "requester_id", Type: proto.ColumnType_INT, Description: "The user who requested this ticket"},
			{Name: "satisfaction_rating_comment", Type: proto.ColumnType_STRING, Transform: transform.FromField("SatisfactionRating.Comment"), Description: "The comment received with this rating, if available"},
			{Name: "satisfaction_rating_id", Type: proto.ColumnType_INT, Transform: transform.FromField("SatisfactionRating.ID"), Description: "Unique identifier for the satisfaction rating on this ticket"},
			{Name: "satisfaction_rating_score", Type: proto.ColumnType_STRING, Transform: transform.FromField("SatisfactionRating.Score"), Description: "The rating \"offered\", \"unoffered\", \"good\" or \"bad\""},
			{Name: "sharing_agreement_ids", Type: proto.ColumnType_JSON, Description: "The ids of the sharing agreements used for this ticket"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The state of the ticket. Allowed values are \"new\", \"open\", \"pending\", \"hold\", \"solved\", or \"closed\"."},
			{Name: "subject", Type: proto.ColumnType_STRING, Description: "The value of the subject field for this ticket"},
			{Name: "submitter_id", Type: proto.ColumnType_INT, Description: "The user who submitted the ticket. The submitter always becomes the author of the first comment on the ticket"},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "The array of tags applied to this ticket"},
			{Name: "ticket_form_id", Type: proto.ColumnType_INT, Description: "Enterprise only. The id of the ticket form to render for the ticket"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of this ticket. Allowed values are \"problem\", \"incident\", \"question\", or \"task\"."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When this record last got updated"},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The API url of this ticket"},
			{Name: "via_channel", Type: proto.ColumnType_STRING, Transform: transform.FromField("Via.Channel"), Description: "How the ticket or event was created. Examples: \"web\", \"mobile\", \"rule\", \"system\""},
			{Name: "via_followup_source_id", Type: proto.ColumnType_STRING, Description: "The id of a closed ticket when creating a follow-up ticket."},
			{Name: "via_source_from", Type: proto.ColumnType_JSON, Transform: transform.FromField("Via.Source.From"), Description: "Source the ticket was sent to"},
			{Name: "via_source_ref", Type: proto.ColumnType_STRING, Transform: transform.FromField("Via.Source.Ref"), Description: "Medium used to raise the ticket"},
			{Name: "via_source_to", Type: proto.ColumnType_JSON, Transform: transform.FromField("Via.Source.From"), Description: "Target that received the ticket"},
		},
	}
}

func listTicket(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	opts := &zendesk.TicketListOptions{
		SortBy:    "created_at",
		SortOrder: "desc",
		PageOptions: zendesk.PageOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	for {
		tickets, page, err := conn.GetTickets(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, t := range tickets {
			d.StreamListItem(ctx, t)
		}
		if !page.HasNext() {
			break
		}
		opts.Page++
	}
	return nil, nil
}

func getTicket(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	result, err := conn.GetTicket(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
