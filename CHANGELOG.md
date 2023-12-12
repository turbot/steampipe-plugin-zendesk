## v0.8.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#55](https://github.com/turbot/steampipe-plugin-zendesk/pull/55))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#55](https://github.com/turbot/steampipe-plugin-zendesk/pull/55))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-zendesk/blob/main/docs/LICENSE). ([#55](https://github.com/turbot/steampipe-plugin-zendesk/pull/55))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#54](https://github.com/turbot/steampipe-plugin-zendesk/pull/54))

## v0.7.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#47](https://github.com/turbot/steampipe-plugin-zendesk/pull/47))

## v0.7.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#45](https://github.com/turbot/steampipe-plugin-zendesk/pull/45))
- Recompiled plugin with Go version `1.21`. ([#45](https://github.com/turbot/steampipe-plugin-zendesk/pull/45))

## v0.6.0 [2023-03-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#38](https://github.com/turbot/steampipe-plugin-zendesk/pull/38))

_Bug fixes_

- Fixed the `external_id` column in `zendesk_ticket` table to be of `STRING` data type instead of `INT`. ([#37](https://github.com/turbot/steampipe-plugin-zendesk/pull/37)) (Thanks [@tylarb](https://github.com/tylarb) for the contribution!)

## v0.5.0 [2022-09-09]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#35](https://github.com/turbot/steampipe-plugin-zendesk/pull/35))
- Recompiled plugin with Go version `1.19`. ([#35](https://github.com/turbot/steampipe-plugin-zendesk/pull/35))

## v0.4.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#31](https://github.com/turbot/steampipe-plugin-zendesk/pull/31))

## v0.4.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#29](https://github.com/turbot/steampipe-plugin-zendesk/pull/29))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#28](https://github.com/turbot/steampipe-plugin-zendesk/pull/28))

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
