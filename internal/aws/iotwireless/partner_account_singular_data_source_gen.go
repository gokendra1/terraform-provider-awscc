// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotwireless_partner_account", partnerAccountDataSource)
}

// partnerAccountDataSource returns the Terraform awscc_iotwireless_partner_account data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTWireless::PartnerAccount resource.
func partnerAccountDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountLinked
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether the partner account is linked to the AWS account.",
		//	  "type": "boolean"
		//	}
		"account_linked": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether the partner account is linked to the AWS account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "PartnerAccount arn. Returned after successful create.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "PartnerAccount arn. Returned after successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Fingerprint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The fingerprint of the Sidewalk application server private key.",
		//	  "type": "string"
		//	}
		"fingerprint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The fingerprint of the Sidewalk application server private key.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PartnerAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The partner account ID to disassociate from the AWS account",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"partner_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The partner account ID to disassociate from the AWS account",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PartnerType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The partner type",
		//	  "enum": [
		//	    "Sidewalk"
		//	  ],
		//	  "type": "string"
		//	}
		"partner_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The partner type",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Sidewalk
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Sidewalk account credentials.",
		//	  "properties": {
		//	    "AppServerPrivateKey": {
		//	      "maxLength": 4096,
		//	      "minLength": 1,
		//	      "pattern": "[a-fA-F0-9]{64}",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "AppServerPrivateKey"
		//	  ],
		//	  "type": "object"
		//	}
		"sidewalk": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AppServerPrivateKey
				"app_server_private_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Sidewalk account credentials.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SidewalkResponse
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Sidewalk account credentials.",
		//	  "properties": {
		//	    "AmazonId": {
		//	      "maxLength": 2048,
		//	      "type": "string"
		//	    },
		//	    "Arn": {
		//	      "type": "string"
		//	    },
		//	    "Fingerprint": {
		//	      "maxLength": 64,
		//	      "minLength": 64,
		//	      "pattern": "[a-fA-F0-9]{64}",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sidewalk_response": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AmazonId
				"amazon_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Arn
				"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Fingerprint
				"fingerprint": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Sidewalk account credentials.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SidewalkUpdate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Sidewalk account credentials.",
		//	  "properties": {
		//	    "AppServerPrivateKey": {
		//	      "maxLength": 4096,
		//	      "minLength": 1,
		//	      "pattern": "[a-fA-F0-9]{64}",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sidewalk_update": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AppServerPrivateKey
				"app_server_private_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Sidewalk account credentials.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the destination.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the destination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTWireless::PartnerAccount",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::PartnerAccount").WithTerraformTypeName("awscc_iotwireless_partner_account")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_linked":         "AccountLinked",
		"amazon_id":              "AmazonId",
		"app_server_private_key": "AppServerPrivateKey",
		"arn":                    "Arn",
		"fingerprint":            "Fingerprint",
		"key":                    "Key",
		"partner_account_id":     "PartnerAccountId",
		"partner_type":           "PartnerType",
		"sidewalk":               "Sidewalk",
		"sidewalk_response":      "SidewalkResponse",
		"sidewalk_update":        "SidewalkUpdate",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
