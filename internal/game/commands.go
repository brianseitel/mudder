package game

import (
	"errors"
	"os"

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
	POS_DEAD = iota
	POS_MORTAL
	POS_INCAP
	POS_STUNNED
	POS_SLEEPING
	POS_RESTING
	POS_FIGHTING
	POS_STANDING
)

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
	{"look", doLook, POS_DEAD, 0, LOG_NORMAL},
	{"scan", doScan, POS_DEAD, 0, LOG_NORMAL},

	// object commands
	{"get", doGet, POS_DEAD, 0, LOG_NORMAL},
	{"drop", doDrop, POS_DEAD, 0, LOG_NORMAL},

	// player commands
	{"inventory", doInventory, POS_DEAD, 0, LOG_NORMAL},

	// Miscellaneous
	{"qui", doQui, POS_DEAD, 0, LOG_NORMAL},
	{"quit", doQuit, POS_DEAD, 0, LOG_NORMAL},
}

func doQui(ch *world.Character, args string) error {
	return errors.New(`You must spell out "quit" in order to leave.`)
}

func doQuit(ch *world.Character, args string) error {
	ch.Send("Seeya!")
	os.Exit(1)
	return nil
}
