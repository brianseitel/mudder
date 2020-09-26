package loader

import (
	"github.com/brianseitel/mudder/internal/lexer"
	"github.com/brianseitel/mudder/internal/world"
)

/*
=== The #MOBILES section

The syntax of this section is:

    #MOBILES
    {
	#<vnum:number>
	<keywords:string>
	<short-description:string>
	<long-description:string>
	<description:string>
	<act-flags:number> <affected-flags:number> <alignment:number> S
	<level:number> <hitroll:number> <armor:number>			\
	<hitnodice:number> d <hitsizedice:number> + <hitplus:number>	\
	<damnodice:number> d <damsizedice:number> + <damroll:number>
	<gold:number> <exp:number>
	<position:number> <position:number> <sex:number>
    }
    #0

The 'vnum' is the virtual number of the mobile.

The 'keywords' are words which can be used in commands to identify the mobile.

The 'short-description' is the description used by the 'act' function and other
functions to identify the mobile.

The 'long-description' is the description used when a character walks in the
room and the mobile is visible.

The 'description' is the longest description.  It is used when a character
explicitly looks at the mobile.

The 'act-flags' define how the mobile acts, and the 'affected-flags' define
more attributes of the mobile.

The 'alignment' of the mobile ranges from -1000 to +1000.  Keep in mind that
certain spells ('protection' and 'dispel evil') give characters fighting evil
monsters an advantage, and that experience earned is influenced by alignment.

The literal letter 'S' must be present after the alignment.  In the original
Diku mob format, 'S' stands for simple.  Merc supports only simple mobs, so the
'S' is redundant.  It is retained not only for compatibility with the Diku
format, but also because it helps the server report errors more accurately.

The 'level' is typically a number from 1 to 35, although there is no upper
limit.

The 'hitroll', 'armor', 'hitnodice, 'hitsizedice', 'hitplus', 'damnodice',
'damsizedice', 'damroll', 'gold', 'exp', 'position', and 'position' fields are
present for compatibility with original Diku mud, but their values are ignored.
Merc generates these values internally based on the level of the mobile.

The 'sex' value may be 0 for neutral, 1 for male, and 2 for female.
*/
func loadMobiles(input string) []*world.Mobile {
	data := lexer.New(input)

	var infos []*world.Mobile

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
		mob := &world.Mobile{}
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
