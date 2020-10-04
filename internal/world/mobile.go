package world

import (
	"math/rand"

	"github.com/brianseitel/mudder/internal/tools"
)

type Mobile struct {
	Raw string // raw data

	// Loaded from file
	VNUM             int
	Keywords         string
	ShortDescription string
	LongDescription  string
	Description      string
	ActFlags         int
	AffectedFlags    int
	Alignment        int
	Level            int
	Sex              int

	// In game stuff
	CurrentRoom *Room
	Position    int

	Hitpoints int
}

// TODO: convert Mobile to Characters
func (m *Mobile) ToCharacter() *Character {
	ch := &Character{}
	ch.IndexData = m

	ch.VNUM = m.VNUM
	ch.Name = m.Keywords
	ch.Keywords = m.Keywords
	ch.ShortDescription = m.ShortDescription
	ch.LongDescription = m.LongDescription
	ch.Sex = m.Sex
	ch.Level = tools.Fuzz(m.Level) // some are weaker, some are stronger

	ch.ActFlags = m.ActFlags
	ch.AffectedBy = m.AffectedFlags
	ch.Position = m.Position

	ch.Armor = 0 // TODO: interpolate level, 100, -100
	hpMin := (ch.Level * ch.Level / 4)
	hpMax := (ch.Level * ch.Level)
	if hpMax == 0 {
		hpMin = 0
		hpMax = 1
	}
	ch.HitPoints = ch.Level*8 + (rand.Intn(hpMax-hpMin) + hpMin)
	ch.MaxHitPoints = ch.HitPoints

	ch.Mana = 0     // TODO: calculate these
	ch.Movement = 0 // TODO: calculate these

	return ch
}
