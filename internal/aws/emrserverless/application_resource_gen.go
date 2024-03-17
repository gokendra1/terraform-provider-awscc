// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package emrserverless

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_emrserverless_application", applicationResource)
}

// applicationResource returns the Terraform awscc_emrserverless_application resource.
// This Terraform resource corresponds to the CloudFormation AWS::EMRServerless::Application resource.
func applicationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the EMR Serverless Application.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"application_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the EMR Serverless Application.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Architecture
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The cpu architecture of an application.",
		//	  "enum": [
		//	    "ARM64",
		//	    "X86_64"
		//	  ],
		//	  "type": "string"
		//	}
		"architecture": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The cpu architecture of an application.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ARM64",
					"X86_64",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the EMR Serverless Application.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the EMR Serverless Application.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoStartConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration for Auto Start of Application.",
		//	  "properties": {
		//	    "Enabled": {
		//	      "default": true,
		//	      "description": "If set to true, the Application will automatically start. Defaults to true.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"auto_start_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If set to true, the Application will automatically start. Defaults to true.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						generic.BoolDefaultValue(true),
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration for Auto Start of Application.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoStopConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration for Auto Stop of Application.",
		//	  "properties": {
		//	    "Enabled": {
		//	      "default": true,
		//	      "description": "If set to true, the Application will automatically stop after being idle. Defaults to true.",
		//	      "type": "boolean"
		//	    },
		//	    "IdleTimeoutMinutes": {
		//	      "description": "The amount of time [in minutes] to wait before auto stopping the Application when idle. Defaults to 15 minutes.",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"auto_stop_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If set to true, the Application will automatically stop after being idle. Defaults to true.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						generic.BoolDefaultValue(true),
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IdleTimeoutMinutes
				"idle_timeout_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The amount of time [in minutes] to wait before auto stopping the Application when idle. Defaults to 15 minutes.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration for Auto Stop of Application.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImageConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The image configuration.",
		//	  "properties": {
		//	    "ImageUri": {
		//	      "description": "The URI of an image in the Amazon ECR registry. This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.",
		//	      "maxLength": 1024,
		//	      "minLength": 1,
		//	      "pattern": "^([a-z0-9]+[a-z0-9-.]*)\\/((?:[a-z0-9]+(?:[._-][a-z0-9]+)*\\/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)(?:\\:([a-zA-Z0-9_][a-zA-Z0-9-._]{0,299})|@(sha256:[0-9a-f]{64}))$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"image_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ImageUri
				"image_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The URI of an image in the Amazon ECR registry. This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 1024),
						stringvalidator.RegexMatches(regexp.MustCompile("^([a-z0-9]+[a-z0-9-.]*)\\/((?:[a-z0-9]+(?:[._-][a-z0-9]+)*\\/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)(?:\\:([a-zA-Z0-9_][a-zA-Z0-9-._]{0,299})|@(sha256:[0-9a-f]{64}))$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The image configuration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InitialCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Initial capacity initialized when an Application is started.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "Worker type for an analytics framework.",
		//	        "maxLength": 50,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z]+[-_]*[a-zA-Z]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "WorkerConfiguration": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Cpu": {
		//	                "description": "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
		//	                "maxLength": 15,
		//	                "minLength": 1,
		//	                "pattern": "^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$",
		//	                "type": "string"
		//	              },
		//	              "Disk": {
		//	                "description": "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
		//	                "maxLength": 15,
		//	                "minLength": 1,
		//	                "pattern": "^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$",
		//	                "type": "string"
		//	              },
		//	              "Memory": {
		//	                "description": "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
		//	                "maxLength": 15,
		//	                "minLength": 1,
		//	                "pattern": "^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Cpu",
		//	              "Memory"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "WorkerCount": {
		//	            "description": "Initial count of workers to be initialized when an Application is started. This count will be continued to be maintained until the Application is stopped",
		//	            "format": "int64",
		//	            "maximum": 1000000,
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "WorkerCount",
		//	          "WorkerConfiguration"
		//	        ],
		//	        "type": "object"
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
		"initial_capacity": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Worker type for an analytics framework.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 50),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z]+[-_]*[a-zA-Z]+$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: WorkerConfiguration
							"worker_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Cpu
									"cpu": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
										Required:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 15),
											stringvalidator.RegexMatches(regexp.MustCompile("^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$"), ""),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
									// Property: Disk
									"disk": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
										Optional:    true,
										Computed:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 15),
											stringvalidator.RegexMatches(regexp.MustCompile("^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$"), ""),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Memory
									"memory": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
										Required:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 15),
											stringvalidator.RegexMatches(regexp.MustCompile("^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$"), ""),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Required: true,
							}, /*END ATTRIBUTE*/
							// Property: WorkerCount
							"worker_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Initial count of workers to be initialized when an Application is started. This count will be continued to be maintained until the Application is stopped",
								Required:    true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(1, 1000000),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Initial capacity initialized when an Application is started.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaximumCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.",
		//	  "properties": {
		//	    "Cpu": {
		//	      "description": "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
		//	      "maxLength": 15,
		//	      "minLength": 1,
		//	      "pattern": "^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$",
		//	      "type": "string"
		//	    },
		//	    "Disk": {
		//	      "description": "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
		//	      "maxLength": 15,
		//	      "minLength": 1,
		//	      "pattern": "^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$",
		//	      "type": "string"
		//	    },
		//	    "Memory": {
		//	      "description": "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
		//	      "maxLength": 15,
		//	      "minLength": 1,
		//	      "pattern": "^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Cpu",
		//	    "Memory"
		//	  ],
		//	  "type": "object"
		//	}
		"maximum_capacity": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Cpu
				"cpu": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 15),
						stringvalidator.RegexMatches(regexp.MustCompile("^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Disk
				"disk": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 15),
						stringvalidator.RegexMatches(regexp.MustCompile("^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Memory
				"memory": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 15),
						stringvalidator.RegexMatches(regexp.MustCompile("^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "User friendly Application name.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9._\\/#-]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "User friendly Application name.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9._\\/#-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Network Configuration for customer VPC connectivity.",
		//	  "properties": {
		//	    "SecurityGroupIds": {
		//	      "description": "The ID of the security groups in the VPC to which you want to connect your job or application.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Identifier of a security group",
		//	        "maxLength": 32,
		//	        "minLength": 1,
		//	        "pattern": "^[-0-9a-zA-Z]+",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "SubnetIds": {
		//	      "description": "The ID of the subnets in the VPC to which you want to connect your job or application.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Identifier of a subnet",
		//	        "maxLength": 32,
		//	        "minLength": 1,
		//	        "pattern": "^[-0-9a-zA-Z]+",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 16,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"network_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The ID of the security groups in the VPC to which you want to connect your job or application.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 5),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 32),
							stringvalidator.RegexMatches(regexp.MustCompile("^[-0-9a-zA-Z]+"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The ID of the subnets in the VPC to which you want to connect your job or application.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 16),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 32),
							stringvalidator.RegexMatches(regexp.MustCompile("^[-0-9a-zA-Z]+"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Network Configuration for customer VPC connectivity.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReleaseLabel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "EMR release label.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9._/-]+$",
		//	  "type": "string"
		//	}
		"release_label": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "EMR release label.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9._/-]+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tag map with key and value",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 128 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^[A-Za-z0-9 /_.:=+@-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^[A-Za-z0-9 /_.:=+@-]*$",
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
						Description: "The value for the tag. You can specify a value that is 1 to 128 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9 /_.:=+@-]+$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9 /_.:=+@-]*$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tag map with key and value",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the application",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the application",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WorkerTypeSpecifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The key-value pairs that specify worker type to WorkerTypeSpecificationInput. This parameter must contain all valid worker types for a Spark or Hive application. Valid worker types include Driver and Executor for Spark applications and HiveDriver and TezTask for Hive applications. You can either set image details in this parameter for each worker type, or in imageConfiguration for all worker types.",
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "description": "The specifications for a worker type.",
		//	      "properties": {
		//	        "ImageConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "The image configuration.",
		//	          "properties": {
		//	            "ImageUri": {
		//	              "description": "The URI of an image in the Amazon ECR registry. This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.",
		//	              "maxLength": 1024,
		//	              "minLength": 1,
		//	              "pattern": "^([a-z0-9]+[a-z0-9-.]*)\\/((?:[a-z0-9]+(?:[._-][a-z0-9]+)*\\/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)(?:\\:([a-zA-Z0-9_][a-zA-Z0-9-._]{0,299})|@(sha256:[0-9a-f]{64}))$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"worker_type_specifications": // Pattern: ""
		schema.MapNestedAttribute{    /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ImageConfiguration
					"image_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ImageUri
							"image_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The URI of an image in the Amazon ECR registry. This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 1024),
									stringvalidator.RegexMatches(regexp.MustCompile("^([a-z0-9]+[a-z0-9-.]*)\\/((?:[a-z0-9]+(?:[._-][a-z0-9]+)*\\/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)(?:\\:([a-zA-Z0-9_][a-zA-Z0-9-._]{0,299})|@(sha256:[0-9a-f]{64}))$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The image configuration.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The key-value pairs that specify worker type to WorkerTypeSpecificationInput. This parameter must contain all valid worker types for a Spark or Hive application. Valid worker types include Driver and Executor for Spark applications and HiveDriver and TezTask for Hive applications. You can either set image details in this parameter for each worker type, or in imageConfiguration for all worker types.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
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
		Description: "Resource schema for AWS::EMRServerless::Application Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMRServerless::Application").WithTerraformTypeName("awscc_emrserverless_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_id":             "ApplicationId",
		"architecture":               "Architecture",
		"arn":                        "Arn",
		"auto_start_configuration":   "AutoStartConfiguration",
		"auto_stop_configuration":    "AutoStopConfiguration",
		"cpu":                        "Cpu",
		"disk":                       "Disk",
		"enabled":                    "Enabled",
		"idle_timeout_minutes":       "IdleTimeoutMinutes",
		"image_configuration":        "ImageConfiguration",
		"image_uri":                  "ImageUri",
		"initial_capacity":           "InitialCapacity",
		"key":                        "Key",
		"maximum_capacity":           "MaximumCapacity",
		"memory":                     "Memory",
		"name":                       "Name",
		"network_configuration":      "NetworkConfiguration",
		"release_label":              "ReleaseLabel",
		"security_group_ids":         "SecurityGroupIds",
		"subnet_ids":                 "SubnetIds",
		"tags":                       "Tags",
		"type":                       "Type",
		"value":                      "Value",
		"worker_configuration":       "WorkerConfiguration",
		"worker_count":               "WorkerCount",
		"worker_type_specifications": "WorkerTypeSpecifications",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
