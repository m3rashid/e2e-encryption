package utils

import "encoding/json"

type RsaPrivateKeyParameters struct {
	D        []byte
	P        []byte
	Q        []byte
	DP       []byte
	DQ       []byte
	InverseQ []byte
	Modulus  []byte
	Exponent []byte
}

func (keyParameters RsaPrivateKeyParameters) ToJson() string {
	jsonBytes, _ := json.Marshal(keyParameters)
	return string(jsonBytes)
}
