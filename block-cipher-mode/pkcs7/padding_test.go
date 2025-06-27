package pkcs7

import (
	"bytes"
	"testing"
)

func TestPKCS7PaddingUnpadding(t *testing.T) {
	blockSize := 16
	original := []byte("YELLOW SUBMARINE")
	padded := Pad(original, blockSize)

	// The padded data should be a multiple of blockSize
	if len(padded)%blockSize != 0 {
		t.Errorf("Padded data length is not multiple of block size")
	}

	unpadded, err := Unpad(padded, blockSize)
	if err != nil {
		t.Errorf("Unpad failed: %v", err)
	}

	if !bytes.Equal(unpadded, original) {
		t.Errorf("Unpadded data does not match original. Got %v, want %v", unpadded, original)
	}
}

func TestInvalidPadding(t *testing.T) {
	// Creating data with invalid padding
	data := []byte{0x01, 0x02, 0x03, 0x04}
	blockSize := 8

	// Force invalid padding by tampering
	dataWithInvalidPadding := append(data, byte(255))
	_, err := Unpad(dataWithInvalidPadding, blockSize)
	if err == nil {
		t.Error("Expected error for invalid padding, but got none")
	}
}
