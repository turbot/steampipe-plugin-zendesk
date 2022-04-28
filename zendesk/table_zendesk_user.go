package zendesk

import (
	"context"

	"github.com/nukosuke/go-zendesk/zendesk"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableZendeskUser() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_user",
		Description: "Zendesk Support has three types of users: end users (your customers), agents, and administrators.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getUser,
		},
		Columns: []*plugin.Column{
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "False if the user has been deleted"},
			{Name: "alias", Type: proto.ColumnType_STRING, Description: "An alias displayed to end users"},
			{Name: "chat_only", Type: proto.ColumnType_BOOL, Description: "Whether or not the user is a chat-only agent"},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time the user was created"},
			{Name: "custom_role_id", Type: proto.ColumnType_INT, Description: "A custom role if the user is an agent on the Enterprise plan"},
			{Name: "default_group_id", Type: proto.ColumnType_INT, Description: "The id of the user's default group"},
			{Name: "details", Type: proto.ColumnType_STRING, Description: "Any details you want to store about the user, such as an address"},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The user's primary email address. *Writeable on create only. On update, a secondary email is added."},
			{Name: "external_id", Type: proto.ColumnType_STRING, Description: "A unique identifier from another system. The API treats the id as case insensitive. Example: \"ian1\" and \"Ian1\" are the same user"},
			{Name: "id", Type: proto.ColumnType_INT, Description: "Automatically assigned when the user is created"},
			{Name: "last_login_at", Type: proto.ColumnType_TIMESTAMP, Description: "The last time the user signed in to Zendesk Support"},
			{Name: "locale", Type: proto.ColumnType_STRING, Description: "The user's locale. A BCP-47 compliant tag for the locale. If both \"locale\" and \"locale_id\" are present on create or update, \"locale_id\" is ignored and only \"locale\" is used."},
			{Name: "locale_id", Type: proto.ColumnType_INT, Description: "The user's language identifier"},
			{Name: "moderator", Type: proto.ColumnType_BOOL, Description: "Designates whether the user has forum moderation capabilities"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The user's name"},
			{Name: "notes", Type: proto.ColumnType_STRING, Description: "Any notes you want to store about the user"},
			{Name: "only_private_comments", Type: proto.ColumnType_BOOL, Description: "true if the user can only create private comments"},
			{Name: "organization_id", Type: proto.ColumnType_INT, Description: "The id of the user's organization. If the user has more than one organization memberships, the id of the user's default organization"},
			{Name: "phone", Type: proto.ColumnType_STRING, Description: "The user's primary phone number."},
			{Name: "photo_content_type", Type: proto.ColumnType_STRING, Description: "The content type of the image. Example value: \"image/png\""},
			{Name: "photo_content_url", Type: proto.ColumnType_STRING, Description: "A full URL where the attachment image file can be downloaded"},
			{Name: "photo_deleted", Type: proto.ColumnType_STRING, Description: "If true, the attachment has been deleted"},
			{Name: "photo_file_name", Type: proto.ColumnType_STRING, Description: "The name of the image file"},
			{Name: "photo_id", Type: proto.ColumnType_INT, Description: "Automatically assigned when created"},
			{Name: "photo_inline", Type: proto.ColumnType_BOOL, Description: "If true, the attachment is excluded from the attachment list and the attachment's URL can be referenced within the comment of a ticket. Default is false"},
			{Name: "photo_size", Type: proto.ColumnType_INT, Description: "The size of the image file in bytes"},
			{Name: "photo_thumbnails", Type: proto.ColumnType_JSON, Description: "An array of attachment objects. Note that photo thumbnails do not have thumbnails"},
			{Name: "report_csv", Type: proto.ColumnType_BOOL, Description: "Whether or not the user can access the CSV report on the Search tab of the Reporting page in the Support admin interface."},
			{Name: "restricted_agent", Type: proto.ColumnType_BOOL, Description: "If the agent has any restrictions; false for admins and unrestricted agents, true for other agents"},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "The user's role. Possible values are \"end-user\", \"agent\", or \"admin\""},
			{Name: "role_type", Type: proto.ColumnType_INT, Description: "The user's role id. 0 for custom agents, 1 for light agent, 2 for chat agent, and 3 for chat agent added to the Support account as a contributor (Chat Phase 4)"},
			{Name: "shared", Type: proto.ColumnType_BOOL, Description: "If the user is shared from a different Zendesk Support instance. Ticket sharing accounts only"},
			{Name: "shared_agent", Type: proto.ColumnType_BOOL, Description: "If the user is a shared agent from a different Zendesk Support instance. Ticket sharing accounts only"},
			{Name: "shared_phone_number", Type: proto.ColumnType_BOOL, Description: "Whether the phone number is shared or not."},
			{Name: "signature", Type: proto.ColumnType_STRING, Description: "The user's signature. Only agents and admins can have signatures"},
			{Name: "suspended", Type: proto.ColumnType_BOOL, Description: "If the agent is suspended. Tickets from suspended users are also suspended, and these users cannot sign in to the end user portal"},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "The user's tags. Only present if your account has user tagging enabled"},
			{Name: "ticket_restriction", Type: proto.ColumnType_STRING, Description: "Specifies which tickets the user has access to. Possible values are: \"organization\", \"groups\", \"assigned\", \"requested\", null"},
			{Name: "timezone", Type: proto.ColumnType_STRING, Description: "The user's time zone."},
			{Name: "two_factor_auth_enabled", Type: proto.ColumnType_BOOL, Description: "If two factor authentication is enabled"},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time the user was last updated"},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The user's API url"},
			{Name: "user_fields", Type: proto.ColumnType_JSON, Description: "Values of custom fields in the user's profile."},
			{Name: "verified", Type: proto.ColumnType_BOOL, Description: "Any of the user's identities is verified."},
		},
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	opts := &zendesk.UserListOptions{
		PageOptions: zendesk.PageOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	for {
		users, page, err := conn.GetUsers(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, t := range users {
			d.StreamListItem(ctx, t)
		}
		if !page.HasNext() {
			break
		}
		opts.Page++
	}
	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	plugin.Logger(ctx).Warn("getUser", "quals", quals)
	id := quals["id"].GetInt64Value()
	plugin.Logger(ctx).Warn("getUser", "id", id)
	result, err := conn.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
