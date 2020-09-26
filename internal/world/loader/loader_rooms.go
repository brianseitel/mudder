package loader

import (
	"github.com/brianseitel/mudder/internal/lexer"
	"github.com/brianseitel/mudder/internal/world"
)

/*
=== The #ROOMS section

The syntax of this section is:

    #ROOMS
    {
	#<vnum:number>
	<name:string>
	<description:string>
	<area:number> <room-flags:number> <sector-type:number>
	{
	    D <door:number>
	    <description:string>
	    <keywords:string>
	    <locks:number> <key:number> <to_room:number>
	}
	{
	    E
	    <keywords:string>
	    <description:string>
	}
	S
    }
    #0

The 'vnum' is the virtual number of the room.

The 'name' is the name of the room.

The 'description' is the long multi-line description of the room.

The 'area' is obsolete and unused.  Rooms belong to whatever area was most
recently defined with #AREA.

The 'room-flags' describe more attributes of the room.

The 'sector-type' identifies the type of terrain.  This affects movement cost
through the room.  Certain sector types (air and boat) require special
capabilities to enter.

Unlike mobiles and objects, rooms don't have any keywords associated with them.
One may not manipulate a room in the same way one manipulates a mobile or
object.

The optional 'D' sections and 'E' sections come after the main data.  A 'D'
section contains a 'door' in the range from 0 to 5:

	0	north
	1	east
	2	south
	3	west
	4	up
	5	down

A 'D' command also contains a 'description' for that direction, and 'keywords'
for manipulating the door.  'Doors' include not just real door, but any kind of
exit from the room.  The 'locks' value is 0 for an unhindered exit, 1 for a
door, and 2 for a pick-proof door.  The 'key' value is the vnum of an object
which locks and unlocks the door.  Lastly, 'to_room' is the vnum of the room to
which this door leads.

You must specify two 'D' sections, one for each side of the door.  If you
specify just one then you'll get a one-way exit.

An 'E' section (extended description) contains a 'keywords' string and a
'description' string.  As you might guess, looking at one of the words in
'keywords' yields the 'description' string.

The 'S' at the end marks the end of the room.  It is not optional.
*/
func loadRooms(input string) []*world.Room {
	data := lexer.New(input)

	var infos []*world.Room

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

		room := &world.Room{}
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
				var d world.Door
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
				var xd world.ExtendedDescription
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
