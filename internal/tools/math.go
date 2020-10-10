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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Range(a, b, c int) int {
	if b < a {
		return a
	}

	if b > c {
		return c
	}

	return b
}
