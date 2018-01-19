package hexToBase64

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexStr string) (base64Str string, err error) {
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return
	}

	base64Str = base64.StdEncoding.EncodeToString(data)
	return
}
