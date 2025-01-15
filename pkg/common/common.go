package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GetQueryConfigMapName(whName string) string {
	return fmt.Sprintf("query-%s", whName)
}

func SHA256String(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
