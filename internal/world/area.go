package world

type Zone struct {
	Area     Area      `json:"areas"`
	Helps    []Help    `json:"helps"`
	Mobiles  []*Mobile `json:"mobiles"`
	Objects  []Object  `json:"objects"`
	Rooms    []*Room   `json:"rooms"`
	Resets   []Reset   `json:"resets"`
	Shops    []Shop    `json:"shops"`
	Specials []Special `json:"specials"`
}

type Area struct {
	Raw string `json:"-"` // raw data

	// parsed out
	Name     string `json:"name"`
	MinLevel int    `json:"min_level"`
	MaxLevel int    `json:"max_level"`
	Author   string `json:"author"`
}

type Help struct {
	Raw string `json:"-"` // raw data

	Level    int    `json:"level"`
	Keywords string `json:"keywords"`
	Text     string `json:"text"`
}

type Mobile struct {
	Raw string `json:"-"` // raw data

	VNUM             int    `json:"vnum"`
	Keywords         string `json:"keywords"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	Description      string `json:"description"`
	ActFlags         int    `json:"act_flags"`
	AffectedFlags    int    `json:"affected_flags"`
	Alignment        int    `json:"alignment"`
	Level            int    `json:"level"`
	Sex              int    `json:"sex"`
}

type Object struct {
	Raw string `json:"-"` // raw data

	VNUM             int                `json:"vnum"`
	Keywords         string             `json:"keywords"`
	ShortDescription string             `json:"short_description"`
	LongDescription  string             `json:"long_description"`
	ItemType         int                `json:"item_type"`
	ExtraFlags       int                `json:"extra_flags"`
	WearFlags        int                `json:"wear_flags"`
	Value0           int                `json:"value0"`
	Value1           int                `json:"value1"`
	Value2           int                `json:"value2"`
	Value3           int                `json:"value3"`
	Weight           int                `json:"weight"`
	ExtraDescription []extraDescription `json:"extra_description"`
	Apply            []apply            `json:"apply"`
}

type extraDescription struct {
	Raw string `json:"-"` // raw data

	Keywords    string `json:"keywords"`
	Description string `json:"descriptions"`
}

type apply struct {
	Raw string `json:"-"` // raw data

	ApplyType  int `json:"apply_type"`
	ApplyValue int `json:"apply_values"`
}

type Room struct {
	VNUM                 int                   `json:"vnum"`
	Name                 string                `json:"name"`
	Description          string                `json:"description"`
	Area                 int                   `json:"area"`
	RoomFlags            int                   `json:"room_flags"`
	SectorType           int                   `json:"sector_type"`
	Doors                []door                `json:"doors"`
	ExtendedDescriptions []extendedDescription `json:"extended_descriptions"`
}

type door struct {
	Door        int    `json:"door"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Locks       int    `json:"locks"`
	Key         int    `json:"key"`
	ToRoom      int    `json:"to_room"`
}

type extendedDescription struct {
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
}

type Reset interface{}

type resetComment struct {
	Comment string `json:"comment"`
}

type resetReadMobile struct {
	VNUM    int    `json:"vnum"`
	Limit   int    `json:"limit"`
	Room    int    `json:"room"`
	Comment string `json:"comment"`
}

type resetReadObject struct {
	VNUM    int    `json:"vnum"`
	Room    int    `json:"room"`
	Comment string `json:"comment"`
}

type resetPutObject struct {
	VNUM          int    `json:"vnum"`
	ContainerItem int    `json:"container_item"`
	Comment       string `json:"comment"`
}

type resetGiveObject struct {
	VNUM    int    `json:"vnum"`
	Comment string `json:"comment"`
}

type resetEquipObject struct {
	VNUM         int    `json:"vnum"`
	WearLocation int    `json:"wear_location"`
	Comment      string `json:"comment"`
}

type resetSetDoorState struct {
	VNUM    int    `json:"vnum"`
	Door    int    `json:"door"`
	State   int    `json:"state"`
	Comment string `json:"comment"`
}

type resetRandomizeExits struct {
	VNUM     int    `json:"vnum"` // room vnum
	LastDoor int    `json:"last_door"`
	Comment  string `json:"comment"`
}

type Shop struct {
	Keeper     int    `json:"keeper"`
	Trade0     int    `json:"trade0"`
	Trade1     int    `json:"trade1"`
	Trade2     int    `json:"trade2"`
	Trade3     int    `json:"trade3"`
	Trade4     int    `json:"trade4"`
	ProfitBuy  int    `json:"profit_buy"`
	ProfitSell int    `json:"profit_sell"`
	OpenHour   int    `json:"open_hour"`
	CloseHour  int    `json:"close_hour"`
	Comment    string `json:"comment"`
}

type Special interface{}

type specialsComment struct {
	Comment string `json:"comment"`
}

type specialsMob struct {
	VNUM    int    `json:"vnum"`
	SpecFun string `json:"spec_fun"`
	Comment string `json:"comment"`
}
