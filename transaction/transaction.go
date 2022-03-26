package transaction

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PMoneda/gcoin/utils"
	"github.com/btcsuite/btcutil/base58"
)

type TxDirection struct {
	TxID     string  `json:"transaction_id"`
	Amount   float64 `json:"amount"`
	Address  string  `json:"address"`
	LockedTo string  `json:"locked_to"`
}

type Transaction struct {
	TxID     string         `json:"transaction_id"`
	Amount   float64        `json:"amount"`
	Inputs   []*TxDirection `json:"inputs"`
	Outputs  []*TxDirection `json:"outputs"`
	AuthHash string         `json:"auth_hash"`
}

func NewTransaction() *Transaction {
	tx := &Transaction{}
	tx.Inputs = make([]*TxDirection, 0)
	tx.Outputs = make([]*TxDirection, 0)
	return tx
}

func (tx *Transaction) GetAmount() float64 {
	return tx.Amount
}

func UnMarshalTransaction(authenticatedTx string) (*Transaction, error) {
	parts := strings.Split(authenticatedTx, ":")
	txStr := parts[0]
	tx := &Transaction{}
	err := json.Unmarshal(base58.Decode(txStr), tx)
	return tx, err
}

func PrepareTransfer(pkOrigin string, destWallet string, amount float64) (*Transaction, error) {
	srcWallet := utils.GetPublicKey(pkOrigin)
	tx := Transaction{}
	tx.Amount = amount
	txId, err := utils.GetRandomByteString(32)
	if err != nil {
		return nil, err
	}
	tx.TxID = txId
	tx.Inputs = GetInputsFromWallet(srcWallet, amount)

	amountTransfer := 0.0
	tx.Outputs = make([]*TxDirection, 1)
	destTx := &TxDirection{
		TxID:    tx.TxID,
		Address: destWallet,
		Amount:  amount,
	}
	tx.Outputs[0] = destTx
	for _, srcUTXO := range tx.Inputs {
		if amountTransfer < amount {
			srcUTXO.LockedTo = destWallet
			amountTransfer += srcUTXO.Amount
		}
	}
	if amountTransfer > amount {
		//tem troco
		returnTx := &TxDirection{
			TxID:    tx.TxID,
			Address: srcWallet,
			Amount:  amountTransfer - amount,
		}
		tx.Outputs = append(tx.Outputs, returnTx)
	} else {
		return nil, fmt.Errorf("no sufficient funds")
	}
	return &tx, nil
}
