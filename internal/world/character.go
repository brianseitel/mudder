package world

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/brianseitel/mudder/internal/positions"
	"github.com/brianseitel/mudder/internal/tools"
	"github.com/fatih/color"
)

const LEVEL_HERO = 50

type Character struct {
	IndexData *Mobile `json:"-"`

	CurrentRoom      *Room `json:"-"`
	CurrentRoomIndex int   `json:"current_room_index"`
	// If this is nil, then it's a mob
	// and not a player
	PCData *PCData `json:"pc_data"`

	// Core character data
	VNUM             int    `json:"vnum"` // used for mobs
	Name             string `json:"name"`
	Keywords         string `json:"keywords"` // "name" for mobs
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	Description      string `json:"description"`
	Sex              int    `json:"sex"`
	Class            int    `json:"class"`
	Race             int    `json:"race"`
	Level            int    `json:"level"`
	Trust            int    `json:"trust"`
	Played           int    `json:"played"`

	// Stuff it's carrying
	Inventory []*Object `json:"inventory"`

	// HP, Mana, Move
	HitPoints    int `json:"hp"`
	MaxHitPoints int `json:"max_hp"`
	Mana         int `json:"mana"`
	MaxMana      int `json:"max_mana"`
	Movement     int `json:"mv"`
	MaxMovement  int `json:"max_mv"`

	Gold       int `json:"gold"`
	Experience int `json:"xp"`

	// Flags
	ActFlags   int `json:"act_flags"`
	AffectedBy int `json:"affected_by"`
	Position   int `json:"position"`

	Practices int `json:"practices"`
	Alignment int `json:"alignment"`

	// Fight stuff
	Fighting *Character `json:"-"`
	Hitroll  int        `json:"-"`
	Damroll  int        `json:"-"`
	Armor    int        `json:"-"`
}

type PCData struct {
	Password string `json:"password"`
	Title    string `json:"title"`

	Level int `json:"level"`

	Bamfin  string `json:"bamfin"`
	Bamfout string `json:"bamfout"`

	Strength     int `json:"str"`
	Intelligence int `json:"int"`
	Wisdom       int `json:"wis"`
	Dexterity    int `json:"dex"`
	Constitution int `json:"con"`

	ModifiedStrength     int `json:"-"`
	ModifiedIntelligence int `json:"-"`
	ModifiedWisdom       int `json:"-"`
	ModifiedDexterity    int `json:"-"`
	ModifiedConstitution int `json:"-"`
}

func (p *Character) Save() error {
	p.CurrentRoomIndex = p.CurrentRoom.VNUM

	f, err := os.Create("players/" + tools.Slug(strings.ToLower(p.Name)) + ".json")
	if err != nil {
		return err
	}
	defer f.Close()

	out, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		panic(err)
	}

	p.Send("Character saved.")
	_, err = f.Write(out)
	return err
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
