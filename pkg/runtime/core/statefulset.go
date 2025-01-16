package core

import (
	appsv1 "k8s.io/api/apps/v1"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

func BuildStatefulSet(tenant *databendv1alpha1.Tenant, warehouse *databendv1alpha1.Warehouse) (*appsv1.StatefulSet, error) {
	_, _ = tenant, warehouse
	return nil, nil
}
