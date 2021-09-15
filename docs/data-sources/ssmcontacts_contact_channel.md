---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ssmcontacts_contact_channel Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SSMContacts::ContactChannel
---

# awscc_ssmcontacts_contact_channel (Data Source)

Data Source schema for AWS::SSMContacts::ContactChannel



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the engagement to a contact channel.
- **channel_address** (String) The details that SSM Incident Manager uses when trying to engage the contact channel.
- **channel_name** (String) The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
- **channel_type** (String) Device type, which specify notification channel. Currently supported values: ?SMS?, ?VOICE?, ?EMAIL?, ?CHATBOT.
- **contact_id** (String) ARN of the contact resource
- **defer_activation** (Boolean) If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.

