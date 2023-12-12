---
title: "Steampipe Table: zendesk_organization - Query Zendesk Organizations using SQL"
description: "Allows users to query Zendesk Organizations, providing insights into the details of organizations, their associated users and tickets."
---

# Table: zendesk_organization - Query Zendesk Organizations using SQL

Zendesk Organizations is a feature within Zendesk that allows you to group users into organizations based on various criteria like domain names, email addresses, or custom rules. It provides a way to manage and track tickets from users who are part of the same organization. Zendesk Organizations helps you organize users for better support and reporting.

## Table Usage Guide

The `zendesk_organization` table provides insights into Zendesk Organizations. As a Support Manager, explore organization-specific details through this table, including associated users, tickets, and other relevant details. Utilize it to uncover information about organizations, such as those with a large number of tickets, the users associated with each organization, and the overall structure of your support environment.

## Examples

### Basic organization info
Explore the basic details of your organization to gain insights into its identity. This can assist in understanding the organization's structure and operations.

```sql+postgres
select
  id,
  name
from
  zendesk_organization;
```

```sql+sqlite
select
  id,
  name
from
  zendesk_organization;
```

### Check if ticket and comment sharing settings match
Explore discrepancies in your organization's sharing settings by identifying instances where ticket and comment sharing settings do not align. This can be beneficial in maintaining consistency and ensuring proper data sharing protocols within your organization.

```sql+postgres
select
  name,
  shared_comments,
  shared_tickets
from
  zendesk_organization
where
  shared_comments != shared_tickets;
```

```sql+sqlite
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
Explore which organizations have the highest number of tickets to identify potential areas of customer service issues or high engagement. This allows for targeted problem-solving or resource allocation.

```sql+postgres
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
  count desc;
```

```sql+sqlite
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
  count(t.id) desc;
```

### Find all users for an organization
Explore which users are part of a specific organization to understand team composition and manage access rights. This is beneficial in a scenario where you need to audit user access or update organization-wide communication.

```sql+postgres
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

```sql+sqlite
select
  u.name,
  u.email
from
  zendesk_user as u,
  zendesk_organization as o
where
  u.organization_id = o.id
and
  lower(o.name) like lower('ACME');
```