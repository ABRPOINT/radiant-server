package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// NewKey generates a new RSA private key with a key size of 4096 bits.
func NewKey() (*rsa.PrivateKey, error) {
	// Use a cryptographically secure random number generator
	// and provide additional entropy from the operating system
	privKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	return privKey, nil
}

// Encrypt encrypts the given data using RSA-OAEP with SHA-256 hash function.
func Encrypt(data []byte, pubKey *rsa.PublicKey) ([]byte, error) {
	if pubKey == nil {
		return nil, fmt.Errorf("public key is nil")
	}

	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, data, nil)
}

// Decrypt decrypts the given encrypted bytes using RSA-OAEP with SHA-256 hash function.
func Decrypt(encryptedBytes []byte, privKey *rsa.PrivateKey) ([]byte, error) {
	if privKey == nil {
		return nil, fmt.Errorf("private key is nil")
	}

	return privKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
}

// Sign signs the given data using RSA-PKCS1v15 with SHA-256 hash function.
func Sign(data []byte, privKey *rsa.PrivateKey) ([]byte, error) {
	if privKey == nil {
		return nil, fmt.Errorf("private key is nil")
	}

	hashed := sha256.Sum256(data)

	return rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed[:])
}

// Validate verifies the signature against the given message using RSA-PKCS1v15 with SHA-256 hash function.
func Validate(message []byte, sig []byte, pubKey *rsa.PublicKey) error {
	if pubKey == nil {
		return fmt.Errorf("public key is nil")
	}

	h := sha256.New()
	h.Write(message)

	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, h.Sum(nil), sig)
}
