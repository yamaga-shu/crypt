package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/yamaga-shu/crypt/block-cipher-mode/pkcs7"
)

// encrypt DES with CBC mode
func encryptDES(key, iv, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = pkcs7.Pad(plaintext, block.BlockSize())

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// decrypt DES with CBC mode
func decryptDES(key, iv, ciphertext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)
	return pkcs7.Unpad(plaintext, block.BlockSize())
}

func main() {
	// Generate a secure random 8-byte key for DES
	key := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	fmt.Printf("key: %x\n", key)
	// Output:
	// key: 5692f2be14ed6cfe

	// IV should be same size as block size
	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	fmt.Printf("iv: %x\n", iv)
	// Output:
	// iv: 9ab302f44b55a6c2

	plain := "Hello, DES!"
	fmt.Printf("Original text: %s\n", plain)
	// Output:
	// Original text: Hello, DES!

	// Encrypt
	encrypted, err := encryptDES(key, iv, []byte(plain))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", encrypted)
	// Output:
	// Encrypted: 3f632d82a60fec682abb2edcfaa5851d

	// Decrypt
	decrypted, err := decryptDES(key, iv, encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
	// Output:
	// Decrypted: Hello, DES!
}
