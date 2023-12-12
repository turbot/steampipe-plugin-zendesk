---
title: "Steampipe Table: zendesk_ticket - Query Zendesk Tickets using SQL"
description: "Allows users to query Zendesk Tickets, specifically providing insights into ticket details, including status, priority, and associated customer information."
---

# Table: zendesk_ticket - Query Zendesk Tickets using SQL

Zendesk Tickets are a key component of Zendesk's customer service platform, representing a communication between a customer and an agent. They contain the entire communication thread and relevant data, offering a comprehensive view of the customer's issue and the steps taken to resolve it. They are instrumental in tracking, prioritizing, and solving customer's support interactions.

## Table Usage Guide

The `zendesk_ticket` table offers insights into the tickets within Zendesk's customer service platform. As a customer support agent or manager, you can explore ticket-specific details through this table, including status, priority, and associated customer information. Use it to manage and prioritize your support interactions, gain a better understanding of customer issues, and track the steps taken towards resolution.

## Examples

### List open tickets
Explore which customer issues are unresolved by identifying open tickets in your customer support system. This allows for efficient prioritization and management of customer requests, ensuring timely responses and improved customer satisfaction.

```sql+postgres
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

```sql+sqlite
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
Gain insights into the distribution of ticket statuses to understand which types are most common in your Zendesk system, aiding in resource allocation and customer service improvement efforts.

```sql+postgres
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

```sql+sqlite
select
  status,
  count(*)
from
  zendesk_ticket
group by
  status
order by
  count(*) desc;
```

### Stale tickets
Identify instances where tickets in 'open', 'pending', or 'hold' status have not been updated in the past week. This can help prioritize ticket resolution and improve customer service response times.
Unsolved tickets that haven't had any update for 7 days.


```sql+postgres
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

```sql+sqlite
select
  id,
  status,
  updated_at,
  subject
from
  zendesk_ticket
where
  updated_at < date('now', '-7 day')
and
  status in ('open', 'pending', 'hold');
```


### Tickets assigned to Jane
Discover the segments that are assigned to a particular user, in this case, Jane. This can be useful to understand Jane's workload and the status of her assigned tasks.

```sql+postgres
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

```sql+sqlite
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
  u.name like '%jane%';
```


### Ticket aging reports
Explore which tickets are still open, pending, or on hold, and assess their age to understand the efficiency of your customer service. This query can help identify potential bottlenecks or delays in your ticket resolution process.
All unsolved tickets, sorted by age in days.


```sql+postgres
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

```sql+sqlite
select
  julianday('now') - julianday(t.created_at) as age,
  t.id,
  t.status,
  u.name as agent,
  o.name as organization,
  substr(t.subject, 1, 40) as ticket
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

```sql+postgres
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

```sql+sqlite
with aging as (
  select
    id,
    created_at,
    status,
    subject,
    julianday('now') - julianday(created_at) as age
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
  date('now') as date,
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