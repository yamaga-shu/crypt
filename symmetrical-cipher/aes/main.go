package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/yamaga-shu/crypt/block-cipher-mode/pkcs7"
)

// encryptAES uses pkcs7.Pad for padding
func encryptAES(key, iv, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = pkcs7.Pad(plaintext, block.BlockSize())

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// decryptAES uses pkcs7.Unpad for unpadding
func decryptAES(key, iv, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
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
	// 16-byte key for AES-128
	key := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	fmt.Printf("key: %x\n", key)
	// Output:
	// key: 734d663ef93ec3c4bf04ecd0f671692e

	// IV should be same size as block size
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	fmt.Printf("iv: %x\n", iv)
	// Output:
	// iv: fe28280ec4c5e08c7af2c0d82294e7b8

	plain := "Hello, AES!"
	fmt.Printf("Original text: %s\n", plain)
	// Output:
	// Original text: Hello, AES!

	// Encrypt
	encrypted, err := encryptAES(key, iv, []byte(plain))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", encrypted)
	// Output:
	// Encrypted: 3c14f20b1285267422d4941a6c5949f3

	// Decrypt
	decrypted, err := decryptAES(key, iv, encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
	// Output:
	// Decrypted: Hello, AES!
}
