package transaction

import (
	"encoding/json"
	"testing"

	"github.com/PMoneda/gcoin/utils"
	"github.com/btcsuite/btcutil/base58"
)

func TestShouldAuthenticateTransaction(t *testing.T) {
	accountA := utils.NewPrivateKey()
	accountB := utils.NewPrivateKey()
	accountC := utils.NewPrivateKey()

	tx, err := PrepareTransfer(accountA, utils.GetPublicKey(accountB), 10)
	if err != nil {
		t.Fail()
	}
	//transfer money from A to B
	authenticated, err := AuthenticateTransaction(tx, accountA, utils.GetPublicKey(accountB))
	if err != nil {
		t.Fail()
	}
	if !VerifyTransaction(authenticated) {
		t.Fail()
	}
	tx2, err := UnMarshalTransaction(authenticated)
	if err != nil {
		t.Fail()
	}
	//emulate fraud
	tx2.Outputs[0].Address = utils.GetPublicKey(accountC)
	bb, _ := json.Marshal(tx2)
	base58.Encode(bb)
	fraud, _ := AuthenticateTransaction(tx2, accountC, utils.GetPublicKey(accountC))
	if VerifyTransaction(fraud) {
		t.Fail()
	}
}
