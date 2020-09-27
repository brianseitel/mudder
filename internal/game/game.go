package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/brianseitel/mudder/internal/world"
	"github.com/brianseitel/mudder/internal/world/loader"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var gameWorld *Game

type Game struct {
	World  *world.World
	Player *world.Player
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	gameWorld = &Game{
		World:  world.New(),
		Player: nil,
	}

	gameWorld.World.Zones = loader.Load()
	gameWorld.World.Populate()

	findRoom(1324236)
	gameWorld.Player = &world.Player{
		CurrentRoom: findRoom(3700),
	}
}

func Start() {
	for {
		input := prompt(gameWorld.Player)

		err := interpret(input)
		if err != nil {
			gameWorld.Player.Send(err.Error())
		}
	}
}

func interpret(input string) error {
	var cmd Command

	found := false
	for _, command := range commandsMap {
		if strings.HasPrefix(command.Keyword, input) {
			cmd = command
			found = true
			break
		}
	}

	if !found {
		for _, social := range socials {
			if strings.HasPrefix(social.Keyword, input) {
				err := doSocial(gameWorld.Player, social, input)
				return err
			}
		}
	}

	if found {
		err := cmd.DoFunc(gameWorld.Player, input)
		return err
	}
	return errors.New("command not found")
}

func findRoom(vnum int) *world.Room {
	for _, zone := range gameWorld.World.Zones {
		for _, room := range zone.Rooms {
			if room.VNUM == vnum {
				return room
			}
		}
	}

	log.Info().Int("vnum", vnum).Msg("Cannot find room")
	return &world.Room{}
}

func prompt(player *world.Player) string {
	err := interpret("look")
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

	player.Send(fmt.Sprintf("[%s]\n", strings.Join(doors, " ")))

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text
}
