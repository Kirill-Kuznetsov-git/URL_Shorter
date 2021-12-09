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

// Encode Function which create a new string with random letters from alphabet. Length of the string is saved in the
// variable lengthNewUrl
func Encode() (string, error) {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(lengthNewUrl)
	// Add letter to string
	for i := 0; i < 10; i++ {
		encodedBuilder.WriteByte(alphabet[rand.Intn(int(length) - 1)])
	}

	return encodedBuilder.String(), nil
}