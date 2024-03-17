// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_transit_gateway_peering_attachment", transitGatewayPeeringAttachmentDataSource)
}

// transitGatewayPeeringAttachmentDataSource returns the Terraform awscc_ec2_transit_gateway_peering_attachment data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::TransitGatewayPeeringAttachment resource.
func transitGatewayPeeringAttachmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time the transit gateway peering attachment was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The time the transit gateway peering attachment was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PeerAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the peer account",
		//	  "type": "string"
		//	}
		"peer_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the peer account",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PeerRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Peer Region",
		//	  "type": "string"
		//	}
		"peer_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Peer Region",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PeerTransitGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the peer transit gateway.",
		//	  "type": "string"
		//	}
		"peer_transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the peer transit gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The status of the transit gateway peering attachment.",
		//	  "properties": {
		//	    "Code": {
		//	      "description": "The status code.",
		//	      "type": "string"
		//	    },
		//	    "Message": {
		//	      "description": "The status message, if applicable.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"status": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Code
				"code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The status code.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Message
				"message": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The status message, if applicable.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The status of the transit gateway peering attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags for the transit gateway peering attachment.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags for the transit gateway peering attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayAttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway peering attachment.",
		//	  "type": "string"
		//	}
		"transit_gateway_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway peering attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway.",
		//	  "type": "string"
		//	}
		"transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::TransitGatewayPeeringAttachment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayPeeringAttachment").WithTerraformTypeName("awscc_ec2_transit_gateway_peering_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"code":                          "Code",
		"creation_time":                 "CreationTime",
		"key":                           "Key",
		"message":                       "Message",
		"peer_account_id":               "PeerAccountId",
		"peer_region":                   "PeerRegion",
		"peer_transit_gateway_id":       "PeerTransitGatewayId",
		"state":                         "State",
		"status":                        "Status",
		"tags":                          "Tags",
		"transit_gateway_attachment_id": "TransitGatewayAttachmentId",
		"transit_gateway_id":            "TransitGatewayId",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
