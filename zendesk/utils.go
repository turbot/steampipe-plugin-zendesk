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

	subdomain := os.Getenv("ZENDESK_SUBDOMAIN")
	user := os.Getenv("ZENDESK_USER")
	token := os.Getenv("ZENDESK_TOKEN")

	zendeskConfig := GetConfig(d.Connection)
	if &zendeskConfig != nil {
		if zendeskConfig.Account != nil {
			subdomain = *zendeskConfig.Account
		}
		if zendeskConfig.Email != nil {
			user = *zendeskConfig.Email
		}
		if zendeskConfig.Token != nil {
			token = *zendeskConfig.Token
		}
	}

	if subdomain == "" {
		return nil, errors.New("ZENDESK_SUBDOMAIN environment variable must be set")
	}

	if user == "" {
		return nil, errors.New("ZENDESK_USER environment variable must be set")
	}

	if token == "" {
		return nil, errors.New("ZENDESK_TOKEN environment variable must be set")
	}

	// example.zendesk.com
	client.SetSubdomain(subdomain)

	// Authenticate with API token
	client.SetCredential(zendesk.NewAPITokenCredential(user, token))

	return client, nil
}
