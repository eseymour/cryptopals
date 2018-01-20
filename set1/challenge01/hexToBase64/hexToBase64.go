package hexToBase64

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 decodes the hex encoded string hexStr returns it as a base 64
// encoded string. HexToBase64 expects hexStr to contain only an even amount
// of hexadecimal characters, failing otherwise.
func HexToBase64(hexStr string) (base64Str string, err error) {
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return
	}

	base64Str = base64.StdEncoding.EncodeToString(data)
	return
}
