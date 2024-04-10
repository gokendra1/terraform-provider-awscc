// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package securityhub

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_securityhub_delegated_admin", delegatedAdminResource)
}

// delegatedAdminResource returns the Terraform awscc_securityhub_delegated_admin resource.
// This Terraform resource corresponds to the CloudFormation AWS::SecurityHub::DelegatedAdmin resource.
func delegatedAdminResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdminAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Web Services account identifier of the account to designate as the Security Hub administrator account",
		//	  "pattern": "^[0-9]{12}$",
		//	  "type": "string"
		//	}
		"admin_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Web Services account identifier of the account to designate as the Security Hub administrator account",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{12}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DelegatedAdminIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the DelegatedAdmin being created and assigned as the unique identifier",
		//	  "pattern": "^[0-9]{12}/[a-zA-Z0-9-]{1,32}$",
		//	  "type": "string"
		//	}
		"delegated_admin_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the DelegatedAdmin being created and assigned as the unique identifier",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current status of the Security Hub administrator account. Indicates whether the account is currently enabled as a Security Hub administrator",
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLE_IN_PROGRESS"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current status of the Security Hub administrator account. Indicates whether the account is currently enabled as a Security Hub administrator",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::SecurityHub::DelegatedAdmin resource represents the AWS Security Hub delegated admin account in your organization. One delegated admin resource is allowed to create for the organization in each region in which you configure the AdminAccountId.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecurityHub::DelegatedAdmin").WithTerraformTypeName("awscc_securityhub_delegated_admin")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_account_id":           "AdminAccountId",
		"delegated_admin_identifier": "DelegatedAdminIdentifier",
		"status":                     "Status",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}