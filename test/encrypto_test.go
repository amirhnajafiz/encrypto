package test

import (
	"regexp"
	"testing"

	"github.com/amirhnajafiz/encrypto"
)

const isBase64 = "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$"

func TestEncryptDecryptMessage(t *testing.T) {
	key := "0123456789abcdef" // must be of 16 bytes for this example to work
	message := "Lorem ipsum dolor sit amet"

	encrypted, err := encrypto.Encrypt(key, message)
	if err != nil {
		t.Error(err)
	}

	flag, _ := regexp.Match(isBase64, []byte(encrypted))
	if !flag {
		t.Error("output is not based64")
	}

	decrypted, err := encrypto.Decrypt(key, encrypted)
	if err != nil {
		t.Error(err)
	}

	if message != decrypted {
		t.Error("output is incorrect!")
	}
}
