// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package amplify

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_amplify_branch", branchResource)
}

// branchResource returns the Terraform awscc_amplify_branch resource.
// This Terraform resource corresponds to the CloudFormation AWS::Amplify::Branch resource.
func branchResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 20,
		//	  "minLength": 1,
		//	  "pattern": "d[a-z0-9]+",
		//	  "type": "string"
		//	}
		"app_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 20),
				stringvalidator.RegexMatches(regexp.MustCompile("d[a-z0-9]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Backend
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "StackArn": {
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"backend": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: StackArn
				"stack_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(20, 2048),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BasicAuthConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EnableBasicAuth": {
		//	      "type": "boolean"
		//	    },
		//	    "Password": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "Username": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Username",
		//	    "Password"
		//	  ],
		//	  "type": "object"
		//	}
		"basic_auth_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableBasicAuth
				"enable_basic_auth": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Password
				"password": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 255),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Username
				"username": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 255),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// BasicAuthConfig is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: BranchName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "(?s).+",
		//	  "type": "string"
		//	}
		"branch_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("(?s).+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BuildSpec
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 25000,
		//	  "minLength": 1,
		//	  "pattern": "(?s).+",
		//	  "type": "string"
		//	}
		"build_spec": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 25000),
				stringvalidator.RegexMatches(regexp.MustCompile("(?s).+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1000),
				stringvalidator.RegexMatches(regexp.MustCompile("(?s).*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableAutoBuild
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_auto_build": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnablePerformanceMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_performance_mode": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnablePullRequestPreview
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_pull_request_preview": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentVariables
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Name": {
		//	        "maxLength": 255,
		//	        "pattern": "(?s).*",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 5500,
		//	        "pattern": "(?s).*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"environment_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile("(?s).*"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(5500),
							stringvalidator.RegexMatches(regexp.MustCompile("(?s).*"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Framework
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"framework": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(255),
				stringvalidator.RegexMatches(regexp.MustCompile("(?s).*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PullRequestEnvironmentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 20,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"pull_request_environment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(20),
				stringvalidator.RegexMatches(regexp.MustCompile("(?s).*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Stage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "EXPERIMENTAL",
		//	    "BETA",
		//	    "PULL_REQUEST",
		//	    "PRODUCTION",
		//	    "DEVELOPMENT"
		//	  ],
		//	  "type": "string"
		//	}
		"stage": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"EXPERIMENTAL",
					"BETA",
					"PULL_REQUEST",
					"PRODUCTION",
					"DEVELOPMENT",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "insertionOrder": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
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
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::Amplify::Branch resource creates a new branch within an app.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Branch").WithTerraformTypeName("awscc_amplify_branch")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_id":                        "AppId",
		"arn":                           "Arn",
		"backend":                       "Backend",
		"basic_auth_config":             "BasicAuthConfig",
		"branch_name":                   "BranchName",
		"build_spec":                    "BuildSpec",
		"description":                   "Description",
		"enable_auto_build":             "EnableAutoBuild",
		"enable_basic_auth":             "EnableBasicAuth",
		"enable_performance_mode":       "EnablePerformanceMode",
		"enable_pull_request_preview":   "EnablePullRequestPreview",
		"environment_variables":         "EnvironmentVariables",
		"framework":                     "Framework",
		"key":                           "Key",
		"name":                          "Name",
		"password":                      "Password",
		"pull_request_environment_name": "PullRequestEnvironmentName",
		"stack_arn":                     "StackArn",
		"stage":                         "Stage",
		"tags":                          "Tags",
		"username":                      "Username",
		"value":                         "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BasicAuthConfig",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
