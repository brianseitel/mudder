package loader

import (
	"github.com/brianseitel/mudder/internal/lexer"
	"github.com/brianseitel/mudder/internal/world"
)

// === The #AREA section
// The syntax of this section is:
//     #AREA	<area-name:string>
// The 'area-name' can be any string.  The 'areas' command provides a list of
// areas, so it's worth while to follow the standard Merc format for this string:
//     #AREA	{ 5 35} Merc    Prototype for New Area~
// The first two numbers are recommended level range.  The name is the name of the
// original author of the area.  The last phrase is the name of the area.
func loadArea(input string) world.Area {
	data := lexer.New(input)

	var info world.Area

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
