package sender

import (
	"crypto/rsa"
	"crypto/sha256"
)

func Sign(data []byte, pk rsa.PrivateKey) {
	hashed := sha256.Sum256(data)
	utils.
}
