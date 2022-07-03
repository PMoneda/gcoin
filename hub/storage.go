package hub

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/PMoneda/gcoin/encoders"
	"github.com/PMoneda/gcoin/utils"
	"github.com/btcsuite/btcutil/base58"
	bolt "go.etcd.io/bbolt"
)

var dbHub *bolt.DB

func init() {
	var err error
	dbHub, err = utils.OpenBoltDb("hub")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
func SaveServiceDescription(desc serviceDescription) error {

	reasons := make([]string, 0)
	if desc.DNS == "" {
		reasons = append(reasons, "DNS is Required")
	}
	if desc.Port == 0 {
		reasons = append(reasons, "PORT is Required")
	}
	if desc.Protocol == "" {
		reasons = append(reasons, "Protocol is Required")
	}
	if desc.PublicKey == "" {
		reasons = append(reasons, "Public Key is Required")
	}
	if desc.Region == "" {
		reasons = append(reasons, "Region is Required")
	}
	if desc.Version == "" {
		reasons = append(reasons, "Version is Required")
	}
	if len(reasons) > 0 {
		errorS := ""
		for _, reason := range reasons {
			errorS = fmt.Sprintf("%s,", reason)
		}
		return fmt.Errorf(errorS)
	}
	return utils.SaveBoltData(dbHub, fmt.Sprintf("service:%s", desc.Type), base58.Encode([]byte(desc.DNS)), encoders.ToJSON(desc))
}

func GetServiceDescription(serviceType string) ([]serviceDescription, error) {
	bucket := fmt.Sprintf("service:%s", serviceType)
	list, err := utils.ListBoltBucketKeys(dbHub, bucket)
	if err != nil {
		return nil, err
	}
	services := make([]serviceDescription, len(list))
	for i, key := range list {
		ds, err := utils.GetBoltData(dbHub, bucket, key)
		if err != nil {
			return nil, err
		}
		serviceDescription := new(serviceDescription)
		if err := json.Unmarshal([]byte(ds), serviceDescription); err != nil {
			return nil, err
		}
		services[i] = *serviceDescription
	}
	return services, nil

}

func DeleteServiceDescription(serviceType string, dns string) error {
	return utils.DeleteBoltData(dbHub, serviceType, dns)
}
