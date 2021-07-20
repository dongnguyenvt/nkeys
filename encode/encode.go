package encode

import (
	"encoding/hex"
	"errors"
)

func Encode(data []byte) []byte {
	buf := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(buf, data)
	return buf
}

func EncodeByte(b byte) []byte {
	return Encode([]byte{b})
}

func DecodeByte(data []byte) (byte, error) {
	decoded, _, err := Decode(data)
	if err != nil {
		return 0, err
	}
	if len(decoded) != 1 {
		return 0, errors.New("invalid data")
	}
	return decoded[0], nil
}

func Decode(data []byte) ([]byte, int, error) {
	raw := make([]byte, hex.DecodedLen(len(data)))
	n, err := hex.Decode(raw, data)
	return raw, n, err
}

func DecodeString(data string) ([]byte, error) {
	return hex.DecodeString(data)
}
