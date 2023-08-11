package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

type AES struct {
	// Create an AES object
}

func (A *AES) Encrypt(key []byte, message string) ([]byte, []byte) {
	plaintext := []byte(message)

	if len(plaintext)%aes.BlockSize != 0 {
		panic("The message is not a multiple of the block size, 16-bit")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// Create a new HMAC with the hash-function SHA256
	hmac := hmac.New(sha256.New, key)
	hmac.Write(ciphertext)

	// Return the HMAC as a byte array
	mac := hmac.Sum(nil)

	return mac, ciphertext
}

func (A *AES) Decrypt(key []byte, mac []byte, ciphertext []byte) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmac_test := hmac.New(sha256.New, key)

	// Write data to it
	hmac_test.Write(ciphertext)

	// Get result and encode as byte array
	mac_test := hmac_test.Sum(nil)

	if !bytes.Equal(mac, mac_test) {
		fmt.Println("Message authentication failed!")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("The ciphertext is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	return string(ciphertext)
}
