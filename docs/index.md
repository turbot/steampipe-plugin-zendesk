---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/zendesk.svg"
brand_color: "#03363D"
display_name: "Zendesk"
name: "zendesk"
description: "Steampipe plugin for querying Tickets, Users and other resources."
social_about: Use SQL to query tickets, users and more from Zendesk. Open source CLI. No DB required. 
social_preview: "/images/plugins/turbot/zendesk-social-graphic.png"
---

# Zendesk + Turbot

[Zendesk](https://www.zendesk.com/) is a customer service SaaS platform with 200,000+ customers. It enables organizations to provide customer service via text, mobile, phone, email, live chat, social media.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  id,
  created_at,
  assignee_id,
  subject
from
  zendesk_ticket
where
  status = 'open';
```
```
+------+---------------------+--------------+----------------------------------+
| id   | created_at          | assignee_id  | subject                          |
+------+---------------------+--------------+----------------------------------+
| 4582 | 2021-04-09 14:53:25 | 383110186421 | Need help with Export            |
| 4579 | 2021-04-08 21:19:23 | 383110186421 | DB and Workspace Scaling Options |
| 4577 | 2021-04-07 23:27:21 | 383110186421 | How do i create a Report?        |
+------+---------------------+--------------+----------------------------------+
```


## Documentation

- **[Table definitions & examples →](/plugins/turbot/zendesk/tables)**

## Get started

### Install

Download and install the latest Zendesk plugin:

```bash
steampipe plugin install zendesk
```

### Credentials

| Item | Description |
| - | - |
| Credentials | Zendesk requires an [API token](https://support.zendesk.com/hc/en-us/articles/226022787-Generating-a-new-API-token-), subdomain and email for all requests. |
| Permissions | You must be an administrator of your domain to create an API token. |
| Radius | A Zendesk connection is scoped to a single Zendesk account, with a single set of credentials. |
| Resolution |  1. Credentials specified in environment variables e.g. `ZENDESK_TOKEN`.<br />2. Credentials in the Steampipe configuration file (`~/.steampipe/config/zendesk.spc`) |

### Configuration

Installing the latest aws plugin will create a config file (`~/.steampipe/config/zendesk.spc`) with a single connection named `zendesk`:

```hcl
connection "zendesk" {
  plugin = "zendesk"
  subdomain  = "dmi"
  email      = "pam@dmi.com"
  token      = "17ImlCYdfZ3WJIrGk96gCpJn1fi1pLwVdrb23kj4"
}
```

- `subdomain` - The subdomain name of your Zendesk account.
- `email` - Email address of agent user who have permission to access the API.
- `token` - [API token ](https://support.zendesk.com/hc/en-us/articles/226022787-Generating-a-new-API-token-) for your Zendesk instance.

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-zendesk
* Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
