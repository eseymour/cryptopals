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
	src := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	want := b.Must(hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"))
	fails := false

	got := make([]byte, len(src))
	err := xor.EncryptRepeatingKey(got, src, key)
	failed := err != nil

	switch {
	case failed != fails:
		t.Errorf("xor.EncryptRepeatingKey(%x, %x, %x) failure: %t, want %t", got, src, key, failed, fails)
	case !fails && !bytes.Equal(got, want):
		t.Errorf("xor.EncryptRepeatingKey(%x, %x, %x) == %x, want %x", got, src, key, got, want)
	}
}
