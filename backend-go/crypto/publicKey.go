package utils

import (
	"crypto/rsa"
	"encoding/binary"
	"encoding/json"
	"math/big"
)

type RsaPublicKeyParameters struct {
	Modulus  []byte
	Exponent []byte
}

func (keyParameters RsaPublicKeyParameters) ToRsaPublicKey() (*rsa.PublicKey, error) {
	modulus := new(big.Int)
	modulus.SetBytes(keyParameters.Modulus)

	buffer := make([]byte, 4)
	copy(buffer[4-len(keyParameters.Exponent):], keyParameters.Exponent)
	e := binary.BigEndian.Uint32(buffer)

	return &rsa.PublicKey{
		N: modulus,
		E: int(e),
	}, nil
}

func (keyParameters RsaPublicKeyParameters) ToJson() string {
	jsonBytes, _ := json.Marshal(keyParameters)
	return string(jsonBytes)
}
