package ingress

import (
	"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"

	"github.com/databendcloud/databend-operator/pkg/common"
	"github.com/databendcloud/databend-operator/pkg/runtime/objectmeta"
)

func BuildIngress(tenant *v1alpha1.Tenant, wh *v1alpha1.Warehouse) *networkingv1.Ingress {
	return &networkingv1.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name:            common.GetQueryIngressName(tenant.Name, wh.Name),
			Namespace:       wh.Namespace,
			Labels:          objectmeta.LabelsFromWarehouse(wh),
			Annotations:     BuildIngressAnnoations(wh),
			OwnerReferences: objectmeta.BuildOwnerReferencesByWarehouse(wh),
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: &wh.Spec.Ingress.IngressClassName,
			Rules: []networkingv1.IngressRule{
				{
					Host: wh.Spec.Ingress.HostName,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									PathType: ptr.To(networkingv1.PathTypePrefix),
									Path:     "/v1/admin",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: common.GetQueryServiceName(tenant.Name, wh.Name),
											Port: networkingv1.ServiceBackendPort{
												Name: string(common.ServiceProtocolAdmin),
											},
										},
									},
								},
								{
									PathType: ptr.To(networkingv1.PathTypePrefix),
									Path:     "/v1/ckhttp",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: common.GetQueryServiceName(tenant.Name, wh.Name),
											Port: networkingv1.ServiceBackendPort{
												Name: string(common.ServiceProtocolCKHttp),
											},
										},
									},
								},
								{
									PathType: ptr.To(networkingv1.PathTypePrefix),
									Path:     "/v1/flight",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: common.GetQueryServiceName(tenant.Name, wh.Name),
											Port: networkingv1.ServiceBackendPort{
												Name: string(common.ServiceProtocolFlight),
											},
										},
									},
								},
								{
									PathType: ptr.To(networkingv1.PathTypePrefix),
									Path:     "/v1/flightsql",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: common.GetQueryServiceName(tenant.Name, wh.Name),
											Port: networkingv1.ServiceBackendPort{
												Name: string(common.ServiceProtocolFlightSQL),
											},
										},
									},
								},
								{
									PathType: ptr.To(networkingv1.PathTypePrefix),
									Path:     "/v1/metrics",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: common.GetQueryServiceName(tenant.Name, wh.Name),
											Port: networkingv1.ServiceBackendPort{
												Name: string(common.ServiceProtocolMetrics),
											},
										},
									},
								},
								{
									PathType: ptr.To(networkingv1.PathTypePrefix),
									Path:     "/v1/mysql",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: common.GetQueryServiceName(tenant.Name, wh.Name),
											Port: networkingv1.ServiceBackendPort{
												Name: string(common.ServiceProtocolMySQL),
											},
										},
									},
								},
								{
									PathType: ptr.To(networkingv1.PathTypePrefix),
									Path:     "/v1/query",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: common.GetQueryServiceName(tenant.Name, wh.Name),
											Port: networkingv1.ServiceBackendPort{
												Name: string(common.ServiceProtocolQuery),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func BuildIngressAnnoations(wh *v1alpha1.Warehouse) map[string]string {
	annotations := map[string]string{
		"external-dns.alpha.kubernetes.io/cloudflare-proxied":      "false",
		"external-dns.alpha.kubernetes.io/hostname":                wh.Spec.Ingress.HostName,
		"external-dns.alpha.kubernetes.io/ingress-hostname-source": "annotation-only",
		"external-dns.alpha.kubernetes.io/ttl":                     "100",
		"nginx.ingress.kubernetes.io/proxy-body-size":              "10240m",
		"nginx.ingress.kubernetes.io/proxy-send-timeout":           "3600s",
	}

	for k, v := range wh.Spec.Ingress.Annotations {
		annotations[k] = v
	}

	if wh.Spec.Ingress.EnableLoadBalance {
		annotations["nginx.ingress.kubernetes.io/upstream-hash-by"] = "$http_x_databend_route_hint"
	}

	return annotations
}
