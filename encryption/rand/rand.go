package random

import (
	"crypto/rand"
	"errors"
)

func GenerateByteArray(length int) ([]byte, error) {
	if length < 0 {
		return nil, errors.New("length should be non-negative")
	}

	ret := make([]byte, length)
	_, err := rand.Read(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
