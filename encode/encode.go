package encode

import "encoding/base32"

// Set our encoding to not include padding '=='
var encoder = base32.StdEncoding.WithPadding(base32.NoPadding)

func Encode(data []byte) []byte {
	buf := make([]byte, encoder.EncodedLen(len(data)))
	encoder.Encode(buf, data)
	return buf
}

func Decode(data []byte) ([]byte, int, error) {
	raw := make([]byte, encoder.DecodedLen(len(data)))
	n, err := encoder.Decode(raw, data)
	return raw, n, err
}

func DecodeString(data string) ([]byte, error) {
	return encoder.DecodeString(data)
}
