package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
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

func main() {
	signature := "2b6bd834bb5c2a35166483c64e654749dd154fe1"
	secret := "secret"
	time := "1610698357000"
	nonce := "1234567890"
	result := VerifyMessageContent(signature, secret, time, nonce)
	fmt.Println(result)
}
