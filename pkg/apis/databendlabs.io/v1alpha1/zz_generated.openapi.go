//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by openapi-gen. DO NOT EDIT.

package v1alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.DiskCacheSpec":        schema_pkg_apis_databendlabsio_v1alpha1_DiskCacheSpec(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.FileLogSpec":          schema_pkg_apis_databendlabsio_v1alpha1_FileLogSpec(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.LogSpec":              schema_pkg_apis_databendlabsio_v1alpha1_LogSpec(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.MetaAuth":             schema_pkg_apis_databendlabsio_v1alpha1_MetaAuth(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.MetaConfig":           schema_pkg_apis_databendlabsio_v1alpha1_MetaConfig(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.OTLPLogSpec":          schema_pkg_apis_databendlabsio_v1alpha1_OTLPLogSpec(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.S3Auth":               schema_pkg_apis_databendlabsio_v1alpha1_S3Auth(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.S3Storage":            schema_pkg_apis_databendlabsio_v1alpha1_S3Storage(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.StderrLogSpec":        schema_pkg_apis_databendlabsio_v1alpha1_StderrLogSpec(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.Storage":              schema_pkg_apis_databendlabsio_v1alpha1_Storage(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.Tenant":               schema_pkg_apis_databendlabsio_v1alpha1_Tenant(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.TenantList":           schema_pkg_apis_databendlabsio_v1alpha1_TenantList(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.TenantSpec":           schema_pkg_apis_databendlabsio_v1alpha1_TenantSpec(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.TenantStatus":         schema_pkg_apis_databendlabsio_v1alpha1_TenantStatus(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.User":                 schema_pkg_apis_databendlabsio_v1alpha1_User(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.Warehouse":            schema_pkg_apis_databendlabsio_v1alpha1_Warehouse(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseIngressSpec": schema_pkg_apis_databendlabsio_v1alpha1_WarehouseIngressSpec(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseList":        schema_pkg_apis_databendlabsio_v1alpha1_WarehouseList(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseSpec":        schema_pkg_apis_databendlabsio_v1alpha1_WarehouseSpec(ref),
		"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseStatus":      schema_pkg_apis_databendlabsio_v1alpha1_WarehouseStatus(ref),
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_DiskCacheSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to enable cache in disk.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "Max size of cache in disk.",
							Ref:         ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path to cache directory in disk. If not set, default to /var/lib/databend/cache.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pvc": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to use PVC as the storage of disk cache.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"storageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "Provide storage class to allocate disk cache automatically. If not set, default to use EmptyDir as disk cache rather than PVC.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/api/resource.Quantity"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_FileLogSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to enable file logging.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"format": {
						SchemaProps: spec.SchemaProps{
							Description: "Log format.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"level": {
						SchemaProps: spec.SchemaProps{
							Description: "Log level.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dir": {
						SchemaProps: spec.SchemaProps{
							Description: "Path to log directory.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_LogSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"file": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifications for logging in files.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.FileLogSpec"),
						},
					},
					"stderr": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifications for stderr logging.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.StderrLogSpec"),
						},
					},
					"query": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifications for query logging.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.OTLPLogSpec"),
						},
					},
					"profile": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifications for profile logging.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.OTLPLogSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.FileLogSpec", "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.OTLPLogSpec", "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.StderrLogSpec"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_MetaAuth(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"user": {
						SchemaProps: spec.SchemaProps{
							Description: "User of Meta cluster.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"password": {
						SchemaProps: spec.SchemaProps{
							Description: "Password of Meta cluster.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"passwordSecretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the secret with User and Password to Meta cluster. Secret can be created in any namespace.",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_MetaConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"user": {
						SchemaProps: spec.SchemaProps{
							Description: "User of Meta cluster.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"password": {
						SchemaProps: spec.SchemaProps{
							Description: "Password of Meta cluster.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"passwordSecretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the secret with User and Password to Meta cluster. Secret can be created in any namespace.",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"endpoints": {
						SchemaProps: spec.SchemaProps{
							Description: "Exposed endpoints of Meta cluster (must list all pod endpoints in the Meta cluster).",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"timeoutInSecond": {
						SchemaProps: spec.SchemaProps{
							Description: "Timeout seconds of connections to Meta cluster.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"autoSyncInterval": {
						SchemaProps: spec.SchemaProps{
							Description: "Interval for warehouse to sync data from Meta cluster.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_OTLPLogSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to enable OTLP logging.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"protocol": {
						SchemaProps: spec.SchemaProps{
							Description: "OpenTelemetry Protocol",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoint for OpenTelemetry Protocol",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels for OpenTelemetry Protocol",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_S3Auth(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"secretKey": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret Access Key of S3 storage.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accessKey": {
						SchemaProps: spec.SchemaProps{
							Description: "Access Key ID of S3 storage.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the secret with SerectKey and AccessKey to S3 storage. Secret can be created in any namespace.",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_S3Storage(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"secretKey": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret Access Key of S3 storage.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accessKey": {
						SchemaProps: spec.SchemaProps{
							Description: "Access Key ID of S3 storage.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the secret with SerectKey and AccessKey to S3 storage. Secret can be created in any namespace.",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"allowInsecure": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to allow insecure connections to S3 storage. If set to true, users can establish HTTP connections to S3 storage. Otherwise, only HTTPS connections are allowed. Default to true.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"rootPath": {
						SchemaProps: spec.SchemaProps{
							Description: "Root path of S3.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"bucketName": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of S3 bucket.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Description: "Region of S3 storage.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoint of S3 storage.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_StderrLogSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to enable stderr logging.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"format": {
						SchemaProps: spec.SchemaProps{
							Description: "Log format.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"level": {
						SchemaProps: spec.SchemaProps{
							Description: "Log level.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_Storage(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"s3": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of S3 storage.",
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.S3Storage"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.S3Storage"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_Tenant(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Tenant is the Schema for the tenants API.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.TenantSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.TenantStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.TenantSpec", "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.TenantStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_TenantList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TenantList contains a list of Tenant.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.Tenant"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.Tenant", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_TenantSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TenantSpec defines the desired state of Tenant.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"s3": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of S3 storage.",
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.S3Storage"),
						},
					},
					"meta": {
						SchemaProps: spec.SchemaProps{
							Description: "Configurations to open connections to a Meta cluster.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.MetaConfig"),
						},
					},
					"users": {
						SchemaProps: spec.SchemaProps{
							Description: "Built-in users in the warehouse created by this tenant. If not set, we'll create \"admin\" user with password \"admin\".",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.User"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.MetaConfig", "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.S3Storage", "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.User"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_TenantStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TenantStatus defines the observed state of Tenant.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"type",
								},
								"x-kubernetes-list-type":       "map",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions for the Tenant.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Condition"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_User(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of warehouse user.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"authType": {
						SchemaProps: spec.SchemaProps{
							Description: "Authentication type of warehouse password. Currently we support: sha256_password, no_password.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"authString": {
						SchemaProps: spec.SchemaProps{
							Description: "Password encrypted with AuthType.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"authStringSecretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the secret with AuthString of user. Secret can be created in any namespace.",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_Warehouse(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Warehouse is the Schema for the warehouses API.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseSpec", "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_WarehouseIngressSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to enable Ingress for Query.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Description: "Annotations for Ingress.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"enableLoadBalance": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to enable load balance for Ingress.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"ingressClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of IngressClass.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"hostName": {
						SchemaProps: spec.SchemaProps{
							Description: "Host name of ingress.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_WarehouseList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WarehouseList contains a list of Warehouse.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.Warehouse"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.Warehouse", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_WarehouseSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WarehouseSpec defines the desired state of Warehouse.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Desired replicas of Query",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"queryImage": {
						SchemaProps: spec.SchemaProps{
							Description: "Image for Query.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tenant": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the Tenant CR, which provides the configuration of storage and Meta cluster. Warehouse must be created in the Tenant's namespace.",
							Ref:         ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
					"cache": {
						SchemaProps: spec.SchemaProps{
							Description: "Configurations of cache in disk.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.DiskCacheSpec"),
						},
					},
					"log": {
						SchemaProps: spec.SchemaProps{
							Description: "Configurations of logging.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.LogSpec"),
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Additional labels added to Query pod.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"resourcesPerNode": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource required for each Query pod.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "Taint tolerations for Query pod.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "Node selector for Query pod.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"ingress": {
						SchemaProps: spec.SchemaProps{
							Description: "Ingress specifications for Query cluster.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseIngressSpec"),
						},
					},
					"settings": {
						SchemaProps: spec.SchemaProps{
							Description: "Custom settings that will append to the config file of Query.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.DiskCacheSpec", "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.LogSpec", "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1.WarehouseIngressSpec", "k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_databendlabsio_v1alpha1_WarehouseStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WarehouseStatus defines the observed state of Warehouse.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"readyReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of the ready Query.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"type",
								},
								"x-kubernetes-list-type":       "map",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions for the Tenant.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Condition"},
	}
}
