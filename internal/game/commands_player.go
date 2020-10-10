package game

import "github.com/brianseitel/mudder/internal/world"

func doInventory(ch *world.Character, args string) error {
	ch.Send("You are carrying:")
	if len(ch.Inventory) == 0 {
		ch.Send("  (empty)")
		return nil
	}

	for _, item := range ch.Inventory {
		ch.Send("   " + item.ShortDescription)
	}
	return nil
}

func doSave(ch *world.Character, args string) error {
	return ch.Save()
}
