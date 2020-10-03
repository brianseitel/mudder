package game

import (
	"testing"

	"github.com/brianseitel/mudder/internal/world"
	"github.com/stretchr/testify/assert"
)

func TestDoLook(t *testing.T) {

	gameWorld = &Game{
		World:     world.New(),
		Character: nil,
	}

	ch := &world.Character{
		CurrentRoom: findRoom(3700),
	}

	err := doLook(ch, "look")

	assert.Nil(t, err)
}
