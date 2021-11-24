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
func Encode(number uint64) (string, error) {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(lengthNewUrl)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
		if encodedBuilder.Len() > 10{
			return "Too big number", errors.New("too big number")
		}
	}

	return encodedBuilder.String(), nil
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
