package zendesk

import (
	"context"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableZendeskBrand() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_brand",
		Description: "Brands are your customer-facing identities. They might represent multiple products or services, or they might literally be multiple brands owned and represented by your company.",
		/*
			List: &plugin.ListConfig{
				Hydrate: listBrand,
			},
		*/
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getBrand,
		},
		Columns: commonColumns([]*plugin.Column{

			// TODO - Change the id from string to int, but only after get calls work for bigint

			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Automatically assigned when the brand is created"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the brand"},
			// Other columns
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "If the brand is set active"},
			{Name: "brand_url", Type: proto.ColumnType_STRING, Description: "The URL of the brand"},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time the brand was created"},
			{Name: "default", Type: proto.ColumnType_BOOL, Description: "Is the brand the default brand for this account"},
			{Name: "has_help_center", Type: proto.ColumnType_BOOL, Description: "If the brand has a Help Center"},
			{Name: "help_center_state", Type: proto.ColumnType_STRING, Description: "The state of the Help Center. Allowed values are \"enabled\", \"disabled\", or \"restricted\"."},
			{Name: "host_mapping", Type: proto.ColumnType_STRING, Description: "The hostmapping to this brand, if any. Only admins can view this property."},
			{Name: "logo_content_type", Type: proto.ColumnType_STRING, Description: "The content type of the image. Example value: \"image/png\""},
			{Name: "logo_content_url", Type: proto.ColumnType_STRING, Description: "A full URL where the attachment image file can be downloaded"},
			{Name: "logo_file_name", Type: proto.ColumnType_STRING, Description: "The name of the image file"},
			{Name: "logo_id", Type: proto.ColumnType_INT, Description: "Automatically assigned when created"},
			{Name: "logo_inline", Type: proto.ColumnType_BOOL, Description: "If true, the attachment is excluded from the attachment list and the attachment's URL can be referenced within the comment of a ticket. Default is false"},
			{Name: "logo_size", Type: proto.ColumnType_INT, Description: "The size of the image file in bytes"},
			{Name: "logo_thumbnails", Type: proto.ColumnType_JSON, Description: "An array of attachment objects"},
			{Name: "signature_template", Type: proto.ColumnType_STRING, Description: "The signature template for a brand"},
			{Name: "ticket_form_ids", Type: proto.ColumnType_JSON, Description: "The ids of ticket forms that are available for use by a brand"},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time of the last update of the brand"},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The API URL of the brand"},
		}),
	}
}

/*
func listBrand(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return nil, nil
}
*/

func getBrand(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := conn.GetBrand(ctx, i)
	if err != nil {
		return nil, err
	}
	return result, nil
}
