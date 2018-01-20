package fixedXor

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}

func TestFixedXor(t *testing.T) {
	cases := []struct {
		plaintext, key, want []byte
		fails                bool
	}{
		{
			[]byte{},
			[]byte{},
			[]byte{},
			false,
		},
		{
			must(hex.DecodeString("1c0111001f010100061a024b53535009181c")),
			must(hex.DecodeString("686974207468652062756c6c277320657965")),
			must(hex.DecodeString("746865206b696420646f6e277420706c6179")),
			false,
		},
	}

	for _, c := range cases {
		got, err := FixedXor(c.plaintext, c.key)
		failed := err != nil
		switch {
		case failed != c.fails:
			t.Errorf("FixedXor(%#v, %#v) failure: %#v, want %#v", c.plaintext, c.key, failed, c.fails)
		case !c.fails && !bytes.Equal(got, c.want):
			t.Errorf("FixedXor(%#v, %#v) == %#v, want %#v", c.plaintext, c.key, got, c.want)
		}
	}
}
