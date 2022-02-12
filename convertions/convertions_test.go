package convertions_test

import (
	convert "roman-numbers/convertions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertNumberToRoman(t *testing.T) {
	t.Run("When roman number is valid, should return valid number", func(t *testing.T) {
		roman := "XII"

		want := 12

		result, _ := convert.ConvertRomanToNumber(roman)

		assert.Equal(t, result, want)
	})

	t.Run("When passed an empty arg, should return an err", func(t *testing.T) {
		_, err := convert.ConvertRomanToNumber("")

		assert.Error(t, err)
	})
}
