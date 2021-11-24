package hasher

import (
	"errors"
	"math"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length = uint64(len(alphabet))
	lengthNewUrl = 10
)

// Encode number - id of the new URL
func Encode(number uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(lengthNewUrl)
	LargerNumber := number * uint64(math.Pow(float64(length - 1), 10-(1/math.Log10(float64(number)))))

	for ; LargerNumber > 0; LargerNumber = LargerNumber / length {
		encodedBuilder.WriteByte(alphabet[(LargerNumber % length)])
		if encodedBuilder.Len() == lengthNewUrl{
			break
		}
	}

	return encodedBuilder.String()
}

func Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
