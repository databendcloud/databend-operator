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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type UserAuthType string

const (
	MD5        UserAuthType = "md5"
	NoPassword UserAuthType = "no_password"
)

type Storage struct {
	// Specification of S3 storage.
	S3 *S3Storage `json:"s3,omitempty"`
}

type S3Storage struct {
	// Authentication configuration of S3 storage.
	S3Auth `json:",inline"`

	// Whether to allow insecure connections to S3 storage.
	// If set to true, users can establish HTTP connections to S3 storage.
	// Otherwise, only HTTPS connections are allowed. Default to true.
	// +kubebuilder:default=true
	AllowInsecure bool `json:"allowInsecure,omitempty"`

	// Root path of S3.
	RootPath string `json:"rootPath,omitempty"`

	// Name of S3 bucket.
	BucketName string `json:"bucketName,omitempty"`

	// Region of S3 storage.
	Region string `json:"region,omitempty"`

	// Endpoint of S3 storage.
	Endpoint string `json:"endpoint,omitempty"`
}

type S3Auth struct {
	// Secret Access Key of S3 storage.
	SecretKey string `json:"secretKey,omitempty"`

	// Access Key ID of S3 storage.
	AccessKey string `json:"accessKey,omitempty"`

	// Reference to the secret with SerectKey and AccessKey to S3 storage.
	// Secret can be created in any namespace.
	SecretRef *corev1.ObjectReference `json:"secretRef,omitempty"`
}

type MetaConfig struct {
	// Authentication configurations to connect to Meta cluster.
	MetaAuth `json:",inline"`

	// Exposed endpoints of Meta cluster (must list all pod endpoints in the Meta cluster).
	Endpoints []string `json:"endpoints,omitempty"`

	// Timeout seconds of connections to Meta cluster.
	// +kubebuilder:default=3
	TimeoutInSeconds int `json:"timeoutInSecond,omitempty"`

	// Interval for warehouse to sync data from Meta cluster.
	// +kubebuilder:default=60
	AutoSyncInterval int `json:"autoSyncInterval,omitempty"`
}

type MetaAuth struct {
	// User of Meta cluster.
	User string `json:"user,omitempty"`

	// Password of Meta cluster.
	Password string `json:"password,omitempty"`

	// Reference to the secret with User and Password to Meta cluster.
	// Secret can be created in any namespace.
	PasswordSecretRef *corev1.ObjectReference `json:"passwordSecretRef,omitempty"`
}

type User struct {
	// Name of warehouse user.
	Name string `json:"name,omitempty"`

	// Authentication type of warehouse password.
	// Currently we support: md5, no_password.
	// +kubebuilder:default="no_password"
	AuthType UserAuthType `json:"authType,omitempty"`

	// Password encrypted with AuthType.
	AuthString string `json:"authString,omitempty"`
}

// TenantSpec defines the desired state of Tenant.
type TenantSpec struct {
	// Object storage specifications. Currently we only support S3.
	Storage `json:",inline"`

	// Configurations to open connections to a Meta cluster.
	Meta MetaConfig `json:"meta,omitempty"`

	// Built-in users in the warehouse created by this tenant.
	// If not set, we'll create "admin" user with password "admin".
	// +listType=map
	// +listMapKey=name
	BuiltinUsers []User `json:"builtinUsers,omitempty"`
}

// TenantStatus defines the observed state of Tenant.
type TenantStatus struct {
	// Conditions for the Tenant.
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.conditions[-1:].type`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Tenant is the Schema for the tenants API.
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantSpec   `json:"spec,omitempty"`
	Status TenantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantList contains a list of Tenant.
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tenant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
