package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/brianseitel/mudder/internal/lexer"
	"github.com/brianseitel/mudder/internal/positions"
	"github.com/brianseitel/mudder/internal/world"
	"github.com/brianseitel/mudder/internal/world/loader"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var gameWorld *Game

type Game struct {
	World     *world.World
	Character *world.Character
}

func bootstrap() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	gameWorld = &Game{
		World:     world.New(),
		Character: nil,
	}

	gameWorld.World.Zones = loader.Load()
	gameWorld.World.Populate()

	gameWorld.Character = &world.Character{
		CurrentRoom: findRoom(3700),

		Position: positions.POS_STANDING,

		Level: 1,

		HitPoints:    100,
		MaxHitPoints: 100,
		Mana:         100,
		MaxMana:      100,
		Movement:     100,
		MaxMovement:  100,
	}
}

func Start() {
	bootstrap()

	// start updates
	go ticker()

	_ = interpret("look")
	for {
		input := prompt(gameWorld.Character)

		err := interpret(input)
		if err != nil {
			gameWorld.Character.Send(err.Error())
		}
	}
}

func interpret(input string) error {
	if len(strings.TrimSpace(input)) == 0 {
		return nil
	}
	var cmd Command

	found := false
	word, args := lexer.SplitArgs(input)
	for _, command := range commandsMap {
		if strings.HasPrefix(command.Keyword, word) {
			cmd = command
			found = true
			break
		}
	}

	if !found {
		for _, social := range socials {
			if strings.HasPrefix(social.Keyword, input) {
				err := doSocial(gameWorld.Character, social, input)
				return err
			}
		}
	}

	if found {
		err := cmd.DoFunc(gameWorld.Character, args)
		return err
	}
	return errors.New("command not found: " + input)
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

var red = color.New(color.FgRed).SprintFunc()
var white = color.New(color.FgWhite).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

func prompt(ch *world.Character) string {
	hp := red(fmt.Sprintf("<%d/%dhp>", ch.HitPoints, ch.MaxHitPoints))
	mana := white(fmt.Sprintf("<%d/%dm>", ch.Mana, ch.MaxMana))
	mv := yellow(fmt.Sprintf("<%d/%dmv>", ch.Movement, ch.MaxMovement))

	ch.Print(fmt.Sprintf("%s%s%s ", hp, mana, mv))

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text
}
