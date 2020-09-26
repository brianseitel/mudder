package world

import "github.com/brianseitel/mudder/internal/lexer"

/*
=== The #SPECIALS section

The syntax of this section is:

    #SPECIALS
    {
	* <comment_to_eol>
    }
    {
	M <mob-vnum:number> <spec-fun:word> <comment:to_eol>
    }
    S

Like the #RESETS section, the #SPECIALS section has one command per line.

This section defines special functions (spec-fun's) for mobiles.  A spec-fun
is a C function which gives additional behavior to all mobiles with a given
vnum, such as the peripatetic mayor or the beholder casting spells in combat.
See 'special.c' for a list of available spec-fun's.

The 'M' command assigns 'spec-fun' to all mobiles of with virtual number
'mob-vnum'.  All spec-fun's are assigned by name.  An 'M' line may have a
comment at the end.

Every three seconds, the server function 'mobile_update' examines every mobile
in the game.  If the mobile has an associated spec-fun, then 'mobile_update'
calls that spec-fun with a single parameter, the 'ch' pointer for that mob.
The spec-fun returns TRUE if the mobile did something, or FALSE if it did not.
If the spec-fun returns TRUE, then further activity by that mobile is
suppressed.

To add a new special function:

(1) Add a DECLARE_SPEC_FUN line to the top of 'special.c'.

(2) Add a line for translating the ascii name of the function into a function
    pointer to the function 'spec_lookup' in 'special.c'.

(3) Write the spec-fun and add it to 'special.c'.  Note that Merc special
    functions take a single parameter, rather than the three parameters of
    Diku.  If you have an Ansi C compiler, you're protected against accidental
    mismatches.

(4) Assign the spec-fun by writing an appropriate line into the #SPECIALS
    section in an area file.  Any number of mobs may have the same spec-fun.

*/
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
