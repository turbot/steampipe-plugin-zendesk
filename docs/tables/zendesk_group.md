# Table: zendesk_group

When support requests arrive in Zendesk Support, they can be assigned to a
Group. Groups serve as the core element of ticket workflow; support agents are
organized into Groups and tickets can be assigned to a Group only, or to an
assigned agent within a Group. A ticket can never be assigned to an agent
without also being assigned to a Group.

## Examples

### Basic group info

```sql
select
  id,
  name
from
  zendesk_group;
```
