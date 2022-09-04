package gorandom

import (
	"math/rand"
	"testing"
	"time"
)

func TestHexGenerator(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 99; i += 1 {
		random := rand.Intn(99)
		str, err := HexString(random)
		if err != nil {
			t.Error("expected not error, got err", err)
		}

		if len(str) != random*2 {
			t.Error("expected not error, got err when creating hex string")
		}
	}
}
