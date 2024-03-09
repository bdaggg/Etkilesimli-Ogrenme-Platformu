package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPass(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	hashInBytes := hasher.Sum(nil)
	hash := hex.EncodeToString(hashInBytes)

	return hash
}
