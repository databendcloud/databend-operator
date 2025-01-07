package object

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

func BuildConfigMap(ctx context.Context, tenant *databendv1alpha1.Tenant, warehouse *databendv1alpha1.Warehouse) (*corev1.ConfigMap, error) {
	_, _, _ = ctx, tenant, warehouse
	return nil, nil
}
