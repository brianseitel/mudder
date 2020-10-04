package game

import (
	"fmt"
	"strings"

	"github.com/brianseitel/mudder/internal/world"
	"github.com/fatih/color"
	"github.com/sanity-io/litter"
)

const (
	ActToCharacter = iota
	ActToVictim
	ActToRoom
	ActToNotVictim
)

// doLook shows the world to the character
func doLook(ch *world.Character, args string) error {
	if args == "" || args == "auto" {
		ch.Send(ch.CurrentRoom.Name)

		var doors []string
		for _, door := range ch.CurrentRoom.Doors {
			var d string
			switch door.Door {
			case 0:
				d = "north"
			case 1:
				d = "east"
			case 2:
				d = "south"
			case 3:
				d = "west"
			case 4:
				d = "up"
			case 5:
				d = "down"
			}
			doors = append(doors, d)
		}

		ch.Send(fmt.Sprintf("[Exits: %s]\n", strings.Join(doors, " ")))
		ch.Send(ch.CurrentRoom.Description)
		ch.ShowList(ch.CurrentRoom.Objects)
		showPeople(ch, ch.CurrentRoom)

	} else {
		// check mobs in room to see if we're looking at that
		for _, mob := range ch.CurrentRoom.People {
			if strings.HasPrefix(mob.Keywords, args) {
				ch.Send(mob.Description)
				return nil
			}
		}

		for _, obj := range ch.CurrentRoom.Objects {
			if strings.HasPrefix(obj.Keywords, args) {
				for _, desc := range obj.ExtraDescription {
					if strings.HasPrefix(desc.Keywords, args) {
						ch.Send(desc.Description)
						break
					}
				}
				ch.Send(obj.LongDescription)
				return nil
			}
		}
		ch.Send("You don't see that here.")
	}
	return nil
}

func doInspect(ch *world.Character, args string) error {
	// check mobs in room to see if we're looking at that
	for _, mob := range ch.CurrentRoom.People {
		if strings.HasPrefix(mob.Keywords, args) {
			litter.Dump(mob)
			return nil
		}
	}
	return nil
}

func doScan(ch *world.Character, args string) error {
	for _, door := range ch.CurrentRoom.Doors {
		room := findRoom(door.ToRoom)
		if room != nil {
			var dir string
			switch door.Door {
			case 0:
				dir = "[north]"
			case 1:
				dir = "[east]"
			case 2:
				dir = "[south]"
			case 3:
				dir = "[west]"
			case 4:
				dir = "[up]"
			case 5:
				dir = "[down]"
			}
			ch.Send(dir)
			if len(room.People) > 0 {
				showPeople(ch, room)
			} else {
				ch.Send(blue("there's no one here"))
			}
		}
	}

	return nil
}

var cyan = color.New(color.FgCyan).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()

func showPeople(ch *world.Character, room *world.Room) {
	for _, mob := range room.People {
		ch.Send(cyan(mob.LongDescription))
	}
}
