package world

import "strings"

type Room struct {
	VNUM                 int
	Name                 string
	Description          string
	Area                 int
	RoomFlags            int
	SectorType           int
	Doors                []Door
	ExtendedDescriptions []ExtendedDescription

	Objects []*Object
	Mobs    []*Mobile
	People  []*Character
}

type Door struct {
	Door        int
	Description string
	Keywords    string
	Locks       int
	Key         int
	ToRoom      int
}

func (r *Room) HasCharacter(ch *Character) bool {
	for _, person := range r.People {
		if person == ch {
			return true
		}
	}
	return false
}

func (r *Room) FindCharacterByName(name string) *Character {
	for _, person := range r.People {
		if strings.HasPrefix(person.Name, name) {
			return person
		}
	}

	return nil
}
