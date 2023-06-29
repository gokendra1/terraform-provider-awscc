---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_quicksight_vpc_connection Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::QuickSight::VPCConnection
---

# awscc_quicksight_vpc_connection (Data Source)

Data Source schema for AWS::QuickSight::VPCConnection



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `availability_status` (String)
- `aws_account_id` (String)
- `created_time` (String)
- `dns_resolvers` (List of String)
- `last_updated_time` (String)
- `name` (String)
- `network_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--network_interfaces))
- `role_arn` (String)
- `security_group_ids` (List of String)
- `status` (String)
- `subnet_ids` (List of String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `vpc_connection_id` (String)
- `vpc_id` (String)

<a id="nestedatt--network_interfaces"></a>
### Nested Schema for `network_interfaces`

Read-Only:

- `availability_zone` (String)
- `error_message` (String)
- `network_interface_id` (String)
- `status` (String)
- `subnet_id` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)