---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_lightsail_database Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Lightsail::Database
---

# awscc_lightsail_database (Data Source)

Data Source schema for AWS::Lightsail::Database



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **availability_zone** (String) The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.
- **backup_retention** (Boolean) When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.
- **ca_certificate_identifier** (String) Indicates the certificate that needs to be associated with the database.
- **database_arn** (String)
- **master_database_name** (String) The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.
- **master_user_password** (String) The password for the master user. The password can include any printable ASCII character except "/", """, or "@". It cannot contain spaces.
- **master_username** (String) The name for the master user.
- **preferred_backup_window** (String) The daily time range during which automated backups are created for your new database if automated backups are enabled.
- **preferred_maintenance_window** (String) The weekly time range during which system maintenance can occur on your new database.
- **publicly_accessible** (Boolean) Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
- **relational_database_blueprint_id** (String) The blueprint ID for your new database. A blueprint describes the major engine version of a database.
- **relational_database_bundle_id** (String) The bundle ID for your new database. A bundle describes the performance specifications for your database.
- **relational_database_name** (String) The name to use for your new Lightsail database resource.
- **relational_database_parameters** (Attributes Set) Update one or more parameters of the relational database. (see [below for nested schema](#nestedatt--relational_database_parameters))
- **rotate_master_user_password** (Boolean) When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.
- **tags** (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--relational_database_parameters"></a>
### Nested Schema for `relational_database_parameters`

Read-Only:

- **allowed_values** (String) Specifies the valid range of values for the parameter.
- **apply_method** (String) Indicates when parameter updates are applied. Can be immediate or pending-reboot.
- **apply_type** (String) Specifies the engine-specific parameter type.
- **data_type** (String) Specifies the valid data type for the parameter.
- **description** (String) Provides a description of the parameter.
- **is_modifiable** (Boolean) A Boolean value indicating whether the parameter can be modified.
- **parameter_name** (String) Specifies the name of the parameter.
- **parameter_value** (String) Specifies the value of the parameter.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

