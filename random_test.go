package gorandom

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 99; i += 1 {
		random := rand.Intn(99)
		str, err := String(true, true, true, random)
		if err != nil {
			t.Error("expected not error, got err", err)
		}

		if len(str) != random {
			t.Error("expected not error, got err when creating random string")
		}
	}
}
