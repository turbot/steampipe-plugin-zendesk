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

Installing the latest zendesk plugin will create a default connection named `zendesk`. This connection will dynamically determine the scope and credentials using the `ZENDESK_SUBDOMAIN`, `ZENDESK_USER` and `ZENDESK_TOKEN` environment variables.

Note that there is nothing special about the default connection, other than that it is created by default on plugin install - You can delete or rename this connection, or modify its configuration options (via the configuration file).

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

### Scope

A Zendesk connection is scoped to a single Zendesk account, with a single set of credentials.

### Configuration Arguments

The Zendesk plugin allows you set credentials static credentials with the following arguments:

- `account` - The account name of your Zendesk instance. If the `account` argument is not specified for a connection, the account will be determined from:
  - The `ZENDESK_SUBDOMAIN` environment variable, if set;
- `email` - Email address of agent user who have permission to access the API. If the `email` argument is not specified in a connection, the email will be taken from:
  - The `ZENDESK_USER` environment variable, if set;
- `token` - API token for your Zendesk instance. If the `token` argument is not specified in a connection, the token can be determined from:
  - Configure the API token by Logging into your Zendesk account and [generate an API token](https://support.zendesk.com/hc/en-us/articles/226022787-Generating-a-new-API-token-), and set
  - The `ZENDESK_TOKEN` environment variable

#### Example configurations

- The default connection.

  ```hcl
  connection "zendesk" {
    plugin = "zendesk"
  }
  ```

- The connection to a specific account using config.

  ```hcl
    connection "zendesk" {
    account    = "dmi"
    email      = "pam@dmi.com"
    token      = "17ImlCYdfZ3WJIrGk96gCpJn1fi1pLwVdrb23kj4"
  }
  ```
