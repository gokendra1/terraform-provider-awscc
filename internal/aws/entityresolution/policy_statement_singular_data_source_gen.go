// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package entityresolution

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_entityresolution_policy_statement", policyStatementDataSource)
}

// policyStatementDataSource returns the Terraform awscc_entityresolution_policy_statement data source.
// This Terraform data source corresponds to the CloudFormation AWS::EntityResolution::PolicyStatement resource.
func policyStatementDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "maxLength": 64,
		//	    "minLength": 3,
		//	    "pattern": "^(entityresolution:[a-zA-Z0-9]+)$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"action": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn of the resource to which the policy statement is being attached.",
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:((schemamapping|matchingworkflow|idmappingworkflow|idnamespace)/[a-zA-Z_0-9-]{1,255})$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn of the resource to which the policy statement is being attached.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Condition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 40960,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"condition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Effect
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "Allow",
		//	    "Deny"
		//	  ],
		//	  "type": "string"
		//	}
		"effect": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Principal
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "maxLength": 64,
		//	    "minLength": 12,
		//	    "pattern": "^(\\\\d{12})|([a-z0-9\\\\.]+)$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"principal": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StatementId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Statement Id of the policy statement that is being attached.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Za-z]+$",
		//	  "type": "string"
		//	}
		"statement_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Statement Id of the policy statement that is being attached.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EntityResolution::PolicyStatement",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EntityResolution::PolicyStatement").WithTerraformTypeName("awscc_entityresolution_policy_statement")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":       "Action",
		"arn":          "Arn",
		"condition":    "Condition",
		"effect":       "Effect",
		"principal":    "Principal",
		"statement_id": "StatementId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}