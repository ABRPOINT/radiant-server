package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// EncodePublicKey encodes an RSA public key to PEM format.
func EncodePublicKey(pubKey *rsa.PublicKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pubKey),
	})
}

// DecodePublicKey decodes a PEM-encoded RSA public key.
func DecodePublicKey(pemKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemKey)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

// EncodePrivateKey encodes an RSA private key to PEM format.
func EncodePrivateKey(privKey *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	})
}

// DecodePrivateKey decodes a PEM-encoded RSA private key.
func DecodePrivateKey(pemKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemKey)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}
