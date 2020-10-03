package game

import (
	"strings"

	"github.com/brianseitel/mudder/internal/world"
)

func doGet(ch *world.Character, args string) error {
	if args == "" {
		ch.Send("Get what?")
	}

	for i, obj := range ch.CurrentRoom.Objects {
		if strings.HasPrefix(obj.Keywords, args) {
			ch.Send("You pick up " + obj.ShortDescription + ".")
			ch.Inventory = append(ch.Inventory, obj)
			ch.CurrentRoom.Objects = ch.CurrentRoom.Objects[:i+copy(ch.CurrentRoom.Objects[i:], ch.CurrentRoom.Objects[i+1:])]
		}
	}
	return nil
}

func doDrop(ch *world.Character, args string) error {
	if args == "" {
		ch.Send("Drop what?")
	}

	if args == "all" {
		for i, obj := range ch.Inventory {
			ch.Send("You drop " + obj.ShortDescription + " to the ground.")
			ch.CurrentRoom.Objects = append(ch.CurrentRoom.Objects, obj)
			ch.Inventory = ch.Inventory[:i+copy(ch.Inventory[i:], ch.Inventory[i+1:])]
		}
		return nil
	} else {
		for i, obj := range ch.Inventory {
			if strings.HasPrefix(obj.Keywords, args) {
				ch.Send("You drop " + obj.ShortDescription + " to the ground.")
				ch.CurrentRoom.Objects = append(ch.CurrentRoom.Objects, obj)
				ch.Inventory = ch.Inventory[:i+copy(ch.Inventory[i:], ch.Inventory[i+1:])]
				return nil
			}
		}
	}

	ch.Send("You're not carrying that.")

	return nil
}
