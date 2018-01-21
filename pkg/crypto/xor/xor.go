package xor

import "strconv"

// KeySizeError records an error when a key with a size not matching function
// constraints is passed.
type KeySizeError int

func (k KeySizeError) Error() string {
	return "crypto/xor: invalid key size " + strconv.Itoa(int(k))
}

// EncryptRepeatingKey returns the result of XORing the byte slice plaintext
// with the byte slice key as a repeating key. Encrypt expects key to have a
// non-zero length. This is a more general XOR cipher and the functions
// EncryptFixedKey and EncryptByteKey are provided as helpers.
func EncryptRepeatingKey(plaintext, key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, KeySizeError(0)
	}

	ciphertext := make([]byte, len(plaintext))
	keyLen := len(key)

	for i := range plaintext {
		ciphertext[i] = plaintext[i] ^ key[i%keyLen]
	}

	return ciphertext, nil
}

// EncryptFixedKey returns the result of XORing the two equal length byte slices
// plaintext and key. EncryptFixedKey expects keys to have non-zero
// lengths and lengths greater than or equal to to plaintext lengths.
func EncryptFixedKey(plaintext, key []byte) ([]byte, error) {
	if len(plaintext) > len(key) {
		return nil, KeySizeError(len(key))
	}

	return EncryptRepeatingKey(plaintext, key)
}

// EncryptByteKey returns the result of XORing the byte slice plaintext with a
// single byte key.
func EncryptByteKey(plaintext []byte, key byte) []byte {
	ciphertext, _ := EncryptRepeatingKey(plaintext, []byte{key})

	return ciphertext
}
