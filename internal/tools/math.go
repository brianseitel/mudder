package tools

import (
	"math/rand"
)

func Fuzz(i int) int {
	mod := rand.Intn(4)

	switch mod {
	case 0:
		i = i - 1
	case 3:
		i = i + 1
	}

	if i <= 0 {
		return 1
	}

	return i
}
