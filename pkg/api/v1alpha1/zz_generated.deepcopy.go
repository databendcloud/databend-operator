//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskCacheSpec) DeepCopyInto(out *DiskCacheSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskCacheSpec.
func (in *DiskCacheSpec) DeepCopy() *DiskCacheSpec {
	if in == nil {
		return nil
	}
	out := new(DiskCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileLogSpec) DeepCopyInto(out *FileLogSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileLogSpec.
func (in *FileLogSpec) DeepCopy() *FileLogSpec {
	if in == nil {
		return nil
	}
	out := new(FileLogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSpec) DeepCopyInto(out *LogSpec) {
	*out = *in
	out.File = in.File
	out.Stderr = in.Stderr
	out.Query = in.Query
	out.Profile = in.Profile
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSpec.
func (in *LogSpec) DeepCopy() *LogSpec {
	if in == nil {
		return nil
	}
	out := new(LogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaAuth) DeepCopyInto(out *MetaAuth) {
	*out = *in
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaAuth.
func (in *MetaAuth) DeepCopy() *MetaAuth {
	if in == nil {
		return nil
	}
	out := new(MetaAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaConfig) DeepCopyInto(out *MetaConfig) {
	*out = *in
	in.MetaAuth.DeepCopyInto(&out.MetaAuth)
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaConfig.
func (in *MetaConfig) DeepCopy() *MetaConfig {
	if in == nil {
		return nil
	}
	out := new(MetaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OTLPLogSpec) DeepCopyInto(out *OTLPLogSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OTLPLogSpec.
func (in *OTLPLogSpec) DeepCopy() *OTLPLogSpec {
	if in == nil {
		return nil
	}
	out := new(OTLPLogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Auth) DeepCopyInto(out *S3Auth) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Auth.
func (in *S3Auth) DeepCopy() *S3Auth {
	if in == nil {
		return nil
	}
	out := new(S3Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Storage) DeepCopyInto(out *S3Storage) {
	*out = *in
	in.S3Auth.DeepCopyInto(&out.S3Auth)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Storage.
func (in *S3Storage) DeepCopy() *S3Storage {
	if in == nil {
		return nil
	}
	out := new(S3Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3Storage)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tenant) DeepCopyInto(out *Tenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tenant.
func (in *Tenant) DeepCopy() *Tenant {
	if in == nil {
		return nil
	}
	out := new(Tenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantList) DeepCopyInto(out *TenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantList.
func (in *TenantList) DeepCopy() *TenantList {
	if in == nil {
		return nil
	}
	out := new(TenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantSpec) DeepCopyInto(out *TenantSpec) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
	in.Meta.DeepCopyInto(&out.Meta)
	if in.BuiltinUsers != nil {
		in, out := &in.BuiltinUsers, &out.BuiltinUsers
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantSpec.
func (in *TenantSpec) DeepCopy() *TenantSpec {
	if in == nil {
		return nil
	}
	out := new(TenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantStatus) DeepCopyInto(out *TenantStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantStatus.
func (in *TenantStatus) DeepCopy() *TenantStatus {
	if in == nil {
		return nil
	}
	out := new(TenantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	if in.AuthStringSecretRef != nil {
		in, out := &in.AuthStringSecretRef, &out.AuthStringSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Warehouse) DeepCopyInto(out *Warehouse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Warehouse.
func (in *Warehouse) DeepCopy() *Warehouse {
	if in == nil {
		return nil
	}
	out := new(Warehouse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Warehouse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseIngressSpec) DeepCopyInto(out *WarehouseIngressSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseIngressSpec.
func (in *WarehouseIngressSpec) DeepCopy() *WarehouseIngressSpec {
	if in == nil {
		return nil
	}
	out := new(WarehouseIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseList) DeepCopyInto(out *WarehouseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Warehouse, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseList.
func (in *WarehouseList) DeepCopy() *WarehouseList {
	if in == nil {
		return nil
	}
	out := new(WarehouseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WarehouseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseServiceSpec) DeepCopyInto(out *WarehouseServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseServiceSpec.
func (in *WarehouseServiceSpec) DeepCopy() *WarehouseServiceSpec {
	if in == nil {
		return nil
	}
	out := new(WarehouseServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseSpec) DeepCopyInto(out *WarehouseSpec) {
	*out = *in
	if in.Tenant != nil {
		in, out := &in.Tenant, &out.Tenant
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	out.Cache = in.Cache
	out.Log = in.Log
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PodResource.DeepCopyInto(&out.PodResource)
	if in.PodTolerations != nil {
		in, out := &in.PodTolerations, &out.PodTolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Service = in.Service
	in.Ingress.DeepCopyInto(&out.Ingress)
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseSpec.
func (in *WarehouseSpec) DeepCopy() *WarehouseSpec {
	if in == nil {
		return nil
	}
	out := new(WarehouseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseStatus) DeepCopyInto(out *WarehouseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseStatus.
func (in *WarehouseStatus) DeepCopy() *WarehouseStatus {
	if in == nil {
		return nil
	}
	out := new(WarehouseStatus)
	in.DeepCopyInto(out)
	return out
}
