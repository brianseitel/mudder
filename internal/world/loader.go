package world

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func New() *World {
	return &World{}
}

func Load() *World {
	body, err := ioutil.ReadFile("areas/area.lst")
	if err != nil {
		panic(err)
	}

	gameWorld := New()
	for _, line := range strings.Split(string(body), "\n") {
		if line == "$" { // end of file
			break
		}
		zone := loadZone(line)
		gameWorld.Zones = append(gameWorld.Zones, zone)
	}

	return gameWorld
}

func loadZone(areaName string) *Zone {
	data := loadFile(areaName)

	zone := &Zone{}
	zone.Area = loadArea(data)
	zone.Helps = loadHelps(data)
	zone.Mobiles = loadMobiles(data)
	zone.Objects = loadObjects(data)
	zone.Rooms = loadRooms(data)
	zone.Resets = loadResets(data)
	zone.Shops = loadShops(data)
	zone.Specials = loadSpecials(data)

	return zone
}

func loadFile(areaName string) string {
	f, err := os.Open(fmt.Sprintf("areas/%s", areaName))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(body)
}
