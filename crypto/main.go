package utils

import (
	"crypto/rsa"
	"encoding/binary"
	"math/big"
)

func (privateKey *RsaPrivateKey) ToRsaPrivateKeyParameters() *RsaPrivateKeyParameters {
	exponent := make([]byte, 4)
	binary.BigEndian.PutUint32(exponent, uint32(privateKey.PublicKey.E))
	for i := range exponent {
		if exponent[i] != 0 {
			exponent = exponent[i:]
			break
		}
	}

	return &RsaPrivateKeyParameters{
		D:        privateKey.D.Bytes(),
		P:        privateKey.Primes[0].Bytes(),
		Q:        privateKey.Primes[1].Bytes(),
		DP:       privateKey.Precomputed.Dp.Bytes(),
		DQ:       privateKey.Precomputed.Dq.Bytes(),
		InverseQ: privateKey.Precomputed.Qinv.Bytes(),
		Modulus:  privateKey.PublicKey.N.Bytes(),
		Exponent: exponent,
	}
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
