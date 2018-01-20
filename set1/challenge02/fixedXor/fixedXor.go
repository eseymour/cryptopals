package fixedXor

import "errors"

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
