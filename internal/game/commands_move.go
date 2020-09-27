package game

import (
	"errors"
	"fmt"

	"github.com/brianseitel/mudder/internal/world"
)

func doNorth(ch *world.Player, args string) error {
	return doMove(ch, "north")
}

func doSouth(ch *world.Player, args string) error {
	return doMove(ch, "south")
}

func doEast(ch *world.Player, args string) error {
	return doMove(ch, "east")
}

func doWest(ch *world.Player, args string) error {
	return doMove(ch, "west")
}

func doUp(ch *world.Player, args string) error {
	return doMove(ch, "up")
}

func doDown(ch *world.Player, args string) error {
	return doMove(ch, "down")
}

func doMove(ch *world.Player, direction string) error {
	var doorCode int
	switch direction {
	case "north", "n":
		doorCode = 0
		direction = "north"
	case "east", "e":
		doorCode = 1
		direction = "east"
	case "south", "s":
		doorCode = 2
		direction = "south"
	case "west", "w":
		doorCode = 3
		direction = "west"
	case "up", "u":
		doorCode = 4
		direction = "up"
	case "down", "d":
		doorCode = 5
		direction = "down"
	default:
		return errors.New("I don't know that direction.")
	}

	if doorCode < 0 {
		return errors.New("HUH?")
	}

	found := false
	for _, d := range ch.CurrentRoom.Doors {
		if d.Door == doorCode {
			ch.CurrentRoom = findRoom(d.ToRoom)
			ch.Send(fmt.Sprintf("You go %s.\n", direction))
			found = true
			break
		}
	}

	if !found {
		return errors.New("That room doesn't exist!")
	}
	return nil
}
