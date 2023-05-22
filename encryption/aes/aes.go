package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

func NewKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func NewIv() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}

func Encrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedData := padData(data, block.BlockSize())
	ciphertext := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedData)

	return ciphertext, nil
}

func Decrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decryptedData := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decryptedData, ciphertext)

	unpaddedData, err := unpadData(decryptedData)
	if err != nil {
		return nil, err
	}

	return unpaddedData, nil
}

func padData(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	pad := byte(padding)
	for i := 0; i < padding; i++ {
		data = append(data, pad)
	}
	return data
}

func unpadData(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, errors.New("invalid padding")
	}
	return data[:length-unpadding], nil
}
