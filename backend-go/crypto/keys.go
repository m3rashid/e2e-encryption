package utils

import (
	"crypto/rsa"
	"encoding/binary"
)

type RsaPrivateKey rsa.PrivateKey
type RsaPublicKey rsa.PublicKey

func (publicKey *RsaPublicKey) ToRsaPublicKeyParameters() (*RsaPublicKeyParameters, error) {
	exponent := make([]byte, 4)
	binary.BigEndian.PutUint32(exponent, uint32(publicKey.E))
	for i := range exponent {
		if exponent[i] != 0 {
			exponent = exponent[i:]
			break
		}
	}

	return &RsaPublicKeyParameters{
		Modulus:  publicKey.N.Bytes(),
		Exponent: exponent,
	}, nil
}

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
