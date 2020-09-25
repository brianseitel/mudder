package world

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/brianseitel/mudder/internal/lexer"
)

func New() *World {
	return &World{}
}

func Load() *World {
	body, err := ioutil.ReadFile("areas/area.lst")
	if err != nil {
		panic(err)
	}

	gameWorld := New()
	for _, line := range strings.Split(string(body), "\n") {
		if line == "$" { // end of file
			break
		}
		zone := loadZone(line)
		gameWorld.Zones = append(gameWorld.Zones, zone)
	}

	return gameWorld
}

func loadZone(areaName string) *Zone {
	data := loadFile(areaName)

	zone := &Zone{}
	zone.Area = loadArea(data)
	zone.Helps = loadHelps(data)
	zone.Mobiles = loadMobiles(data)
	zone.Objects = loadObjects(data)
	zone.Rooms = loadRooms(data)
	zone.Resets = loadResets(data)
	zone.Shops = loadShops(data)
	zone.Specials = loadSpecials(data)

	return zone
}

func loadArea(input string) Area {
	data := lexer.New(input)

	var info Area

	// if no areas, get outta here
	if err := data.Jump("#AREA"); err != nil {
		return info
	}

	for data.Next() != '{' {
		continue
	}

	data.Letter()
	data.Gobble()

	switch data.Current() {
	case 'N': // None
		info.MinLevel = -1
		info.MaxLevel = -1
		data.Advance(4)
	case 'A': // All
		info.MinLevel = 1
		info.MaxLevel = 9999
		data.Advance(3)
	default:
		info.MinLevel = data.Number()
		info.MaxLevel = data.Number()
	}
	data.Letter() // } -- unused

	info.Author = data.Word()
	info.Name = data.String()

	return info
}

func loadHelps(input string) []Help {
	data := lexer.New(input)

	var infos []Help

	// If no helps, just get outta here.
	if err := data.Jump("#HELPS"); err != nil {
		return infos
	}

	for {
		// Grab this help section
		help := Help{}
		help.Level = data.Number()
		help.Keywords = data.String()
		if help.Keywords == "$" {
			break
		}
		help.Text = data.String()
		// if help text begins with a period, strip it off
		// this is present to avoid the stripping of leading  whitespace
		if help.Text[0] == '.' {
			help.Text = help.Text[1:]
		}

		infos = append(infos, help)

		data.Gobble()
	}

	return infos
}

func loadShops(input string) []Shop {
	data := lexer.New(input)

	var infos []Shop

	// if no shops, get outta here
	if err := data.Jump("#SHOPS"); err != nil {
		return infos
	}

	for {
		data.Gobble()
		if data.Current() == '0' {
			// we're done
			break
		}

		var shop Shop
		shop.Keeper = data.Number()
		shop.Trade1 = data.Number()
		shop.Trade2 = data.Number()
		shop.Trade3 = data.Number()
		shop.Trade4 = data.Number()
		shop.ProfitBuy = data.Number()
		shop.ProfitSell = data.Number()
		shop.OpenHour = data.Number()
		shop.CloseHour = data.Number()
		shop.Comment = data.EOL()
		infos = append(infos, shop)

		data.Gobble()
	}
	return infos
}

func loadResets(input string) []Reset {
	data := lexer.New(input)

	var infos []Reset

	// if no resets, get outta here
	if err := data.Jump("#RESETS"); err != nil {
		return infos
	}

	done := false
	for !done {
		// grab extraneous whitespace
		data.Gobble()

		// figure out type
		switch data.Current() {
		case 'M':
			data.Letter() // grab the M
			reset := resetReadMobile{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			reset.Limit = data.Number()
			reset.Room = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'O':
			data.Letter() // grab the D
			reset := resetReadObject{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			data.Number() // third number -- unused
			reset.Room = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'P':
			data.Letter() // grab the P
			reset := resetPutObject{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			data.Number() // third number -- unused
			reset.ContainerItem = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'G':
			data.Letter() // grab the G
			reset := resetGiveObject{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			data.Number() // third number -- unused
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'E':
			data.Letter() // grab the E
			reset := resetEquipObject{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			data.Number() // third number -- unused
			reset.WearLocation = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'D':
			data.Letter() // grab the D
			reset := resetSetDoorState{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			reset.Door = data.Number()
			reset.State = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'R':
			data.Letter() // grab the D
			reset := resetRandomizeExits{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			reset.LastDoor = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case '*':
			data.Letter() // grab the D
			reset := resetComment{}
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'S':
			done = true
		default:
		}

	}

	return infos
}

func loadSpecials(input string) []Special {
	data := lexer.New(input)

	var infos []Special

	// if no resets, get outta here
	if err := data.Jump("#SPECIALS"); err != nil {
		return infos
	}

	done := false
	for !done {
		// grab extraneous whitespace
		data.Gobble()

		// figure out type
		switch data.Current() {
		case 'M':
			data.Letter()
			special := specialsMob{}
			special.VNUM = data.Number()
			special.SpecFun = data.Word()
			special.Comment = data.EOL()
			infos = append(infos, special)
		case '*':
			data.Letter()
			special := specialsComment{}
			special.Comment = data.EOL()
			infos = append(infos, special)
		case 'S':
			done = true
		default:
			panic("unrecognized type: " + data.Snapshot())
		}

	}

	return infos
}

func loadMobiles(input string) []*Mobile {
	data := lexer.New(input)

	var infos []*Mobile

	// if no mobs, get out
	if err := data.Jump("#MOBILES"); err != nil {
		return infos
	}

	for {
		// grab any extraneous whitespace
		data.Gobble()

		// if we hit a #0, that's the end
		if data.Current() == '#' && data.Peek() == '0' {
			break
		}

		// Grab this mob
		data.Letter() // Gobble initial '#'
		mob := &Mobile{}
		mob.VNUM = data.Number()
		mob.Keywords = data.String()
		mob.ShortDescription = data.String()
		mob.LongDescription = data.String()
		mob.Description = data.String()
		mob.ActFlags = data.Number()
		mob.AffectedFlags = data.Number()
		mob.Alignment = data.Number()

		data.Word() // grab the next S - unused

		mob.Level = data.Number()

		data.Number() // hitroll -- unused
		data.Number() // armor -- unused
		data.Number() // hit dice -- unused
		data.Letter() // d -- unused
		data.Number() // hitsize -- unused
		data.Number() // hitplus -- unused
		data.Number() // damage dice -- unused
		data.Letter() // d -- unused
		data.Number() // damage size -- unused
		data.Number() // damage roll -- unused

		data.Number() // gold -- unused
		data.Number() // experience -- unused
		data.Number() // position 1 -- unused
		data.Number() // position 2 -- unused
		mob.Sex = data.Number()

		infos = append(infos, mob)
	}
	return infos
}

func loadRooms(input string) []*Room {
	data := lexer.New(input)

	var infos []*Room

	// if no rooms, get outta here
	if err := data.Jump("#ROOMS"); err != nil {
		return infos
	}

	data.Gobble()

	for {
		// end of section
		if data.Current() == '#' && data.Peek() == '0' {
			break
		} else if data.Current() == '#' {
			data.Letter()
		}

		room := &Room{}
		room.VNUM = data.Number()
		room.Name = data.String()
		room.Description = data.String()
		room.Area = data.Number()
		room.RoomFlags = data.Number()
		room.SectorType = data.Number()

		// get rid of extra whitespace
		data.Gobble()

		for {
			if data.Current() == 'D' {
				var d door
				// get the D
				data.Letter()
				d.Door = data.Number()
				d.Description = data.String()
				d.Keywords = data.String()
				d.Locks = data.Number()
				d.Key = data.Number()
				d.ToRoom = data.Number()

				room.Doors = append(room.Doors, d)
				data.Gobble()
			}

			if data.Current() == 'E' {
				data.Letter() // 'E' -- unused
				var xd extendedDescription
				xd.Keywords = data.String()
				xd.Description = data.String()
				room.ExtendedDescriptions = append(room.ExtendedDescriptions, xd)
				data.Gobble()
			}

			// a standalone 'S' means this room is done
			if data.Current() == 'S' {
				infos = append(infos, room)
				// skip everything until the next '#', which is
				// either the next room or the end of the list
				break
			}
		}

		for data.Next() != '#' {
			continue
		}
	}
	return infos
}

func loadObjects(input string) []Object {
	data := lexer.New(input)

	var infos []Object

	// if no objects, get outta here
	if err := data.Jump("#OBJECTS"); err != nil {
		return infos
	}

	// grab any extraneous whitespace
	data.Gobble()
	for {

		// if we hit a #0, that's the end
		if data.Current() == '#' && data.Peek() == '0' {
			break
		}

		// Grab this object section
		object := Object{}
		data.Letter() // '#' -- unused
		object.VNUM = data.Number()
		object.Keywords = data.String()
		object.ShortDescription = data.String()
		object.LongDescription = data.String()

		data.String() // action descriotion -- unused

		object.ItemType = data.Number()
		object.ExtraFlags = data.Number()
		object.WearFlags = data.Number()
		object.Value0 = data.Number()
		object.Value1 = data.Number()
		object.Value2 = data.Number()
		object.Value3 = data.Number()
		object.Weight = data.Number()

		data.Number() // cost -- unused
		data.Number() // cost per day -- unused

		// get rid of some whitespace
		data.Gobble()

		// get extra descriptions
		for {
			if data.Current() == 'E' {
				var xtra extraDescription
				data.Letter() // consume the E
				xtra.Keywords = data.String()
				xtra.Description = data.String()
				object.ExtraDescription = append(object.ExtraDescription, xtra)
				data.Gobble()
			} else if data.Current() == 'A' {
				var app apply
				data.Letter() // consume the A
				app.ApplyType = data.Number()
				app.ApplyValue = data.Number()
				object.Apply = append(object.Apply, app)
				data.Gobble()
			} else {
				break
			}
		}

		infos = append(infos, object)
	}
	return infos
}

func loadFile(areaName string) string {
	f, err := os.Open(fmt.Sprintf("areas/%s", areaName))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(body)
}
