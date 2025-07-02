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
	// pkcs7.Padを使用
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
	// pkcs7.Unpadを使用
	return pkcs7.Unpad(plaintext, block.BlockSize())
}

func main() {
	// 8-byte key for DES
	key := []byte("12345678")
	// IV should be same size as block size
	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	plain := "Hello, DES!"

	fmt.Printf("Original text: %s\n", plain)
	// Output:
	// Original text: Hello, DES!

	// Encrypt
	ciphertext, err := encryptDES(key, iv, []byte(plain))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", ciphertext)
	// Output:
	// Encrypted: 5417c391196d2fd170f54219849d67e5

	// Decrypt
	decrypted, err := decryptDES(key, iv, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
	// Output:
	// Decrypted: Hello, DES!
}
