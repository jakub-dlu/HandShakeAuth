package util

import (
	"crypto/rand"
	"io"
)

// generates random bytes (n is size)
func RandBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
