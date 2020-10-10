package game

import (
	"errors"
	"strings"

	"github.com/brianseitel/mudder/internal/positions"
	"github.com/brianseitel/mudder/internal/world"
)

func doKill(ch *world.Character, args string) error {
	if args == "" {
		return errors.New("Kill whom?")
	}

	if args == "self" {
		ch.Send("You hit yourself. Ouch!")
		multiHit(ch, nil)
		return nil
	}

	// find mob in room
	var target *world.Character
	for _, mob := range ch.CurrentRoom.People {
		if strings.HasPrefix(mob.Keywords, args) {
			target = mob
			break
		}
	}

	if target == nil {
		return errors.New("They aren't here!")
	}

	if ch.Position == positions.POS_FIGHTING && ch.Fighting != target {
		ch.Send("You're already fighting someone else!")
	}

	ch.Fighting = target
	ch.Position = positions.POS_FIGHTING
	ch.Send("You start fighting " + target.ShortDescription + ".")
	multiHit(ch, target)

	return nil
}
