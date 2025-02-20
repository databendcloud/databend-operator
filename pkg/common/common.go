package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

func GetQueryConfigMapName(tnName, whName string) string {
	return fmt.Sprintf("query-cm-%s-%s", tnName, whName)
}

func GetQueryStatefulSetName(tnName, whName string) string {
	return fmt.Sprintf("query-sts-%s-%s", tnName, whName)
}

func GetQueryImage(wh *v1alpha1.Warehouse) string {
	if wh.Spec.QueryImage != "" {
		return wh.Spec.QueryImage
	}
	return "databend/databend-query:latest"
}

func GetTenantServiceAccountName(tnName string) string {
	return fmt.Sprintf("databend-tenant-%s", tnName)
}

func SHA256String(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
