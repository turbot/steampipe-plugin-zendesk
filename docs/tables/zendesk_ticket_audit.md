# Table: zendesk_ticket_audit

Each update to a ticket creates a row in the `zendesk_ticket_audit` table. Each
update can also cause a number of events, stored as a `jsonb` array in the `events`
column.

It's common to expand the ticket audit data with the events when working on this
table. [Event details are outlined here](https://develop.zendesk.com/hc/en-us/articles/360059038133).

## Examples

### List all ticket audit rows

```sql
select
  ticket_id,
  id,
  created_at,
  author_id,
  jsonb_array_length(events)
from
  zendesk_ticket_audit;
```

### List all events associated with ticket audits

```sql
select
  ta.ticket_id,
  ta.id,
  e ->> 'id' as event_id,
  e ->> 'type' as type,
  e ->> 'field_name' as field_name,
  e ->> 'previous_value' as previous_value,
  e ->> 'value' as value
from
  zendesk_ticket_audit as ta,
lateral
  jsonb_array_elements(ta.events) as e;
```

### List all events that changed tags

```sql
select
  ta.ticket_id,
  ta.id,
  e ->> 'id' as event_id,
  e ->> 'type' as type,
  e ->> 'field_name' as field_name,
  e ->> 'previous_value' as previous_value,
  e ->> 'value' as value
from
  zendesk_ticket_audit as ta,
lateral
  jsonb_array_elements(ta.events) as e
where
  e ->> 'type' = 'Change'
and
  e ->> 'field_name' = 'tags';
```

### List all satisfaction rating events

```sql
select
  ta.ticket_id,
  ta.id,
  e ->> 'id' as event_id,
  e ->> 'type' as type,
  e ->> 'score' as score
from
  zendesk_ticket_audit as ta,
lateral
  jsonb_array_elements(ta.events) as e
where
  e ->> 'type' = 'SatisfactionRating'
```
