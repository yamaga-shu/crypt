package pkcs7

import (
	"bytes"
	"errors"
)

// Pad adds PKCS#7 padding to the input data to make its length a multiple of blockSize.
func Pad(data []byte, blockSize int) []byte {
	if blockSize <= 0 || 255 < blockSize {
		panic("Invalid block size")
	}
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Unpad removes PKCS#7 padding from the input data.
// It returns an error if the padding is invalid.
func Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, errors.New("invalid data length")
	}
	padding := int(data[len(data)-1])
	if padding == 0 || padding > blockSize {
		return nil, errors.New("invalid padding size")
	}
	for i := 0; i < padding; i++ {
		if data[len(data)-1-i] != byte(padding) {
			return nil, errors.New("invalid padding")
		}
	}
	return data[:len(data)-padding], nil
}
