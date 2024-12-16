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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	databendlabsiov1alpha1 "github.com/databendcloud/databend-operator/pkg/client/applyconfiguration/databendlabs.io/v1alpha1"
	internal "github.com/databendcloud/databend-operator/pkg/client/applyconfiguration/internal"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=databendlabs.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("DiskCacheSpec"):
		return &databendlabsiov1alpha1.DiskCacheSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FileLogSpec"):
		return &databendlabsiov1alpha1.FileLogSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("LogSpec"):
		return &databendlabsiov1alpha1.LogSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MetaAuth"):
		return &databendlabsiov1alpha1.MetaAuthApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MetaConfig"):
		return &databendlabsiov1alpha1.MetaConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OTLPLogSpec"):
		return &databendlabsiov1alpha1.OTLPLogSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("S3Auth"):
		return &databendlabsiov1alpha1.S3AuthApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("S3Storage"):
		return &databendlabsiov1alpha1.S3StorageApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Storage"):
		return &databendlabsiov1alpha1.StorageApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Tenant"):
		return &databendlabsiov1alpha1.TenantApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TenantSpec"):
		return &databendlabsiov1alpha1.TenantSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TenantStatus"):
		return &databendlabsiov1alpha1.TenantStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("User"):
		return &databendlabsiov1alpha1.UserApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Warehouse"):
		return &databendlabsiov1alpha1.WarehouseApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WarehouseIngressSpec"):
		return &databendlabsiov1alpha1.WarehouseIngressSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WarehouseServiceSpec"):
		return &databendlabsiov1alpha1.WarehouseServiceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WarehouseSpec"):
		return &databendlabsiov1alpha1.WarehouseSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WarehouseStatus"):
		return &databendlabsiov1alpha1.WarehouseStatusApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
