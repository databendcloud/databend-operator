package utils

import "fmt"

func GetQueryConfigMapName(whName string) string {
	return fmt.Sprintf("query-%s", whName)
}
