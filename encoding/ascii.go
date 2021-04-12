package encoding

import "fmt"

type encType int

const (
	asciiTypeDefault encType = iota
	asciiTypeAlpha
	asciiTypeNumeric
	asciiTypeAlphaNumeric
)

var (
	ASCII        = &asciiEncoder{asciiTypeDefault}
	Alpha        = &asciiEncoder{asciiTypeAlpha}
	Numeric      = &asciiEncoder{asciiTypeNumeric}
	AlphaNumeric = &asciiEncoder{asciiTypeAlphaNumeric}
)

type asciiEncoder struct {
	t encType
}

func (e asciiEncoder) Encode(data []byte) ([]byte, error) {
	out := []byte{}
	for _, r := range data {
		switch e.t {
		default:
			if r > 127 {
				return nil, fmt.Errorf("invalid ASCII char: '%s'", string(r))
			}
		case asciiTypeAlpha:
			switch {
			case r >= 65 && r <= 90, r >= 97 && r <= 122:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII alpha char: '%s'", string(r))
			}
		case asciiTypeNumeric:
			if r < 48 || r > 57 {
				return nil, fmt.Errorf("invalid ASCII numeric char: '%s'", string(r))
			}
		case asciiTypeAlphaNumeric:
			switch {
			case r >= 48 && r <= 57, r >= 65 && r <= 90, r >= 97 && r <= 122:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII alphanumeric char: '%s'", string(r))
			}
		}
		out = append(out, r)
	}

	return out, nil
}

func (e asciiEncoder) Decode(data []byte, _ int) ([]byte, error) {
	out := []byte{}
	for _, r := range data {
		switch e.t {
		default:
			if r > 127 {
				return nil, fmt.Errorf("invalid ASCII char: '%s'", string(r))
			}
		case asciiTypeAlpha:
			switch {
			case r >= 65 && r <= 90, r >= 97 && r <= 122:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII alpha char: '%s'", string(r))
			}
		case asciiTypeNumeric:
			if r < 48 || r > 57 {
				return nil, fmt.Errorf("invalid ASCII numeric char: '%s'", string(r))
			}
		case asciiTypeAlphaNumeric:
			switch {
			case r >= 48 && r <= 57, r >= 65 && r <= 90, r >= 97 && r <= 122:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII alphanumeric char: '%s'", string(r))
			}
		}
		out = append(out, r)
	}

	return out, nil
}
