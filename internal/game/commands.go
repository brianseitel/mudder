package game

import (
	"errors"
	"fmt"
	"os"

	"github.com/brianseitel/mudder/internal/world"
)

type commandFunc func(ch *world.Player, args string) error

type Command struct {
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

var commandsMap = map[string]Command{
	// Common movement commands
	"north": {doNorth, 0, 0, 0},
	"south": {doSouth, 0, 0, 0},
	"west":  {doWest, 0, 0, 0},
	"east":  {doEast, 0, 0, 0},
	"up":    {doUp, 0, 0, 0},
	"down":  {doDown, 0, 0, 0},

	// Miscellaneous
	"look": {doLook, POS_DEAD, 0, LOG_NORMAL},
	"qui":  {doQui, POS_DEAD, 0, LOG_NORMAL},
	"quit": {doQuit, POS_DEAD, 0, LOG_NORMAL},
}

func doQui(ch *world.Player, args string) error {
	return errors.New(`You must spell out "quit" in order to leave.`)
}

func doQuit(ch *world.Player, args string) error {
	fmt.Println("Seeya!")
	os.Exit(1)
	return nil
}
