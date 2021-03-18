# Table: zendesk_search

The Search API is a unified search API that returns tickets, users, and
organizations. You can define filters to narrow your search results according
to resource type, dates, and object properties, such as ticket requester or
tag.

[Zendesk search reference](https://support.zendesk.com/hc/en-us/articles/203663226).

## Examples

### Find a ticket number 123

```sql
select
  result_number,
  jsonb_pretty(result)
from
  zendesk_search
where
  query = '123';
```

### Find information about the user Jane

```sql
select
  result_number,
  jsonb_pretty(result)
from
  zendesk_search
where
  query = 'type:user jane';
```

### Find information about the organization ACME

```sql
select
  result_number,
  jsonb_pretty(result)
from
  zendesk_search
where
  query = 'type:organization acme';
```


### Consolidated results from multiple searches
```sql
select
  result_number,
  jsonb_pretty(result)
from
  zendesk_search
where
  query in ('type:user jane', 'type:organization acme');
```