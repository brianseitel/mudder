package game

import (
	"errors"
	"os"

	"github.com/brianseitel/mudder/internal/positions"
	"github.com/brianseitel/mudder/internal/world"
)

type commandFunc func(ch *world.Character, args string) error

type Command struct {
	Keyword         string
	DoFunc          commandFunc
	MinimumPosition int
	MinimumLevel    int
	LoggingLevel    int
}

const (
	LOG_NORMAL = iota
	LOG_ALWAYS
	LOG_DEAD
)

var commandsMap = []Command{
	// Common movement commands
	{"north", doNorth, 0, 0, 0},
	{"south", doSouth, 0, 0, 0},
	{"west", doWest, 0, 0, 0},
	{"east", doEast, 0, 0, 0},
	{"up", doUp, 0, 0, 0},
	{"down", doDown, 0, 0, 0},

	// info commands
	{"look", doLook, positions.POS_DEAD, 0, LOG_NORMAL},
	{"scan", doScan, positions.POS_DEAD, 0, LOG_NORMAL},
	{"inspect", doInspect, positions.POS_DEAD, 0, LOG_NORMAL}, // TODO: make this admin only

	// object commands
	{"get", doGet, positions.POS_DEAD, 0, LOG_NORMAL},
	{"drop", doDrop, positions.POS_DEAD, 0, LOG_NORMAL},

	// player commands
	{"inventory", doInventory, positions.POS_DEAD, 0, LOG_NORMAL},
	{"save", doSave, positions.POS_DEAD, 0, LOG_NORMAL},

	// battle commands
	{"kill", doKill, positions.POS_DEAD, 0, LOG_NORMAL},

	// Miscellaneous
	{"qui", doQui, positions.POS_DEAD, 0, LOG_NORMAL},
	{"quit", doQuit, positions.POS_DEAD, 0, LOG_NORMAL},
}

func doQui(ch *world.Character, args string) error {
	return errors.New(`You must spell out "quit" in order to leave.`)
}

func doQuit(ch *world.Character, args string) error {
	ch.Send("Seeya!")
	os.Exit(1)
	return nil
}
