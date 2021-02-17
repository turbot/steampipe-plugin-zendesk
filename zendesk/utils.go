package zendesk

import (
	"context"
	"errors"
	"os"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*zendesk.Client, error) {
	// You can set custom *http.Client here
	client, err := zendesk.NewClient(nil)
	if err != nil {
		return nil, err
	}

	zendeskConfig := GetConfig(d.Connection)
	if &zendeskConfig != nil {
		if zendeskConfig.Account != nil {
			os.Setenv("ZENDESK_SUBDOMAIN", *zendeskConfig.Account)
		}
		if zendeskConfig.Email != nil {
			os.Setenv("ZENDESK_USER", *zendeskConfig.Email)
		}
		if zendeskConfig.Token != nil {
			os.Setenv("ZENDESK_TOKEN", *zendeskConfig.Token)
		}
	}

	subdomain, ok := os.LookupEnv("ZENDESK_SUBDOMAIN")
	if !ok || subdomain == "" {
		return nil, errors.New("ZENDESK_SUBDOMAIN environment variable must be set")
	}

	user, ok := os.LookupEnv("ZENDESK_USER")
	if !ok || user == "" {
		return nil, errors.New("ZENDESK_USER environment variable must be set")
	}

	token, ok := os.LookupEnv("ZENDESK_TOKEN")
	if !ok || token == "" {
		return nil, errors.New("ZENDESK_TOKEN environment variable must be set")
	}

	// example.zendesk.com
	client.SetSubdomain(subdomain)

	// Authenticate with API token
	client.SetCredential(zendesk.NewAPITokenCredential(user, token))

	return client, nil
}
