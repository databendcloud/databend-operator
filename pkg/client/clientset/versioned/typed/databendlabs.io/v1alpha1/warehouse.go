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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	databendlabsiov1alpha1 "github.com/databendcloud/databend-operator/pkg/client/applyconfiguration/databendlabs.io/v1alpha1"
	scheme "github.com/databendcloud/databend-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// WarehousesGetter has a method to return a WarehouseInterface.
// A group's client should implement this interface.
type WarehousesGetter interface {
	Warehouses(namespace string) WarehouseInterface
}

// WarehouseInterface has methods to work with Warehouse resources.
type WarehouseInterface interface {
	Create(ctx context.Context, warehouse *v1alpha1.Warehouse, opts v1.CreateOptions) (*v1alpha1.Warehouse, error)
	Update(ctx context.Context, warehouse *v1alpha1.Warehouse, opts v1.UpdateOptions) (*v1alpha1.Warehouse, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, warehouse *v1alpha1.Warehouse, opts v1.UpdateOptions) (*v1alpha1.Warehouse, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Warehouse, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WarehouseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Warehouse, err error)
	Apply(ctx context.Context, warehouse *databendlabsiov1alpha1.WarehouseApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Warehouse, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, warehouse *databendlabsiov1alpha1.WarehouseApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Warehouse, err error)
	WarehouseExpansion
}

// warehouses implements WarehouseInterface
type warehouses struct {
	*gentype.ClientWithListAndApply[*v1alpha1.Warehouse, *v1alpha1.WarehouseList, *databendlabsiov1alpha1.WarehouseApplyConfiguration]
}

// newWarehouses returns a Warehouses
func newWarehouses(c *DatabendlabsV1alpha1Client, namespace string) *warehouses {
	return &warehouses{
		gentype.NewClientWithListAndApply[*v1alpha1.Warehouse, *v1alpha1.WarehouseList, *databendlabsiov1alpha1.WarehouseApplyConfiguration](
			"warehouses",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.Warehouse { return &v1alpha1.Warehouse{} },
			func() *v1alpha1.WarehouseList { return &v1alpha1.WarehouseList{} }),
	}
}
