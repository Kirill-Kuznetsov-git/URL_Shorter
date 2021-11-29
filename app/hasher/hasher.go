package hasher

import (
	"math/rand"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length = uint64(len(alphabet))
	lengthNewUrl = 10
)

func Encode() (string, error) {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(lengthNewUrl)

	for i := 0; i < 10; i++ {
		encodedBuilder.WriteByte(alphabet[rand.Intn(int(length) - 1)])
	}

	return encodedBuilder.String(), nil
}