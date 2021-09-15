---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudfront_realtime_log_config Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::CloudFront::RealtimeLogConfig
---

# awscc_cloudfront_realtime_log_config (Resource)

Resource Type definition for AWS::CloudFront::RealtimeLogConfig



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **end_points** (Attributes List, Min: 1) (see [below for nested schema](#nestedatt--end_points))
- **fields** (List of String)
- **name** (String)
- **sampling_rate** (Number)

### Read-Only

- **arn** (String)
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--end_points"></a>
### Nested Schema for `end_points`

Required:

- **kinesis_stream_config** (Attributes) (see [below for nested schema](#nestedatt--end_points--kinesis_stream_config))
- **stream_type** (String)

<a id="nestedatt--end_points--kinesis_stream_config"></a>
### Nested Schema for `end_points.kinesis_stream_config`

Required:

- **role_arn** (String)
- **stream_arn** (String)

