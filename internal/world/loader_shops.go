package world

import "github.com/brianseitel/mudder/internal/lexer"

/*
=== The #SHOPS section

The syntax of this section is:

    #SHOPS
    {
	<keeper:number>						\
	<trade-0:number> <trade-1:number> <trade-2:number>	\
	<trade-3:number> <trade-4:number>			\
	<profit-buy:number> <profit-sell:number>		\
	<open-hour:number> <close-hour:number>			\
	<comment:to_eol>
    }
    0

Like the #RESETS section, the #SHOPS section has one command per line.

The 'keeper' is the vnum of the mobile who is the shopkeeper.  All mobiles
with that vnum will be shopkeepers.

The 'trade-0' through 'trade-5' numbers are item types which the shopkeeper
will buy.  Unused slots should have a '0' in them; for instance, a shopkeeper
who doesn't buy anything would have five zeroes.

The 'profit-buy' number is a markup for players buying the item, in percentage
points.  100 is nominal price; 150 is 50% markup, and so on.  The 'profit-sell'
number is a markdown for players selling the item, in percentage points.
100 is nominal price; 75 is a 25% markdown, and so on.  The buying markup
should be at least 100, and the selling markdown should be at most 100.

The 'open-hour' and 'close-hour' numbers define the hours when the shopkeeper
will do business.  For a 24-hour shop, these numbers would be 0 and 23.

Everything beyond 'close-hour' to the end of the line is taken to be a comment.

Note that there is no room number for a shop.  Just load the shopkeeper mobile
into the room of your choice, and make it a sentinel.  Or, for a roving
shopkeeper, just make it non-sentinel.

The objects a shopkeeper sells are exactly those loaded by 'G' reset commands
for that shopkeeper.  These items replenish automatically.  If a player sells
an object to a shopkeeper, the shopkeeper will keep it for resale if he, she,
or it doesn't already have an identical object.  These items do not replenish.
*/

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
