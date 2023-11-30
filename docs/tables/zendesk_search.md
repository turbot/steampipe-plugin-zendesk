---
title: "Steampipe Table: zendesk_search - Query Zendesk Search Results using SQL"
description: "Allows users to query Zendesk Search Results, specifically retrieving tickets, users, organizations, and other entities based on search criteria."
---

# Table: zendesk_search - Query Zendesk Search Results using SQL

Zendesk Search is a feature within Zendesk Support that allows you to perform advanced searches across your Zendesk data. It enables you to find tickets, users, organizations, and other entities based on a wide range of search criteria. Zendesk Search is a powerful tool for filtering and sorting your Zendesk data, providing a comprehensive view of your customer support activities.

## Table Usage Guide

The `zendesk_search` table provides insights into search results within Zendesk Support. As a Zendesk administrator or support agent, you can explore this table to retrieve detailed information based on specific search criteria, including tickets, users, organizations, and other entities. Use it to uncover insights, track support activities, and optimize your customer support operations.

## Examples

### Find a ticket number 123
Discover the specifics of a particular ticket, such as its details and associated data, which can be useful when needing to quickly access information about a specific customer issue or request.

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
Explore the details related to a specific user in a Zendesk account. This is useful for gaining a comprehensive understanding of a user's interactions and activities within the system.

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
Discover the segments that pertain to a specific organization, ACME, to gain insights into relevant details. This can be useful for understanding the organization's interactions and engagements.

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
Explore consolidated information from multiple searches by using this query. This is beneficial when you want to simultaneously examine different types of data, such as user and organization details, based on specific search terms.
```sql
select
  result_number,
  jsonb_pretty(result)
from
  zendesk_search
where
  query in ('type:user jane', 'type:organization acme');
```