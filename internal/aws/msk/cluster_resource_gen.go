// Code generated by generators/resource/main.go; DO NOT EDIT.

package msk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_msk_cluster", clusterResource)
}

// clusterResource returns the Terraform awscc_msk_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::MSK::Cluster resource.
func clusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"broker_node_group_info": {
			// Property: BrokerNodeGroupInfo
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BrokerAZDistribution": {
			//       "maxLength": 9,
			//       "minLength": 6,
			//       "type": "string"
			//     },
			//     "ClientSubnets": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "ConnectivityInfo": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "PublicAccess": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Type": {
			//               "maxLength": 23,
			//               "minLength": 7,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "InstanceType": {
			//       "maxLength": 32,
			//       "minLength": 5,
			//       "type": "string"
			//     },
			//     "SecurityGroups": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "StorageInfo": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "EBSStorageInfo": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ProvisionedThroughput": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Enabled": {
			//                   "type": "boolean"
			//                 },
			//                 "VolumeThroughput": {
			//                   "type": "integer"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "VolumeSize": {
			//               "maximum": 16384,
			//               "minimum": 1,
			//               "type": "integer"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "ClientSubnets",
			//     "InstanceType"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"broker_az_distribution": {
						// Property: BrokerAZDistribution
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(6, 9),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"client_subnets": {
						// Property: ClientSubnets
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							resource.RequiresReplace(),
						},
					},
					"connectivity_info": {
						// Property: ConnectivityInfo
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"public_access": {
									// Property: PublicAccess
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"type": {
												// Property: Type
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(7, 23),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"instance_type": {
						// Property: InstanceType
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(5, 32),
						},
					},
					"security_groups": {
						// Property: SecurityGroups
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"storage_info": {
						// Property: StorageInfo
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"ebs_storage_info": {
									// Property: EBSStorageInfo
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"provisioned_throughput": {
												// Property: ProvisionedThroughput
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"enabled": {
															// Property: Enabled
															Type:     types.BoolType,
															Optional: true,
															Computed: true,
															PlanModifiers: []tfsdk.AttributePlanModifier{
																resource.UseStateForUnknown(),
															},
														},
														"volume_throughput": {
															// Property: VolumeThroughput
															Type:     types.Int64Type,
															Optional: true,
															Computed: true,
															PlanModifiers: []tfsdk.AttributePlanModifier{
																resource.UseStateForUnknown(),
															},
														},
													},
												),
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"volume_size": {
												// Property: VolumeSize
												Type:     types.Int64Type,
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.IntBetween(1, 16384),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
		},
		"client_authentication": {
			// Property: ClientAuthentication
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Sasl": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Iam": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Enabled": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         },
			//         "Scram": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Enabled": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Tls": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CertificateAuthorityArnList": {
			//           "insertionOrder": false,
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         },
			//         "Enabled": {
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Unauthenticated": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Enabled": {
			//           "type": "boolean"
			//         }
			//       },
			//       "required": [
			//         "Enabled"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sasl": {
						// Property: Sasl
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"iam": {
									// Property: Iam
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"enabled": {
												// Property: Enabled
												Type:     types.BoolType,
												Required: true,
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"scram": {
									// Property: Scram
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"enabled": {
												// Property: Enabled
												Type:     types.BoolType,
												Required: true,
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"tls": {
						// Property: Tls
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"certificate_authority_arn_list": {
									// Property: CertificateAuthorityArnList
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
										resource.UseStateForUnknown(),
									},
								},
								"enabled": {
									// Property: Enabled
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"unauthenticated": {
						// Property: Unauthenticated
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Type:     types.BoolType,
									Required: true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"cluster_name": {
			// Property: ClusterName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"configuration_info": {
			// Property: ConfigurationInfo
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Arn": {
			//       "type": "string"
			//     },
			//     "Revision": {
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "Revision",
			//     "Arn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Required: true,
					},
					"revision": {
						// Property: Revision
						Type:     types.Int64Type,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"current_version": {
			// Property: CurrentVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The current version of the MSK cluster",
			//   "type": "string"
			// }
			Description: "The current version of the MSK cluster",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"encryption_info": {
			// Property: EncryptionInfo
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "EncryptionAtRest": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "DataVolumeKMSKeyId": {
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "DataVolumeKMSKeyId"
			//       ],
			//       "type": "object"
			//     },
			//     "EncryptionInTransit": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ClientBroker": {
			//           "enum": [
			//             "TLS",
			//             "TLS_PLAINTEXT",
			//             "PLAINTEXT"
			//           ],
			//           "type": "string"
			//         },
			//         "InCluster": {
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"encryption_at_rest": {
						// Property: EncryptionAtRest
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"data_volume_kms_key_id": {
									// Property: DataVolumeKMSKeyId
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"encryption_in_transit": {
						// Property: EncryptionInTransit
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"client_broker": {
									// Property: ClientBroker
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"TLS",
											"TLS_PLAINTEXT",
											"PLAINTEXT",
										}),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"in_cluster": {
									// Property: InCluster
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
										resource.RequiresReplace(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"enhanced_monitoring": {
			// Property: EnhancedMonitoring
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DEFAULT",
			//     "PER_BROKER",
			//     "PER_TOPIC_PER_BROKER",
			//     "PER_TOPIC_PER_PARTITION"
			//   ],
			//   "maxLength": 23,
			//   "minLength": 7,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(7, 23),
				validate.StringInSlice([]string{
					"DEFAULT",
					"PER_BROKER",
					"PER_TOPIC_PER_BROKER",
					"PER_TOPIC_PER_PARTITION",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"kafka_version": {
			// Property: KafkaVersion
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
		},
		"logging_info": {
			// Property: LoggingInfo
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BrokerLogs": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CloudWatchLogs": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Enabled": {
			//               "type": "boolean"
			//             },
			//             "LogGroup": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         },
			//         "Firehose": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "DeliveryStream": {
			//               "type": "string"
			//             },
			//             "Enabled": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         },
			//         "S3": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Bucket": {
			//               "type": "string"
			//             },
			//             "Enabled": {
			//               "type": "boolean"
			//             },
			//             "Prefix": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "BrokerLogs"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"broker_logs": {
						// Property: BrokerLogs
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"cloudwatch_logs": {
									// Property: CloudWatchLogs
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"enabled": {
												// Property: Enabled
												Type:     types.BoolType,
												Required: true,
											},
											"log_group": {
												// Property: LogGroup
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"firehose": {
									// Property: Firehose
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"delivery_stream": {
												// Property: DeliveryStream
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"enabled": {
												// Property: Enabled
												Type:     types.BoolType,
												Required: true,
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"s3": {
									// Property: S3
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"enabled": {
												// Property: Enabled
												Type:     types.BoolType,
												Required: true,
											},
											"prefix": {
												// Property: Prefix
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"number_of_broker_nodes": {
			// Property: NumberOfBrokerNodes
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Required: true,
		},
		"open_monitoring": {
			// Property: OpenMonitoring
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Prometheus": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "JmxExporter": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "EnabledInBroker": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "EnabledInBroker"
			//           ],
			//           "type": "object"
			//         },
			//         "NodeExporter": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "EnabledInBroker": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "EnabledInBroker"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Prometheus"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"prometheus": {
						// Property: Prometheus
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"jmx_exporter": {
									// Property: JmxExporter
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"enabled_in_broker": {
												// Property: EnabledInBroker
												Type:     types.BoolType,
												Required: true,
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"node_exporter": {
									// Property: NodeExporter
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"enabled_in_broker": {
												// Property: EnabledInBroker
												Type:     types.BoolType,
												Required: true,
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"storage_mode": {
			// Property: StorageMode
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "LOCAL",
			//     "TIERED"
			//   ],
			//   "maxLength": 6,
			//   "minLength": 5,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(5, 6),
				validate.StringInSlice([]string{
					"LOCAL",
					"TIERED",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A key-value pair to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A key-value pair to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::MSK::Cluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MSK::Cluster").WithTerraformTypeName("awscc_msk_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                            "Arn",
		"broker_az_distribution":         "BrokerAZDistribution",
		"broker_logs":                    "BrokerLogs",
		"broker_node_group_info":         "BrokerNodeGroupInfo",
		"bucket":                         "Bucket",
		"certificate_authority_arn_list": "CertificateAuthorityArnList",
		"client_authentication":          "ClientAuthentication",
		"client_broker":                  "ClientBroker",
		"client_subnets":                 "ClientSubnets",
		"cloudwatch_logs":                "CloudWatchLogs",
		"cluster_name":                   "ClusterName",
		"configuration_info":             "ConfigurationInfo",
		"connectivity_info":              "ConnectivityInfo",
		"current_version":                "CurrentVersion",
		"data_volume_kms_key_id":         "DataVolumeKMSKeyId",
		"delivery_stream":                "DeliveryStream",
		"ebs_storage_info":               "EBSStorageInfo",
		"enabled":                        "Enabled",
		"enabled_in_broker":              "EnabledInBroker",
		"encryption_at_rest":             "EncryptionAtRest",
		"encryption_in_transit":          "EncryptionInTransit",
		"encryption_info":                "EncryptionInfo",
		"enhanced_monitoring":            "EnhancedMonitoring",
		"firehose":                       "Firehose",
		"iam":                            "Iam",
		"in_cluster":                     "InCluster",
		"instance_type":                  "InstanceType",
		"jmx_exporter":                   "JmxExporter",
		"kafka_version":                  "KafkaVersion",
		"log_group":                      "LogGroup",
		"logging_info":                   "LoggingInfo",
		"node_exporter":                  "NodeExporter",
		"number_of_broker_nodes":         "NumberOfBrokerNodes",
		"open_monitoring":                "OpenMonitoring",
		"prefix":                         "Prefix",
		"prometheus":                     "Prometheus",
		"provisioned_throughput":         "ProvisionedThroughput",
		"public_access":                  "PublicAccess",
		"revision":                       "Revision",
		"s3":                             "S3",
		"sasl":                           "Sasl",
		"scram":                          "Scram",
		"security_groups":                "SecurityGroups",
		"storage_info":                   "StorageInfo",
		"storage_mode":                   "StorageMode",
		"tags":                           "Tags",
		"tls":                            "Tls",
		"type":                           "Type",
		"unauthenticated":                "Unauthenticated",
		"volume_size":                    "VolumeSize",
		"volume_throughput":              "VolumeThroughput",
	})

	opts = opts.WithCreateTimeoutInMinutes(120).WithDeleteTimeoutInMinutes(30)

	opts = opts.WithUpdateTimeoutInMinutes(720)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
