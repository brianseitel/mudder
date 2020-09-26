package world

import "github.com/brianseitel/mudder/internal/lexer"

/*
=== The #OBJECTS section

The syntax of this section is:

    #OBJECTS
    {
	#<vnum:number>
	<keywords:string>
	<short-description:string>
	<long-description:string>
	<action-description:string>
	<item-type:number> <extra-flags:number> <wear-flags:number>
	<value-0:number> <value-1:number> <value-2:number> <value-3:number>
	<weight:number> <cost:number> <cost-per-day:number>
	{
	    E
	    <keyword:string>
	    <description:string>
	}
	{
	    A
	    <apply-type:number> <apply-value:number>
	}
    }
    #0

The 'vnum' is the virtual number of the object.

The 'keywords' are words which can be used in commands to identify the object.

The 'short-description' is the description used by the 'act' function and other
functions to identify the object.  The first character of the short-description
should be lower case, because this description is used in the middle of
sentences.

The 'long-description' is the description used when a character walks in the
room and the object is visible.

The 'action-description' is not used.

The 'item-type' is the type of the item (weapon, armor, potion, et cetera).

The 'extra-flags' describe more attributes of the object.  The 'wear-flags'
describe whether the item can be picked up, and if so, what bodily locations
can wear it.

The interpretation of the four 'value' numbers depends upon the type of the
object.  Interpretations are given below.

The 'weight' of the object is just that.

'Cost' and 'cost-per-day' are ignored.  'Cost' is generated internally based on
the level of the object.  Because Merc has no rent, 'cost-per-day' is
completely ignored.

The optional 'E' sections and 'A' sections come after the main data.
An 'E' section ('extra description') contains a keyword-list and a string
associated with those keywords.  This description string is used when a
character looks at a word on the keyword list.

An 'A' section ('apply') contains an apply-type and an apply-value.  When a
character uses this object as equipment (holds, wields, or wears it), then
the value of 'apply-value' is added to the character attribute identified by
'apply-type'.  Not all 'apply-types' are implemented; you have to read the
function 'affect_modify' in handler.c to see exactly which ones are.

An object may have an unlimited number of 'E' and 'A' sections.
*/
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
