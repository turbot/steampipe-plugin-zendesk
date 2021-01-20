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

Download and install the latest Zendesk plugin:

```bash
$ steampipe plugin install zendesk
Installing plugin zendesk...
$
```

## Configure API Token

Login to your Zendesk account and [generate an API token](https://support.zendesk.com/hc/en-us/articles/226022787-Generating-a-new-API-token-).

Set Zendesk API credentials as environment variables (Mac, Linux):

```bash
export ZENDESK_SUBDOMAIN=dmi
export ZENDESK_USER=pam@dmi.com
export ZENDESK_TOKEN=17ImlCYdfZ3WJIrGk96gCpJn1fi1pLwVdrb23kj4
```

Run a query:

```bash
$ steampipe query
Welcome to Steampipe v0.0.11
Type ".inspect" for more information.
> select id, subject from zendesk_ticket limit 3;
+------+----------------------------------------------+
|  id  |                   subject                    |
+------+----------------------------------------------+
| 4387 | Paper order was meant to be recycled         |
| 4389 | Inappropriate behavior from sales rep        |
| 4385 | Late delivery                                |
+------+----------------------------------------------+
>
```
