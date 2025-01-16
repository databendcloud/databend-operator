package statefulset

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
	"github.com/databendcloud/databend-operator/pkg/runtime/resource"
)

type StatefulSetBuilder struct {
	tenant             *databendv1alpha1.Tenant
	warehouse          *databendv1alpha1.Warehouse
}

func NewStatefulSetBuilder(
	tenant *databendv1alpha1.Tenant,
	warehouse *databendv1alpha1.Warehouse,
) *StatefulSetBuilder {
	return &StatefulSetBuilder{
		tenant:             tenant,
		warehouse:          warehouse,
	}
}

func (b *StatefulSetBuilder) Build() *appsv1.StatefulSet {
	meta := b.buildMeta()
	sts := &appsv1.StatefulSet{
		ObjectMeta: meta,
		Spec: appsv1.StatefulSetSpec{
			Replicas:            ptr.To(int32(b.warehouse.Spec.Replicas)),
			PodManagementPolicy: appsv1.ParallelPodManagement,
			Selector: &metav1.LabelSelector{
				MatchLabels: meta.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: meta,
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						FSGroup:      ptr.To(int64(1000)),
						RunAsNonRoot: ptr.To(true),
					},
					NodeSelector: copyMap(b.warehouse.Spec.NodeSelector),
					Tolerations:  copyTolerations(b.warehouse.Spec.PodTolerations),
					Containers:   b.buildPodContainers(),
					Volumes:      b.buildPodVolumes(),
					Affinity: &corev1.Affinity{
						PodAffinity: b.buildPodAffinity(),
					},
					TerminationGracePeriodSeconds: ptr.To(int64(30)),
				},
			},
		},
	}
	sts = sts.DeepCopy()
	patchQueryPodWithCache(&sts.Spec.Template, b.tenant, b.warehouse)
	return sts
}

func (b *StatefulSetBuilder) statefulSetName() string {
	return common.GetQueryStatefulSetName(b.tenant.Name, b.warehouse.Name)
}

func (b *StatefulSetBuilder) buildAnnotations() map[string]string {
	annotations := map[string]string{
		common.KeyTenant:        b.tenant.Name,
		common.KeyWarehouse:     b.warehouse.Name,
		common.KeyWarehouseSize: fmt.Sprint(b.warehouse.Spec.Replicas),
	}
	return annotations
}

func (b *StatefulSetBuilder) buildLabels() map[string]string {
	lbs := map[string]string{
		common.KeyTenant:        b.tenant.Name,
		common.KeyWarehouse:     b.warehouse.Name,
		common.KeyWarehouseSize: fmt.Sprint(b.warehouse.Spec.Replicas),
		common.KeyApp:           common.ValueAppWarehouse,
	}
	return lbs
}

func (b *StatefulSetBuilder) buildPodContainers() []corev1.Container {
	image := common.GetQueryImage(b.warehouse)

	command := []string{
		"/databend-query",
		"--config-file=/etc/config/config.toml",
		"--flight-api-address=$(POD_IP):9191",
		"--cluster-id=" + b.warehouse.Name,
	}
	q := []corev1.Container{
		{
			Name:  common.DatabendQueryContainerName,
			Image: image,
			SecurityContext: &corev1.SecurityContext{
				AllowPrivilegeEscalation: ptr.To(false),
				Capabilities: &corev1.Capabilities{
					Drop: []corev1.Capability{"ALL"},
				},
				ReadOnlyRootFilesystem: ptr.To(true),
				RunAsNonRoot:           ptr.To(true),
				RunAsUser:              ptr.To(int64(1000)),
				RunAsGroup:             ptr.To(int64(1000)),
			},
			Command:      command,
			Resources:    b.warehouse.Spec.PodResource,
			Ports:        getPorts(),
			Env:          getEnvs(),
			VolumeMounts: getVolumeMounts(),
		},
	}
	return q
}

func (b *StatefulSetBuilder) buildMeta() metav1.ObjectMeta {
	meta := metav1.ObjectMeta{
		Name:        b.statefulSetName(),
		Namespace:   b.warehouse.Namespace,
		Labels:      b.buildLabels(),
		Annotations: b.buildAnnotations(),
	}

	var (
		apiVersion = b.warehouse.APIVersion
		kind       = b.warehouse.Kind
	)
	if len(apiVersion) == 0 || len(kind) == 0 {
		apiVersion = databendv1alpha1.GroupVersion.String()
		kind = databendv1alpha1.WarehouseKind
	}
	meta.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion: apiVersion,
			Kind:       kind,
			Name:       b.warehouse.Name,
			UID:        b.warehouse.UID,
		},
	}
	return meta
}

func (b *StatefulSetBuilder) buildPodVolumes() []corev1.Volume {
	volumes := []corev1.Volume{
		{
			Name: "query-config",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: common.GetQueryConfigMapName(b.tenant.Name, b.warehouse.Name),
					},
				},
			},
		},
		{
			Name: "tmp",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/tmp",
				},
			},
		},
	}

	return volumes
}

func (b *StatefulSetBuilder) buildPodAffinity() *corev1.PodAffinity {
	podAffinity := &corev1.PodAffinity{
		PreferredDuringSchedulingIgnoredDuringExecution: []corev1.WeightedPodAffinityTerm{
			{
				Weight: 10,
				PodAffinityTerm: corev1.PodAffinityTerm{
					TopologyKey:       "kubernetes.io/hostname",
					NamespaceSelector: &metav1.LabelSelector{},
					LabelSelector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      common.KeyApp,
								Operator: metav1.LabelSelectorOpIn,
								Values: []string{
									common.ValueAppWarehouse,
								},
							},
						},
					},
				},
			},
		},
	}
	// ensure all instances are in the same zone
	if b.warehouse.Spec.Replicas > 1 {
		podAffinity.RequiredDuringSchedulingIgnoredDuringExecution = []corev1.PodAffinityTerm{
			{
				LabelSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						common.KeyTenant:    b.tenant.Name,
						common.KeyWarehouse: b.warehouse.Name,
					},
				},
				TopologyKey: "topology.kubernetes.io/zone",
			},
		}
	}
	return podAffinity
}

func getEnvs() []corev1.EnvVar {
	envs := []corev1.EnvVar{
		{
			Name: "POD_IP",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "status.podIP",
				},
			},
		},
	}
	return envs
}

func getPorts() []corev1.ContainerPort {
	return []corev1.ContainerPort{
		{
			Name:          string(common.ServiceProtocolFlight),
			ContainerPort: int32(common.ServicePortFlight),
		},
		{
			Name:          string(common.ServiceProtocolAdmin),
			ContainerPort: int32(common.ServicePortAdmin),
		},
		{
			Name:          string(common.ServiceProtocolMetrics),
			ContainerPort: int32(common.ServicePortMetrics),
		},

		{
			Name:          string(common.ServiceProtocolMySQL),
			ContainerPort: int32(common.ServicePortMySQL),
		},
		{
			Name:          string(common.ServiceProtocolCKHttp),
			ContainerPort: int32(common.ServicePortCKHttp),
		},
		{
			Name:          string(common.ServiceProtocolQuery),
			ContainerPort: int32(common.ServicePortQuery),
		},
		{
			Name:          string(common.ServiceProtocolFlightSQL),
			ContainerPort: int32(common.ServicePortFlightSQL),
		},
	}
}

func getVolumeMounts() []corev1.VolumeMount {
	mnts := []corev1.VolumeMount{
		{
			Name:      "query-config",
			MountPath: "/etc/config",
		},
		{
			Name:      "tmp",
			MountPath: "/tmp",
		},
	}
	return mnts
}

func patchQueryPodWithCache(tpl *corev1.PodTemplateSpec, tn *databendv1alpha1.Tenant, wh *databendv1alpha1.Warehouse) {
	settings := resource.GetCacheSettings(tn, wh)
	if settings == nil {
		return
	}
	sizeLimit := kresource.MustParse(settings.K8sResourceLimit)
	cacheVolume := corev1.Volume{
		Name: settings.VolumeName,
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{
				SizeLimit: &sizeLimit,
			},
		},
	}
	tpl.Spec.Volumes = append(tpl.Spec.Volumes, cacheVolume)
	for idx := range tpl.Spec.Containers {
		container := tpl.Spec.Containers[idx]
		if container.Name != common.DatabendQueryContainerName {
			continue
		}
		container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
			Name:      settings.VolumeName,
			MountPath: settings.Path,
		})
		tpl.Spec.Containers[idx] = container
	}
}

func copyMap(src map[string]string) map[string]string {
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func copyTolerations(src []corev1.Toleration) []corev1.Toleration {
	dst := make([]corev1.Toleration, len(src))
	copy(dst, src)
	return dst
}
