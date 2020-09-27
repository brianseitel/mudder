package world

type Mobile struct {
	Raw string // raw data

	// Loaded from file
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
