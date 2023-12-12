---
title: "Steampipe Table: zendesk_user - Query Zendesk Users using SQL"
description: "Allows users to query Users in Zendesk, specifically the details of users including their roles, emails, and active status, providing insights into user management and activity."
---

# Table: zendesk_user - Query Zendesk Users using SQL

Zendesk is a customer service software company that provides a cloud-based customer support platform improving the relationship between businesses and their customers. It offers a suite of support apps that helps improve customer service and having better customer engagement and relationships. The Users in Zendesk are the agents, administrators, or customers who can create and manage tickets.

## Table Usage Guide

The `zendesk_user` table provides insights into Users within Zendesk. As a customer service manager, explore user-specific details through this table, including roles, emails, and active status. Utilize it to uncover information about users, such as their roles, email addresses, and whether they are active or not.

## Examples

### Basic user info
Determine the active status and last login details of users to better understand their engagement with your platform. This information could be useful in identifying patterns of usage or detecting inactive users.

```sql+postgres
select
  id,
  name,
  email,
  active,
  last_login_at
from
  zendesk_user;
```

```sql+sqlite
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
Explore which users in your Zendesk account have administrative privileges. This is useful for auditing account access and ensuring only the appropriate users have high-level permissions.

```sql+postgres
select
  name,
  email
from
  zendesk_user
where
  role = 'admin';
```

```sql+sqlite
select
  name,
  email
from
  zendesk_user
where
  role = 'admin';
```

### Agents and admins (paid seats) who have not logged in for 30 days
Determine the agents and administrators who haven't accessed the system in the last 30 days. This could be useful in assessing user engagement levels or identifying inactive accounts for potential follow-up or account management actions.

```sql+postgres
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

```sql+sqlite
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
  last_login_at < date('now','-30 days');
```

### Number of users per organization
Explore which organizations have the most users, allowing you to understand user distribution and identify high-usage organizations. This can help in resource allocation and strategic planning.

```sql+postgres
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

```sql+sqlite
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
  count(*) desc;
```