---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/zendesk.svg"
brand_color: "#03363D"
display_name: "Zendesk"
name: "zendesk"
description: "Steampipe plugin for querying Tickets, Users and other resources."
---

# Zendesk

The Zendesk plugin is used to query tickets, users and other data.

## Installation

To download and install the latest zendesk plugin:

```bash
$ steampipe plugin install zendesk
Installing plugin zendesk...
$
```

### Scope

A Zendesk connection is scoped to a single Zendesk account, with a single set of credentials.

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest zendesk plugin will create a connection file (`~/.steampipe/config/zendesk.spc`) with a single connection named `zendesk`. You must modify this connection to include your user email, organization name, and API token.

```hcl
connection "zendesk" {
  plugin = "zendesk"
  subdomain  = "dmi"
  email      = "pam@dmi.com"
  token      = "17ImlCYdfZ3WJIrGk96gCpJn1fi1pLwVdrb23kj4"
}
```

### Configuration Arguments

- `subdomain` - The subdomain name of your Zendesk account.
- `email` - Email address of agent user who have permission to access the API.
- `token` - [API token ](https://support.zendesk.com/hc/en-us/articles/226022787-Generating-a-new-API-token-) for your Zendesk instance.

For backward compatibility, you may instead authenticate via environment variables, however this behavior is deprecated.

- If the `subdomain` argument is not specified for a connection, the subdomain will be determined from the `ZENDESK_SUBDOMAIN` environment variable, if set.
- If the `email` argument is not specified in a connection, the email will be taken from the `ZENDESK_USER` environment variable, if set.
- If the `token` argument is not specified in a connection, the token will be determined from the `ZENDESK_TOKEN` environment variable, if set.
