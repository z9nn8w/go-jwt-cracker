package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// ParseJWT 解析 JWT Token，输出header和payload信息，并返回header中的alg
func ParseJWT(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid JWT format or alg is none")
	}

	header := parts[0]
	payload := parts[1]
	/* signature := parts[2] */

	// Base64URL解码 header和payload
	headerDecoded, err := base64Decode(header)
	if err != nil {
		return "", err
	}

	payloadDecoded, err := base64Decode(payload)
	if err != nil {
		return "", err
	}

	fmt.Printf("Header: %s\n", string(headerDecoded))
	fmt.Printf("Payload: %s\n", string(payloadDecoded))

	// 返回解码后header中的alg
	var headerStruct struct {
		Alg string `json:"alg"`
	}
	if err := json.Unmarshal(headerDecoded, &headerStruct); err != nil {
		return "", fmt.Errorf("failed to parse header JSON: %v", err)
	}

	if headerStruct.Alg == "" {
		return "", fmt.Errorf("alg field not found in header")
	}

	return headerStruct.Alg, nil
}

func base64Decode(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
