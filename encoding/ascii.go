package encoding

import "fmt"

type encType int

const (
	asciiTypeDefault encType = iota
	asciiTypeAlpha
	asciiTypeNumeric
	asciiTypeAlphaNumeric
	asciiTypeNonAlpha
	asciiTypeNonNumeric
	asciiTypeNonAlphaNumeric
	asciiTypeAll
)

var (
	ASCII           = &asciiEncoder{asciiTypeDefault}         // Alpha, Numeric, Special Characters (ANS)
	Alpha           = &asciiEncoder{asciiTypeAlpha}           // Alpha (A)
	Numeric         = &asciiEncoder{asciiTypeNumeric}         // Numeric (N)
	AlphaNumeric    = &asciiEncoder{asciiTypeAlphaNumeric}    // AlphaNumric (AN)
	NonAlpha        = &asciiEncoder{asciiTypeNonAlpha}        // Numeric, Special Characters (NS)
	NonNumeric      = &asciiEncoder{asciiTypeNonNumeric}      // Alpha, Special Characters (AS)
	NonAlphaNumeric = &asciiEncoder{asciiTypeNonAlphaNumeric} // Special Characters (S)
	All             = &asciiEncoder{asciiTypeAll}             // All type of characters
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
				return nil, fmt.Errorf("invalid ASCII char: %#X", r)
			}
		case asciiTypeAlpha:
			switch {
			case r >= 65 && r <= 90, r >= 97 && r <= 122:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII alpha char: %#X", r)
			}
		case asciiTypeNumeric:
			if r < 48 || r > 57 {
				return nil, fmt.Errorf("invalid ASCII numeric char: %#X", r)
			}
		case asciiTypeAlphaNumeric:
			switch {
			case r >= 48 && r <= 57, r >= 65 && r <= 90, r >= 97 && r <= 122:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII alphanumeric char: %#X", r)
			}
		case asciiTypeNonAlpha:
			switch {
			case r <= 64, r >= 91 && r <= 96, r >= 123:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII non alpha char: %#X", r)
			}
		case asciiTypeNonNumeric:
			if r >= 48 && r <= 57 {
				return nil, fmt.Errorf("invalid ASCII non numeric char: %#X", r)
			}
		case asciiTypeNonAlphaNumeric:
			switch {
			case r <= 47, r >= 58 && r <= 64, r >= 91 && r <= 96, r >= 123:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII non alphanumeric char: %#X", r)
			}
		case asciiTypeAll:
			break
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
				return nil, fmt.Errorf("invalid ASCII char: %#X", r)
			}
		case asciiTypeAlpha:
			switch {
			case r >= 65 && r <= 90, r >= 97 && r <= 122:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII alpha char: %#X", r)
			}
		case asciiTypeNumeric:
			if r < 48 || r > 57 {
				return nil, fmt.Errorf("invalid ASCII numeric char: %#X", r)
			}
		case asciiTypeAlphaNumeric:
			switch {
			case r >= 48 && r <= 57, r >= 65 && r <= 90, r >= 97 && r <= 122:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII alphanumeric char: %#X", r)
			}
		case asciiTypeNonAlpha:
			switch {
			case r <= 64, r >= 91 && r <= 96, r >= 123:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII non alpha char: %#X", r)
			}
		case asciiTypeNonNumeric:
			if r >= 48 && r <= 57 {
				return nil, fmt.Errorf("invalid ASCII non numeric char: %#X", r)
			}
		case asciiTypeNonAlphaNumeric:
			switch {
			case r <= 47, r >= 58 && r <= 64, r >= 91 && r <= 96, r >= 123:
				break
			default:
				return nil, fmt.Errorf("invalid ASCII non alphanumeric char: %#X", r)
			}
		}
		out = append(out, r)
	}

	return out, nil
}
