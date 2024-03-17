// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_route53resolver_firewall_rule_group", firewallRuleGroupResource)
}

// firewallRuleGroupResource returns the Terraform awscc_route53resolver_firewall_rule_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::Route53Resolver::FirewallRuleGroup resource.
func firewallRuleGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn",
		//	  "maxLength": 600,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Rfc3339TimeString",
		//	  "maxLength": 40,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Rfc3339TimeString",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatorRequestId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the creator request.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"creator_request_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the creator request.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FirewallRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FirewallRules",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Firewall Rule associating the Rule Group to a Domain List",
		//	    "properties": {
		//	      "Action": {
		//	        "description": "Rule Action",
		//	        "enum": [
		//	          "ALLOW",
		//	          "BLOCK",
		//	          "ALERT"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "BlockOverrideDnsType": {
		//	        "description": "BlockOverrideDnsType",
		//	        "enum": [
		//	          "CNAME"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "BlockOverrideDomain": {
		//	        "description": "BlockOverrideDomain",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "BlockOverrideTtl": {
		//	        "description": "BlockOverrideTtl",
		//	        "maximum": 604800,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "BlockResponse": {
		//	        "description": "BlockResponse",
		//	        "enum": [
		//	          "NODATA",
		//	          "NXDOMAIN",
		//	          "OVERRIDE"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "FirewallDomainListId": {
		//	        "description": "ResourceId",
		//	        "maxLength": 64,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Priority": {
		//	        "description": "Rule Priority",
		//	        "type": "integer"
		//	      },
		//	      "Qtype": {
		//	        "description": "Qtype",
		//	        "maxLength": 16,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "FirewallDomainListId",
		//	      "Priority",
		//	      "Action"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"firewall_rules": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Action
					"action": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Rule Action",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"ALLOW",
								"BLOCK",
								"ALERT",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: BlockOverrideDnsType
					"block_override_dns_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "BlockOverrideDnsType",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"CNAME",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: BlockOverrideDomain
					"block_override_domain": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "BlockOverrideDomain",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 255),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: BlockOverrideTtl
					"block_override_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "BlockOverrideTtl",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(0, 604800),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: BlockResponse
					"block_response": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "BlockResponse",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"NODATA",
								"NXDOMAIN",
								"OVERRIDE",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: FirewallDomainListId
					"firewall_domain_list_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "ResourceId",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 64),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Priority
					"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "Rule Priority",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Qtype
					"qtype": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Qtype",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 16),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "FirewallRules",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResourceId",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResourceId",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModificationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Rfc3339TimeString",
		//	  "maxLength": 40,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"modification_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Rfc3339TimeString",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FirewallRuleGroupName",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FirewallRuleGroupName",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AccountId",
		//	  "maxLength": 32,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AccountId",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RuleCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Count",
		//	  "type": "integer"
		//	}
		"rule_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Count",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ShareStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
		//	  "enum": [
		//	    "NOT_SHARED",
		//	    "SHARED_WITH_ME",
		//	    "SHARED_BY_ME"
		//	  ],
		//	  "type": "string"
		//	}
		"share_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
		//	  "enum": [
		//	    "COMPLETE",
		//	    "DELETING",
		//	    "UPDATING",
		//	    "INACTIVE_OWNER_ACCOUNT_CLOSED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StatusMessage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FirewallRuleGroupStatus",
		//	  "type": "string"
		//	}
		"status_message": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FirewallRuleGroupStatus",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 255,
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 127),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 255),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource schema for AWS::Route53Resolver::FirewallRuleGroup.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::FirewallRuleGroup").WithTerraformTypeName("awscc_route53resolver_firewall_rule_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                  "Action",
		"arn":                     "Arn",
		"block_override_dns_type": "BlockOverrideDnsType",
		"block_override_domain":   "BlockOverrideDomain",
		"block_override_ttl":      "BlockOverrideTtl",
		"block_response":          "BlockResponse",
		"creation_time":           "CreationTime",
		"creator_request_id":      "CreatorRequestId",
		"firewall_domain_list_id": "FirewallDomainListId",
		"firewall_rules":          "FirewallRules",
		"id":                      "Id",
		"key":                     "Key",
		"modification_time":       "ModificationTime",
		"name":                    "Name",
		"owner_id":                "OwnerId",
		"priority":                "Priority",
		"qtype":                   "Qtype",
		"rule_count":              "RuleCount",
		"share_status":            "ShareStatus",
		"status":                  "Status",
		"status_message":          "StatusMessage",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
