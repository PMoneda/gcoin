package utils

import (
	"testing"
)

func TestDigitalSignature(t *testing.T) {
	privateKey := NewPrivateKey()

	signature, err := Sign([]byte("Hello World"), privateKey)
	if err != nil {
		t.Fail()
	}
	if !Verify([]byte("Hello World"), signature) {
		t.Fail()
	}
}
