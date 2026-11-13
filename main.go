package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-zendesk/zendesk"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: zendesk.Plugin})
}
