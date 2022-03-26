package utils

import "testing"

func TestShouldGetRandomBytes(t *testing.T) {
	byteString, err := GetRandomByteString(32)
	if err != nil {
		t.Fail()
	}
	if len(byteString) != 64 {
		t.Fail()
	}
}
