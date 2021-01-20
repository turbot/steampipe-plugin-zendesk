# Table: zendesk_trigger

A trigger consists of one or more actions performed when a ticket is created or
updated. The actions are performed only if certain conditions are met. For
example, a trigger can notify the customer when an agent changes the status of
a ticket to Solved.

## Examples

### List triggers in order

```sql
select
  position,
  title
from
  zendesk_trigger
where
  active
order by
  position;
```

### List all inactive triggers

```sql
select
  title
from
  zendesk_trigger
where
  not active;
```

### Find triggers that only work on high priority tickets

Triggers include both all (`conditions_all`) and any (`conditions_any`) fields.
[Read more here](https://support.zendesk.com/hc/en-us/articles/203662246-About-triggers-and-how-they-work#h_81700717131513292855843).

This test checks if `conditions_all` requires the ticket to be in a high
priority for the trigger to run. We rely on JSON submatching in postgres to
find the condition among the array.

```sql
select
  title,
  conditions_all
from
  zendesk_trigger
where
  conditions_all @> '[{"field":"priority","operator":"value","value":"high"}]';
```

# Expand the actions for each trigger

```sql
select
  title,
  jsonb_path_query(actions, '$[*].field') as actions,
  jsonb_path_query(actions, '$[*].value')
from
  zendesk_trigger;
```
