// Code generated by generators/resource/main.go; DO NOT EDIT.

package kendra

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_kendra_index", indexResourceType)
}

// indexResourceType returns the Terraform awscc_kendra_index resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Kendra::Index resource type.
func indexResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1000),
			},
			Computed: true,
		},
		"capacity_units": {
			// Property: CapacityUnits
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "QueryCapacityUnits": {
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "StorageCapacityUnits": {
			//       "minimum": 0,
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "StorageCapacityUnits",
			//     "QueryCapacityUnits"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"query_capacity_units": {
						// Property: QueryCapacityUnits
						Type:     types.NumberType,
						Required: true,
					},
					"storage_capacity_units": {
						// Property: StorageCapacityUnits
						Type:     types.NumberType,
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1000),
			},
			Optional: true,
		},
		"document_metadata_configurations": {
			// Property: DocumentMetadataConfigurations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 30,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Relevance": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Duration": {
			//             "maxLength": 10,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Freshness": {
			//             "type": "boolean"
			//           },
			//           "Importance": {
			//             "maximum": 10,
			//             "minimum": 1,
			//             "type": "integer"
			//           },
			//           "RankOrder": {
			//             "enum": [
			//               "ASCENDING",
			//               "DESCENDING"
			//             ],
			//             "type": "string"
			//           },
			//           "ValueImportanceItems": {
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Key": {
			//                   "maxLength": 50,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "maximum": 10,
			//                   "minimum": 1,
			//                   "type": "integer"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Search": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Displayable": {
			//             "type": "boolean"
			//           },
			//           "Facetable": {
			//             "type": "boolean"
			//           },
			//           "Searchable": {
			//             "type": "boolean"
			//           },
			//           "Sortable": {
			//             "type": "boolean"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Type": {
			//         "enum": [
			//           "STRING_VALUE",
			//           "STRING_LIST_VALUE",
			//           "LONG_VALUE",
			//           "DATE_VALUE"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Name",
			//       "Type"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 500,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type: types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 30),
						},
						Required: true,
					},
					"relevance": {
						// Property: Relevance
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"duration": {
									// Property: Duration
									Type: types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 10),
									},
									Optional: true,
								},
								"freshness": {
									// Property: Freshness
									Type:     types.BoolType,
									Optional: true,
								},
								"importance": {
									// Property: Importance
									Type: types.NumberType,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(1, 10),
									},
									Optional: true,
								},
								"rank_order": {
									// Property: RankOrder
									Type:     types.StringType,
									Optional: true,
								},
								"value_importance_items": {
									// Property: ValueImportanceItems
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type: types.StringType,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 50),
												},
												Optional: true,
											},
											"value": {
												// Property: Value
												Type: types.NumberType,
												Validators: []tfsdk.AttributeValidator{
													validate.IntBetween(1, 10),
												},
												Optional: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"search": {
						// Property: Search
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"displayable": {
									// Property: Displayable
									Type:     types.BoolType,
									Optional: true,
								},
								"facetable": {
									// Property: Facetable
									Type:     types.BoolType,
									Optional: true,
								},
								"searchable": {
									// Property: Searchable
									Type:     types.BoolType,
									Optional: true,
								},
								"sortable": {
									// Property: Sortable
									Type:     types.BoolType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 500,
				},
			),
			Optional: true,
		},
		"edition": {
			// Property: Edition
			// CloudFormation resource type schema:
			// {
			//   "description": "Edition of index",
			//   "enum": [
			//     "DEVELOPER_EDITION",
			//     "ENTERPRISE_EDITION"
			//   ],
			//   "type": "string"
			// }
			Description: "Edition of index",
			Type:        types.StringType,
			Required:    true,
			// Edition is a force-new attribute.
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique ID of index",
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "type": "string"
			// }
			Description: "Unique ID of index",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(36, 36),
			},
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of index",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of index",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1000),
			},
			Required: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn",
			//   "maxLength": 1284,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Role Arn",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1284),
			},
			Required: true,
		},
		"server_side_encryption_configuration": {
			// Property: ServerSideEncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "KmsKeyId": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						Type: types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// ServerSideEncryptionConfiguration is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "List of tags",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Kendra resources",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for the tag",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array"
			// }
			Description: "List of tags",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
						Required: true,
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
						Type:        types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 200,
				},
			),
			Optional: true,
		},
		"user_context_policy": {
			// Property: UserContextPolicy
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ATTRIBUTE_FILTER",
			//     "USER_TOKEN"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"user_token_configurations": {
			// Property: UserTokenConfigurations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "JsonTokenTypeConfiguration": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "GroupAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "UserNameAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "UserNameAttributeField",
			//           "GroupAttributeField"
			//         ],
			//         "type": "object"
			//       },
			//       "JwtTokenTypeConfiguration": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ClaimRegex": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "GroupAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Issuer": {
			//             "maxLength": 65,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "KeyLocation": {
			//             "enum": [
			//               "URL",
			//               "SECRET_MANAGER"
			//             ],
			//             "type": "string"
			//           },
			//           "SecretManagerArn": {
			//             "description": "Role Arn",
			//             "maxLength": 1284,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "URL": {
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "UserNameAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "KeyLocation"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"json_token_type_configuration": {
						// Property: JsonTokenTypeConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"group_attribute_field": {
									// Property: GroupAttributeField
									Type: types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
									Required: true,
								},
								"user_name_attribute_field": {
									// Property: UserNameAttributeField
									Type: types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"jwt_token_type_configuration": {
						// Property: JwtTokenTypeConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"claim_regex": {
									// Property: ClaimRegex
									Type: types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
									Optional: true,
								},
								"group_attribute_field": {
									// Property: GroupAttributeField
									Type: types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
									Optional: true,
								},
								"issuer": {
									// Property: Issuer
									Type: types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 65),
									},
									Optional: true,
								},
								"key_location": {
									// Property: KeyLocation
									Type:     types.StringType,
									Required: true,
								},
								"secret_manager_arn": {
									// Property: SecretManagerArn
									Description: "Role Arn",
									Type:        types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 1284),
									},
									Optional: true,
								},
								"url": {
									// Property: URL
									Type: types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2048),
									},
									Optional: true,
								},
								"user_name_attribute_field": {
									// Property: UserNameAttributeField
									Type: types.StringType,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 1,
				},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "A Kendra index",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kendra::Index").WithTerraformTypeName("awscc_kendra_index")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                  "Arn",
		"capacity_units":                       "CapacityUnits",
		"claim_regex":                          "ClaimRegex",
		"description":                          "Description",
		"displayable":                          "Displayable",
		"document_metadata_configurations":     "DocumentMetadataConfigurations",
		"duration":                             "Duration",
		"edition":                              "Edition",
		"facetable":                            "Facetable",
		"freshness":                            "Freshness",
		"group_attribute_field":                "GroupAttributeField",
		"id":                                   "Id",
		"importance":                           "Importance",
		"issuer":                               "Issuer",
		"json_token_type_configuration":        "JsonTokenTypeConfiguration",
		"jwt_token_type_configuration":         "JwtTokenTypeConfiguration",
		"key":                                  "Key",
		"key_location":                         "KeyLocation",
		"kms_key_id":                           "KmsKeyId",
		"name":                                 "Name",
		"query_capacity_units":                 "QueryCapacityUnits",
		"rank_order":                           "RankOrder",
		"relevance":                            "Relevance",
		"role_arn":                             "RoleArn",
		"search":                               "Search",
		"searchable":                           "Searchable",
		"secret_manager_arn":                   "SecretManagerArn",
		"server_side_encryption_configuration": "ServerSideEncryptionConfiguration",
		"sortable":                             "Sortable",
		"storage_capacity_units":               "StorageCapacityUnits",
		"tags":                                 "Tags",
		"type":                                 "Type",
		"url":                                  "URL",
		"user_context_policy":                  "UserContextPolicy",
		"user_name_attribute_field":            "UserNameAttributeField",
		"user_token_configurations":            "UserTokenConfigurations",
		"value":                                "Value",
		"value_importance_items":               "ValueImportanceItems",
	})

	opts = opts.WithCreateTimeoutInMinutes(240).WithDeleteTimeoutInMinutes(720)

	opts = opts.WithUpdateTimeoutInMinutes(240)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_kendra_index", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
