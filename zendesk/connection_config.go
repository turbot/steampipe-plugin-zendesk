package zendesk

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type zendeskConfig struct {
	SubDomain *string `cty:"subdomain"`
	Email     *string `cty:"email"`
	Token     *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"subdomain": {
		Type: schema.TypeString,
	},
	"email": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &zendeskConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) zendeskConfig {
	if connection == nil || connection.Config == nil {
		return zendeskConfig{}
	}
	config, _ := connection.Config.(zendeskConfig)
	return config
}
