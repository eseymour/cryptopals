package repeatingXOR

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/eseymour/cryptopals/pkg/crypto/xor"

	b "github.com/eseymour/cryptopals/pkg/util/bytes"
)

// TestRepeatingXOR only tests the case from the challenge. The package
// github.com/eseymour/cryptopals/pkg/crypto/xor has more extensive tests.
func TestRepeatingXOR(t *testing.T) {
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	want := b.Must(hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"))
	fails := false

	got, err := xor.EncryptRepeatingKey(plaintext, key)
	failed := err != nil

	switch {
	case failed != fails:
		t.Errorf("xor.EncryptRepeatingKey(%#v, %#v) failure: %#v, want %#v", plaintext, key, failed, fails)
	case !fails && !bytes.Equal(got, want):
		t.Errorf("xor.EncryptRepeatingKey(%#v, %#v) == %#v, want %#v", plaintext, key, got, want)
	}
}
