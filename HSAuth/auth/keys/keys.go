package keys

import (
	"HSAuth/util"
	"golang.org/x/crypto/curve25519"
)

// generates key pair
func GenerateKeyPair() ([]byte, []byte) {
	privateKey, err := util.RandBytes(32)
	if err != nil {
		panic(err)
	}

	publicKey, err := curve25519.X25519(privateKey[:], curve25519.Basepoint)
	if err != nil {
		panic(err)
	}

	return publicKey, privateKey
}
