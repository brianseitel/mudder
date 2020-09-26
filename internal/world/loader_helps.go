package world

import "github.com/brianseitel/mudder/internal/lexer"

// === The #HELPS section

// The syntax of this section is:

//     #HELPS
//     {
// 	<level:number> <keywords:string>
// 	<help-text:string>
//     }
//     0 $~

// The 'level' number is the minimum character level needed to read this section.
// This allows for immortal-only help text.

// The 'keywords' are a set of keywords for this help text.

// The 'help-text' is the help text itself.

// Normally when a player uses 'help', both the keywords and the help-text are
// shown.  If the 'level' is negative, however, the keywords are suppressed.  This
// allows the help file mechanism to be used for certain other commands, such as
// the initial 'greetings' text.

// If a 'help-text' begins with a leading '.', the leading '.' is stripped off.
// This provides for an escape mechanism from the usual leading-blank stripping of
// strings, so that picturesque greeting screens may be used.
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
