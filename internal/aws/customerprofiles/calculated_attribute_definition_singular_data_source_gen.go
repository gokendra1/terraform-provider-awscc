// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package customerprofiles

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_customerprofiles_calculated_attribute_definition", calculatedAttributeDefinitionDataSource)
}

// calculatedAttributeDefinitionDataSource returns the Terraform awscc_customerprofiles_calculated_attribute_definition data source.
// This Terraform data source corresponds to the CloudFormation AWS::CustomerProfiles::CalculatedAttributeDefinition resource.
func calculatedAttributeDefinitionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttributeDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Mathematical expression and a list of attribute items specified in that expression.",
		//	  "properties": {
		//	    "Attributes": {
		//	      "description": "A list of attribute items specified in the mathematical expression.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "The details of a single attribute item specified in the mathematical expression.",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "The name of an attribute defined in a profile object type.",
		//	            "maxLength": 64,
		//	            "minLength": 1,
		//	            "pattern": "^[a-zA-Z0-9_.-]+$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 2,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "Expression": {
		//	      "description": "Mathematical expression that is performed on attribute items provided in the attribute list. Each element in the expression should follow the structure of \"{ObjectTypeName.AttributeName}\".",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Attributes",
		//	    "Expression"
		//	  ],
		//	  "type": "object"
		//	}
		"attribute_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Attributes
				"attributes": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of an attribute defined in a profile object type.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "A list of attribute items specified in the mathematical expression.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Expression
				"expression": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Mathematical expression that is performed on attribute items provided in the attribute list. Each element in the expression should follow the structure of \"{ObjectTypeName.AttributeName}\".",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Mathematical expression and a list of attribute items specified in that expression.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CalculatedAttributeName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name of the calculated attribute.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z_][a-zA-Z_0-9-]*$",
		//	  "type": "string"
		//	}
		"calculated_attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name of the calculated attribute.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Conditions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The conditions including range, object count, and threshold for the calculated attribute.",
		//	  "properties": {
		//	    "ObjectCount": {
		//	      "description": "The number of profile objects used for the calculated attribute.",
		//	      "maximum": 100,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "Range": {
		//	      "additionalProperties": false,
		//	      "description": "The relative time period over which data is included in the aggregation.",
		//	      "properties": {
		//	        "Unit": {
		//	          "description": "The unit of time.",
		//	          "enum": [
		//	            "DAYS"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "Value": {
		//	          "description": "The amount of time of the specified unit.",
		//	          "maximum": 366,
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "Value",
		//	        "Unit"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Threshold": {
		//	      "additionalProperties": false,
		//	      "description": "The threshold for the calculated attribute.",
		//	      "properties": {
		//	        "Operator": {
		//	          "description": "The operator of the threshold.",
		//	          "enum": [
		//	            "EQUAL_TO",
		//	            "GREATER_THAN",
		//	            "LESS_THAN",
		//	            "NOT_EQUAL_TO"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "Value": {
		//	          "description": "The value of the threshold.",
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Value",
		//	        "Operator"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"conditions": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ObjectCount
				"object_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of profile objects used for the calculated attribute.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Range
				"range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Unit
						"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The unit of time.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Value
						"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The amount of time of the specified unit.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The relative time period over which data is included in the aggregation.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Threshold
				"threshold": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Operator
						"operator": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The operator of the threshold.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Value
						"value": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The value of the threshold.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The threshold for the calculated attribute.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The conditions including range, object count, and threshold for the calculated attribute.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the calculated attribute definition was created.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp of when the calculated attribute definition was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the calculated attribute.",
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the calculated attribute.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The display name of the calculated attribute.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z_][a-zA-Z_0-9-\\s]*$",
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The display name of the calculated attribute.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name of the domain.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name of the domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the calculated attribute definition was most recently edited.",
		//	  "type": "string"
		//	}
		"last_updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp of when the calculated attribute definition was most recently edited.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Statistic
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The aggregation operation to perform for the calculated attribute.",
		//	  "enum": [
		//	    "FIRST_OCCURRENCE",
		//	    "LAST_OCCURRENCE",
		//	    "COUNT",
		//	    "SUM",
		//	    "MINIMUM",
		//	    "MAXIMUM",
		//	    "AVERAGE",
		//	    "MAX_OCCURRENCE"
		//	  ],
		//	  "type": "string"
		//	}
		"statistic": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The aggregation operation to perform for the calculated attribute.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CustomerProfiles::CalculatedAttributeDefinition",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::CalculatedAttributeDefinition").WithTerraformTypeName("awscc_customerprofiles_calculated_attribute_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attribute_details":         "AttributeDetails",
		"attributes":                "Attributes",
		"calculated_attribute_name": "CalculatedAttributeName",
		"conditions":                "Conditions",
		"created_at":                "CreatedAt",
		"description":               "Description",
		"display_name":              "DisplayName",
		"domain_name":               "DomainName",
		"expression":                "Expression",
		"key":                       "Key",
		"last_updated_at":           "LastUpdatedAt",
		"name":                      "Name",
		"object_count":              "ObjectCount",
		"operator":                  "Operator",
		"range":                     "Range",
		"statistic":                 "Statistic",
		"tags":                      "Tags",
		"threshold":                 "Threshold",
		"unit":                      "Unit",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
