package gorandom

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func HexString(str_length int) (string, error) {
	rand.Seed(time.Now().Unix())
	token := make([]byte, str_length)

	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(token), nil
}
