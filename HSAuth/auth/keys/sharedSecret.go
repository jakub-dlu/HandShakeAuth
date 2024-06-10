package keys

import (
	"HSAuth/models"
	"bytes"
	"golang.org/x/crypto/curve25519"
)

// compares sharedSecrets using parameters:
// user shared secret, user public key and key pair from data base (using one of them)
func CompareSharedSecrets(sharedSecretV1 []byte, usersPublicKey []byte, k models.KeysModel) (bool, error) {
	sharedSecretV2, err := curve25519.X25519(k.PrivateKey[:], usersPublicKey[:])
	if err != nil {
		return false, err
	}

	return bytes.Equal(sharedSecretV1, sharedSecretV2), nil
}
