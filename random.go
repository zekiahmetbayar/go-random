package gorandom

import (
	"errors"
	"math/rand"
	"time"

	"github.com/zekiahmetbayar/go-random/internal"
)

// String returns random string of given length.
//
// Parameters:
//
//	numbers: Include numbers to random string
//	letters: Include english letters to random string
//	specials: Include special characters to random string
//	length: Random string length
func String(numbers, letters, specials bool, str_length int) (string, error) {
	var CHARSET []string
	var random string

	if str_length == 0 || str_length < 0 {
		return "", errors.New("string length must be greater than zero")
	}

	if !(numbers || letters || specials) {
		return "", errors.New("at least one of letters, special characters or numbers must be included")
	}

	CHARSET = internal.GetCharset(numbers, letters, specials) // Get charset
	rand.Seed(time.Now().UnixNano())                          // Set start point randomly

	for i := 0; i < str_length; i += 1 {
		random += CHARSET[rand.Intn(len(CHARSET))] // Generate random string
	}

	return random, nil
}
