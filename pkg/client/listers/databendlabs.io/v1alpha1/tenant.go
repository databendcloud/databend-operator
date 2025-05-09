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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// TenantLister helps list Tenants.
// All objects returned here must be treated as read-only.
type TenantLister interface {
	// List lists all Tenants in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Tenant, err error)
	// Tenants returns an object that can list and get Tenants.
	Tenants(namespace string) TenantNamespaceLister
	TenantListerExpansion
}

// tenantLister implements the TenantLister interface.
type tenantLister struct {
	listers.ResourceIndexer[*v1alpha1.Tenant]
}

// NewTenantLister returns a new TenantLister.
func NewTenantLister(indexer cache.Indexer) TenantLister {
	return &tenantLister{listers.New[*v1alpha1.Tenant](indexer, v1alpha1.Resource("tenant"))}
}

// Tenants returns an object that can list and get Tenants.
func (s *tenantLister) Tenants(namespace string) TenantNamespaceLister {
	return tenantNamespaceLister{listers.NewNamespaced[*v1alpha1.Tenant](s.ResourceIndexer, namespace)}
}

// TenantNamespaceLister helps list and get Tenants.
// All objects returned here must be treated as read-only.
type TenantNamespaceLister interface {
	// List lists all Tenants in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Tenant, err error)
	// Get retrieves the Tenant from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Tenant, error)
	TenantNamespaceListerExpansion
}

// tenantNamespaceLister implements the TenantNamespaceLister
// interface.
type tenantNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.Tenant]
}
