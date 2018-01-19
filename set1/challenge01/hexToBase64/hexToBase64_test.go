package hexToBase64

import "testing"

func TestHexToBase64(t *testing.T) {
	cases := []struct {
		in, want string
		fails    bool
	}{
		{
			"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
			false,
		},
		{
			"This is not valid hex",
			"This should be ignored",
			true,
		},
	}

	for _, c := range cases {
		got, err := HexToBase64(c.in)
		switch {
		case (err != nil) != c.fails:
			t.Errorf("HexToBase64(%#v) errors: %#v, want %#v", c.in, err != nil, c.fails)
		case !c.fails && got != c.want:
			t.Errorf("HexToBase64(%#v) == %#v,  want %#v", c.in, got, c.want)
		}
	}
}
