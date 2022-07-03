package wallet

import (
	"github.com/PMoneda/gcoin/utils"
	bolt "go.etcd.io/bbolt"
)

var dbWallet *bolt.DB

func openWallet() (*bolt.DB, error) {
	if dbWallet != nil {
		return dbWallet, nil
	}
	dbWallet, err := utils.OpenBoltDb("wallet")
	return dbWallet, err

}

func SaveContact(alias string, data string) error {
	return saveData("contacts", alias, data)
}

func saveData(bucketName string, key string, value string) error {
	wallet, err := openWallet()
	if err != nil {
		return err
	}
	tx, err := wallet.Begin(true)
	if err != nil {
		return err
	}
	bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
	if err != nil {
		return err
	}
	err = bucket.Put([]byte(key), []byte(value))
	if err != nil {
		return err
	}
	return tx.Commit()
}

func getData(bucketName string, key string) (string, error) {
	wallet, err := openWallet()
	if err != nil {
		return "", err
	}
	return utils.GetBoltData(wallet, bucketName, key)
}

func ListContacts() ([]string, error) {
	wallet, err := openWallet()
	if err != nil {
		return nil, err
	}
	return utils.ListBoltBucket(wallet, "contacts")
}

func DeleteAliasContact(alias string) error {
	wallet, err := openWallet()
	if err != nil {
		return err
	}
	return utils.DeleteBoltData(wallet, "contacts", alias)
}
