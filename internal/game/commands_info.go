package game

import "github.com/brianseitel/mudder/internal/world"

const (
	ActToCharacter = iota
	ActToVictim
	ActToRoom
	ActToNotVictim
)

// doLook shows the world to the character
func doLook(ch *world.Character, args string) error {
	ch.Send(ch.CurrentRoom.Name)
	ch.Send("")

	ch.Send(ch.CurrentRoom.Description)

	ch.ShowList(ch.CurrentRoom.Objects)
	ch.ShowPeople(ch.CurrentRoom.People)
	ch.ShowMobs(ch.CurrentRoom.Mobs)
	return nil
}
