// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_connect_hours_of_operation", hoursOfOperationResourceType)
}

// hoursOfOperationResourceType returns the Terraform awscc_connect_hours_of_operation resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Connect::HoursOfOperation resource type.
func hoursOfOperationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"config": {
			// Property: Config
			// CloudFormation resource type schema:
			// {
			//   "description": "Configuration information for the hours of operation: day, start time, and end time.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Contains information about the hours of operation.",
			//     "properties": {
			//       "Day": {
			//         "description": "The day that the hours of operation applies to.",
			//         "enum": [
			//           "SUNDAY",
			//           "MONDAY",
			//           "TUESDAY",
			//           "WEDNESDAY",
			//           "THURSDAY",
			//           "FRIDAY",
			//           "SATURDAY"
			//         ],
			//         "type": "string"
			//       },
			//       "EndTime": {
			//         "additionalProperties": false,
			//         "description": "The end time that your contact center closes.",
			//         "properties": {
			//           "Hours": {
			//             "description": "The hours.",
			//             "maximum": 23,
			//             "minimum": 0,
			//             "type": "integer"
			//           },
			//           "Minutes": {
			//             "description": "The minutes.",
			//             "maximum": 59,
			//             "minimum": 0,
			//             "type": "integer"
			//           }
			//         },
			//         "required": [
			//           "Hours",
			//           "Minutes"
			//         ],
			//         "type": "object"
			//       },
			//       "StartTime": {
			//         "additionalProperties": false,
			//         "description": "The start time that your contact center opens.",
			//         "properties": {
			//           "Hours": {
			//             "description": "The hours.",
			//             "maximum": 23,
			//             "minimum": 0,
			//             "type": "integer"
			//           },
			//           "Minutes": {
			//             "description": "The minutes.",
			//             "maximum": 59,
			//             "minimum": 0,
			//             "type": "integer"
			//           }
			//         },
			//         "required": [
			//           "Hours",
			//           "Minutes"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "required": [
			//       "Day",
			//       "StartTime",
			//       "EndTime"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 100,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Configuration information for the hours of operation: day, start time, and end time.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"day": {
						// Property: Day
						Description: "The day that the hours of operation applies to.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SUNDAY",
								"MONDAY",
								"TUESDAY",
								"WEDNESDAY",
								"THURSDAY",
								"FRIDAY",
								"SATURDAY",
							}),
						},
					},
					"end_time": {
						// Property: EndTime
						Description: "The end time that your contact center closes.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"hours": {
									// Property: Hours
									Description: "The hours.",
									Type:        types.Int64Type,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(0, 23),
									},
								},
								"minutes": {
									// Property: Minutes
									Description: "The minutes.",
									Type:        types.Int64Type,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(0, 59),
									},
								},
							},
						),
						Required: true,
					},
					"start_time": {
						// Property: StartTime
						Description: "The start time that your contact center opens.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"hours": {
									// Property: Hours
									Description: "The hours.",
									Type:        types.Int64Type,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(0, 23),
									},
								},
								"minutes": {
									// Property: Minutes
									Description: "The minutes.",
									Type:        types.Int64Type,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(0, 59),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(100),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the hours of operation.",
			//   "maxLength": 250,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The description of the hours of operation.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 250),
			},
		},
		"hours_of_operation_arn": {
			// Property: HoursOfOperationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the hours of operation.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/operating-hours/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the hours of operation.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the Amazon Connect instance.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the Amazon Connect instance.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$"), ""),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the hours of operation.",
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the hours of operation.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 127),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more tags.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more tags.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
		"time_zone": {
			// Property: TimeZone
			// CloudFormation resource type schema:
			// {
			//   "description": "The time zone of the hours of operation.",
			//   "type": "string"
			// }
			Description: "The time zone of the hours of operation.",
			Type:        types.StringType,
			Required:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Connect::HoursOfOperation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::HoursOfOperation").WithTerraformTypeName("awscc_connect_hours_of_operation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"config":                 "Config",
		"day":                    "Day",
		"description":            "Description",
		"end_time":               "EndTime",
		"hours":                  "Hours",
		"hours_of_operation_arn": "HoursOfOperationArn",
		"instance_arn":           "InstanceArn",
		"key":                    "Key",
		"minutes":                "Minutes",
		"name":                   "Name",
		"start_time":             "StartTime",
		"tags":                   "Tags",
		"time_zone":              "TimeZone",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
