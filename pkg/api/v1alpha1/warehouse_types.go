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

type DiskCacheSpec struct {
    Enabled      bool   `json:"enabled,omitempty"`
    MaxBytes     int    `json:"size,omitempty"`
    Path         string `json:"path,omitempty"`
    StorageClass string `json:"storageClass,omitempty"`
}

type LogSpec struct {
    File    FileLogSpec `json:"file,omitempty"`
    Stderr  FileLogSpec `json:"stderr,omitempty"`
    Query   OTLPLogSpec `json:"query,omitempty"`
    Profile OTLPLogSpec `json:"profile,omitempty"`
}

type FileLogSpec struct {
    Enabled   bool   `json:"enabled,omitempty"`
    Level     string `json:"level,omitempty"`
    Directory string `json:"directory,omitempty"`
}

type OTLPLogSpec struct {
    Enabled  bool   `json:"enabled,omitempty"`
    Endpoint string `json:"endpoint,omitempty"`
    Protocol string `json:"protocol,omitempty"`
}

type WarehouseServiceSpec struct {
    Type         string `json:"type,omitempty"`
    ExternalName string `json:"externalName,omitempty"`
}

type WarehouseIngressSpec struct {
    Annotations      map[string]string `json:"annotations,omitempty"`
    IngressClassName string            `json:"ingressClassName,omitempty"`
    HostName         string            `json:"hostName,omitempty"`
}

// WarehouseSpec defines the desired state of Warehouse.
type WarehouseSpec struct {
	Replicas             int    `json:"replicas,omitempty"`
	AutoSuspendAfterSecs int    `json:"autoSuspendAfterSecs,omitempty"`
    QueryImage           string `json:"queryImage,omitempty"`
    
    Cache DiskCacheSpec `json:"diskCacheSize,omitempty"`
    Log   LogSpec       `json:"log,omitempty"`
    
    PodResource corev1.ResourceRequirements `json:"resourcesPerNode,omitempty"`
    
    Tenant         corev1.LocalObjectReference `json:"tenant,omitempty"`
    PodTolerations []corev1.Toleration         `json:"tolerations,omitempty"`
    NodeSelector   map[string]string           `json:"nodeSelector,omitempty"`
    
    Service WarehouseServiceSpec `json:"service,omitempty"`
    Ingress WarehouseIngressSpec `json:"ingress,omitempty"`
    
    PodLabels map[string]string `json:"labels,omitempty"`
    
    Settings map[string]string `json:"settings,omitempty"`

}

// WarehouseStatus defines the observed state of Warehouse.
type WarehouseStatus struct {
    ReadyReplicas int              `json:"readyReplicas,omitempty"`
	Conditions    metav1.Condition `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

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
