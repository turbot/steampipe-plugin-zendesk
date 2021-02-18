package zendesk

import (
	"context"

	"github.com/nukosuke/go-zendesk/zendesk"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	//"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type searchResult struct {
	Query        string
	ResultNumber int
	Result       interface{}
}

func tableZendeskSearch() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_search",
		Description: "Search results for the given query.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("query"),
			Hydrate:    listSearch,
		},
		Columns: []*plugin.Column{
			{Name: "query", Type: proto.ColumnType_STRING},
			{Name: "result_number", Type: proto.ColumnType_INT},
			// Mixture of types, including tickets and users
			{Name: "result", Type: proto.ColumnType_JSON},
		},
	}
}

func listSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	q := quals["query"].GetStringValue()
	opts := &zendesk.SearchOptions{
		PageOptions: zendesk.PageOptions{
			Page:    1,
			PerPage: 100,
		},
		Query: q,
	}
	n := 0
	for true {
		results, page, err := conn.Search(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, i := range results.List() {
			n++
			d.StreamListItem(ctx, searchResult{
				Query:        q,
				ResultNumber: n,
				Result:       i,
			})
		}
		if !page.HasNext() {
			break
		}
		opts.Page++
	}
	return nil, nil
}
