// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package databrew

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_databrew_job", jobDataSourceType)
}

// jobDataSourceType returns the Terraform awscc_databrew_job data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DataBrew::Job resource type.
func jobDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"data_catalog_outputs": {
			// Property: DataCatalogOutputs
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "CatalogId": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "DatabaseName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "DatabaseOptions": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "TableName": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "TempDirectory": {
			//             "additionalProperties": false,
			//             "description": "S3 Output location",
			//             "properties": {
			//               "Bucket": {
			//                 "type": "string"
			//               },
			//               "Key": {
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Bucket"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "TableName"
			//         ],
			//         "type": "object"
			//       },
			//       "Overwrite": {
			//         "type": "boolean"
			//       },
			//       "S3Options": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Location": {
			//             "additionalProperties": false,
			//             "description": "S3 Output location",
			//             "properties": {
			//               "Bucket": {
			//                 "type": "string"
			//               },
			//               "Key": {
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Bucket"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Location"
			//         ],
			//         "type": "object"
			//       },
			//       "TableName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "DatabaseName",
			//       "TableName"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"catalog_id": {
						// Property: CatalogId
						Type:     types.StringType,
						Computed: true,
					},
					"database_name": {
						// Property: DatabaseName
						Type:     types.StringType,
						Computed: true,
					},
					"database_options": {
						// Property: DatabaseOptions
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"table_name": {
									// Property: TableName
									Type:     types.StringType,
									Computed: true,
								},
								"temp_directory": {
									// Property: TempDirectory
									Description: "S3 Output location",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Computed: true,
											},
											"key": {
												// Property: Key
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"overwrite": {
						// Property: Overwrite
						Type:     types.BoolType,
						Computed: true,
					},
					"s3_options": {
						// Property: S3Options
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"location": {
									// Property: Location
									Description: "S3 Output location",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Computed: true,
											},
											"key": {
												// Property: Key
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"table_name": {
						// Property: TableName
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"database_outputs": {
			// Property: DatabaseOutputs
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "DatabaseOptions": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "TableName": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "TempDirectory": {
			//             "additionalProperties": false,
			//             "description": "S3 Output location",
			//             "properties": {
			//               "Bucket": {
			//                 "type": "string"
			//               },
			//               "Key": {
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Bucket"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "TableName"
			//         ],
			//         "type": "object"
			//       },
			//       "DatabaseOutputMode": {
			//         "description": "Database table name",
			//         "enum": [
			//           "NEW_TABLE"
			//         ],
			//         "type": "string"
			//       },
			//       "GlueConnectionName": {
			//         "description": "Glue connection name",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "GlueConnectionName",
			//       "DatabaseOptions"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"database_options": {
						// Property: DatabaseOptions
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"table_name": {
									// Property: TableName
									Type:     types.StringType,
									Computed: true,
								},
								"temp_directory": {
									// Property: TempDirectory
									Description: "S3 Output location",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Computed: true,
											},
											"key": {
												// Property: Key
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"database_output_mode": {
						// Property: DatabaseOutputMode
						Description: "Database table name",
						Type:        types.StringType,
						Computed:    true,
					},
					"glue_connection_name": {
						// Property: GlueConnectionName
						Description: "Glue connection name",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"dataset_name": {
			// Property: DatasetName
			// CloudFormation resource type schema:
			// {
			//   "description": "Dataset name",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Dataset name",
			Type:        types.StringType,
			Computed:    true,
		},
		"encryption_key_arn": {
			// Property: EncryptionKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Encryption Key Arn",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Encryption Key Arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"encryption_mode": {
			// Property: EncryptionMode
			// CloudFormation resource type schema:
			// {
			//   "description": "Encryption mode",
			//   "enum": [
			//     "SSE-KMS",
			//     "SSE-S3"
			//   ],
			//   "type": "string"
			// }
			Description: "Encryption mode",
			Type:        types.StringType,
			Computed:    true,
		},
		"job_sample": {
			// Property: JobSample
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Job Sample",
			//   "properties": {
			//     "Mode": {
			//       "description": "Sample configuration mode for profile jobs.",
			//       "enum": [
			//         "FULL_DATASET",
			//         "CUSTOM_ROWS"
			//       ],
			//       "type": "string"
			//     },
			//     "Size": {
			//       "description": "Sample configuration size for profile jobs.",
			//       "format": "int64",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Job Sample",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"mode": {
						// Property: Mode
						Description: "Sample configuration mode for profile jobs.",
						Type:        types.StringType,
						Computed:    true,
					},
					"size": {
						// Property: Size
						Description: "Sample configuration size for profile jobs.",
						Type:        types.NumberType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"log_subscription": {
			// Property: LogSubscription
			// CloudFormation resource type schema:
			// {
			//   "description": "Log subscription",
			//   "enum": [
			//     "ENABLE",
			//     "DISABLE"
			//   ],
			//   "type": "string"
			// }
			Description: "Log subscription",
			Type:        types.StringType,
			Computed:    true,
		},
		"max_capacity": {
			// Property: MaxCapacity
			// CloudFormation resource type schema:
			// {
			//   "description": "Max capacity",
			//   "type": "integer"
			// }
			Description: "Max capacity",
			Type:        types.NumberType,
			Computed:    true,
		},
		"max_retries": {
			// Property: MaxRetries
			// CloudFormation resource type schema:
			// {
			//   "description": "Max retries",
			//   "type": "integer"
			// }
			Description: "Max retries",
			Type:        types.NumberType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Job name",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Job name",
			Type:        types.StringType,
			Computed:    true,
		},
		"output_location": {
			// Property: OutputLocation
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Output location",
			//   "properties": {
			//     "Bucket": {
			//       "type": "string"
			//     },
			//     "Key": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Bucket"
			//   ],
			//   "type": "object"
			// }
			Description: "Output location",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"bucket": {
						// Property: Bucket
						Type:     types.StringType,
						Computed: true,
					},
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"outputs": {
			// Property: Outputs
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "CompressionFormat": {
			//         "enum": [
			//           "GZIP",
			//           "LZ4",
			//           "SNAPPY",
			//           "BZIP2",
			//           "DEFLATE",
			//           "LZO",
			//           "BROTLI",
			//           "ZSTD",
			//           "ZLIB"
			//         ],
			//         "type": "string"
			//       },
			//       "Format": {
			//         "enum": [
			//           "CSV",
			//           "JSON",
			//           "PARQUET",
			//           "GLUEPARQUET",
			//           "AVRO",
			//           "ORC",
			//           "XML"
			//         ],
			//         "type": "string"
			//       },
			//       "FormatOptions": {
			//         "additionalProperties": false,
			//         "description": "Format options for job Output",
			//         "properties": {
			//           "Csv": {
			//             "additionalProperties": false,
			//             "description": "Output Csv options",
			//             "properties": {
			//               "Delimiter": {
			//                 "maxLength": 1,
			//                 "minLength": 1,
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Location": {
			//         "additionalProperties": false,
			//         "description": "S3 Output location",
			//         "properties": {
			//           "Bucket": {
			//             "type": "string"
			//           },
			//           "Key": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Bucket"
			//         ],
			//         "type": "object"
			//       },
			//       "Overwrite": {
			//         "type": "boolean"
			//       },
			//       "PartitionColumns": {
			//         "insertionOrder": true,
			//         "items": {
			//           "type": "string"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       }
			//     },
			//     "required": [
			//       "Location"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"compression_format": {
						// Property: CompressionFormat
						Type:     types.StringType,
						Computed: true,
					},
					"format": {
						// Property: Format
						Type:     types.StringType,
						Computed: true,
					},
					"format_options": {
						// Property: FormatOptions
						Description: "Format options for job Output",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"csv": {
									// Property: Csv
									Description: "Output Csv options",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"delimiter": {
												// Property: Delimiter
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"location": {
						// Property: Location
						Description: "S3 Output location",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket": {
									// Property: Bucket
									Type:     types.StringType,
									Computed: true,
								},
								"key": {
									// Property: Key
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"overwrite": {
						// Property: Overwrite
						Type:     types.BoolType,
						Computed: true,
					},
					"partition_columns": {
						// Property: PartitionColumns
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"profile_configuration": {
			// Property: ProfileConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ColumnStatisticsConfigurations": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Selectors": {
			//             "insertionOrder": true,
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Name": {
			//                   "maxLength": 255,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Regex": {
			//                   "maxLength": 255,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "minItems": 1,
			//             "type": "array"
			//           },
			//           "Statistics": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "IncludedStatistics": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "maxLength": 128,
			//                   "minLength": 1,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "minItems": 1,
			//                 "type": "array"
			//               },
			//               "Overrides": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "additionalProperties": false,
			//                   "properties": {
			//                     "Parameters": {
			//                       "additionalProperties": false,
			//                       "patternProperties": {
			//                         "": {
			//                           "type": "string"
			//                         }
			//                       },
			//                       "type": "object"
			//                     },
			//                     "Statistic": {
			//                       "maxLength": 128,
			//                       "minLength": 1,
			//                       "pattern": "",
			//                       "type": "string"
			//                     }
			//                   },
			//                   "required": [
			//                     "Statistic",
			//                     "Parameters"
			//                   ],
			//                   "type": "object"
			//                 },
			//                 "minItems": 1,
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Statistics"
			//         ],
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "DatasetStatisticsConfiguration": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "IncludedStatistics": {
			//           "insertionOrder": true,
			//           "items": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "Overrides": {
			//           "insertionOrder": true,
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Parameters": {
			//                 "additionalProperties": false,
			//                 "patternProperties": {
			//                   "": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "Statistic": {
			//                 "maxLength": 128,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Statistic",
			//               "Parameters"
			//             ],
			//             "type": "object"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ProfileColumns": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Name": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Regex": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"column_statistics_configurations": {
						// Property: ColumnStatisticsConfigurations
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"selectors": {
									// Property: Selectors
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"name": {
												// Property: Name
												Type:     types.StringType,
												Computed: true,
											},
											"regex": {
												// Property: Regex
												Type:     types.StringType,
												Computed: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Computed: true,
								},
								"statistics": {
									// Property: Statistics
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"included_statistics": {
												// Property: IncludedStatistics
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"overrides": {
												// Property: Overrides
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"parameters": {
															// Property: Parameters
															// Pattern: ""
															Type:     types.MapType{ElemType: types.StringType},
															Computed: true,
														},
														"statistic": {
															// Property: Statistic
															Type:     types.StringType,
															Computed: true,
														},
													},
													tfsdk.ListNestedAttributesOptions{},
												),
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
					"dataset_statistics_configuration": {
						// Property: DatasetStatisticsConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"included_statistics": {
									// Property: IncludedStatistics
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"overrides": {
									// Property: Overrides
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"parameters": {
												// Property: Parameters
												// Pattern: ""
												Type:     types.MapType{ElemType: types.StringType},
												Computed: true,
											},
											"statistic": {
												// Property: Statistic
												Type:     types.StringType,
												Computed: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"profile_columns": {
						// Property: ProfileColumns
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
								"regex": {
									// Property: Regex
									Type:     types.StringType,
									Computed: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"project_name": {
			// Property: ProjectName
			// CloudFormation resource type schema:
			// {
			//   "description": "Project name",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Project name",
			Type:        types.StringType,
			Computed:    true,
		},
		"recipe": {
			// Property: Recipe
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Name": {
			//       "description": "Recipe name",
			//       "type": "string"
			//     },
			//     "Version": {
			//       "description": "Recipe version",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Name"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "Recipe name",
						Type:        types.StringType,
						Computed:    true,
					},
					"version": {
						// Property: Version
						Description: "Recipe version",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role arn",
			//   "type": "string"
			// }
			Description: "Role arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
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
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"timeout": {
			// Property: Timeout
			// CloudFormation resource type schema:
			// {
			//   "description": "Timeout",
			//   "type": "integer"
			// }
			Description: "Timeout",
			Type:        types.NumberType,
			Computed:    true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "Job type",
			//   "enum": [
			//     "PROFILE",
			//     "RECIPE"
			//   ],
			//   "type": "string"
			// }
			Description: "Job type",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::DataBrew::Job",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataBrew::Job").WithTerraformTypeName("awscc_databrew_job")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":                           "Bucket",
		"catalog_id":                       "CatalogId",
		"column_statistics_configurations": "ColumnStatisticsConfigurations",
		"compression_format":               "CompressionFormat",
		"csv":                              "Csv",
		"data_catalog_outputs":             "DataCatalogOutputs",
		"database_name":                    "DatabaseName",
		"database_options":                 "DatabaseOptions",
		"database_output_mode":             "DatabaseOutputMode",
		"database_outputs":                 "DatabaseOutputs",
		"dataset_name":                     "DatasetName",
		"dataset_statistics_configuration": "DatasetStatisticsConfiguration",
		"delimiter":                        "Delimiter",
		"encryption_key_arn":               "EncryptionKeyArn",
		"encryption_mode":                  "EncryptionMode",
		"format":                           "Format",
		"format_options":                   "FormatOptions",
		"glue_connection_name":             "GlueConnectionName",
		"included_statistics":              "IncludedStatistics",
		"job_sample":                       "JobSample",
		"key":                              "Key",
		"location":                         "Location",
		"log_subscription":                 "LogSubscription",
		"max_capacity":                     "MaxCapacity",
		"max_retries":                      "MaxRetries",
		"mode":                             "Mode",
		"name":                             "Name",
		"output_location":                  "OutputLocation",
		"outputs":                          "Outputs",
		"overrides":                        "Overrides",
		"overwrite":                        "Overwrite",
		"parameters":                       "Parameters",
		"partition_columns":                "PartitionColumns",
		"profile_columns":                  "ProfileColumns",
		"profile_configuration":            "ProfileConfiguration",
		"project_name":                     "ProjectName",
		"recipe":                           "Recipe",
		"regex":                            "Regex",
		"role_arn":                         "RoleArn",
		"s3_options":                       "S3Options",
		"selectors":                        "Selectors",
		"size":                             "Size",
		"statistic":                        "Statistic",
		"statistics":                       "Statistics",
		"table_name":                       "TableName",
		"tags":                             "Tags",
		"temp_directory":                   "TempDirectory",
		"timeout":                          "Timeout",
		"type":                             "Type",
		"value":                            "Value",
		"version":                          "Version",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}