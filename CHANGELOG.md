## v0.3.0 [2021-11-23]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#20](https://github.com/turbot/steampipe-plugin-zendesk/pull/20))

_Bug fixes_

  - Fixed example query in `zendesk_organization` table ([#22](https://github.com/turbot/steampipe-plugin-zendesk/pull/22))

## v0.2.4 [2021-09-22]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#16](https://github.com/turbot/steampipe-plugin-zendesk/pull/16))
- Changed plugin license to Apache 2.0 per [turbot/steampipe](https://github.com/turbot/steampipe/issues/488) ([#14](https://github.com/turbot/steampipe-plugin-zendesk/pull/14))

## v0.2.3 [2021-05-06]

_What's new?_

_Documentation_

- Updated README.md and index.md with latest standards ([#12](https://github.com/turbot/steampipe-plugin-zendesk/pull/12))

## v0.2.2 [2021-03-18]

_Enhancements_

- Update examples for `zendesk_search` table ([#10](https://github.com/turbot/steampipe-plugin-zendesk/pull/10))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v024-2021-03-16)

## v0.2.1 [2021-02-25]

_Bug fixes_

- Recompiled plugin with latest [steampipe-plugin-sdk](https://github.com/turbot/steampipe-plugin-sdk) to resolve SDK issues:
  - Fix error for missing required quals [#40](https://github.com/turbot/steampipe-plugin-sdk/issues/42).
  - Queries fail with error socket: too many open files [#190](https://github.com/turbot/steampipe/issues/190)

## v0.2.0 [2021-02-18]

_What's new?_

- Added support for [connection configuration](https://github.com/turbot/steampipe-plugin-zendesk/blob/main/docs/index.md#connection-configuration). You may specify zendesk `subdomain`, `email` and `token` for each connection in a configuration file. You can have multiple zendesk connections, each configured for a different zendesk account.
