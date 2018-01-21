package xor

import (
	"bytes"
	"encoding/hex"
	"testing"
)

// must is a helper that wraps calls returning ([]byte, error), panicking if the
// error is non-nil.
func must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}

	return data
}

// variableLengthKeyCase is a shared type between repeating key and fixed key
// cases, and used to add cases from other case types.
type variableLengthKeyCase struct {
	plaintext, key, want []byte
	err                  error
}

// Test cases
var (
	repeatingKeyCases = []variableLengthKeyCase{
		{ // Challenge 5
			[]byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"),
			[]byte("ICE"),
			must(hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")),
			nil,
		},
		{ // Challenge 5 decrypt
			must(hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")),
			[]byte("ICE"),
			[]byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"),
			nil,
		},
		{ // Failure path
			[]byte{},
			[]byte{},
			nil,
			KeySizeError(0),
		},
	}
	fixedKeyCases = []variableLengthKeyCase{
		{ // Challenge 2
			must(hex.DecodeString("1c0111001f010100061a024b53535009181c")),
			must(hex.DecodeString("686974207468652062756c6c277320657965")),
			must(hex.DecodeString("746865206b696420646f6e277420706c6179")),
			nil,
		},
		{ // Challenge 2 transposed
			must(hex.DecodeString("746865206b696420646f6e277420706c6179")),
			must(hex.DecodeString("1c0111001f010100061a024b53535009181c")),
			must(hex.DecodeString("686974207468652062756c6c277320657965")),
			nil,
		},
		{ // Longer key than plaintext
			must(hex.DecodeString("1c0111001f010100061a")),
			must(hex.DecodeString("686974207468652062756c6c277320657965")),
			must(hex.DecodeString("746865206b696420646f")),
			nil,
		},
		{ // Failure path
			must(hex.DecodeString("deadbeef")),
			must(hex.DecodeString("cafe")),
			nil,
			KeySizeError(2),
		},
	}
	byteKeyCases = []struct {
		plaintext []byte
		key       byte
		want      []byte
	}{
		{ // Challenge 3 spoilers
			must(hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")),
			0x58,
			[]byte("Cooking MC's like a pound of bacon"),
		},
		{ // Challenge 3 encrypt
			[]byte("Cooking MC's like a pound of bacon"),
			0x58,
			must(hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")),
		},
	}
)

// TestEncryptRepeatingKey uses not only the repeatingKeyCases, but also tests
// cases from fixed key and byte key functions since they are subsets of
// repeating key XOR. Failure cases from other functions are skipped.
func TestEncryptRepeatingKey(t *testing.T) {
	cases := repeatingKeyCases

	// Add fixed key cases, skipping cases that are expected to fail
	for _, c := range fixedKeyCases {
		if c.err != nil {
			continue
		}
		cases = append(cases, c)
	}

	// Add single byte key cases
	for _, c := range byteKeyCases {
		cases = append(cases, variableLengthKeyCase{c.plaintext, []byte{c.key}, c.want, nil})
	}

	// Run tests
	for _, c := range cases {
		got, err := EncryptRepeatingKey(c.plaintext, c.key)
		switch {
		case err != c.err:
			t.Errorf("EncryptRepeatingKey(%x, %x) error == %v, want %v", c.plaintext, c.key, err, c.err)
		case err == nil && !bytes.Equal(got, c.want):
			t.Errorf("EncryptRepeatingKey(%x, %x) == %x, want %x", c.plaintext, c.key, got, c.want)
		}
	}
}

func TestEncryptFixedKey(t *testing.T) {
	for _, c := range fixedKeyCases {
		got, err := EncryptFixedKey(c.plaintext, c.key)
		switch {
		case err != c.err:
			t.Errorf("EncryptFixedKey(%x, %x) error == %v, want %v", c.plaintext, c.key, err, c.err)
		case err == nil && !bytes.Equal(got, c.want):
			t.Errorf("EncryptFixedKey(%x, %x) == %x, want %x", c.plaintext, c.key, got, c.want)
		}
	}
}

func TestEncryptByteKey(t *testing.T) {
	for _, c := range byteKeyCases {
		var got = EncryptByteKey(c.plaintext, c.key)
		if !bytes.Equal(got, c.want) {
			t.Errorf("EncryptByteKey(%x, %x) == %x, want %x", c.plaintext, c.key, got, c.want)
		}
	}
}
