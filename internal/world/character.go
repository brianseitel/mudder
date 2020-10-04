package world

import (
	"fmt"

	"github.com/fatih/color"
)

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
