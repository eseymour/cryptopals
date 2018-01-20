package fixedXor

import "errors"

// FixedXor returns a byte slice containing the result of xoring the byte slices
// plaintext and key. FixedXor expects plaintext and key to have the same
// length, failing otherwise.
func FixedXor(plaintext, key []byte) (ciphertext []byte, err error) {
	if len(plaintext) != len(key) {
		err = errors.New("fixedXor: key and plaintext are not the same length")
		return
	}

	ciphertext = make([]byte, len(plaintext))
	for i := range plaintext {
		ciphertext[i] = plaintext[i] ^ key[i]
	}

	return
}
