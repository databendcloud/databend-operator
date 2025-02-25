package ingress

import (
	"github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
