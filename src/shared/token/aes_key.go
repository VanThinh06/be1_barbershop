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
	// Chuyển key thành mảng byte với độ dài là 32 (AES-256)
	keyBytes := []byte(key)

	// Tạo block cipher với key đã cho
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Thêm đệm cho plaintext nếu cần
	plaintextBytes := []byte(plaintext)
	blockSize := block.BlockSize()
	padding := blockSize - len(plaintextBytes)%blockSize
	padText := append(plaintextBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// Tạo IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Tạo một chế độ mã hóa CBC với IV
	mode := cipher.NewCBCEncrypter(block, iv)

	// Mã hóa plaintext
	ciphertext := make([]byte, len(padText))
	mode.CryptBlocks(ciphertext, padText)

	// Ghép IV và ciphertext để tạo chuỗi mã hóa cuối cùng
	result := append(iv, ciphertext...)

	// Chuyển đổi sang base64 để có chuỗi có thể đọc được
	return base64.StdEncoding.EncodeToString(result), nil
}

func DecodeAESString(key string, encryptedText string) (string, error) {
	keyBytes := []byte(key)

	encBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// Tạo block cipher với key đã cho
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Tạo IV từ 16 byte đầu tiên của chuỗi giải mã
	iv := encBytes[:aes.BlockSize]

	// Tạo một chế độ giải mã CBC với IV
	mode := cipher.NewCBCDecrypter(block, iv)

	// Giải mã ciphertext
	decrypted := make([]byte, len(encBytes)-aes.BlockSize)
	mode.CryptBlocks(decrypted, encBytes[aes.BlockSize:])

	// Loại bỏ đệm
	padding := decrypted[len(decrypted)-1]
	decrypted = decrypted[:len(decrypted)-int(padding)]

	return string(decrypted), nil
}