package utils

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Int64ToByteArray(value int64) []byte {
	arr := make([]byte, 8)
	binary.LittleEndian.PutUint64(arr, uint64(value))
	return arr
}

func Int64ToByteArrayNBytes(value int64, n int) []byte {
	arr := make([]byte, n)
	binary.LittleEndian.PutUint64(arr, uint64(value))
	return arr
}

func Int32ToByteArray(value int32) []byte {
	arr := make([]byte, 4)
	binary.LittleEndian.PutUint32(arr, uint32(value))
	return arr
}

func Int32ToByteArrayNBytes(value int32, n int32) []byte {
	arr := make([]byte, n)
	binary.LittleEndian.PutUint32(arr, uint32(value))
	return arr
}

func Int64FromArray(array []byte) int64 {
	return int64(binary.LittleEndian.Uint64(array))
}

func Int32FromArray(array []byte) int32 {
	return int32(binary.LittleEndian.Uint32(array))
}

func GetRandomByteString(size int) (string, error) {
	client := http.Client{}
	resp, err := client.Get(fmt.Sprintf("https://www.random.org/cgi-bin/randbyte?nbytes=%d&format=h", size))
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	key := string(body)
	key = strings.ReplaceAll(key, " ", "")
	key = strings.ReplaceAll(key, "\n", "")
	return key, nil
}
