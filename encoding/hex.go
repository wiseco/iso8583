package encoding

import (
	"encoding/hex"
	"strings"
)

var (
	Hex      Encoder = &hexEncoder{}
	HexUpper Encoder = &hexEncoder{true}
)

type hexEncoder struct {
	upper bool
}

func (e hexEncoder) Encode(data []byte) ([]byte, error) {
	out := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(out, data)

	if e.upper {
		// Convert to uppercase
		return []byte(strings.ToUpper(string(out))), nil
	}

	return out, nil
}

func (e hexEncoder) Decode(data []byte, _ int) ([]byte, error) {
	out := make([]byte, hex.DecodedLen(len(data)))
	_, err := hex.Decode(out, data)
	if err != nil {
		return nil, err
	}

	return out, nil
}
