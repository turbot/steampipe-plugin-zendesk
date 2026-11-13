package zendesk

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type zendeskConfig struct {
	SubDomain *string `hcl:"subdomain"`
	Email     *string `hcl:"email"`
	Token     *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &zendeskConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) zendeskConfig {
	if connection == nil || connection.GetConfig() == nil {
		return zendeskConfig{}
	}
	config, _ := connection.GetConfig().(zendeskConfig)
	return config
}
