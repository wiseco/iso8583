package field

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wiseco/iso8583/encoding"
	"github.com/wiseco/iso8583/padding"
	"github.com/wiseco/iso8583/prefix"
)

func TestStringField(t *testing.T) {
	field := NewString(&Spec{
		Length:      10,
		Description: "Field",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
		Pad:         padding.Left(' '),
	})

	str := field.(*String)

	field.SetBytes([]byte("hello"))
	require.Equal(t, "hello", str.Value)

	packed, err := field.Pack()

	require.NoError(t, err)
	require.Equal(t, "     hello", string(packed))

	length, err := field.Unpack([]byte("     olleh"))

	require.NoError(t, err)
	require.Equal(t, 10, length)
	require.Equal(t, "olleh", string(field.Bytes()))
	require.Equal(t, "olleh", str.Value)
}
