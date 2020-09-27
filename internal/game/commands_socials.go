package game

import "github.com/brianseitel/mudder/internal/world"

type Social struct {
	Keyword string
	Actions [7]string
}

func doSocial(ch *world.Player, social Social, args string) error {
	ch.Send(social.Actions[0])
	return nil
}

var socials = []Social{
	{
		Keyword: "accuse",
		Actions: [7]string{
			"Accuse whom?",
			"$n is in an accusing mood.",
			"You look accusingly at $M.",
			"$n looks accusingly at $N.",
			"$n looks accusingly at you.",
			"You accuse yourself.",
			"$n seems to have a bad conscience.",
		},
	},
	{
		Keyword: "applaud",
		Actions: [7]string{
			"Clap, clap, clap.",
			"$n gives a round of applause.",
			"You clap at $S actions.",
			"$n claps at $N's actions.",
			"$n gives you a round of applause.  You MUST'VE done something good!",
			"You applaud at yourself.  Boy, are we conceited!",
			"$n applauds at $mself.  Boy, are we conceited!",
		},
	},

	{
		Keyword: "bark",
		Actions: [7]string{
			"Woof!  Woof!",
			"$n barks like a dog.",
			"You bark at $M.",
			"$n barks at $N.",
			"$n barks at you.",
			"You bark at yourself.  Woof!  Woof!",
			"$n barks at $mself.  Woof!  Woof!",
		},
	},

	{
		Keyword: "beer",
		Actions: [7]string{
			"You down a cold, frosty beer.",
			"$n downs a cold, frosty beer.",
			"You draw a cold, frosty beer for $N.",
			"$n draws a cold, frosty beer for $N.",
			"$n draws a cold, frosty beer for you.",
			"You draw yourself a beer.",
			"$n draws $mself a beer.",
		},
	},

	{
		Keyword: "beg",
		Actions: [7]string{
			"You beg the gods for mercy.",
			"The gods fall down laughing at $n's request for mercy.",
			"You desperately try to squeeze a few coins from $M.",
			"$n begs you for money.",
			"$n begs $N for a gold piece!",
			"Begging yourself for money doesn't help.",
			"$n begs himself for money.",
		},
	},

	{
		Keyword: "blush",
		Actions: [7]string{
			"Your cheeks are burning.",
			"$n blushes.",
			"You get all flustered up seeing $M.",
			"$n blushes as $e sees $N here.",
			"$n blushes as $e sees you here.  Such an effect on people!",
			"You blush at your own folly.",
			"$n blushes as $e notices $s boo-boo.",
		},
	},

	{
		Keyword: "bounce",
		Actions: [7]string{
			"BOIINNNNNNGG!",
			"$n bounces around.",
			"You bounce onto $S lap.",
			"$n bounces onto $N's lap.",
			"$n bounces onto your lap.",
			"You bounce your head like a basketball.",
			"$n plays basketball with $s head.",
		},
	},

	{
		Keyword: "bow",
		Actions: [7]string{
			"You bow deeply.",
			"$n bows deeply.",
			"You bow before $M.",
			"$n bows before $N.",
			"$n bows before you.",
			"You kiss your toes.",
			"$n folds up like a jack knife and kisses $s own toes.",
		},
	},

	{
		Keyword: "burp",
		Actions: [7]string{
			"You burp loudly.",
			"$n burps loudly.",
			"You burp loudly to $M in response.",
			"$n burps loudly in response to $N's remark.",
			"$n burps loudly in response to your remark.",
			"You burp at yourself.",
			"$n burps at $mself.  What a sick sight.",
		},
	},

	{
		Keyword: "cackle",
		Actions: [7]string{
			"You throw back your head and cackle with insane glee!",
			"$n throws back $s head and cackles with insane glee!",
			"You cackle gleefully at $N",
			"$n cackles gleefully at $N.",
			"$n cackles gleefully at you.  Better keep your distance from $m.",
			"You cackle at yourself.  Now, THAT'S strange!",
			"$n is really crazy now!  $e cackles at $mself.",
		},
	},

	{
		Keyword: "chuckle",
		Actions: [7]string{
			"You chuckle politely.",
			"$n chuckles politely.",
			"You chuckle at $S joke.",
			"$n chuckles at $N's joke.",
			"$n chuckles at your joke.",
			"You chuckle at your own joke, since no one else would.",
			"$n chuckles at $s own joke, since none of you would.",
		},
	},

	{
		Keyword: "clap",
		Actions: [7]string{
			"You clap your hands together.",
			"$n shows $s approval by clapping $s hands together.",
			"You clap at $S performance.",
			"$n claps at $N's performance.",
			"$n claps at your performance.",
			"You clap at your own performance.",
			"$n claps at $s own performance.",
		},
	},

	{
		Keyword: "comb",
		Actions: [7]string{
			"You comb your hair - perfect.",
			"$n combs $s hair, how dashing!",
			"You patiently untangle $N's hair - what a mess!",
			"$n tries patiently to untangle $N's hair.",
			"$n pulls your hair in an attempt to comb it.",
			"You pull your hair, but it will not be combed.",
			"$n tries to comb $s tangled hair.",
		},
	},

	{
		Keyword: "comfort",
		Actions: [7]string{
			"Do you feel uncomfortable?",
			"",
			"You comfort $M.",
			"$n comforts $N.",
			"$n comforts you.",
			"You make a vain attempt to comfort yourself.",
			"$n has no one to comfort $m but $mself.",
		},
	},

	{
		Keyword: "cringe",
		Actions: [7]string{
			"You cringe in terror.",
			"$n cringes in terror!",
			"You cringe away from $M.",
			"$n cringes away from $N in mortal terror.",
			"$n cringes away from you.",
			"I beg your pardon?",
			"",
		},
	},

	{
		Keyword: "cry",
		Actions: [7]string{
			"Waaaaah ...",
			"$n bursts into tears.",
			"You cry on $S shoulder.",
			"$n cries on $N's shoulder.",
			"$n cries on your shoulder.",
			"You cry to yourself.",
			"$n sobs quietly to $mself.",
		},
	},

	{
		Keyword: "cuddle",
		Actions: [7]string{
			"Whom do you feel like cuddling today?",
			"",
			"You cuddle $M.",
			"$n cuddles $N.",
			"$n cuddles you.",
			"You must feel very cuddly indeed ... :)",
			"$n cuddles up to $s shadow.  What a sorry sight.",
		},
	},

	{
		Keyword: "curse",
		Actions: [7]string{
			"You swear loudly for a long time.",
			"$n swears: @*&^%@*&!",
			"You swear at $M.",
			"$n swears at $N.",
			"$n swears at you!  Where are $s manners?",
			"You swear at your own mistakes.",
			"$n starts swearing at $mself.  Why don't you help?",
		},
	},

	{
		Keyword: "curtsey",
		Actions: [7]string{
			"You curtsey to your audience.",
			"$n curtseys gracefully.",
			"You curtsey to $M.",
			"$n curtseys gracefully to $N.",
			"$n curtseys gracefully for you.",
			"You curtsey to your audience (yourself).",
			"$n curtseys to $mself, since no one is paying attention to $m.",
		},
	},

	{
		Keyword: "dance",
		Actions: [7]string{
			"Feels silly, doesn't it?",
			"$n tries to break dance, but nearly breaks $s neck!",
			"You sweep $M into a romantic waltz.",
			"$n sweeps $N into a romantic waltz.",
			"$n sweeps you into a romantic waltz.",
			"You skip and dance around by yourself.",
			"$n dances a pas-de-une.",
		},
	},

	/*
	 * This one's for Baka, Penn, and Onethumb!
	 */
	{
		Keyword: "drool",
		Actions: [7]string{
			"You drool on yourself.",
			"$n drools on $mself.",
			"You drool all over $N.",
			"$n drools all over $N.",
			"$n drools all over you.",
			"You drool on yourself.",
			"$n drools on $mself.",
		},
	},

	{
		Keyword: "fart",
		Actions: [7]string{
			"Where are your manners?",
			"$n lets off a real rip-roarer ... a greenish cloud envelops $n!",
			"You fart at $M.  Boy, you are sick.",
			"$n farts in $N's direction.  Better flee before $e turns to you!",
			"$n farts in your direction.  You gasp for air.",
			"You fart at yourself.  You deserve it.",
			"$n farts at $mself.  Better $m than you.",
		},
	},

	{
		Keyword: "flip",
		Actions: [7]string{
			"You flip head over heels.",
			"$n flips head over heels.",
			"You flip $M over your shoulder.",
			"$n flips $N over $s shoulder.",
			"$n flips you over $s shoulder.  Hmmmm.",
			"You tumble all over the room.",
			"$n does some nice tumbling and gymnastics.",
		},
	},

	{
		Keyword: "fondle",
		Actions: [7]string{
			"Who needs to be fondled?",
			"",
			"You fondly fondle $M.",
			"$n fondly fondles $N.",
			"$n fondly fondles you.",
			"You fondly fondle yourself, feels funny doesn't it ?",
			"$n fondly fondles $mself - this is going too far !!",
		},
	},

	{
		Keyword: "french",
		Actions: [7]string{
			"Kiss whom?",
			"",
			"You give $N a long and passionate kiss.",
			"$n kisses $N passionately.",
			"$n gives you a long and passionate kiss.",
			"You gather yourself in your arms and try to kiss yourself.",
			"$n makes an attempt at kissing $mself.",
		},
	},

	{
		Keyword: "frown",
		Actions: [7]string{
			"What's bothering you ?",
			"$n frowns.",
			"You frown at what $E did.",
			"$n frowns at what $E did.",
			"$n frowns at what you did.",
			"You frown at yourself.  Poor baby.",
			"$n frowns at $mself.  Poor baby.",
		},
	},

	{
		Keyword: "fume",
		Actions: [7]string{
			"You grit your teeth and fume with rage.",
			"$n grits $s teeth and fumes with rage.",
			"You stare at $M, fuming.",
			"$n stares at $N, fuming with rage.",
			"$n stares at you, fuming with rage!",
			"That's right - hate yourself!",
			"$n clenches $s fists and stomps his feet, fuming with anger.",
		},
	},

	{
		Keyword: "gasp",
		Actions: [7]string{
			"You gasp in astonishment.",
			"$n gasps in astonishment.",
			"You gasp as you realize what $e did.",
			"$n gasps as $e realizes what $N did.",
			"$n gasps as $e realizes what you did.",
			"You look at yourself and gasp!",
			"$n takes one look at $mself and gasps in astonisment!",
		},
	},

	{
		Keyword: "giggle",
		Actions: [7]string{
			"You giggle.",
			"$n giggles.",
			"You giggle in $S's presence.",
			"$n giggles at $N's actions.",
			"$n giggles at you.  Hope it's not contagious!",
			"You giggle at yourself.  You must be nervous or something.",
			"$n giggles at $mself.  $e must be nervous or something.",
		},
	},

	{
		Keyword: "glare",
		Actions: [7]string{
			"You glare at nothing in particular.",
			"$n glares around $m.",
			"You glare icily at $M.",
			"$n glares at $N.",
			"$n glares icily at you, you feel cold to your bones.",
			"You glare icily at your feet, they are suddenly very cold.",
			"$n glares at $s feet, what is bothering $m?",
		},
	},

	{
		Keyword: "grin",
		Actions: [7]string{
			"You grin evilly.",
			"$n grins evilly.",
			"You grin evilly at $M.",
			"$n grins evilly at $N.",
			"$n grins evilly at you.  Hmmm.  Better keep your distance.",
			"You grin at yourself.  You must be getting very bad thoughts.",
			"$n grins at $mself.  You must wonder what's in $s mind.",
		},
	},

	{
		Keyword: "groan",
		Actions: [7]string{
			"You groan loudly.",
			"$n groans loudly.",
			"You groan at the sight of $M.",
			"$n groans at the sight of $N.",
			"$n groans at the sight of you.",
			"You groan as you realize what you have done.",
			"$n groans as $e realizes what $e has done.",
		},
	},

	{
		Keyword: "grope",
		Actions: [7]string{
			"Whom do you wish to grope?",
			"",
			"Well, what sort of noise do you expect here?",
			"$n gropes $N.",
			"$n gropes you.",
			"You grope yourself - YUCK.",
			"$n gropes $mself - YUCK.",
		},
	},

	{
		Keyword: "grovel",
		Actions: [7]string{
			"You grovel in the dirt.",
			"$n grovels in the dirt.",
			"You grovel before $M.",
			"$n grovels in the dirt before $N.",
			"$n grovels in the dirt before you.",
			"That seems a little silly to me.",
			"",
		},
	},

	{
		Keyword: "growl",
		Actions: [7]string{
			"Grrrrrrrrrr ...",
			"$n growls.",
			"Grrrrrrrrrr ... take that, $N!",
			"$n growls at $N.  Better leave the room before the fighting starts.",
			"$n growls at you.  Hey, two can play it that way!",
			"You growl at yourself.  Boy, do you feel bitter!",
			"$n growls at $mself.  This could get interesting...",
		},
	},

	{
		Keyword: "grumble",
		Actions: [7]string{
			"You grumble.",
			"$n grumbles.",
			"You grumble to $M.",
			"$n grumbles to $N.",
			"$n grumbles to you.",
			"You grumble under your breath.",
			"$n grumbles under $s breath.",
		},
	},

	{
		Keyword: "grunt",
		Actions: [7]string{
			"GRNNNHTTTT.",
			"$n grunts like a pig.",
			"GRNNNHTTTT.",
			"$n grunts to $N.  What a pig!",
			"$n grunts to you.  What a pig!",
			"GRNNNHTTTT.",
			"$n grunts to nobody in particular.  What a pig!",
		},
	},

	{
		Keyword: "hand",
		Actions: [7]string{
			"Kiss whose hand?",
			"",
			"You kiss $S hand.",
			"$n kisses $N's hand.  How continental!",
			"$n kisses your hand.  How continental!",
			"You kiss your own hand.",
			"$n kisses $s own hand.",
		},
	},

	{
		Keyword: "hop",
		Actions: [7]string{
			"You hop around like a little kid.",
			"",
			"",
			"",
			"",
			"",
			"",
		},
	},

	{
		Keyword: "hug",
		Actions: [7]string{
			"Hug whom?",
			"",
			"You hug $M.",
			"$n hugs $N.",
			"$n hugs you.",
			"You hug yourself.",
			"$n hugs $mself in a vain attempt to get friendship.",
		},
	},

	{
		Keyword: "kiss",
		Actions: [7]string{
			"Isn't there someone you want to kiss?",
			"",
			"You kiss $M.",
			"$n kisses $N.",
			"$n kisses you.",
			"All the lonely people :(",
			"",
		},
	},

	{
		Keyword: "laugh",
		Actions: [7]string{
			"You laugh.",
			"$n laughs.",
			"You laugh at $N mercilessly.",
			"$n laughs at $N mercilessly.",
			"$n laughs at you mercilessly.  Hmmmmph.",
			"You laugh at yourself.  I would, too.",
			"$n laughs at $mself.  Let's all join in!!!",
		},
	},

	{
		Keyword: "lick",
		Actions: [7]string{
			"You lick your lips and smile.",
			"$n licks $s lips and smiles.",
			"You lick $M.",
			"$n licks $N.",
			"$n licks you.",
			"You lick yourself.",
			"$n licks $mself - YUCK.",
		},
	},

	{
		Keyword: "love",
		Actions: [7]string{
			"You love the whole world.",
			"$n loves everybody in the world.",
			"You tell your true feelings to $N.",
			"$n whispers softly to $N.",
			"$n whispers to you sweet words of love.",
			"Well, we already know you love yourself (lucky someone does!)",
			"$n loves $mself, can you believe it ?",
		},
	},

	{
		Keyword: "massage",
		Actions: [7]string{
			"Massage what?  Thin air?",
			"",
			"You gently massage $N's shoulders.",
			"$n massages $N's shoulders.",
			"$n gently massages your shoulders.  Ahhhhhhhhhh ...",
			"You practice yoga as you try to massage yourself.",
			"$n gives a show on yoga positions, trying to massage $mself.",
		},
	},

	{
		Keyword: "moan",
		Actions: [7]string{
			"You start to moan.",
			"$n starts moaning.",
			"You moan for the loss of $m.",
			"$n moans for the loss of $N.",
			"$n moans at the sight of you.  Hmmmm.",
			"You moan at yourself.",
			"$n makes $mself moan.",
		},
	},

	{
		Keyword: "nibble",
		Actions: [7]string{
			"Nibble on whom?",
			"",
			"You nibble on $N's ear.",
			"$n nibbles on $N's ear.",
			"$n nibbles on your ear.",
			"You nibble on your OWN ear.",
			"$n nibbles on $s OWN ear.",
		},
	},

	{
		Keyword: "nod",
		Actions: [7]string{
			"You nod your silly head off.",
			"$n nods $s silly head off.",
			"You nod in recognition to $M.",
			"$n nods in recognition to $N.",
			"$n nods in recognition to you.  You DO know $m, right?",
			"You nod at yourself.  Are you getting senile?",
			"$n nods at $mself.  $e must be getting senile.",
		},
	},

	{
		Keyword: "nudge",
		Actions: [7]string{
			"Nudge whom?",
			"",
			"You nudge $M.",
			"$n nudges $N.",
			"$n nudges you.",
			"You nudge yourself, for some strange reason.",
			"$n nudges $mself, to keep $mself awake.",
		},
	},

	{
		Keyword: "nuzzle",
		Actions: [7]string{
			"Nuzzle whom?",
			"",
			"You nuzzle $S neck softly.",
			"$n softly nuzzles $N's neck.",
			"$n softly nuzzles your neck.",
			"I'm sorry, friend, but that's impossible.",
			"",
		},
	},

	{
		Keyword: "pat",
		Actions: [7]string{
			"Pat whom?",
			"",
			"You pat $N on $S ass.",
			"$n pats $N on $S ass.",
			"$n pats you on your ass.",
			"You pat yourself on your ass, very sensual.",
			"$n pats $mself on the ass.",
		},
	},

	{
		Keyword: "point",
		Actions: [7]string{
			"Point at whom?",
			"",
			"You point at $M accusingly.",
			"$n points at $N accusingly.",
			"$n points at you accusingly.",
			"You point proudly at yourself.",
			"$n points proudly at $mself.",
		},
	},

	{
		Keyword: "poke",
		Actions: [7]string{
			"Poke whom?",
			"",
			"You poke $M in the ribs.",
			"$n pokes $N in the ribs.",
			"$n pokes you in the ribs.",
			"You poke yourself in the ribs, feeling very silly.",
			"$n pokes $mself in the ribs, looking very sheepish.",
		},
	},

	{
		Keyword: "ponder",
		Actions: [7]string{
			"You ponder the question.",
			"$n sits down and thinks deeply.",
			"",
			"",
			"",
			"",
			"",
		},
	},

	{
		Keyword: "pout",
		Actions: [7]string{
			"Ah, don't take it so hard.",
			"$n pouts.",
			"",
			"",
			"",
			"",
			"",
		},
	},

	{
		Keyword: "pray",
		Actions: [7]string{
			"You feel righteous, and maybe a little foolish.",
			"$n begs and grovels to the powers that be.",
			"You crawl in the dust before $M.",
			"$n falls down and grovels in the dirt before $N.",
			"$n kisses the dirt at your feet.",
			"Talk about narcissism ...",
			"$n mumbles a prayer to $mself.",
		},
	},

	{
		Keyword: "puke",
		Actions: [7]string{
			"You puke ... chunks everywhere!",
			"$n pukes.",
			"You puke on $M.",
			"$n pukes on $N.",
			"$n spews vomit and pukes all over your clothing!",
			"You puke on yourself.",
			"$n pukes on $s clothes.",
		},
	},

	{
		Keyword: "punch",
		Actions: [7]string{
			"Punch whom?",
			"",
			"You punch $M playfully.",
			"$n punches $N playfully.",
			"$n punches you playfully.  OUCH!",
			"You punch yourself.  You deserve it.",
			"$n punches $mself.  Why don't you join in?",
		},
	},

	{
		Keyword: "purr",
		Actions: [7]string{
			"MMMMEEEEEEEEOOOOOOOOOWWWWWWWWWWWW.",
			"$n purrs contentedly.",
			"You purr contentedly in $M lap.",
			"$n purrs contentedly in $N's lap.",
			"$n purrs contentedly in your lap.",
			"You purr at yourself.",
			"$n purrs at $mself.  Must be a cat thing.",
		},
	},

	{
		Keyword: "ruffle",
		Actions: [7]string{
			"You've got to ruffle SOMEONE.",
			"",
			"You ruffle $N's hair playfully.",
			"$n ruffles $N's hair playfully.",
			"$n ruffles your hair playfully.",
			"You ruffle your hair.",
			"$n ruffles $s hair.",
		},
	},

	{
		Keyword: "scream",
		Actions: [7]string{
			"ARRRRRRRRRRGH!!!!!",
			"$n screams loudly!",
			"ARRRRRRRRRRGH!!!!!  Yes, it MUST have been $S fault!!!",
			"$n screams loudly at $N.  Better leave before $n blames you, too!!!",
			"$n screams at you!  That's not nice!  *sniff*",
			"You scream at yourself.  Yes, that's ONE way of relieving tension!",
			"$n screams loudly at $mself!  Is there a full moon up?",
		},
	},

	{
		Keyword: "shake",
		Actions: [7]string{
			"You shake your head.",
			"$n shakes $s head.",
			"You shake $S hand.",
			"$n shakes $N's hand.",
			"$n shakes your hand.",
			"You are shaken by yourself.",
			"$n shakes and quivers like a bowl full of jelly.",
		},
	},

	{
		Keyword: "shiver",
		Actions: [7]string{
			"Brrrrrrrrr.",
			"$n shivers uncomfortably.",
			"You shiver at the thought of fighting $M.",
			"$n shivers at the thought of fighting $N.",
			"$n shivers at the suicidal thought of fighting you.",
			"You shiver to yourself?",
			"$n scares $mself to shivers.",
		},
	},

	{
		Keyword: "shrug",
		Actions: [7]string{
			"You shrug.",
			"$n shrugs helplessly.",
			"You shrug in response to $s question.",
			"$n shrugs in response to $N's question.",
			"$n shrugs in respopnse to your question.",
			"You shrug to yourself.",
			"$n shrugs to $mself.  What a strange person.",
		},
	},

	{
		Keyword: "sigh",
		Actions: [7]string{
			"You sigh.",
			"$n sighs loudly.",
			"You sigh as you think of $M.",
			"$n sighs at the sight of $N.",
			"$n sighs as $e thinks of you.  Touching, huh?",
			"You sigh at yourself.  You MUST be lonely.",
			"$n sighs at $mself.  What a sorry sight.",
		},
	},

	{
		Keyword: "sing",
		Actions: [7]string{
			"You raise your clear voice towards the sky.",
			"$n has begun to sing.",
			"You sing a ballad to $m.",
			"$n sings a ballad to $N.",
			"$n sings a ballad to you!  How sweet!",
			"You sing a little ditty to yourself.",
			"$n sings a little ditty to $mself.",
		},
	},

	{
		Keyword: "smile",
		Actions: [7]string{
			"You smile happily.",
			"$n smiles happily.",
			"You smile at $M.",
			"$n beams a smile at $N.",
			"$n smiles at you.",
			"You smile at yourself.",
			"$n smiles at $mself.",
		},
	},

	{
		Keyword: "smirk",
		Actions: [7]string{
			"You smirk.",
			"$n smirks.",
			"You smirk at $S saying.",
			"$n smirks at $N's saying.",
			"$n smirks at your saying.",
			"You smirk at yourself.  Okay ...",
			"$n smirks at $s own 'wisdom'.",
		},
	},

	{
		Keyword: "snap",
		Actions: [7]string{
			"PRONTO ! You snap your fingers.",
			"$n snaps $s fingers.",
			"You snap back at $M.",
			"$n snaps back at $N.",
			"$n snaps back at you!",
			"You snap yourself to attention.",
			"$n snaps $mself to attention.",
		},
	},

	{
		Keyword: "snarl",
		Actions: [7]string{
			"You grizzle your teeth and look mean.",
			"$n snarls angrily.",
			"You snarl at $M.",
			"$n snarls at $N.",
			"$n snarls at you, for some reason.",
			"You snarl at yourself.",
			"$n snarls at $mself.",
		},
	},

	{
		Keyword: "sneeze",
		Actions: [7]string{
			"Gesundheit!",
			"$n sneezes.",
			"",
			"",
			"",
			"",
			"",
		},
	},

	{
		Keyword: "snicker",
		Actions: [7]string{
			"You snicker softly.",
			"$n snickers softly.",
			"You snicker with $M about your shared secret.",
			"$n snickers with $N about their shared secret.",
			"$n snickers with you about your shared secret.",
			"You snicker at your own evil thoughts.",
			"$n snickers at $s own evil thoughts.",
		},
	},

	{
		Keyword: "sniff",
		Actions: [7]string{
			"You sniff sadly. *SNIFF*",
			"$n sniffs sadly.",
			"You sniff sadly at the way $E is treating you.",
			"$n sniffs sadly at the way $N is treating $m.",
			"$n sniffs sadly at the way you are treating $m.",
			"You sniff sadly at your lost opportunities.",
			"$n sniffs sadly at $mself.  Something MUST be bothering $m.",
		},
	},

	{
		Keyword: "snore",
		Actions: [7]string{
			"Zzzzzzzzzzzzzzzzz.",
			"$n snores loudly.",
			"",
			"",
			"",
			"",
			"",
		},
	},

	{
		Keyword: "snowball",
		Actions: [7]string{
			"Whom do you want to throw a snowball at?",
			"",
			"You throw a snowball in $N's face.",
			"$n throws a snowball at $N.",
			"$n throws a snowball at you.",
			"You throw a snowball at yourself.",
			"$n throws a snowball at $mself.",
		},
	},

	{
		Keyword: "snuggle",
		Actions: [7]string{
			"Who?",
			"",
			"you snuggle $M.",
			"$n snuggles up to $N.",
			"$n snuggles up to you.",
			"You snuggle up, getting ready to sleep.",
			"$n snuggles up, getting ready to sleep.",
		},
	},

	{
		Keyword: "spank",
		Actions: [7]string{
			"Spank whom?",
			"",
			"You spank $M playfully.",
			"$n spanks $N playfully.",
			"$n spanks you playfully.  OUCH!",
			"You spank yourself.  Kinky!",
			"$n spanks $mself.  Kinky!",
		},
	},

	{
		Keyword: "squeeze",
		Actions: [7]string{
			"Where, what, how, whom?",
			"",
			"You squeeze $M fondly.",
			"$n squeezes $N fondly.",
			"$n squeezes you fondly.",
			"You squeeze yourself - try to relax a little!",
			"$n squeezes $mself.",
		},
	},

	{
		Keyword: "stare",
		Actions: [7]string{
			"You stare at the sky.",
			"$n stares at the sky.",
			"You stare dreamily at $N, completely lost in $S eyes..",
			"$n stares dreamily at $N.",
			"$n stares dreamily at you, completely lost in your eyes.",
			"You stare dreamily at yourself - enough narcissism for now.",
			"$n stares dreamily at $mself - NARCISSIST!",
		},
	},

	{
		Keyword: "strut",
		Actions: [7]string{
			"Strut your stuff.",
			"$n struts proudly.",
			"You strut to get $S attention.",
			"$n struts, hoping to get $N's attention.",
			"$n struts, hoping to get your attention.",
			"You strut to yourself, lost in your own world.",
			"$n struts to $mself, lost in $s own world.",
		},
	},

	{
		Keyword: "sulk",
		Actions: [7]string{
			"You sulk.",
			"$n sulks in the corner.",
			"",
			"",
			"",
			"",
			"",
		},
	},

	{
		Keyword: "thank",
		Actions: [7]string{
			"Thank you too.",
			"",
			"You thank $N heartily.",
			"$n thanks $N heartily.",
			"$n thanks you heartily.",
			"You thank yourself since nobody else wants to !",
			"$n thanks $mself since you won't.",
		},
	},

	{
		Keyword: "tickle",
		Actions: [7]string{
			"Whom do you want to tickle?",
			"",
			"You tickle $N.",
			"$n tickles $N.",
			"$n tickles you - hee hee hee.",
			"You tickle yourself, how funny!",
			"$n tickles $mself.",
		},
	},

	{
		Keyword: "twiddle",
		Actions: [7]string{
			"You patiently twiddle your thumbs.",
			"$n patiently twiddles $s thumbs.",
			"You twiddle $S ears.",
			"$n twiddles $N's ears.",
			"$n twiddles your ears.",
			"You twiddle your ears like Dumbo.",
			"$n twiddles $s own ears like Dumbo.",
		},
	},

	{
		Keyword: "wave",
		Actions: [7]string{
			"You wave.",
			"$n waves happily.",
			"You wave goodbye to $N.",
			"$n waves goodbye to $N.",
			"$n waves goodbye to you.  Have a good journey.",
			"Are you going on adventures as well?",
			"$n waves goodbye to $mself.",
		},
	},

	{
		Keyword: "whistle",
		Actions: [7]string{
			"You whistle appreciatively.",
			"$n whistles appreciatively.",
			"You whistle at the sight of $M.",
			"$n whistles at the sight of $N.",
			"$n whistles at the sight of you.",
			"You whistle a little tune to yourself.",
			"$n whistles a little tune to $mself.",
		},
	},

	{
		Keyword: "wiggle",
		Actions: [7]string{
			"Your wiggle your bottom.",
			"$n wiggles $s bottom.",
			"You wiggle your bottom toward $M.",
			"$n wiggles $s bottom toward $N.",
			"$n wiggles his bottom toward you.",
			"You wiggle about like a fish.",
			"$n wiggles about like a fish.",
		},
	},

	{
		Keyword: "wince",
		Actions: [7]string{
			"You wince.  Ouch!",
			"$n winces.  Ouch!",
			"You wince at $M.",
			"$n winces at $N.",
			"$n winces at you.",
			"You wince at yourself.  Ouch!",
			"$n winces at $mself.  Ouch!",
		},
	},

	{
		Keyword: "wink",
		Actions: [7]string{
			"You wink suggestively.",
			"$n winks suggestively.",
			"You wink suggestively at $N.",
			"$n winks at $N.",
			"$n winks suggestively at you.",
			"You wink at yourself ?? - what are you up to ?",
			"$n winks at $mself - something strange is going on...",
		},
	},

	{
		Keyword: "yawn",
		Actions: [7]string{
			"You must be tired.",
			"$n yawns.",
			"",
			"",
			"",
			"",
			"",
		},
	},
}
