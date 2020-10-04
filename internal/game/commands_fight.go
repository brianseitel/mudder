package game

import (
	"errors"
	"fmt"
	"math/rand"
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

func stopFighting(ch *world.Character, victim *world.Character) {
	ch.Position = positions.POS_STANDING
}

// TODO: make this a player character
func multiHit(ch *world.Character, victim *world.Character) {
	oneHit(ch, victim)

	// if the character isn't fighting the victim, just do one hit and leave
	// return

	chance := ch.Level
	if ch.PCData != nil {
		chance = ch.PCData.Level / 2
	}

	if rand.Intn(100) < chance {
		oneHit(ch, victim)
		// if ch stops fighting victim, return
	}

	chance = ch.Level
	if ch.PCData != nil {
		chance = ch.PCData.Level / 4
	}

	if rand.Intn(100) < chance {
		oneHit(ch, victim)
		// if ch stops fighting victim, return
	}

	chance = ch.Level / 2
	if ch.PCData != nil {
		chance = 0
	}

	if rand.Intn(100) < chance {
		oneHit(ch, victim)
	}
}

func oneHit(ch *world.Character, victim *world.Character) {
	if victim.Position == positions.POS_DEAD || victim.CurrentRoom != ch.CurrentRoom {
		return
	}

	// figure out thac0
	thac000 := 20
	thac032 := 0
	// TODO: figure out thac0 for players
	if ch.PCData != nil {
		thac000 = 18
		thac032 = 6
	}

	thac0 := thac032 - thac000 - ch.Hitroll
	victimAC := 0 // TODO: give mobiles armor

	// roll some dice
	roll := rand.Intn(20) // roll a 20-sided die
	if roll == 0 || (roll != 19 && roll < thac0-victimAC) {
		// miss
		ch.Send("You miss!")
		return
	}

	damage := rand.Intn(ch.Level*3/2) + (ch.Level / 2)
	// TODO: change this to handle weapons

	damage += ch.Damroll
	if victim.Position == positions.POS_SLEEPING {
		damage *= 2
	}

	if damage <= 0 {
		damage = 1
	}

	ch.Send(fmt.Sprintf("You strike %s for %d damage!", victim.ShortDescription, damage))
	victim.HitPoints -= damage
	if victim.HitPoints <= 0 {
		victim.Position = positions.POS_DEAD
		ch.Send("You have KILLED " + victim.ShortDescription + "!")
		victim.Fighting = nil
		ch.Fighting = nil
		ch.Position = positions.POS_STANDING
		victim.CurrentRoom.RemovePerson(victim)
	}
}
