---
title: "Steampipe Table: zendesk_ticket_audit - Query Zendesk Ticket Audits using SQL"
description: "Allows users to query Zendesk Ticket Audits, specifically providing detailed information about each change to a ticket, including the author, the timestamp, and the specific changes made."
---

# Table: zendesk_ticket_audit - Query Zendesk Ticket Audits using SQL

A Zendesk Ticket Audit is a record of all updates and changes made to a given ticket within the Zendesk Support Suite. Each audit contains detailed information about the changes, including the author of the change, the timestamp of the change, and the specific changes made to the ticket. Zendesk Ticket Audits provide a comprehensive history of a ticket's lifecycle, facilitating transparency and accountability in customer support processes.

## Table Usage Guide

The `zendesk_ticket_audit` table provides insights into the changes and updates made to tickets within the Zendesk Support Suite. As a customer support representative or manager, explore audit-specific details through this table, including the author, the timestamp, and the specific changes made to a ticket. Utilize it to track ticket history, monitor changes, and ensure accountability in your customer support processes.

## Examples

### List all ticket audit rows
Explore which ticket audits have been performed, enabling you to understand when and by whom each ticket was modified, as well as the number of events associated with each modification. This can be useful for tracking changes and maintaining accountability in customer support situations.

```sql
select
  ticket_id,
  id,
  created_at,
  author_id,
  jsonb_array_length(events)
from
  zendesk_ticket_audit;
```

### List all events associated with ticket audits
Explore all changes related to ticket audits, including the type of change and the old and new values, to better understand the sequence of events and actions taken. This can be useful for troubleshooting, auditing, or understanding the history of a ticket.

```sql
select
  ta.ticket_id,
  ta.id,
  e ->> 'id' as event_id,
  e ->> 'type' as type,
  e ->> 'field_name' as field_name,
  e ->> 'previous_value' as previous_value,
  e ->> 'value' as value
from
  zendesk_ticket_audit as ta,
lateral
  jsonb_array_elements(ta.events) as e;
```

### List all events that changed tags
This query provides a way to track the changes made to event tags. It is useful in monitoring and auditing purposes, allowing users to identify and analyze alterations made to specific event tags.

```sql
select
  ta.ticket_id,
  ta.id,
  e ->> 'id' as event_id,
  e ->> 'type' as type,
  e ->> 'field_name' as field_name,
  e ->> 'previous_value' as previous_value,
  e ->> 'value' as value
from
  zendesk_ticket_audit as ta,
lateral
  jsonb_array_elements(ta.events) as e
where
  e ->> 'type' = 'Change'
and
  e ->> 'field_name' = 'tags';
```

### List all satisfaction rating events
Identify instances where customer satisfaction ratings have been recorded. This is useful for monitoring customer feedback and improving service quality.

```sql
select
  ta.ticket_id,
  ta.id,
  e ->> 'id' as event_id,
  e ->> 'type' as type,
  e ->> 'score' as score
from
  zendesk_ticket_audit as ta,
lateral
  jsonb_array_elements(ta.events) as e
where
  e ->> 'type' = 'SatisfactionRating'
```