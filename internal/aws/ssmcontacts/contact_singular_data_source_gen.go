// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssmcontacts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ssmcontacts_contact", contactDataSource)
}

// contactDataSource returns the Terraform awscc_ssmcontacts_contact data source.
// This Terraform data source corresponds to the CloudFormation AWS::SSMContacts::Contact resource.
func contactDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Alias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_\\-\\.]*$",
		//	  "type": "string"
		//	}
		"alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the contact.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the contact.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\s]*$",
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Plan
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A set amount of time that an escalation plan or engagement plan engages the specified contacts or contact methods.",
		//	    "oneOf": [
		//	      {
		//	        "required": [
		//	          "DurationInMinutes"
		//	        ]
		//	      },
		//	      {
		//	        "required": [
		//	          "RotationIds"
		//	        ]
		//	      }
		//	    ],
		//	    "properties": {
		//	      "DurationInMinutes": {
		//	        "description": "The time to wait until beginning the next stage.",
		//	        "type": "integer"
		//	      },
		//	      "RotationIds": {
		//	        "description": "List of Rotation Ids to associate with Contact",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "Targets": {
		//	        "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
		//	          "oneOf": [
		//	            {
		//	              "required": [
		//	                "ChannelTargetInfo"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "ContactTargetInfo"
		//	              ]
		//	            }
		//	          ],
		//	          "properties": {
		//	            "ChannelTargetInfo": {
		//	              "additionalProperties": false,
		//	              "description": "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
		//	              "properties": {
		//	                "ChannelId": {
		//	                  "description": "The Amazon Resource Name (ARN) of the contact channel.",
		//	                  "type": "string"
		//	                },
		//	                "RetryIntervalInMinutes": {
		//	                  "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "required": [
		//	                "ChannelId",
		//	                "RetryIntervalInMinutes"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "ContactTargetInfo": {
		//	              "additionalProperties": false,
		//	              "description": "The contact that SSM Incident Manager is engaging during an incident.",
		//	              "properties": {
		//	                "ContactId": {
		//	                  "description": "The Amazon Resource Name (ARN) of the contact.",
		//	                  "type": "string"
		//	                },
		//	                "IsEssential": {
		//	                  "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "required": [
		//	                "ContactId",
		//	                "IsEssential"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"plan": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DurationInMinutes
					"duration_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The time to wait until beginning the next stage.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: RotationIds
					"rotation_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "List of Rotation Ids to associate with Contact",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Targets
					"targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ChannelTargetInfo
								"channel_target_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ChannelId
										"channel_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The Amazon Resource Name (ARN) of the contact channel.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: RetryIntervalInMinutes
										"retry_interval_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Description: "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: ContactTargetInfo
								"contact_target_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ContactId
										"contact_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The Amazon Resource Name (ARN) of the contact.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: IsEssential
										"is_essential": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Description: "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "The contact that SSM Incident Manager is engaging during an incident.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
		//	  "enum": [
		//	    "PERSONAL",
		//	    "CUSTOM",
		//	    "SERVICE",
		//	    "ESCALATION",
		//	    "ONCALL_SCHEDULE"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SSMContacts::Contact",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMContacts::Contact").WithTerraformTypeName("awscc_ssmcontacts_contact")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":                     "Alias",
		"arn":                       "Arn",
		"channel_id":                "ChannelId",
		"channel_target_info":       "ChannelTargetInfo",
		"contact_id":                "ContactId",
		"contact_target_info":       "ContactTargetInfo",
		"display_name":              "DisplayName",
		"duration_in_minutes":       "DurationInMinutes",
		"is_essential":              "IsEssential",
		"plan":                      "Plan",
		"retry_interval_in_minutes": "RetryIntervalInMinutes",
		"rotation_ids":              "RotationIds",
		"targets":                   "Targets",
		"type":                      "Type",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
