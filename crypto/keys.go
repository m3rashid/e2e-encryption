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
