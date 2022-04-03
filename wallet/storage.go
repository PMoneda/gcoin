package wallet

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

var dbWallet *bolt.DB

func openWallet() (*bolt.DB, error) {
	if dbWallet != nil {
		return dbWallet, nil
	}
	dbWallet, err := bolt.Open("./wallet.db", 0666, nil)
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
	tx, err := wallet.Begin(false)
	if err != nil {
		return "", err
	}
	bucket := tx.Bucket([]byte(bucketName))
	value := bucket.Get([]byte(key))
	return string(value), nil
}

func ListContacts() ([]string, error) {
	return list("contacts")
}

func list(bucketName string) ([]string, error) {
	contactList := make([]string, 0)
	wallet, err := openWallet()
	if err != nil {
		return nil, err
	}
	tx, err := wallet.Begin(false)
	if err != nil {
		return nil, err
	}
	bucket := tx.Bucket([]byte(bucketName))
	bucket.ForEach(func(k, v []byte) error {
		contactList = append(contactList, fmt.Sprintf("%s:%s", string(k), string(v)))
		return nil
	})
	return contactList, nil
}

func DeleteAliasContact(alias string) error {
	wallet, err := openWallet()
	if err != nil {
		return err
	}
	tx, err := wallet.Begin(true)
	if err != nil {
		return err
	}
	bucket := tx.Bucket([]byte("contacts"))
	err = bucket.Delete([]byte(alias))
	if err != nil {
		return err
	}
	return tx.Commit()
}
