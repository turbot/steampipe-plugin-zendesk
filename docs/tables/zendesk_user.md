# Table: zendesk_user

Zendesk Support has three types of users: end users (your customers), agents, and administrators.

## Examples

### Basic user info

```sql
select
  id,
  name,
  email,
  active,
  last_login_at
from
  zendesk_user;
```

### List administrators

```sql
select
  name,
  email
from
  zendesk_user
where
  role = 'admin';
```

### Agents and admins (paid seats) who have not logged in for 30 days

```sql
select
  name,
  email,
  role,
  last_login_at
from
  zendesk_user
where
  role in ('admin', 'agent')
and
  last_login_at < current_date - interval '30 days';
```

### Number of users per organization

```sql
select
  o.name,
  count(*)
from
  zendesk_user as u,
  zendesk_organization as o
where
  u.organization_id = o.id
group by
  o.id,
  o.name
order by
  count desc;
```
