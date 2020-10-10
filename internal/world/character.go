package world

import (
	"fmt"

	"github.com/brianseitel/mudder/internal/positions"
	"github.com/brianseitel/mudder/internal/tools"
	"github.com/fatih/color"
)

const LEVEL_HERO = 50

type Character struct {
	IndexData *Mobile

	CurrentRoom *Room

	// If this is nil, then it's a mob
	// and not a player
	PCData *PCData

	// Core character data
	VNUM             int // used for mobs
	Name             string
	Keywords         string // "name" for mobs
	ShortDescription string
	LongDescription  string
	Description      string
	Sex              int
	Class            int
	Race             int
	Level            int
	Trust            int
	Played           int

	// Stuff it's carrying
	Inventory []*Object

	// HP, Mana, Move
	HitPoints    int
	MaxHitPoints int
	Mana         int
	MaxMana      int
	Movement     int
	MaxMovement  int

	Gold       int
	Experience int

	// Flags
	ActFlags   int
	AffectedBy int
	Position   int

	Practices int
	Alignment int

	// Fight stuff
	Fighting *Character
	Hitroll  int
	Damroll  int
	Armor    int
}

type PCData struct {
	Password string
	Title    string

	Level int

	Bamfin  string
	Bamfout string

	Strength     int
	Intelligence int
	Wisdom       int
	Dexterity    int
	Constitution int

	ModifiedStrength     int
	ModifiedIntelligence int
	ModifiedWisdom       int
	ModifiedDexterity    int
	ModifiedConstitution int
}

func (p *Character) Send(str interface{}) {
	fmt.Println(str)
}

func (p *Character) Print(str interface{}) {
	fmt.Print(str)
}

func (p *Character) IsNPC() bool {
	return p.PCData != nil
}

func (p *Character) SetFighting(opponent *Character) {
	if p.Fighting != nil {
		// TODO: log error, already fighting
	}
	p.Fighting = opponent
	p.Position = positions.POS_FIGHTING
}

func (p *Character) ShowList(things []*Object) {
	for _, thing := range things {
		p.Send(cyan(thing.LongDescription))
	}
}

func (p *Character) ShowPeople(people []*Character) {
	for _, ch := range people {
		p.Send(ch.ShortDescription)
	}
}

var cyan = color.New(color.FgCyan).SprintFunc()

func (p *Character) UpdatePosition() {
	if p.HitPoints > 0 {
		if p.Position <= positions.POS_STUNNED {
			p.Position = positions.POS_STANDING
			return
		}
	}

	if p.PCData == nil || p.HitPoints <= -11 {
		p.Position = positions.POS_DEAD
		return
	}

	if p.HitPoints <= -6 {
		p.Position = positions.POS_MORTAL
	} else if p.HitPoints <= -3 {
		p.Position = positions.POS_INCAP
	} else {
		p.Position = positions.POS_STUNNED
	}
}

func (p *Character) GainExperience(xp int) {
	if p.IsNPC() {
		return
	}

	// TODO: handle heroes
	p.Experience += tools.Max(1000, p.Experience+xp)
	for p.Level <= LEVEL_HERO && p.Experience >= 1000*(p.Level+1) {
		p.Send("You raise a level!")
		p.Level += 1
		// TODO: advance level
	}
}
