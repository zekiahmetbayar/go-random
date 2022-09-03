package gorandom

import (
	"errors"
)

func Generate(numbers, letters, specials bool, str_length int) (string, error) {
	if str_length == 0 || str_length < 0 {
		return "", errors.New("string length must be greater than zero")
	}

	if !(numbers || letters || specials) {
		return "", errors.New("at least one of letters, special characters or numbers must be included")
	}

	//	var CHARSET []string
	//	CHARSET = internal.GetCharset(numbers, letters, specials)

	return "", nil
}
