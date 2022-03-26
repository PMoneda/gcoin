package transaction

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PMoneda/gcoin/utils"
	"github.com/btcsuite/btcutil/base58"
)

func AuthenticateTransaction(transaction *Transaction, pk string, pubKey string) (string, error) {
	originWallet := utils.GetPublicKey(pk)
	for _, in := range transaction.Inputs {
		if in.Address != originWallet {
			return "", fmt.Errorf("cannot authenticate transactions with input adresses different from signer")
		}
	}
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s:%s", pk, pubKey)))
	authHash, err := utils.Sign(sum[:], pk)
	if err != nil {
		return "", err
	}
	transaction.AuthHash = base58.Encode([]byte(authHash))
	bytearray, err := json.Marshal(transaction)
	if err != nil {
		return "", err
	}
	b := base58.Encode(bytearray)
	txHash := sha256.Sum256(bytearray)
	signature, err := utils.Sign([]byte(txHash[:]), pk)
	if err != nil {
		return "", err
	}
	block := fmt.Sprintf("%s:%s", b, signature)
	return block, nil
}

func VerifyTransaction(transaction string) bool {
	if transaction == "" {
		return false
	}
	parts := strings.Split(transaction, ":")
	txData := parts[0]
	bytearray := base58.Decode(txData)
	txHash := sha256.Sum256(bytearray)
	signature := parts[1]
	return utils.Verify(txHash[:], signature)
}
