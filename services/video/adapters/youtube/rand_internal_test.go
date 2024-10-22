package youtube

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandEnStringRunes(t *testing.T) {
	t.Parallel()

	length := 10
	result, err := randEnStringRunes(uint(length))

	require.NoError(t, err, "There should be no error to get random en string")
	assert.Equal(t,
		len([]rune(result)), length,
		"Length of result string should match input length",
	)

	for _, r := range result {
		assert.True(
			t, unicode.Is(unicode.Latin, r), "All characters should be English letters")
	}
}

func TestRandUaStringRunes(t *testing.T) {
	t.Parallel()

	length := 10
	result, err := randUaStringRunes(uint(length))

	require.NoError(t, err, "There should be no error to get random ua string")
	assert.Equal(t,
		len([]rune(result)), length,
		"Length of result string should match input length",
	)

	for _, r := range result {
		assert.True(
			t,
			unicode.Is(unicode.Cyrillic, r),
			"All characters should be Ukrainian letters",
		)
	}
}

func TestRandString(t *testing.T) {
	t.Parallel()

	length := 10
	result, err := randString(uint(length))

	require.NoError(t, err, "There should be no error to get random string")
	assert.Equal(t,
		len([]rune(result)), length,
		"Length of result string should match input length",
	)
}

func TestRandOrder(t *testing.T) {
	t.Parallel()

	result, err := randOrder()

	require.NoError(t, err, "There should be no error to get random order")
	assert.IsType(t, "", result, "RandOrder should return a string")
	assert.NotEmpty(t, result, "RandOrder should return a non-empty string")
}
