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
		if zendeskConfig.SubDomain != nil {
			subdomain = *zendeskConfig.SubDomain
		}
		if zendeskConfig.Email != nil {
			user = *zendeskConfig.Email
		}
		if zendeskConfig.Token != nil {
			token = *zendeskConfig.Token
		}
	}

	if subdomain == "" {
		return nil, errors.New("'subdomain' must be set in the connection configuration")
	}

	if user == "" {
		return nil, errors.New("'email' must be set in the connection configuration")
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration")
	}

	// example.zendesk.com
	client.SetSubdomain(subdomain)

	// Authenticate with API token
	client.SetCredential(zendesk.NewAPITokenCredential(user, token))

	return client, nil
}
