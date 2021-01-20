# Table: zendesk_ticket

Work with [tickets in Zendesk](https://developer.zendesk.com/rest_api/docs/support/tickets).

## Examples

### List open tickets

```sql
select
  id,
  created_at,
  assignee_id,
  organization_id,
  subject
from
  zendesk_ticket
where
  status = 'open';
```

### Ticket status summary

```sql
select
  status,
  count(*)
from
  zendesk_ticket
group by
  status
order by
  count desc;
```

### Stale tickets

Unsolved tickets that haven't had any update for 7 days.

```sql
select
  id,
  status,
  updated_at,
  subject
from
  zendesk_ticket
where
  updated_at < current_date - interval '7 days'
and
  status in ('open', 'pending', 'hold');
```


### Tickets assigned to Jane

```sql
select
  u.name,
  t.status,
  t.subject
from
  zendesk_ticket as t,
  zendesk_user as u
where
  t.assignee_id = u.id
and
  u.name ilike '%jane%';
```


### Ticket aging reports

All unsolved tickets, sorted by age in days.

```sql
select
  date_part('day', now() - t.created_at) as age,
  t.id,
  t.status,
  u.name as agent,
  o.name as organization,
  substring(t.subject for 40) as ticket
from
  zendesk_ticket as t,
  zendesk_user as u,
  zendesk_organization as o
where
  t.assignee_id = u.id
and
  t.organization_id = o.id
and
  t.status in ('open', 'pending', 'hold')
order by
  t.id asc;
```

Summary of ticket age, used for daily snapshots of progress.

```sql
with aging as (
  select
    id,
    created_at,
    status,
    subject,
    date_part('day', now() - created_at) as age
  from
    zendesk.zendesk_ticket
  where
    status in ('open', 'pending', 'hold')
  order by
    id asc
),
stats as (
  select
    status,
    sum(age)
  from
    aging
  group by
    status
)
select
  current_date as date,
  (
    select
      sum
    from
      stats
    where
      status = 'open'
  ) as open,
  (
    select
      sum
    from
      stats
    where
      status = 'pending'
  ) as pending,
  (
    select
      sum
    from
      stats
    where
      status = 'hold'
  ) as hold,
  (
    select
      sum(sum)
    from
      stats
  ) as total;
```
