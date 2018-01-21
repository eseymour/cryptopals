package fixedXOR

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/eseymour/cryptopals/pkg/crypto/xor"
	b "github.com/eseymour/cryptopals/pkg/util/bytes"
)

// TestFixedXOR only tests the case from the challenge. The package
// github.com/eseymour/cryptopals/pkg/crypto/xor has more extensive tests.
func TestFixedXOR(t *testing.T) {

	plaintext := b.Must(hex.DecodeString("1c0111001f010100061a024b53535009181c"))
	key := b.Must(hex.DecodeString("686974207468652062756c6c277320657965"))
	want := b.Must(hex.DecodeString("746865206b696420646f6e277420706c6179"))
	fails := false

	got, err := xor.EncryptFixedKey(plaintext, key)
	failed := err != nil

	switch {
	case failed != fails:
		t.Errorf("xor.EncryptFixedKey(%#v, %#v) failure: %#v, want %#v", plaintext, key, failed, fails)
	case !failed && !bytes.Equal(got, want):
		t.Errorf("xor.EncryptFixedKey(%#v, %#v) == %#v, want %#v", plaintext, key, got, want)
	}
}
