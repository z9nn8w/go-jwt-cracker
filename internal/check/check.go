package check

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"
)

func CheckHMAC(alg, key, token string) bool {
	parts := strings.Split(token, ".")
	message := parts[0] + "." + parts[1]
	signature := parts[2]

	// 解码 Base64URL 格式的签名
	sigBytes, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("Error decoding signature:", err)
		return false
	}

	// 根据算法计算 HMAC
	switch alg {
	case "HS256":
		mac := hmac.New(sha256.New, []byte(key))
		mac.Write([]byte(message))
		calculatedSignature := mac.Sum(nil)
		return subtle.ConstantTimeCompare(calculatedSignature, sigBytes) == 1
	case "HS384":
		mac := hmac.New(sha512.New384, []byte(key))
		mac.Write([]byte(message))
		calculatedSignature := mac.Sum(nil)
		return subtle.ConstantTimeCompare(calculatedSignature, sigBytes) == 1
	case "HS512":
		mac := hmac.New(sha512.New, []byte(key))
		mac.Write([]byte(message))
		calculatedSignature := mac.Sum(nil)
		return subtle.ConstantTimeCompare(calculatedSignature, sigBytes) == 1
	default:
		fmt.Println("the algorithm is not supported")
		return false
	}
}
