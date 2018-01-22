package xor

import (
	"fmt"
)

const (
	keyParam = "key"
	dstParam = "dst"
)

// SizeError records an error when the size of a parameter does not match the
// the constraints of the function
type SizeError struct {
	size  int
	param string
}

func (s SizeError) Error() string {
	return fmt.Sprintf("crypto/xor: invalid size %d for parameter %s", s.size, s.param)
}

// EncryptRepeatingKey XORs each byte from src with bytes from key as a
// repeating key, storing the result in dst. Dst and src may point to the same
// memory. EncryptRepeatingKey expects key to have a non-zero length and dst to
// be at least as long as src, retuning errors otherwise. EncryptRepeatingKey is
// a more general XOR cipher and the functions EncryptFixedKey and
// EncryptByteKey are provided as helpers.
func EncryptRepeatingKey(dst, src, key []byte) error {
	if len(key) == 0 {
		return SizeError{0, keyParam}
	}
	if len(dst) < len(src) {
		return SizeError{len(dst), dstParam}
	}

	keyLen := len(key)
	for i := range src {
		dst[i] = src[i] ^ key[i%keyLen]
	}

	return nil
}

// EncryptFixedKey is the same as EncryptRepeatingKey with the added constraint
// that key is at least as long as src.
func EncryptFixedKey(dst, src, key []byte) error {
	if len(dst) > len(key) {
		return SizeError{len(key), keyParam}
	}

	return EncryptRepeatingKey(dst, src, key)
}

// EncryptByteKey is the same as EncryptRepeatingKey, execpt key is passed in as
// a byte slice with a length of 1
func EncryptByteKey(dst, src []byte, key byte) error {
	return EncryptRepeatingKey(dst, src, []byte{key})
}
