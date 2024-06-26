// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3outposts_endpoint", endpointDataSource)
}

// endpointDataSource returns the Terraform awscc_s3outposts_endpoint data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3Outposts::Endpoint resource.
func endpointDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "Private",
		//	  "description": "The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.",
		//	  "enum": [
		//	    "CustomerOwnedIp",
		//	    "Private"
		//	  ],
		//	  "type": "string"
		//	}
		"access_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the endpoint.",
		//	  "maxLength": 500,
		//	  "minLength": 5,
		//	  "pattern": "^arn:[^:]+:s3-outposts:[a-zA-Z0-9\\-]+:\\d{12}:outpost\\/[^:]+\\/endpoint/[a-zA-Z0-9]{19}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The VPC CIDR committed by this endpoint.",
		//	  "maxLength": 20,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The VPC CIDR committed by this endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time the endpoint was created.",
		//	  "pattern": "^([0-2]\\d{3})-(0[0-9]|1[0-2])-([0-2]\\d|3[01])T([01]\\d|2[0-4]):([0-5]\\d):([0-6]\\d)((\\.\\d{3})?)Z$",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time the endpoint was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomerOwnedIpv4Pool
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.",
		//	  "pattern": "^ipv4pool-coip-([0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"customer_owned_ipv_4_pool": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FailedReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The failure reason, if any, for a create or delete endpoint operation.",
		//	  "properties": {
		//	    "ErrorCode": {
		//	      "description": "The failure code, if any, for a create or delete endpoint operation.",
		//	      "type": "string"
		//	    },
		//	    "Message": {
		//	      "description": "Additional error details describing the endpoint failure and recommended action.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"failed_reason": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ErrorCode
				"error_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The failure code, if any, for a create or delete endpoint operation.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Message
				"message": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Additional error details describing the endpoint failure and recommended action.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The failure reason, if any, for a create or delete endpoint operation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the endpoint.",
		//	  "maxLength": 500,
		//	  "minLength": 5,
		//	  "pattern": "^[a-zA-Z0-9]{19}$",
		//	  "type": "string"
		//	}
		"endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaces
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The network interfaces of the endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The container for the network interface.",
		//	    "properties": {
		//	      "NetworkInterfaceId": {
		//	        "maxLength": 100,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "NetworkInterfaceId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"network_interfaces": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: NetworkInterfaceId
					"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The network interfaces of the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the customer outpost on which the bucket resides.",
		//	  "pattern": "^(op-[a-f0-9]{17}|\\d{12}|ec2)$",
		//	  "type": "string"
		//	}
		"outpost_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the customer outpost on which the bucket resides.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the security group to use with the endpoint.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^sg-([0-9a-f]{8}|[0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the security group to use with the endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "Available",
		//	    "Pending",
		//	    "Deleting",
		//	    "Create_Failed",
		//	    "Delete_Failed"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^subnet-([0-9a-f]{8}|[0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3Outposts::Endpoint",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::Endpoint").WithTerraformTypeName("awscc_s3outposts_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_type":               "AccessType",
		"arn":                       "Arn",
		"cidr_block":                "CidrBlock",
		"creation_time":             "CreationTime",
		"customer_owned_ipv_4_pool": "CustomerOwnedIpv4Pool",
		"endpoint_id":               "Id",
		"error_code":                "ErrorCode",
		"failed_reason":             "FailedReason",
		"message":                   "Message",
		"network_interface_id":      "NetworkInterfaceId",
		"network_interfaces":        "NetworkInterfaces",
		"outpost_id":                "OutpostId",
		"security_group_id":         "SecurityGroupId",
		"status":                    "Status",
		"subnet_id":                 "SubnetId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
