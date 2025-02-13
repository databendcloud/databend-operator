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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// WarehouseKind is the Kind name for the Warehouse.
	WarehouseKind string = "Warehouse"
)

const (
	// WarehouseCreated means that the Warehouse creation has succeeded.
	WarehouseCreated string = "Created"

	// WarehouseSuspended means that the Warehouse is suspended.
	WarehouseSuspended string = "Suspended"

	// WarehouseFailed means that the Warehouse have failed due to some reasons.
	WarehouseFailed string = "Failed"

	// WarehouseRunning means that the Warehouse is running.
	WarehouseRunning string = "Running"
)

const (
	// WarehouseSuspendedReason is the "Suspended" condition reason.
	// When the Warehouse is suspended, this is added.
	WarehouseSuspendedReason string = "Suspended"

	// WarehouseResumeReason is the "Suspended" condition reason.
	// When the Warehouse suspension is changed from True to False, this is added.
	WarehouseResumedReason string = "Resumed"

	// WarehouseRunningReason is the "Running" condition reason.
	// When the created objects succeeded after building succeeded, this is added.
	WarehouseRunningReason string = "WarehouseRunning"

	// WarehouseCreatedReason is the "Created" condition reason.
	// When the Warehouse creation succeeded and related objects are not ready, this is added.
	WarehouseCreatedReason string = "WarehouseCreated"

	// WarehouseBuildFailedReason is the "Failed" condition reason.
	// When the Warehouse building failed in the reconciling loop, this is added.
	WarehouseBuildFailedReason string = "WarehouseBuildFailed"

	// WarehouseRunFailedReason is the "Failed" condition reason.
	// When the Warehouse failed outside the building stage, this is added.
	WarehouseRunFailedReason string = "WarehouseRunFailed"
)

type DiskCacheSpec struct {
	// Whether to enable cache in disk.
	// +kubebuilder:default=false
	Enabled bool `json:"enabled,omitempty"`

	// Max size of cache in disk.
	MaxSize resource.Quantity `json:"size,omitempty"`

	// Path to cache directory in disk.
	// If not set, default to /var/lib/databend/cache.
	Path string `json:"path,omitempty"`

	// Provide storage class to allocate disk cache automatically.
	StorageClass string `json:"storageClass,omitempty"`
}

type LogSpec struct {
	// Specifications for logging in files.
	File FileLogSpec `json:"file,omitempty"`

	// Specifications for stderr logging.
	Stderr FileLogSpec `json:"stderr,omitempty"`

	// Specifications for query logging.
	Query OTLPLogSpec `json:"query,omitempty"`

	// Specifications for profile logging.
	Profile OTLPLogSpec `json:"profile,omitempty"`
}

type FileLogSpec struct {
	// Whether to enable file logging.
	// +kubebuilder:default=false
	Enabled bool `json:"enabled,omitempty"`

	// Log level.
	Level string `json:"level,omitempty"`

	// Path to log directory.
	Directory string `json:"directory,omitempty"`
}

type OTLPLogSpec struct {
	// Whether to enable OTLP logging.
	// +kubebuilder:default=false
	Enabled bool `json:"enabled,omitempty"`

	// OpenTelemetry Protocol
	// +kubebuilder:default="http"
	Protocol string `json:"protocol,omitempty"`

	// Endpoint for OpenTelemetry Protocol
	Endpoint string `json:"endpoint,omitempty"`
}

type WarehouseServiceSpec struct {
	// Type of service [ClusterIP | NodePort | ExternalName | LoadBalance].
	// +kubebuilder:default="ClusterIP"
	Type string `json:"type,omitempty"`

	// External name is needed when Type is set to "ExternalName"
	ExternalName string `json:"externalName,omitempty"`
}

type WarehouseIngressSpec struct {
	// Annotations for Ingress.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Name of IngressClass.
	IngressClassName string `json:"ingressClassName,omitempty"`

	// Host name of ingress.
	HostName string `json:"hostName,omitempty"`
}

// WarehouseSpec defines the desired state of Warehouse.
type WarehouseSpec struct {
	// Desired replicas of Query
	// +kubebuilder:default=1
	Replicas int `json:"replicas,omitempty"`

	// Image for Query.
	QueryImage string `json:"queryImage,omitempty"`

	// Reference to the Tenant CR, which provides the configuration of storage and Meta cluster.
	// Warehouse must be created in the Tenant's namespace.
	Tenant *corev1.LocalObjectReference `json:"tenant,omitempty"`

	// Configurations of cache in disk.
	Cache DiskCacheSpec `json:"diskCacheSize,omitempty"`

	// Configurations of logging.
	Log LogSpec `json:"log,omitempty"`

	// Additional labels added to Query pod.
	PodLabels map[string]string `json:"labels,omitempty"`

	// Resource required for each Query pod.
	PodResource corev1.ResourceRequirements `json:"resourcesPerNode,omitempty"`

	// Taint tolerations for Query pod.
	PodTolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// Node selector for Query pod.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Service specifications for Query cluster.
	Service WarehouseServiceSpec `json:"service,omitempty"`

	// Ingress specifications for Query cluster.
	Ingress WarehouseIngressSpec `json:"ingress,omitempty"`

	// Custom settings that will append to the config file of Query.
	Settings map[string]string `json:"settings,omitempty"`
}

// WarehouseStatus defines the observed state of Warehouse.
type WarehouseStatus struct {
	// Number of the ready Query.
	ReadyReplicas int `json:"readyReplicas,omitempty"`

	// Conditions for the Tenant.
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type=number,JSONPath=`.status.readyReplicas`
// +kubebuilder:printcolumn:name="Replicas",type=number,JSONPath=`.spec.replicas`
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.conditions[-1:].type`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Warehouse is the Schema for the warehouses API.
type Warehouse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WarehouseSpec   `json:"spec,omitempty"`
	Status WarehouseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WarehouseList contains a list of Warehouse.
type WarehouseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Warehouse `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Warehouse{}, &WarehouseList{})
}
