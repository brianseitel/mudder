package loader

import (
	"github.com/brianseitel/mudder/internal/lexer"
	"github.com/brianseitel/mudder/internal/world"
)

/*
=== The #RESETS section

The syntax of this section is:

    #RESETS
    {
	* <comment:to_eol>
    }
    {
	M <:number> <mob-vnum:number> <limit:number> <room-vnum:number>	\
	<comment:to_eol>
    }
    {
	O <:number> <obj-vnum:number> <:number> <room-vnum:number>	\
	<comment:to_eol>
    }
    {
	P <:number> <obj-vnum:number> <:number> <obj-vnum:number>	\
	<comment:to_eol>
    }
    {
	G <:number> <obj-vnum:number> <:number>				\
	<comment:to_eol>
    }
    {
	E <:number> <obj-vnum:number> <:number> <wear_loc:number>	\
	<comment:to_eol>
    }
    {
	D <:number> <room-vnum:number> <door:number> <state:number>	\
	<comment:to_eol>
    }
    {
	R <:number> <room-vnum:number> <last-door:number>		\
	<comment:to_eol>
    }
    S

To reset an area, the server executes each command in the list of reset
commands once.  Each area is reset once when the server loads, and again
periodically as it ages.  An area is reset if it is at least 3 area-minutes old
and is empty of players, or if it is 15 area-minutes old.  At the 14
area-minute mark, each (awake) player in the area is warned of the impending
reset.  These values are coded into the function 'reset_area' in 'db.c'.

An 'area-minute' varies between 30 and 90 seconds of real time, with an
average of 60 seconds.  The variation defeats area timekeepers.

The 'resets' section contains a series of single lines.  The backslashes and
line splitting above are for readability; they are not part of the file format.
Because of the end-of-line comments, this section is not as free-format as
other sections.

The reset commands are:

    *	comment
    M	read a mobile
    O	read an object
    P	put object in object
    G	give object to mobile
    E	equip object to mobile
    D	set state of door
    R	randomize room exits
    S	stop (end of list)

The '*' lines contain comments.  The 'S' line is the last line of the section.

Every other command contains four numbers (three for the 'G' command).  The
first number is ignored.  The next three (or two) numbers are interpreted as
follows:

For the 'M' command, the second number is the vnum of a mobile to load.  The
third number is the limit of how many of this mobile may be present in the
world.  The fourth number is the vnum of the room where the mobile is loaded.

For the 'O', 'P', 'G', and 'E' commands, the second number is the vnum of an
object to load.  The third number is ignored.

For the 'O' command, the fourth number is the vnum of the room where the object
is loaded.  The object is not loaded if the target room already contains any
objects with this vnum.  The object is also not loaded if any players are
present in the area.

For the 'P' command, the fourth number is the vnum of a container object where
the object will be loaded.  The actual container used is the most recently
loaded object with the right vnum; for best results, there should be only one
such container in the world.  The object is not loaded if no container object
exists, or if someone is carrying it, or if it already contains one of the
to-be-loaded object.

For the 'G' command, there is no fourth number.  If the most recent 'M' command
succeeded (e.g. the mobile limit wasn't exceeded), the object is given to that
mobile.  If the most recent 'M' command failed (due to hitting mobile limit),
then the object is not loaded.

For the 'E' command, the fourth number is an equipment location.  If the most
recent 'M' command succeeded, that mobile is equipped with the object.  If the
most recent 'M' command failed, then the object is not loaded.

All objects have a level limit, which is computed by inheritance from the most
recently read 'M' command (whether it succeeded or not) in 'area_update' in
'db.c'.  As distributed, an object's level equals the mobile level minus 2,
clipped to the range 0 to 35.

For the 'D' command, the second number is the vnum of a room.  The third number
is a door number from 0 to 5.  The fourth number indicates how to set the door:
0 for open and unlocked; 1 for closed and unlocked; 2 for closed and locked.

Room exits must be coherent: if room 1 has an exit to room 2, and room 2 has an
exit in the reverse direction, that exit must go back to room 1.  This doesn't
prevent one-way exits; room 2 doesn't HAVE to have an exit in the reverse
direction.

For the 'R' command, the second number is the vnum of a room.  The third number
is a door number.  When this command, the doors from 0 to the indicated door
number are shuffled.  The room will still have the same exits leading to the
same other rooms as before, but the directions will be different.  Thus, a door
number of 4 makes a two-dimensional maze room; a door number of 6 makes a
three-dimensional maze room.

Use of both the 'D' and 'R' commands on the same room will yield unpredicatable
results.

Any line (except an 'S' line) may have a comment at the end.
*/
func loadResets(input string) []world.Reset {
	data := lexer.New(input)

	var infos []world.Reset

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
			reset := world.ResetReadMobile{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			reset.Limit = data.Number()
			reset.Room = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'O':
			data.Letter() // grab the O
			reset := world.ResetReadObject{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			data.Number() // third number -- unused
			reset.Room = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'P':
			data.Letter() // grab the P
			reset := world.ResetPutObject{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			data.Number() // third number -- unused
			reset.ContainerItem = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'G':
			data.Letter() // grab the G
			reset := world.ResetGiveObject{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			data.Number() // third number -- unused
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'E':
			data.Letter() // grab the E
			reset := world.ResetEquipObject{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			data.Number() // third number -- unused
			reset.WearLocation = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'D':
			data.Letter() // grab the D
			reset := world.ResetSetDoorState{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			reset.Door = data.Number()
			reset.State = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'R':
			data.Letter() // grab the D
			reset := world.ResetRandomizeExits{}
			data.Number() // first number -- unused
			reset.VNUM = data.Number()
			reset.LastDoor = data.Number()
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case '*':
			data.Letter() // grab the D
			reset := world.ResetComment{}
			reset.Comment = data.EOL()
			infos = append(infos, reset)
		case 'S':
			done = true
		default:
		}

	}

	return infos
}
