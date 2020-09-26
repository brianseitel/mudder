package world

import "fmt"

type Player struct {
	Name string

	CurrentRoom *Room
}

func (p *Player) Send(str interface{}) {
	fmt.Println(str)
}

func (p *Player) ShowList(things []*Object) {
	for _, thing := range things {
		p.Send(thing.ShortDescription)
	}
}

func (p *Player) ShowPeople(people []*Player) {
	for _, player := range people {
		p.Send(player.Name + " is here.")
	}
}

func (p *Player) ShowMobs(mobs []*Mobile) {
	for _, mob := range mobs {
		p.Send(mob.LongDescription)
	}
}
