package utils

import (
	"crypto/rsa"
	"encoding/binary"
	"encoding/json"
	"math/big"
)

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

func (keyParameters RsaPrivateKeyParameters) ToRsaPrivateKey() (*rsa.PrivateKey, error) {
	d, p, q := new(big.Int), new(big.Int), new(big.Int)
	d.SetBytes(keyParameters.D)
	p.SetBytes(keyParameters.P)
	q.SetBytes(keyParameters.Q)
	dp, dq, inverseQ, modulus := new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	dp.SetBytes(keyParameters.DP)
	dq.SetBytes(keyParameters.DQ)
	inverseQ.SetBytes(keyParameters.InverseQ)
	modulus.SetBytes(keyParameters.Modulus)

	buffer := make([]byte, 4)
	copy(buffer[4-len(keyParameters.Exponent):], keyParameters.Exponent)
	e := binary.BigEndian.Uint32(buffer)

	return &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: modulus,
			E: int(e),
		},
		D:      d,
		Primes: []*big.Int{p, q},
		Precomputed: rsa.PrecomputedValues{
			Dp:   dp,
			Dq:   dq,
			Qinv: inverseQ,
		},
	}, nil
}
