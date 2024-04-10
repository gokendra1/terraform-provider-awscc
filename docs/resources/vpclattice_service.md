---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_vpclattice_service Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A service is any software application that can run on instances containers, or serverless functions within an account or virtual private cloud (VPC).
---

# awscc_vpclattice_service (Resource)

A service is any software application that can run on instances containers, or serverless functions within an account or virtual private cloud (VPC).



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `auth_type` (String)
- `certificate_arn` (String)
- `custom_domain_name` (String)
- `dns_entry` (Attributes) (see [below for nested schema](#nestedatt--dns_entry))
- `name` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `created_at` (String)
- `id` (String) Uniquely identifies the resource.
- `last_updated_at` (String)
- `service_id` (String)
- `status` (String)

<a id="nestedatt--dns_entry"></a>
### Nested Schema for `dns_entry`

Read-Only:

- `domain_name` (String)
- `hosted_zone_id` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_vpclattice_service.example <resource ID>
```