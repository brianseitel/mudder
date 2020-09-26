package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/brianseitel/mudder/internal/world"
	"github.com/brianseitel/mudder/internal/world/loader"
)

var gameWorld *Game

type Game struct {
	World  *world.World
	Player *world.Player
}

func init() {
	gameWorld = &Game{
		World:  world.New(),
		Player: nil,
	}

	gameWorld.World.Zones = loader.Load()
	gameWorld.World.Populate()

	gameWorld.Player = &world.Player{
		CurrentRoom: findRoom(3700),
	}
}

func Start() {
	for {
		input := prompt(gameWorld.Player)
		if cmd, ok := commandsMap[input]; ok {
			err := cmd.DoFunc(gameWorld.Player, input)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func findRoom(vnum int) *world.Room {
	for _, zone := range gameWorld.World.Zones {
		for _, room := range zone.Rooms {
			if room.VNUM == vnum {
				return room
			}
		}
	}

	panic("shit! can't find room")
}

func prompt(player *world.Player) string {
	cmd, ok := commandsMap["look"]
	_ = ok

	err := cmd.DoFunc(player, "")
	if err != nil {
		panic(err)
	}

	var doors []string
	for _, door := range player.CurrentRoom.Doors {
		var d string
		switch door.Door {
		case 0:
			d = "north"
		case 1:
			d = "east"
		case 2:
			d = "south"
		case 3:
			d = "west"
		case 4:
			d = "up"
		case 5:
			d = "down"
		}
		doors = append(doors, d)
	}

	fmt.Printf("[%s]\n", strings.Join(doors, " "))

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text
}
