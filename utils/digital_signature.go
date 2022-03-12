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

func NewPrivateKey() string {
	pk, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return base58.Encode(pk.D.Bytes())
}

func GetPublicKey(privateKey string) string {
	var pri ecdsa.PrivateKey
	pri.D = new(big.Int).SetBytes(base58.Decode(privateKey))
	pri.PublicKey.Curve = elliptic.P256()
	pri.PublicKey.X, pri.PublicKey.Y = pri.PublicKey.Curve.ScalarBaseMult(pri.D.Bytes())
	return base58.Encode(elliptic.MarshalCompressed(elliptic.P256(), pri.PublicKey.X, pri.PublicKey.Y))
}

func Sign(data []byte, privateKey string) (string, error) {
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
	if len(parts) < 3 {
		return false
	}
	r := new(big.Int).SetBytes(base58.Decode(parts[0]))
	s := new(big.Int).SetBytes(base58.Decode(parts[1]))
	publicKey := parts[2]
	byteArray := base58.Decode(publicKey)
	var pubKey ecdsa.PublicKey
	pubKey.Curve = elliptic.P256()
	pubKey.X, pubKey.Y = elliptic.UnmarshalCompressed(elliptic.P256(), byteArray)
	return ecdsa.Verify(&pubKey, data, r, s)
}
