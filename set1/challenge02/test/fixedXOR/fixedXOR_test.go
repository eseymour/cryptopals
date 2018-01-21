package fixedXOR

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/eseymour/cryptopals/pkg/crypto/xor"
)

// must is a helper that wraps calls returning ([]byte, error), panicking if the
// error is non-nil.
func must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}

// TestFixedXOR only tests the case from the challenge. The
// package github.com/eseymour/cryptopals/pkg/crypto/xor has more extensive
// tests.
func TestFixedXOR(t *testing.T) {
	cases := []struct {
		plaintext, key, want []byte
		fails                bool
	}{
		{ // Challenge case
			must(hex.DecodeString("1c0111001f010100061a024b53535009181c")),
			must(hex.DecodeString("686974207468652062756c6c277320657965")),
			must(hex.DecodeString("746865206b696420646f6e277420706c6179")),
			false,
		},
	}

	for _, c := range cases {
		got, err := xor.EncryptFixedKey(c.plaintext, c.key)
		failed := err != nil
		switch {
		case failed != c.fails:
			t.Errorf("xor.EncryptFixedKey(%#v, %#v) failure: %#v, want %#v", c.plaintext, c.key, failed, c.fails)
		case !c.fails && !bytes.Equal(got, c.want):
			t.Errorf("xor.EncryptFixedKey(%#v, %#v) == %#v, want %#v", c.plaintext, c.key, got, c.want)
		}
	}
}
