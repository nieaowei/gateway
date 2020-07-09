package public

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenSha256BySecret(str, secret string) string {
	hash := sha256.New()
	hash.Write([]byte(str))

	hash1 := sha256.New()
	hash1.Write(append(hash.Sum(nil), []byte(secret)...))

	return hex.EncodeToString(hash1.Sum(nil))
}
