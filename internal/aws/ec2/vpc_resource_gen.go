// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_vpc", vPCResourceType)
}

// vPCResourceType returns the Terraform awscc_ec2_vpc resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::VPC resource type.
func vPCResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cidr_block": {
			// Property: CidrBlock
			// CloudFormation resource type schema:
			// {
			//   "description": "The primary IPv4 CIDR block for the VPC.",
			//   "type": "string"
			// }
			Description: "The primary IPv4 CIDR block for the VPC.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"cidr_block_associations": {
			// Property: CidrBlockAssociations
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of IPv4 CIDR block association IDs for the VPC.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of IPv4 CIDR block association IDs for the VPC.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"default_network_acl": {
			// Property: DefaultNetworkAcl
			// CloudFormation resource type schema:
			// {
			//   "description": "The default network ACL ID that is associated with the VPC.",
			//   "insertionOrder": false,
			//   "type": "string"
			// }
			Description: "The default network ACL ID that is associated with the VPC.",
			Type:        types.StringType,
			Computed:    true,
		},
		"default_security_group": {
			// Property: DefaultSecurityGroup
			// CloudFormation resource type schema:
			// {
			//   "description": "The default security group ID that is associated with the VPC.",
			//   "insertionOrder": false,
			//   "type": "string"
			// }
			Description: "The default security group ID that is associated with the VPC.",
			Type:        types.StringType,
			Computed:    true,
		},
		"enable_dns_hostnames": {
			// Property: EnableDnsHostnames
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"enable_dns_support": {
			// Property: EnableDnsSupport
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range \"plus two\" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range \"plus two\" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The Id for the model.",
			//   "type": "string"
			// }
			Description: "The Id for the model.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_tenancy": {
			// Property: InstanceTenancy
			// CloudFormation resource type schema:
			// {
			//   "description": "The allowed tenancy of instances launched into the VPC.\n\n\"default\": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.\n\n\"dedicated\": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.\n\nUpdating InstanceTenancy requires no replacement only if you are updating its value from \"dedicated\" to \"default\". Updating InstanceTenancy from \"default\" to \"dedicated\" requires replacement.",
			//   "type": "string"
			// }
			Description: "The allowed tenancy of instances launched into the VPC.\n\n\"default\": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.\n\n\"dedicated\": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.\n\nUpdating InstanceTenancy requires no replacement only if you are updating its value from \"dedicated\" to \"default\". Updating InstanceTenancy from \"default\" to \"dedicated\" requires replacement.",
			Type:        types.StringType,
			Optional:    true,
		},
		"ipv_6_cidr_blocks": {
			// Property: Ipv6CidrBlocks
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of IPv6 CIDR blocks that are associated with the VPC.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of IPv6 CIDR blocks that are associated with the VPC.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the VPC.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "The tags for the VPC.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::EC2::VPC",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPC").WithTerraformTypeName("awscc_ec2_vpc")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cidr_block":              "CidrBlock",
		"cidr_block_associations": "CidrBlockAssociations",
		"default_network_acl":     "DefaultNetworkAcl",
		"default_security_group":  "DefaultSecurityGroup",
		"enable_dns_hostnames":    "EnableDnsHostnames",
		"enable_dns_support":      "EnableDnsSupport",
		"id":                      "Id",
		"instance_tenancy":        "InstanceTenancy",
		"ipv_6_cidr_blocks":       "Ipv6CidrBlocks",
		"key":                     "Key",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}