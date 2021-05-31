package field

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wiseco/iso8583/encoding"
	"github.com/wiseco/iso8583/padding"
	"github.com/wiseco/iso8583/prefix"
)

func TestNumericField(t *testing.T) {
	field := NewNumeric(&Spec{
		Length:      10,
		Description: "Field",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
		Pad:         padding.Left(' '),
	})

	num := field.(*Numeric)

	field.SetBytes([]byte("100"))
	require.Equal(t, 100, num.Value)

	packed, err := field.Pack()

	require.NoError(t, err)
	require.Equal(t, "       100", string(packed))

	length, err := field.Unpack([]byte("      9876"))

	require.NoError(t, err)
	require.Equal(t, 10, length)
	require.Equal(t, "9876", string(field.Bytes()))
	require.Equal(t, 9876, num.Value)
}

func TestNumericFieldWithNotANumber(t *testing.T) {
	field := NewNumeric(&Spec{
		Length:      10,
		Description: "Field",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
		Pad:         padding.Left(' '),
	})

	num := field.(*Numeric)

	field.SetBytes([]byte("hello"))
	require.Equal(t, 0, num.Value)

	packed, err := field.Pack()

	require.NoError(t, err)
	require.Equal(t, "         0", string(packed))

	_, err = field.Unpack([]byte("hhhhhhhhhh"))

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to convert into number")
}
