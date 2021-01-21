# Table: zendesk_organization

Just as agents can be segmented into groups in Zendesk Support, your customers
(end-users) can be segmented into organizations. You can manually assign
customers to an organization or automatically assign them to an organization by
their email address domain. Organizations can be used in business rules to
route tickets to groups of agents or to send email notifications.

## Examples

### Basic organization info

```sql
select
  id,
  name
from
  zendesk_organization;
```

### Check if ticket and comment sharing settings match

```sql
select
  name,
  shared_comments,
  shared_tickets
from
  zendesk_organization
where
  shared_comments != shared_tickets;
```

### Get ticket counts by organization

```sql
select
  o.name,
  count(t.id)
from
  zendesk_organization as o,
  zendesk_ticket as t
where
  o.id = t.organization_id
group by
  o.name
order by
  count desc;`
```

### Find all users for an organization

```sql
select
  u.name,
  u.email
from
  zendesk_user as u,
  zendesk_organization as o
where
  u.organization_id = o.id
and
  o.name ilike 'ACME';
```

