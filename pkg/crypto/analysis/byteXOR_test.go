package analysis

import (
	"encoding/hex"
	"testing"

	"github.com/eseymour/cryptopals/pkg/util/bytes"
)

func TestBreakXOREncryptByteKey(t *testing.T) {
	cases := []struct {
		ciphertext []byte
		want       byte
	}{
		{ // Challenge 3 spoilers
			bytes.Must(hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")),
			0x58,
		},
	}

	for _, c := range cases {
		got, _ := BreakXOREncryptByteKey(c.ciphertext)
		if got != c.want {
			t.Errorf("BreakXOREncryptByteKey(%x) == %x, want %x", c.ciphertext, got, c.want)
		}
	}
}
