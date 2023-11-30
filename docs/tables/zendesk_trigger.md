---
title: "Steampipe Table: zendesk_trigger - Query Zendesk Triggers using SQL"
description: "Allows users to query Zendesk Triggers, specifically to retrieve and manage information about automated actions that are performed on tickets based on defined conditions."
---

# Table: zendesk_trigger - Query Zendesk Triggers using SQL

Zendesk Triggers are automated actions that occur when tickets meet predefined conditions. Triggers are used to streamline customer support workflows by automatically updating ticket properties and sending notifications to customers and agents. They play a crucial role in the automation and efficiency of ticket handling within the Zendesk Support Suite.

## Table Usage Guide

The `zendesk_trigger` table provides insights into triggers within Zendesk Support Suite. As a support manager or system administrator, explore trigger-specific details through this table, including conditions, actions, and associated metadata. Utilize it to uncover information about triggers, such as those with specific conditions or actions, the order of triggers, and the verification of trigger effectiveness.

## Examples

### List triggers in order
Identify the sequence of active triggers to understand their order of execution. This is useful for troubleshooting and optimizing workflow processes.

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
Explore which triggers are currently inactive in your Zendesk account. This can be useful for cleaning up your configuration and reactivating or deleting unused triggers to improve system performance.

```sql
select
  title
from
  zendesk_trigger
where
  not active;
```

### Find triggers that only work on high priority tickets
Explore which triggers are set to activate only on high-priority tickets, allowing you to prioritize and manage your customer service resources efficiently. Additionally, gain insights into the specific actions each trigger initiates, helping you understand and optimize your automated support processes.
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