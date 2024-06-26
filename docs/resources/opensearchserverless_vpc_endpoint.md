---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_opensearchserverless_vpc_endpoint Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Amazon OpenSearchServerless vpc endpoint resource
---

# awscc_opensearchserverless_vpc_endpoint (Resource)

Amazon OpenSearchServerless vpc endpoint resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the VPC Endpoint
- `subnet_ids` (List of String) The ID of one or more subnets in which to create an endpoint network interface
- `vpc_id` (String) The ID of the VPC in which the endpoint will be used.

### Optional

- `security_group_ids` (List of String) The ID of one or more security groups to associate with the endpoint network interface

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `vpc_endpoint_id` (String) The identifier of the VPC Endpoint

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_opensearchserverless_vpc_endpoint.example <resource ID>
```
