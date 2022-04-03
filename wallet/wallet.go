package wallet

import (
	"github.com/PMoneda/gcoin/utils"
)

type Wallet struct {
	PrivateKey string
	PublicKey  string
}

type WalletStatus struct {
	PublicKey string
}

func CreateNewWallet() Wallet {
	pk := utils.NewPrivateKey()
	w := Wallet{
		PrivateKey: pk,
		PublicKey:  utils.GetPublicKey(pk),
	}
	return w
}

//SignIn user's pk on the wallet
func SignIn(pk string) error {
	return saveData("me", "private_key", pk)
}

//AddTrustedContact add contact to wallet storage
func AddTrustedContact(pubKey string, alias string) error {
	return SaveContact(alias, pubKey)
}

func ListAllContacts() ([]string, error) {
	return ListContacts()
}

func DeleteContact(alias string) error {
	return DeleteAliasContact(alias)
}

func Status() (*WalletStatus, error) {
	ws := &WalletStatus{}
	privK, err := getData("me", "private_key")
	if err != nil {
		return nil, err
	}
	ws.PublicKey = utils.GetPublicKey(privK)
	return ws, nil
}
