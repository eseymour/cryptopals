package xor

import (
	"fmt"
)

// KeySizeError records an error when a key with a size not matching function
// constraints is passed.
type KeySizeError struct {
	actual, expected int
}

func (k KeySizeError) Error() string {
	return fmt.Sprintf("crypto/xor: key length %d and plaintext length %d do not match",
		k.actual, k.expected)
}

// Encrypt returns the result of XORing the byte slice plaintext with the byte
// slice key as a repeating key. This is a more general XOR cipher and the
// functions XORBytesFixedLengthKey and XORBytesSingleByteKey are provided as
// helpers.
func Encrypt(plaintext, key []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))
	keyLen := len(key)

	for i := range plaintext {
		ciphertext[i] = plaintext[i] ^ key[i%keyLen]
	}

	return
}

// EncryptFixedLengthKey returns the result of XORing the two equal length byte
// slices plaintext and key. XORBytesFixedLengthKey fails and retunrs a non-nil
// error err when the lengths of plaintext and key do not match.
func EncryptFixedLengthKey(plaintext, key []byte) (ciphertext []byte, err error) {
	if len(plaintext) != len(key) {
		err = KeySizeError{len(key), len(plaintext)}
		return
	}

	ciphertext = Encrypt(plaintext, key)
	return
}

// EncryptSingleByteKey returns the result of XORing the byte slice plaintext
// with a single byte key.
func EncryptSingleByteKey(plaintext []byte, key byte) (ciphertext []byte) {
	ciphertext = Encrypt(plaintext, []byte{key})
	return
}
