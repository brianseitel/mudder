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

type Mobile struct {
	Raw string // raw data

	VNUM             int
	Keywords         string
	ShortDescription string
	LongDescription  string
	Description      string
	ActFlags         int
	AffectedFlags    int
	Alignment        int
	Level            int
	Sex              int
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
	ExtraDescription []extraDescription
	Apply            []apply
}

type extraDescription struct {
	Raw string // raw data

	Keywords    string
	Description string
}

type apply struct {
	Raw string // raw data

	ApplyType  int
	ApplyValue int
}

type Room struct {
	VNUM                 int
	Name                 string
	Description          string
	Area                 int
	RoomFlags            int
	SectorType           int
	Doors                []door
	ExtendedDescriptions []extendedDescription
}

type door struct {
	Door        int
	Description string
	Keywords    string
	Locks       int
	Key         int
	ToRoom      int
}

type extendedDescription struct {
	Keywords    string
	Description string
}

type Reset interface{}

type resetComment struct {
	Comment string
}

type resetReadMobile struct {
	VNUM    int
	Limit   int
	Room    int
	Comment string
}

type resetReadObject struct {
	VNUM    int
	Room    int
	Comment string
}

type resetPutObject struct {
	VNUM          int
	ContainerItem int
	Comment       string
}

type resetGiveObject struct {
	VNUM    int
	Comment string
}

type resetEquipObject struct {
	VNUM         int
	WearLocation int
	Comment      string
}

type resetSetDoorState struct {
	VNUM    int
	Door    int
	State   int
	Comment string
}

type resetRandomizeExits struct {
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

type specialsComment struct {
	Comment string
}

type specialsMob struct {
	VNUM    int
	SpecFun string
	Comment string
}
