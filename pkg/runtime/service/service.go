package service

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
	"github.com/databendcloud/databend-operator/pkg/runtime/objectmeta"
)

func BuildService(tenant *v1alpha1.Tenant, wh *v1alpha1.Warehouse) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: *objectmeta.BuildObjectMetaUnderWarehouse(
			wh, common.GetQueryServiceName(tenant.Name, wh.Name),
		),
		Spec: corev1.ServiceSpec{
			Selector: objectmeta.LabelsFromWarehouse(wh),
			Ports: []corev1.ServicePort{
				{
					Port: int32(common.ServicePortFlight),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(common.ServicePortFlight),
					},
					Name: string(common.ServiceProtocolFlight),
				},
				{
					Port: int32(common.ServicePortAdmin),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(common.ServicePortAdmin),
					},
					Name: string(common.ServiceProtocolAdmin),
				},
				{
					Port: int32(common.ServicePortMetrics),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(common.ServicePortMetrics),
					},
					Name: string(common.ServiceProtocolMetrics),
				},

				{
					Port: int32(common.ServicePortMySQL),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(common.ServicePortMySQL),
					},
					Name: string(common.ServiceProtocolMySQL),
				},
				{
					Port: int32(common.ServicePortCKHttp),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(common.ServicePortCKHttp),
					},
					Name: string(common.ServiceProtocolCKHttp),
				},
				{
					Port: int32(common.ServicePortQuery),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(common.ServicePortQuery),
					},
					Name: string(common.ServiceProtocolQuery),
				},
				{
					Port: int32(common.ServicePortFlightSQL),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(common.ServicePortFlightSQL),
					},
					Name: string(common.ServiceProtocolFlightSQL),
				},
			},
		},
	}
}
