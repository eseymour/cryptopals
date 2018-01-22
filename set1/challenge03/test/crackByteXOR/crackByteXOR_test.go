package crackByteXOR

import (
	"encoding/hex"
	"testing"

	"gonum.org/v1/gonum/stat/distuv"

	"github.com/eseymour/cryptopals/pkg/crypto/analysis"
	"github.com/eseymour/cryptopals/pkg/util/bytes"
)

func TestCrackByteXOR(t *testing.T) {
	ciphertext := bytes.Must(hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	confidence := 0.99

	_, chiSquare := analysis.BreakXOREncryptByteKey(ciphertext)

	chiSquaredDist := distuv.ChiSquared{25, nil}

	p := chiSquaredDist.CDF(chiSquare)
	if p > confidence {
		t.Errorf("BreakXOREncryptByteKey(%x) key is likely not correct with p-value %f", ciphertext, p)
	}
}
