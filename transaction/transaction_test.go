package transaction

import (
	"math/rand"
	"testing"

	"github.com/PMoneda/gcoin/utils"
)

func TestShouldCreateNewTransaction(t *testing.T) {
	accountA := utils.NewPrivateKey()
	accountDestiny := utils.GetPublicKey(utils.NewPrivateKey())
	//Quero transferir 1.5 para outra wallet
	amount := rand.Float64() * 10
	tx, err := PrepareTransfer(accountA, accountDestiny, amount)
	if err != nil {
		if err.Error() != "no sufficient funds" {
			t.Fail()
		}
		return
	}
	totalTransferOutput := 0.0
	totalTransferInput := 0.0
	for _, out := range tx.Outputs {
		totalTransferOutput += out.Amount
	}
	for _, in := range tx.Inputs {
		totalTransferInput += in.Amount
		if in.LockedTo == "" {
			t.Fail()
		}
	}
	if totalTransferOutput != totalTransferInput {
		t.Fail()
	}
	hasNoDestWallet := true
	for _, out := range tx.Outputs {
		if out.Address == accountDestiny {
			hasNoDestWallet = false
			if out.Amount != amount {
				t.Fail()
			}
		}
	}
	if hasNoDestWallet {
		t.Fail()
	}
}
