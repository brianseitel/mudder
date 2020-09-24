package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/brianseitel/mudder/internal/world"
)

var gameWorld *Game

type Game struct {
	World *world.World

	CurrentRoom *world.Room
}

func init() {
	gameWorld = &Game{
		World: world.Load(),
	}

	gameWorld.CurrentRoom = findRoom(3700)
}

func Start() {
	for {
		cmd := prompt(gameWorld.CurrentRoom)
		switch cmd {
		case "north", "south", "west", "east", "up", "down",
			"n", "s", "e", "w", "u", "d":
			gameWorld.move(cmd)
		}
	}
}

func (g *Game) move(direction string) {
	doorCode := -1
	switch direction {
	case "north", "n":
		doorCode = 0
		direction = "north"
	case "east", "e":
		doorCode = 1
		direction = "east"
	case "south", "s":
		doorCode = 2
		direction = "south"
	case "west", "w":
		doorCode = 3
		direction = "west"
	case "up", "u":
		doorCode = 4
		direction = "up"
	case "down", "d":
		doorCode = 5
		direction = "down"
	default:
		fmt.Println("i don't know direction", direction)
	}

	if doorCode < 0 {
		fmt.Println("HUH?")
		return
	}

	found := false
	for _, d := range g.CurrentRoom.Doors {
		if d.Door == doorCode {
			g.CurrentRoom = findRoom(d.ToRoom)
			fmt.Println("You go", direction)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("That exit doesn't exist!")
	}
}

func findRoom(vnum int) *world.Room {
	for _, area := range gameWorld.World.Zones {
		for _, room := range area.Rooms {
			if room.VNUM == vnum {
				return room
			}
		}
	}

	panic("shit! can't find room")
}

func prompt(room *world.Room) string {
	fmt.Println("["+strconv.Itoa(room.VNUM)+"]", room.Name)
	fmt.Println(room.Description)

	var doors []string
	for _, door := range room.Doors {
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
