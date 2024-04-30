package zendesk

import (
	"context"
	"errors"
	"os"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*zendesk.Client, error) {
	// You can set custom *http.Client here
	client, err := zendesk.NewClient(nil)
	if err != nil {
		return nil, err
	}

	subdomain := os.Getenv("ZENDESK_SUBDOMAIN")
	user := os.Getenv("ZENDESK_USER")
	token := os.Getenv("ZENDESK_TOKEN")

	zendeskConfig := GetConfig(d.Connection)
	if zendeskConfig.SubDomain != nil {
		subdomain = *zendeskConfig.SubDomain
	}
	if zendeskConfig.Email != nil {
		user = *zendeskConfig.Email
	}
	if zendeskConfig.Token != nil {
		token = *zendeskConfig.Token
	}

	if subdomain == "" {
		return nil, errors.New("'subdomain' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if user == "" {
		return nil, errors.New("'email' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// example.zendesk.com
	err = client.SetSubdomain(subdomain)
	if err != nil {
		return nil, err
	}

	// Authenticate with API token
	client.SetCredential(zendesk.NewAPITokenCredential(user, token))

	return client, nil
}

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "subdomain",
			Description: "The organization subdomain name of the Zendesk instance.",
			Hydrate:     getSubdomain,
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

var getSubdomainMemoize = plugin.HydrateFunc(getSubdomainUncached).Memoize(memoize.WithCacheKeyFunction(getSubdomainCacheKey))

func getSubdomainCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheKey := "getSubdomain"
	return cacheKey, nil
}

func getSubdomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	config, err := getSubdomainMemoize(ctx, d, h)
	if err != nil {
		return nil, err
	}

	c := config.(zendeskConfig)

	return c.SubDomain, nil
}

func getSubdomainUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheKey := "getSubdomain"

	var consultData zendeskConfig

	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		consultData = cachedData.(zendeskConfig)
	} else {
		consultData = GetConfig(d.Connection)

		d.ConnectionManager.Cache.Set(cacheKey, consultData)
	}

	return consultData, nil
}
