package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func VerifyMessageContent(signature string, secret string, time string, nonce string) bool {
	h := sha1.New()
	h.Write([]byte(secret))
	h.Write([]byte(time))
	h.Write([]byte(nonce))
	calc := hex.EncodeToString(h.Sum(nil))
	return strings.EqualFold(calc, signature)
}
