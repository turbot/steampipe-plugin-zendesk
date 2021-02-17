package zendesk

import (
	"context"

	"github.com/nukosuke/go-zendesk/zendesk"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableZendeskOrganization() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_organization",
		Description: "Just as agents can be segmented into groups in Zendesk Support, your customers (end-users) can be segmented into organizations. You can manually assign customers to an organization or automatically assign them to an organization by their email address domain. Organizations can be used in business rules to route tickets to groups of agents or to send email notifications.",
		List: &plugin.ListConfig{
			Hydrate: listOrganization,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getOrganization,
		},
		Columns: []*plugin.Column{
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time the organization was created"},
			//{Name: "details", Type: proto.ColumnType_STRING, Description: "Any details obout the organization, such as the address"},
			{Name: "domain_names", Type: proto.ColumnType_JSON, Description: "An array of domain names associated with this organization"},
			//{Name: "external_id", Type: proto.ColumnType_STRING, Description: "A unique external id to associate organizations to an external record"},
			{Name: "group_id", Type: proto.ColumnType_INT, Description: "New tickets from users in this organization are automatically put in this group"},
			{Name: "id", Type: proto.ColumnType_INT, Description: "Automatically assigned when the organization is created"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "A unique name for the organization"},
			//{Name: "notes", Type: proto.ColumnType_STRING, Description: "Any notes you have about the organization"},
			{Name: "organization_fields", Type: proto.ColumnType_JSON, Description: "Custom fields for this organization"},
			{Name: "shared_comments", Type: proto.ColumnType_BOOL, Description: "End users in this organization are able to see each other's comments on tickets"},
			{Name: "shared_tickets", Type: proto.ColumnType_BOOL, Description: "End users in this organization are able to see each other's tickets"},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "The tags of the organization"},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time of the last update of the organization"},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The API url of this organization"},
		},
	}
}

func listOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	opts := &zendesk.OrganizationListOptions{
		PageOptions: zendesk.PageOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	for true {
		organizations, page, err := conn.GetOrganizations(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, t := range organizations {
			d.StreamListItem(ctx, t)
		}
		if !page.HasNext() {
			break
		}
		opts.Page++
	}
	return nil, nil
}

func getOrganization(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	result, err := conn.GetOrganization(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
