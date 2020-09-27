package world

type Zone struct {
	Area     Area
	Helps    []Help
	Mobiles  []*Mobile
	Objects  []Object
	Rooms    []*Room
	Resets   []Reset
	Shops    []Shop
	Specials []Special
}

type Area struct {
	Raw string // raw data

	// parsed out
	Name     string
	MinLevel int
	MaxLevel int
	Author   string
}

type Help struct {
	Raw string // raw data

	Level    int
	Keywords string
	Text     string
}

type Object struct {
	Raw string // raw data

	VNUM             int
	Keywords         string
	ShortDescription string
	LongDescription  string
	ItemType         int
	ExtraFlags       int
	WearFlags        int
	Value0           int
	Value1           int
	Value2           int
	Value3           int
	Weight           int
	ExtraDescription []ExtraDescription
	Apply            []Apply
}

type ExtraDescription struct {
	Raw string // raw data

	Keywords    string
	Description string
}

type Apply struct {
	Raw string // raw data

	ApplyType  int
	ApplyValue int
}

type ExtendedDescription struct {
	Keywords    string
	Description string
}

type Reset interface{}

type ResetComment struct {
	Comment string
}

type ResetReadMobile struct {
	VNUM    int
	Limit   int
	Room    int
	Comment string
}

type ResetReadObject struct {
	VNUM    int
	Room    int
	Comment string
}

type ResetPutObject struct {
	VNUM          int
	ContainerItem int
	Comment       string
}

type ResetGiveObject struct {
	VNUM    int
	Comment string
}

type ResetEquipObject struct {
	VNUM         int
	WearLocation int
	Comment      string
}

type ResetSetDoorState struct {
	VNUM    int
	Door    int
	State   int
	Comment string
}

type ResetRandomizeExits struct {
	VNUM     int // room vnum
	LastDoor int
	Comment  string
}

type Shop struct {
	Keeper     int
	Trade0     int
	Trade1     int
	Trade2     int
	Trade3     int
	Trade4     int
	ProfitBuy  int
	ProfitSell int
	OpenHour   int
	CloseHour  int
	Comment    string
}

type Special interface{}

type SpecialsComment struct {
	Comment string
}

type SpecialsMob struct {
	VNUM    int
	SpecFun string
	Comment string
}
