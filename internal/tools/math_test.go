package tools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuzz(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(Fuzz(10))
	}
	assert.True(t, false)
}
