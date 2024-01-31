package token

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)


func GenerateAESString(key string, plaintext string) (string, error) {

	keyBytes := []byte(key)


	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	plaintextBytes := []byte(plaintext)
	blockSize := block.BlockSize()
	padding := blockSize - len(plaintextBytes)%blockSize
	padText := append(plaintextBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)


	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(padText))
	mode.CryptBlocks(ciphertext, padText)

	result := append(iv, ciphertext...)

	return base64.StdEncoding.EncodeToString(result), nil
}

func DecodeAESString(key string, encryptedText string) (string, error) {
	keyBytes := []byte(key)

	encBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	iv := encBytes[:aes.BlockSize]

	mode := cipher.NewCBCDecrypter(block, iv)

	decrypted := make([]byte, len(encBytes)-aes.BlockSize)
	mode.CryptBlocks(decrypted, encBytes[aes.BlockSize:])

	padding := decrypted[len(decrypted)-1]
	decrypted = decrypted[:len(decrypted)-int(padding)]

	return string(decrypted), nil
}