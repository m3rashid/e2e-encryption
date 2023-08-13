package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
)

type RsaCrypto struct{}

func (crypto RsaCrypto) GenerateKeyPair(keySize int) (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return "", "", err
	}

	var rsaPrivateKey RsaPrivateKey = RsaPrivateKey(*privateKey)
	rsaPrivateKeyParameters := rsaPrivateKey.ToRsaPrivateKeyParameters()

	var rsaPublicKey RsaPublicKey = RsaPublicKey(privateKey.PublicKey)
	rsaPublicKeyParameters, err := rsaPublicKey.ToRsaPublicKeyParameters()
	if err != nil {
		return "", "", err
	}

	return rsaPrivateKeyParameters.ToJson(), rsaPublicKeyParameters.ToJson(), nil
}

func (crypto RsaCrypto) Encrypt(plainText string, publicKeyJson string) (string, error) {
	var rsaPublicKeyParameters RsaPublicKeyParameters
	jsonBytes := []byte(publicKeyJson)
	err := json.Unmarshal(jsonBytes, &rsaPublicKeyParameters)
	if err != nil {
		return "", err
	}
	publicKey, err := rsaPublicKeyParameters.ToRsaPublicKey()
	if err != nil {
		return "", err
	}

	hash := sha256.New()
	plainTextBytes := []byte(plainText)
	cipherText, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, plainTextBytes, nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (crypto RsaCrypto) Decrypt(cipherText string, privateKeyJson string, provider string) (string, error) {
	var rsaPrivateKeyParameters RsaPrivateKeyParameters
	jsonBytes := []byte(privateKeyJson)
	err := json.Unmarshal(jsonBytes, &rsaPrivateKeyParameters)
	if err != nil {
		return "", err
	}
	privateKey, err := rsaPrivateKeyParameters.ToRsaPrivateKey()
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	hash := sha256.New()
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, data, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func (rsaCrypto RsaCrypto) SignData(data string, privateKeyJson string) (string, error) {
	var rsaPrivateKeyParameters RsaPrivateKeyParameters
	jsonBytes := []byte(privateKeyJson)
	err := json.Unmarshal(jsonBytes, &rsaPrivateKeyParameters)
	if err != nil {
		return "", err
	}
	signatureKey, err := rsaPrivateKeyParameters.ToRsaPrivateKey()
	if err != nil {
		return "", err
	}

	dataToSign := []byte(data)
	hashed := sha512.Sum512(dataToSign)
	signature, err := rsa.SignPKCS1v15(rand.Reader, signatureKey, crypto.SHA512, hashed[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func (rsaCrypto RsaCrypto) VerifySignature(data string, signature string, publicKeyJson string) (bool, error) {
	var rsaPublicKeyParameters RsaPublicKeyParameters
	jsonBytes := []byte(publicKeyJson)
	err := json.Unmarshal(jsonBytes, &rsaPublicKeyParameters)
	if err != nil {
		return false, err
	}
	signatureKey, err := rsaPublicKeyParameters.ToRsaPublicKey()
	if err != nil {
		return false, err
	}

	dataToVerify := []byte(data)
	hashed := sha512.Sum512(dataToVerify)
	binarySignature, _ := base64.StdEncoding.DecodeString(signature)

	verifyErr := rsa.VerifyPKCS1v15(signatureKey, crypto.SHA512, hashed[:], binarySignature)
	if verifyErr != nil {
		return false, verifyErr
	}

	return true, nil
}
