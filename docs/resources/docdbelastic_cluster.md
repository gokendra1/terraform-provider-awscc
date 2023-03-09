---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_docdbelastic_cluster Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::DocDBElastic::Cluster Amazon DocumentDB (with MongoDB compatibility) Elastic Scale resource describes a Cluster
---

# awscc_docdbelastic_cluster (Resource)

The AWS::DocDBElastic::Cluster Amazon DocumentDB (with MongoDB compatibility) Elastic Scale resource describes a Cluster



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_user_name` (String)
- `auth_type` (String)
- `cluster_name` (String)
- `shard_capacity` (Number)
- `shard_count` (Number)

### Optional

- `admin_user_password` (String)
- `kms_key_id` (String)
- `preferred_maintenance_window` (String)
- `subnet_ids` (List of String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- `vpc_security_group_ids` (List of String)

### Read-Only

- `cluster_arn` (String)
- `cluster_endpoint` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_docdbelastic_cluster.example <resource ID>
```