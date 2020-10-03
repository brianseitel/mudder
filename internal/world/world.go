package world

import (
	"github.com/rs/zerolog/log"
)

type World struct {
	Zones   []*Zone
	Mobs    map[int]Mobile
	Objects map[int]Object
}

var world *World

func New() *World {
	if world == nil {
		world = &World{}
	}

	return world
}

func (w *World) Populate() {
	w.Mobs = make(map[int]Mobile)
	w.Objects = make(map[int]Object)

	for _, zone := range w.Zones {
		for _, mob := range zone.Mobiles {
			w.Mobs[mob.VNUM] = *mob
		}

		for _, object := range zone.Objects {
			w.Objects[object.VNUM] = object
		}
	}

	// go through all the zones again and get the resets and execute them
	for _, zone := range w.Zones {
		// do this last
		for _, reset := range zone.Resets {
			switch res := reset.(type) {
			case ResetReadMobile:
				count := 0
				room := findRoom(w, res.Room)
				add := true
				for _, mob := range room.Mobs {
					if mob.VNUM == res.VNUM {
						count++
						if count == res.Limit {
							add = false
							break
						}
					}
				}

				if add {
					mob, ok := w.Mobs[res.VNUM]
					if ok {
						room.Mobs = append(room.Mobs, &mob)
					}
				}
			case ResetReadObject:
				room := findRoom(w, res.Room)

				// See if this object already exists in the room
				// if it does, don't add it.
				add := true
				for _, obj := range room.Objects {
					if obj.VNUM == res.VNUM {
						add = false
						break
					}
				}

				if add {
					obj, ok := w.Objects[res.VNUM]
					if ok {
						room.Objects = append(room.Objects, &obj)
					}
				}
			}
		}
	}
}

func findRoom(w *World, vnum int) *Room {
	for _, zone := range w.Zones {
		for _, room := range zone.Rooms {
			if room.VNUM == vnum {
				return room
			}
		}
	}

	log.Info().Int("vnum", vnum).Msg("Cannot find room")
	return &Room{}
}
