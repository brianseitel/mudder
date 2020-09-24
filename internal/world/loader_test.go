package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type areaCase struct {
	Input  string
	Output areaCaseOutput
}

type areaCaseOutput struct {
	MinLevel int
	MaxLevel int
	Author   string
	Name     string
}

func TestLoadArea(t *testing.T) {
	cases := []areaCase{
		{
			Input:  `#SHOPS`,
			Output: areaCaseOutput{},
		},
		{
			Input: `#AREA	{ 5 35} Merc    Prototype for New Area~`,
			Output: areaCaseOutput{
				MinLevel: 5,
				MaxLevel: 35,
				Author:   "Merc",
				Name:     "Prototype for New Area",
			},
		},
		{
			Input: `#AREA	{ 1  5} Hatchet Mud School~
`,
			Output: areaCaseOutput{
				MinLevel: 1,
				MaxLevel: 5,
				Author:   "Hatchet",
				Name:     "Mud School",
			},
		},
		{
			Input: `#AREA	{ None} Test A Test Area~
`,
			Output: areaCaseOutput{
				MinLevel: -1,
				MaxLevel: -1,
				Author:   "Test",
				Name:     "A Test Area",
			},
		},
		{
			Input: `#AREA	{ All} Foo Bar Mud~
`,
			Output: areaCaseOutput{
				MinLevel: 1,
				MaxLevel: 9999,
				Author:   "Foo",
				Name:     "Bar Mud",
			},
		},
	}

	for _, c := range cases {
		info := loadArea(c.Input)

		assert.Equal(t, c.Output.MinLevel, info.MinLevel)
		assert.Equal(t, c.Output.MaxLevel, info.MaxLevel)
		assert.Equal(t, c.Output.Author, info.Author)
		assert.Equal(t, c.Output.Name, info.Name)
	}

}

func TestLoadHelps(t *testing.T) {
	data := `#HELPS



-1 DIKU~
.                    Original game idea, concept, and design:

          Katja Nyboe               [Superwoman] (katz@freja.diku.dk)
          Tom Madsen              [Stormbringer] (noop@freja.diku.dk)
          Hans Henrik Staerfeldt           [God] (bombman@freja.diku.dk)
          Michael Seifert                 [Papi] (seifert@freja.diku.dk)
          Sebastian Hammer               [Quinn] (quinn@freja.diku.dk)

                     Additional contributions from:

Michael Curran  - the player title collection and additional locations.
Ragnar Loenn    - the bulletin board.
Bill Wisner     - for being the first to successfully port the game,
                  uncovering several old bugs, uh, inconsistencies,
                  in the process.

And: Mads Haar and Stephan Dahl for additional locations.

Developed at: DIKU -- The Department of Computer Science
                      at the University of Copenhagen.

~

0 $~

#$`

	helps := loadHelps(data)

	assert.Equal(t, 1, len(helps))
	assert.Equal(t, -1, helps[0].Level)
	assert.Equal(t, "DIKU", helps[0].Keywords)
	assert.Equal(t, `                    Original game idea, concept, and design:

          Katja Nyboe               [Superwoman] (katz@freja.diku.dk)
          Tom Madsen              [Stormbringer] (noop@freja.diku.dk)
          Hans Henrik Staerfeldt           [God] (bombman@freja.diku.dk)
          Michael Seifert                 [Papi] (seifert@freja.diku.dk)
          Sebastian Hammer               [Quinn] (quinn@freja.diku.dk)

                     Additional contributions from:

Michael Curran  - the player title collection and additional locations.
Ragnar Loenn    - the bulletin board.
Bill Wisner     - for being the first to successfully port the game,
                  uncovering several old bugs, uh, inconsistencies,
                  in the process.

And: Mads Haar and Stephan Dahl for additional locations.

Developed at: DIKU -- The Department of Computer Science
                      at the University of Copenhagen.

`, helps[0].Text)
}

func TestLoadShops(t *testing.T) {
	type testCases struct {
		Input         string
		ExpectResults bool
		Output        struct {
			Keeper int
		}
	}
	cases := []testCases{
		{
			Input: `#SHOPS
 3717	 1 15 17 19  0	 100  100	 0 23	; Adept of Frag
0`,
			ExpectResults: true,
			Output: struct {
				Keeper int
			}{
				Keeper: 3717,
			},
		},
		{
			Input:         `#OBJECTS `,
			ExpectResults: false,
			Output: struct{ Keeper int }{
				Keeper: 0,
			},
		},
	}

	for _, c := range cases {
		shops := loadShops(c.Input)
		if c.ExpectResults {
			assert.Equal(t, c.Output.Keeper, shops[0].Keeper)
		}
	}
}

func TestLoadResets(t *testing.T) {
	data := `#RESETS
M 0 3700 1 3701           (acolyte of FRAG)
D 0 3700 2 1
O 0 3700 0 123 object read
P 0 3700 0 123 put in a container
G 0 3700 0 give an object
E 0 3700 0 123 equip an item
R 0 3700 99 a door
* this is a comment
S`

	resets := loadResets(data)
	// Type M
	assert.Equal(t, 3700, resets[0].(resetReadMobile).VNUM)
	assert.Equal(t, 3701, resets[0].(resetReadMobile).Room)
	assert.Equal(t, "(acolyte of FRAG)", resets[0].(resetReadMobile).Comment)

	// Type D
	assert.Equal(t, 3700, resets[1].(resetSetDoorState).VNUM)
	assert.Equal(t, 2, resets[1].(resetSetDoorState).Door)
	assert.Equal(t, 1, resets[1].(resetSetDoorState).State)

	// Type O
	assert.Equal(t, 3700, resets[2].(resetReadObject).VNUM)
	assert.Equal(t, 123, resets[2].(resetReadObject).Room)
	assert.Equal(t, "object read", resets[2].(resetReadObject).Comment)

	// Type P
	assert.Equal(t, 3700, resets[3].(resetPutObject).VNUM)
	assert.Equal(t, 123, resets[3].(resetPutObject).ContainerItem)
	assert.Equal(t, "put in a container", resets[3].(resetPutObject).Comment)

	// Type G
	assert.Equal(t, 3700, resets[4].(resetGiveObject).VNUM)
	assert.Equal(t, "give an object", resets[4].(resetGiveObject).Comment)

	// Type E
	assert.Equal(t, 3700, resets[5].(resetEquipObject).VNUM)
	assert.Equal(t, 123, resets[5].(resetEquipObject).WearLocation)
	assert.Equal(t, "equip an item", resets[5].(resetEquipObject).Comment)

	// Type R
	assert.Equal(t, 3700, resets[6].(resetRandomizeExits).VNUM)
	assert.Equal(t, 99, resets[6].(resetRandomizeExits).LastDoor)
	assert.Equal(t, "a door", resets[6].(resetRandomizeExits).Comment)

	// Type * (Comment)
	assert.Equal(t, "this is a comment", resets[7].(resetComment).Comment)
}

func TestLoadResetsNoResets(t *testing.T) {
	data := `#SHOPS`

	resets := loadResets(data)

	assert.Empty(t, resets)
}

func TestLoadSpecialsNoSpecials(t *testing.T) {
	data := `#SHOPS`

	specials := loadSpecials(data)

	assert.Empty(t, specials)
}

func TestLoadSpecials(t *testing.T) {
	data := `#SPECIALS
M  3707 spec_cast_adept
M  3714 spec_fido
*  this is a test fun
S`

	specials := loadSpecials(data)

	assert.Len(t, specials, 3)
	assert.Equal(t, 3707, specials[0].(specialsMob).VNUM)
	assert.Equal(t, "spec_cast_adept", specials[0].(specialsMob).SpecFun)

	assert.Equal(t, 3714, specials[1].(specialsMob).VNUM)
	assert.Equal(t, "spec_fido", specials[1].(specialsMob).SpecFun)

	assert.Equal(t, "this is a test fun", specials[2].(specialsComment).Comment)
}

func TestLoadMobiles(t *testing.T) {
	data := `#MOBILES
#3700
acolyte cleric~
the acolyte of Frag~
An acolyte of Overlord Frag is here, grinning at you.
~
He is big and bad.  Don't mess with him.
~
2|64 8|128 1000 S
30 -10 -20 1d1+29999 2d4+30
0 -99000
8 8 1
#3701
blob~
the blob~
The blob is here, waiting to eat you up.
~
He is big, he is bad.  You don't want to fight him when he isn't chained up.
Perhaps now would be a good time to flee from him!!!  If you don't flee, you
would really suffer the consequences.
~
2|64 8 0 S
5 20 -99 1d1+50 1d1+0
0 -9999
8 8 0
#0`

	mobs := loadMobiles(data)

	assert.Len(t, mobs, 2)

	assert.Equal(t, 3700, mobs[0].VNUM)
	assert.Equal(t, "acolyte cleric", mobs[0].Keywords)
	assert.Equal(t, "the acolyte of Frag", mobs[0].ShortDescription)
	assert.Equal(t, "An acolyte of Overlord Frag is here, grinning at you.\n", mobs[0].LongDescription)
	assert.Equal(t, "He is big and bad.  Don't mess with him.\n", mobs[0].Description)
	assert.Equal(t, 66, mobs[0].ActFlags)
	assert.Equal(t, 136, mobs[0].AffectedFlags)
	assert.Equal(t, 1000, mobs[0].Alignment)
	assert.Equal(t, 30, mobs[0].Level)
	assert.Equal(t, 1, mobs[0].Sex)

	assert.Equal(t, 3701, mobs[1].VNUM)
	assert.Equal(t, "blob", mobs[1].Keywords)
	assert.Equal(t, "the blob", mobs[1].ShortDescription)
	assert.Equal(t, "The blob is here, waiting to eat you up.\n", mobs[1].LongDescription)
	assert.Equal(t, `He is big, he is bad.  You don't want to fight him when he isn't chained up.
Perhaps now would be a good time to flee from him!!!  If you don't flee, you
would really suffer the consequences.
`, mobs[1].Description)
	assert.Equal(t, 66, mobs[1].ActFlags)
	assert.Equal(t, 8, mobs[1].AffectedFlags)
	assert.Equal(t, 0, mobs[1].Alignment)
	assert.Equal(t, 5, mobs[1].Level)
	assert.Equal(t, 0, mobs[1].Sex)
}

func TestLoadMobilesNoMobs(t *testing.T) {
	data := `#OBJECTS`

	mobs := loadMobiles(data)

	assert.Len(t, mobs, 0)
}

func TestLoadRooms(t *testing.T) {
	data := `
#ROOMS
#3700
Entrance to Mud School~
  This is the entrance to the Merc Mud School.  Go north to go through mud
school.  If you have been here before and want to go directly to the arena,
go south.
~
2 4|8 0
D5
You see the Temple of Midgaard.
~
~
0 0 3001
D0
You see the doorway into the Mud School Building.
~
~
0 0 3757
D2
You see the one way door into the Arena of Mud School.
~
door~
1 -1 3744
S
#3701
A Room in Mud School~
You are in a square white room.  The walls are all blank, with no windows.
Light fluoresces off the ceiling in soft white tones.  Of course, there is a
sign on the wall.  The exits are west and south.  A small plaque is on the
wall.
~
2 4|8 0
D2
You see a path into another room.
~
~
0 -1 3757
D3
You see a path into another room.
~
~
0 -1 3702
E
plaque~
This zone (Mud School) is created by Hatchet for any Merc Mud.
Copyright 1992, 1993.
~
E
sign~
Equipment check time!  Type 'EQUIPMENT' to see your current equipment.  Right
now, you have just a vest, a shield, and a weapon.  As you go through Mud
School you will acquire a complete set of equipment.

To pick up items on the ground as you see them, type 'GET item'.
To wear a piece of armor, type 'WEAR item'.
To hold a light source, type 'HOLD item'.
To wield a weapon, type 'WIELD weapon'.
To wear, hold, and wield everything you have, type 'WEAR ALL'.
To stop using a piece of equipment, type 'REMOVE item'.
Finally, to see the items in your inventory, type 'INVENTORY'.

When you are ready to continue, go west.
~
S
#0`

	rooms := loadRooms(data)

	assert.Len(t, rooms, 2)
	assert.Equal(t, 3700, rooms[0].VNUM)
	assert.Equal(t, 3701, rooms[1].VNUM)
}
func TestLoadRoomsNoRooms(t *testing.T) {
	data := `
#SHOPS`

	rooms := loadRooms(data)

	assert.Len(t, rooms, 0)
}

func TestLoadObjects(t *testing.T) {
	data := `#OBJECTS
#3700
mace~
a sub issue mace~
You see a sub issue mace here.~
~
5 64 1|8192
0 2 3 7
5 0 5
E
mace~
You see a mace of great but cheap craftsmanship.  Imprinted on the side is:
Merc Industries
~
A
18 1
A
19 1
#0`

	objs := loadObjects(data)

	assert.Len(t, objs, 1)
	assert.Equal(t, 3700, objs[0].VNUM)
	assert.Equal(t, "mace", objs[0].Keywords)
	assert.Equal(t, "a sub issue mace", objs[0].ShortDescription)
	assert.Equal(t, "You see a sub issue mace here.", objs[0].LongDescription)
	assert.Equal(t, 5, objs[0].ItemType)
	assert.Equal(t, 64, objs[0].ExtraFlags)
	assert.Equal(t, 8193, objs[0].WearFlags)
	assert.Equal(t, 0, objs[0].Value0)
	assert.Equal(t, 2, objs[0].Value1)
	assert.Equal(t, 3, objs[0].Value2)
	assert.Equal(t, 7, objs[0].Value3)
	assert.Equal(t, 5, objs[0].Weight)
}

func TestLoadObjectsNoObjects(t *testing.T) {
	data := `#ROOM`

	objs := loadObjects(data)

	assert.Len(t, objs, 0)
}
