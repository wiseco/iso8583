package encoding

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestASCII(t *testing.T) {
	enc := ASCII

	t.Run("Decode", func(t *testing.T) {
		res, err := enc.Decode([]byte("hello!"), 0)

		require.NoError(t, err)
		require.Equal(t, []byte("hello!"), res)

		_, err = enc.Decode([]byte("hello, 世界!"), 0)
		require.Error(t, err)
	})

	t.Run("Encode", func(t *testing.T) {
		res, err := enc.Encode([]byte("hello!"))

		require.NoError(t, err)
		require.Equal(t, []byte("hello!"), res)

		_, err = enc.Encode([]byte("hello, 世界!"))
		require.Error(t, err)
	})
}

func TestAlpha(t *testing.T) {
	enc := Alpha

	t.Run("Decode", func(t *testing.T) {
		res, err := enc.Decode([]byte("hello"), 0)

		require.NoError(t, err)
		require.Equal(t, []byte("hello"), res)

		_, err = enc.Decode([]byte("Hello09"), 0)
		require.Error(t, err)
	})

	t.Run("Encode", func(t *testing.T) {
		res, err := enc.Encode([]byte("hello"))

		require.NoError(t, err)
		require.Equal(t, []byte("hello"), res)

		_, err = enc.Encode([]byte("Hello09"))
		require.Error(t, err)
	})
}

func TestNumeric(t *testing.T) {
	enc := Numeric

	t.Run("Decode", func(t *testing.T) {
		res, err := enc.Decode([]byte("01234"), 0)

		require.NoError(t, err)
		require.Equal(t, []byte("01234"), res)

		_, err = enc.Decode([]byte("Hello09"), 0)
		require.Error(t, err)
	})

	t.Run("Encode", func(t *testing.T) {
		res, err := enc.Encode([]byte("01234"))

		require.NoError(t, err)
		require.Equal(t, []byte("01234"), res)

		_, err = enc.Encode([]byte("Hello09"))
		require.Error(t, err)
	})
}

func TestAlphaNumeric(t *testing.T) {
	enc := AlphaNumeric

	t.Run("Decode", func(t *testing.T) {
		res, err := enc.Decode([]byte("Hello09"), 0)

		require.NoError(t, err)
		require.Equal(t, []byte("Hello09"), res)

		_, err = enc.Decode([]byte("Hello09!"), 0)
		require.Error(t, err)
	})

	t.Run("Encode", func(t *testing.T) {
		res, err := enc.Encode([]byte("Hello09"))

		require.NoError(t, err)
		require.Equal(t, []byte("Hello09"), res)

		_, err = enc.Encode([]byte("Hello09!"))
		require.Error(t, err)
	})
}
