// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iotwireless_task_definition", taskDefinitionDataSourceType)
}

// taskDefinitionDataSourceType returns the Terraform awscc_iotwireless_task_definition data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoTWireless::TaskDefinition resource type.
func taskDefinitionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "TaskDefinition arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "TaskDefinition arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_create_tasks": {
			// Property: AutoCreateTasks
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
			//   "type": "boolean"
			// }
			Description: "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the new wireless gateway task definition",
			//   "pattern": "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}",
			//   "type": "string"
			// }
			Description: "The ID of the new wireless gateway task definition",
			Type:        types.StringType,
			Computed:    true,
		},
		"lo_ra_wan_update_gateway_task_entry": {
			// Property: LoRaWANUpdateGatewayTaskEntry
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The list of task definitions.",
			//   "properties": {
			//     "CurrentVersion": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Model": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "PackageVersion": {
			//           "maxLength": 32,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Station": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UpdateVersion": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Model": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "PackageVersion": {
			//           "maxLength": 32,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Station": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The list of task definitions.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"current_version": {
						// Property: CurrentVersion
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"model": {
									// Property: Model
									Type:     types.StringType,
									Computed: true,
								},
								"package_version": {
									// Property: PackageVersion
									Type:     types.StringType,
									Computed: true,
								},
								"station": {
									// Property: Station
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"update_version": {
						// Property: UpdateVersion
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"model": {
									// Property: Model
									Type:     types.StringType,
									Computed: true,
								},
								"package_version": {
									// Property: PackageVersion
									Type:     types.StringType,
									Computed: true,
								},
								"station": {
									// Property: Station
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the new resource.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the new resource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the destination.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the destination.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"task_definition_type": {
			// Property: TaskDefinitionType
			// CloudFormation resource type schema:
			// {
			//   "description": "A filter to list only the wireless gateway task definitions that use this task definition type",
			//   "enum": [
			//     "UPDATE"
			//   ],
			//   "type": "string"
			// }
			Description: "A filter to list only the wireless gateway task definitions that use this task definition type",
			Type:        types.StringType,
			Computed:    true,
		},
		"update": {
			// Property: Update
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about the gateways to update.",
			//   "properties": {
			//     "LoRaWAN": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CurrentVersion": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Model": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "PackageVersion": {
			//               "maxLength": 32,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Station": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "SigKeyCrc": {
			//           "format": "int64",
			//           "type": "integer"
			//         },
			//         "UpdateSignature": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "UpdateVersion": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Model": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "PackageVersion": {
			//               "maxLength": 32,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Station": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UpdateDataRole": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "UpdateDataSource": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Information about the gateways to update.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"lo_ra_wan": {
						// Property: LoRaWAN
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"current_version": {
									// Property: CurrentVersion
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"model": {
												// Property: Model
												Type:     types.StringType,
												Computed: true,
											},
											"package_version": {
												// Property: PackageVersion
												Type:     types.StringType,
												Computed: true,
											},
											"station": {
												// Property: Station
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"sig_key_crc": {
									// Property: SigKeyCrc
									Type:     types.Int64Type,
									Computed: true,
								},
								"update_signature": {
									// Property: UpdateSignature
									Type:     types.StringType,
									Computed: true,
								},
								"update_version": {
									// Property: UpdateVersion
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"model": {
												// Property: Model
												Type:     types.StringType,
												Computed: true,
											},
											"package_version": {
												// Property: PackageVersion
												Type:     types.StringType,
												Computed: true,
											},
											"station": {
												// Property: Station
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"update_data_role": {
						// Property: UpdateDataRole
						Type:     types.StringType,
						Computed: true,
					},
					"update_data_source": {
						// Property: UpdateDataSource
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoTWireless::TaskDefinition",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::TaskDefinition").WithTerraformTypeName("awscc_iotwireless_task_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                 "Arn",
		"auto_create_tasks":                   "AutoCreateTasks",
		"current_version":                     "CurrentVersion",
		"id":                                  "Id",
		"key":                                 "Key",
		"lo_ra_wan":                           "LoRaWAN",
		"lo_ra_wan_update_gateway_task_entry": "LoRaWANUpdateGatewayTaskEntry",
		"model":                               "Model",
		"name":                                "Name",
		"package_version":                     "PackageVersion",
		"sig_key_crc":                         "SigKeyCrc",
		"station":                             "Station",
		"tags":                                "Tags",
		"task_definition_type":                "TaskDefinitionType",
		"update":                              "Update",
		"update_data_role":                    "UpdateDataRole",
		"update_data_source":                  "UpdateDataSource",
		"update_signature":                    "UpdateSignature",
		"update_version":                      "UpdateVersion",
		"value":                               "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
