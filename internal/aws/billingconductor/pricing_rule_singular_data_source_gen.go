// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package billingconductor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_billingconductor_pricing_rule", pricingRuleDataSourceType)
}

// pricingRuleDataSourceType returns the Terraform awscc_billingconductor_pricing_rule data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::BillingConductor::PricingRule resource type.
func pricingRuleDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Pricing rule ARN",
			//   "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:pricingrule/[a-zA-Z0-9]{10}",
			//   "type": "string"
			// }
			Description: "Pricing rule ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"associated_pricing_plan_count": {
			// Property: AssociatedPricingPlanCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of pricing plans associated with pricing rule",
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Description: "The number of pricing plans associated with pricing rule",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Creation timestamp in UNIX epoch time format",
			//   "type": "integer"
			// }
			Description: "Creation timestamp in UNIX epoch time format",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Pricing rule description",
			//   "maxLength": 1024,
			//   "type": "string"
			// }
			Description: "Pricing rule description",
			Type:        types.StringType,
			Computed:    true,
		},
		"last_modified_time": {
			// Property: LastModifiedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Latest modified timestamp in UNIX epoch time format",
			//   "type": "integer"
			// }
			Description: "Latest modified timestamp in UNIX epoch time format",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"modifier_percentage": {
			// Property: ModifierPercentage
			// CloudFormation resource type schema:
			// {
			//   "description": "Pricing rule modifier percentage",
			//   "minimum": 0,
			//   "type": "number"
			// }
			Description: "Pricing rule modifier percentage",
			Type:        types.Float64Type,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Pricing rule name",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9_\\+=\\.\\-@]+",
			//   "type": "string"
			// }
			Description: "Pricing rule name",
			Type:        types.StringType,
			Computed:    true,
		},
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			// {
			//   "description": "A term used to categorize the granularity of a Pricing Rule.",
			//   "enum": [
			//     "GLOBAL",
			//     "SERVICE"
			//   ],
			//   "type": "string"
			// }
			Description: "A term used to categorize the granularity of a Pricing Rule.",
			Type:        types.StringType,
			Computed:    true,
		},
		"service": {
			// Property: Service
			// CloudFormation resource type schema:
			// {
			//   "description": "The service which a pricing rule is applied on",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9\\.\\-]+",
			//   "type": "string"
			// }
			Description: "The service which a pricing rule is applied on",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
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
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "One of MARKUP or DISCOUNT that describes the direction of the rate that is applied to a pricing plan.",
			//   "enum": [
			//     "MARKUP",
			//     "DISCOUNT"
			//   ],
			//   "type": "string"
			// }
			Description: "One of MARKUP or DISCOUNT that describes the direction of the rate that is applied to a pricing plan.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::BillingConductor::PricingRule",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::BillingConductor::PricingRule").WithTerraformTypeName("awscc_billingconductor_pricing_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                           "Arn",
		"associated_pricing_plan_count": "AssociatedPricingPlanCount",
		"creation_time":                 "CreationTime",
		"description":                   "Description",
		"key":                           "Key",
		"last_modified_time":            "LastModifiedTime",
		"modifier_percentage":           "ModifierPercentage",
		"name":                          "Name",
		"scope":                         "Scope",
		"service":                       "Service",
		"tags":                          "Tags",
		"type":                          "Type",
		"value":                         "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
