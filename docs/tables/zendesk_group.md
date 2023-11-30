---
title: "Steampipe Table: zendesk_group - Query Zendesk Groups using SQL"
description: "Allows users to query Groups in Zendesk, specifically providing details about group name, group members, and related information."
---

# Table: zendesk_group - Query Zendesk Groups using SQL

Zendesk Groups are a feature within Zendesk Support that allow you to organize your agents into different groups. This helps in managing and routing tickets effectively. Groups can be based on the agent's expertise, the team they belong to, the type of issue, or any other criteria.

## Table Usage Guide

The `zendesk_group` table provides insights into Groups within Zendesk Support. As a support manager or team lead, explore group-specific details through this table, including group members, associated tickets, and related information. Utilize it to manage workload distribution, ensure efficient ticket routing, and improve overall support operations.

## Examples

### Basic group info
Explore the different groups available within your Zendesk account to better manage customer interactions and support workflows. This can be particularly useful for organizing your support team and streamlining communication.

```sql
select
  id,
  name
from
  zendesk_group;
```