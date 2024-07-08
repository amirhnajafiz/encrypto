package encrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func Encrypt(key string, message string) (string, error) {
	byteKey := []byte(key)
	byteMsg := []byte(message)

	block, err := aes.NewCipher(byteKey)
	if err != nil {
		return err.Error(), ErrCipherCreation
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return err.Error(), ErrEncryption
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(key string, message string) (string, error) {
	byteKey := []byte(key)

	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return err.Error(), ErrBase64Decode
	}

	block, err := aes.NewCipher(byteKey)
	if err != nil {
		return err.Error(), ErrCipherCreation
	}

	if len(cipherText) < aes.BlockSize {
		return "", ErrBlockSize
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
