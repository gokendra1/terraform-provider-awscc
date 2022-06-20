// Code generated by generators/resource/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_macie_findings_filter", findingsFilterResourceType)
}

// findingsFilterResourceType returns the Terraform awscc_macie_findings_filter resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Macie::FindingsFilter resource type.
func findingsFilterResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"action": {
			// Property: Action
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter action.",
			//   "enum": [
			//     "ARCHIVE",
			//     "NOOP"
			//   ],
			//   "type": "string"
			// }
			Description: "Findings filter action.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ARCHIVE",
					"NOOP",
				}),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter ARN.",
			//   "type": "string"
			// }
			Description: "Findings filter ARN.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter description",
			//   "type": "string"
			// }
			Description: "Findings filter description",
			Type:        types.StringType,
			Optional:    true,
		},
		"finding_criteria": {
			// Property: FindingCriteria
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter criteria.",
			//   "properties": {
			//     "Criterion": {
			//       "description": "Map of filter criteria.",
			//       "patternProperties": {
			//         "": {
			//           "properties": {
			//             "eq": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             },
			//             "gt": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "gte": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "lt": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "lte": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "neq": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Findings filter criteria.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"criterion": {
						// Property: Criterion
						Description: "Map of filter criteria.",
						// Pattern: ""
						Attributes: tfsdk.MapNestedAttributes(
							map[string]tfsdk.Attribute{
								"eq": {
									// Property: eq
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"gt": {
									// Property: gt
									Type:     types.Int64Type,
									Optional: true,
								},
								"gte": {
									// Property: gte
									Type:     types.Int64Type,
									Optional: true,
								},
								"lt": {
									// Property: lt
									Type:     types.Int64Type,
									Optional: true,
								},
								"lte": {
									// Property: lte
									Type:     types.Int64Type,
									Optional: true,
								},
								"neq": {
									// Property: neq
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"findings_filter_list_items": {
			// Property: FindingsFilterListItems
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filters list.",
			//   "items": {
			//     "description": "Returned by ListHandler representing filter name and ID.",
			//     "properties": {
			//       "Id": {
			//         "type": "string"
			//       },
			//       "Name": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Findings filters list.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"id": {
						// Property: Id
						Type:     types.StringType,
						Optional: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter ID.",
			//   "type": "string"
			// }
			Description: "Findings filter ID.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter name",
			//   "type": "string"
			// }
			Description: "Findings filter name",
			Type:        types.StringType,
			Required:    true,
		},
		"position": {
			// Property: Position
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter position.",
			//   "type": "integer"
			// }
			Description: "Findings filter position.",
			Type:        types.Int64Type,
			Optional:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Macie FindingsFilter resource schema.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::FindingsFilter").WithTerraformTypeName("awscc_macie_findings_filter")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                     "Action",
		"arn":                        "Arn",
		"criterion":                  "Criterion",
		"description":                "Description",
		"eq":                         "eq",
		"finding_criteria":           "FindingCriteria",
		"findings_filter_list_items": "FindingsFilterListItems",
		"gt":                         "gt",
		"gte":                        "gte",
		"id":                         "Id",
		"lt":                         "lt",
		"lte":                        "lte",
		"name":                       "Name",
		"neq":                        "neq",
		"position":                   "Position",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
