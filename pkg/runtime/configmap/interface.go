package configmap

import (
	corev1 "k8s.io/api/core/v1"
)

type TomlConfig interface {
	BuildConfigMap() (*corev1.ConfigMap, error)
}
