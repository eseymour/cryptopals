package crackByteXOR

import (
	"encoding/hex"
	"testing"

	"github.com/eseymour/cryptopals/pkg/crypto/analysis"
	"github.com/eseymour/cryptopals/pkg/util/bytes"
)

func TestCrackByteXOR(t *testing.T) {
	ciphertext := bytes.Must(hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	threshold := 0.10

	_, pValue := analysis.BreakXOREncryptByteKey(ciphertext)

	if pValue < threshold {
		t.Errorf("BreakXOREncryptByteKey(%x) key is likely not correct with p-value %f", ciphertext, pValue)
	}
}
