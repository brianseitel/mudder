package game

import (
	"fmt"
	"math/rand"

	"github.com/brianseitel/mudder/internal/positions"
	"github.com/brianseitel/mudder/internal/tools"
	"github.com/brianseitel/mudder/internal/world"
)

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
		damage(ch, victim, 0, 0) // TODO: define damageType for last param
		return
	}

	dam := rand.Intn(ch.Level*3/2) + (ch.Level / 2)
	// TODO: change this to handle weapons

	dam += ch.Damroll
	if victim.Position == positions.POS_SLEEPING {
		dam *= 2
	}

	if dam <= 0 {
		dam = 1
	}

	damage(ch, victim, dam, 0)            // TODO: define damageType
	sendDamageMessage(ch, victim, dam, 0) // TODO: define damagetype
}

func damage(ch *world.Character, victim *world.Character, damage int, damageType int) {
	// if victim is dead, don't bother doing anything else
	if victim.Position == positions.POS_DEAD {
		return
	}

	if victim.Position > positions.POS_STUNNED {
		if victim.Fighting == nil {
			victim.SetFighting(ch)
			victim.Position = positions.POS_FIGHTING
		}
	}

	if victim.Position > positions.POS_STUNNED {
		if ch.Fighting == nil {
			ch.SetFighting(victim)
		}
	}

	// TODO: strip invisibility
	// TODO: handle charmed creatures
	// TODO: handle damage modifiers - SANCTUARY 50%, PROTECT 25%

	if damage < 0 {
		damage = 0
	}

	// TODO: check hit type

	victim.HitPoints -= damage
	if victim.Level > 1000 { // Immmortals don't die
		victim.HitPoints = 1
	}

	victim.UpdatePosition()

	// TODO: update when we have a descriptor
	if victim.PCData != nil {
		switch victim.Position {
		case positions.POS_MORTAL:
			// TODO: victim is mortally wounded
			victim.Send("You are mortally wounded and will die soon, if not aided.\n\r")
		case positions.POS_INCAP:
			// TODO: victim is incapacitated
			victim.Send("You are incapacitated and will slowly die, if not aided.\n\r")
		case positions.POS_STUNNED:
			// TODO: victim is stunned
			victim.Send("You are stunned but will probably recover.\n\r")
		case positions.POS_DEAD:
			// TODO: victim is dead
			victim.Send("You have been KILLED!!\n\r\n\r")
		}
	}

	// You win! Experience and stuff
	if victim.Position == positions.POS_DEAD {
		gainExp(ch, victim)
		stopFighting(victim, false)
		deathCry(ch, victim)
		// TODO: handle corpse creation
		// TODO: handle auto looting

		// Kill 'em dead
		victim.AffectedBy = 0
		victim.Armor = 100
		victim.Position = positions.POS_RESTING
		victim.HitPoints = victim.MaxHitPoints
		victim.Mana = victim.MaxMana
		victim.Movement = victim.MaxMovement
		victim.CurrentRoom.RemovePerson(victim)
		victim.CurrentRoom = nil
		// TODO: save char
	}

	if victim == ch {
		return
	}

	// TODO: handle wimpy
}

func deathCry(ch *world.Character, victim *world.Character) {
	if victim.IsNPC() {
		ch.Send(fmt.Sprintf("%s falls to the floor DEAD!", victim.ShortDescription))
		return
	}

	ch.Send(fmt.Sprintf("%s falls to the floor DEAD!", ch.Name))

	// let the surrounding rooms know
	for _, thisRoom := range ch.CurrentRoom.Doors {
		otherRoom := findRoom(thisRoom.ToRoom)
		if otherRoom != nil {
			for _, person := range otherRoom.People {
				if victim.IsNPC() {
					person.Send("You hear something's death cry.")
				} else {
					person.Send("You hear someone's death cry.")
				}
			}
		}
	}
}

func sendDamageMessage(ch *world.Character, victim *world.Character, dam int, damageType int) {
	var verb, plural string

	punctuation := "."
	if dam > 24 {
		punctuation = "!"
	}

	switch {
	case dam == 0:
		verb = "miss"
		plural = "misses"
	case dam <= 4:
		verb = "scratch"
		plural = "scratches"
	case dam <= 8:
		verb = "graze"
		plural = "grazes"
	case dam <= 12:
		verb = "hit"
		plural = "hits"
	case dam <= 16:
		verb = "injure"
		plural = "injures"
	case dam <= 20:
		verb = "wound"
		plural = "wounds"
	case dam <= 24:
		verb = "maul"
		plural = "mauls"
	case dam <= 28:
		verb = "decimate"
		plural = "decimates"
	case dam <= 32:
		verb = "devastate"
		plural = "devastates"
	case dam <= 36:
		verb = "maim"
		plural = "maims"
	case dam <= 40:
		verb = "MUTILATE"
		plural = "MUTILATES"
	case dam <= 44:
		verb = "DISEMBOWEL"
		plural = "DISEMBOWELS"
	case dam <= 48:
		verb = "EVISCERATE"
		plural = "EVISCERATES"
	case dam <= 52:
		verb = "MASSACRE"
		plural = "MASSACRES"
	case dam <= 100:
		verb = "*** DEMOLISH ***"
		plural = "*** DEMOLISHES ***"
	case dam > 100:
		verb = "*** ANNIHILATE ***"
		plural = "*** ANNIHILATES ***"
	default:
		verb = "hit"
		plural = "hits"
	}

	if ch.PCData != nil {
		ch.Send(fmt.Sprintf("You %s %s%s (%d dmg)", verb, victim.ShortDescription, punctuation, dam))
	} else {
		ch.Send(fmt.Sprintf("%s %s you%s (%d dmg)", ch.ShortDescription, plural, punctuation, dam))
	}
}

func stopFighting(ch *world.Character, both bool) {
	if both {
		ch.Fighting.Fighting = nil
		ch.Fighting.Position = positions.POS_STANDING
	}
	ch.Fighting = nil
	ch.Position = positions.POS_STANDING
}

func gainExp(ch *world.Character, victim *world.Character) {
	if ch.PCData == nil {
		return // NPCs don't gain xp
	}

	if victim == ch {
		return // no xp for killing yourself
	}

	xp := computeExperience(ch, victim)
	ch.Send(fmt.Sprintf("You receive %d experience points.", xp))
	ch.GainExperience(xp)
}

func computeExperience(ch *world.Character, victim *world.Character) int {
	xp := 300 - tools.Range(-3, ch.Level-victim.Level, 6)*50
	align := ch.Alignment - victim.Alignment

	if align > 500 {
		ch.Alignment = tools.Min(ch.Alignment+(align-500)/4.0, 1000.)
		xp = 5 * xp / 4
	} else if align < -500 {
		ch.Alignment = tools.Max(ch.Alignment+(align-500)/4.0, 1000.)
	} else {
		ch.Alignment -= ch.Alignment / 4
		xp = 3 * xp / 4
	}

	// TODO: adjust for popularity of the target, so more popular gets less xp
	minXP := int(xp * 3 / 4)
	maxXP := int(xp * 5 / 4)
	xp = rand.Intn(maxXP-minXP) + minXP
	xp = tools.Max(0, xp)

	return xp
}
