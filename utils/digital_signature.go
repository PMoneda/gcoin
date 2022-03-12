package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

type Pair struct {
	PrivateKey string
	PublicKey  string
}

func NewKeyPair() Pair {
	var pair Pair
	pk, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pair.PrivateKey = base58.Encode(pk.D.Bytes())
	pair.PublicKey = base58.Encode(elliptic.MarshalCompressed(elliptic.P256(), pk.PublicKey.X, pk.PublicKey.Y))
	return pair
}

func Sign(data []byte, privateKey string) (string, error) {
	//mount ecdsa private key based on D value
	var pri ecdsa.PrivateKey
	pri.D = new(big.Int).SetBytes(base58.Decode(privateKey))
	pri.PublicKey.Curve = elliptic.P256()
	pri.PublicKey.X, pri.PublicKey.Y = pri.PublicKey.Curve.ScalarBaseMult(pri.D.Bytes())

	r, s, err := ecdsa.Sign(rand.Reader, &pri, data)
	if err != nil {
		return "", err
	}
	pubKeyStr := base58.Encode(elliptic.MarshalCompressed(elliptic.P256(), pri.PublicKey.X, pri.PublicKey.Y))
	sig2 := fmt.Sprintf("%s.%s.%s", base58.Encode(r.Bytes()), base58.Encode(s.Bytes()), pubKeyStr)
	return sig2, nil
}

func Verify(data []byte, signature string) bool {
	parts := strings.Split(signature, ".")
	r := new(big.Int).SetBytes(base58.Decode(parts[0]))
	s := new(big.Int).SetBytes(base58.Decode(parts[1]))
	publicKey := parts[2]
	byteArray := base58.Decode(publicKey)
	var pubKey ecdsa.PublicKey
	pubKey.Curve = elliptic.P256()
	pX, pY := elliptic.UnmarshalCompressed(elliptic.P256(), byteArray)
	pubKey.X = pX
	pubKey.Y = pY

	return ecdsa.Verify(&pubKey, data, r, s)
}
