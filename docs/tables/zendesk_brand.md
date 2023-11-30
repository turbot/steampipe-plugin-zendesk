---
title: "Steampipe Table: zendesk_brand - Query Zendesk Brands using SQL"
description: "Allows users to query Brands within Zendesk, specifically providing insights into brand-specific information such as brand name, subdomain, host mapping and more."
---

# Table: zendesk_brand - Query Zendesk Brands using SQL

A Brand in Zendesk is a unique identity that you can create within your account. It includes a brand name, host mapping, subdomain, and other details. Brands in Zendesk allow you to provide customer service for multiple brands, products, or services using a single Zendesk account.

## Table Usage Guide

The `zendesk_brand` table provides insights into Brands within Zendesk. As a customer service manager, explore brand-specific details through this table, including brand name, host mapping, subdomain, and more. Utilize it to manage and monitor multiple brands, products, or services using a single Zendesk account.

## Examples

### Get brand information
Explore the specific details of a particular brand in your Zendesk account to better understand its characteristics and settings. This is particularly useful when you need to assess the brand's configuration for troubleshooting or optimization purposes.

```sql
select
  *
from
  zendesk_brand
where
  id = '1234';
```