package utils

import (
	"testing"
)

func TestDigitalSignature(t *testing.T) {
	pair := NewKeyPair()

	signature, err := Sign([]byte("Hello World"), pair.PrivateKey)
	if err != nil {
		t.Fail()
	}
	if !Verify([]byte("Hello World"), signature) {
		t.Fail()
	}
}
