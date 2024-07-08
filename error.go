package encrypto

import "errors"

var (
	ErrCipherCreation = errors.New("failed to create a new cipher")
	ErrEncryption     = errors.New("failed to encrypt")
	ErrBase64Decode   = errors.New("failed to decode to base64")
	ErrBlockSize      = errors.New("invalid ciphertext block size")
)
