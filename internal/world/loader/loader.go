package loader

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/brianseitel/mudder/internal/world"
)

func Load() []*world.Zone {
	body, err := ioutil.ReadFile("areas/area.lst")
	if err != nil {
		panic(err)
	}

	zones := []*world.Zone{}
	for _, line := range strings.Split(string(body), "\n") {
		if line == "$" { // end of file
			break
		}
		zone := loadZone(line)
		zones = append(zones, zone)
	}

	return zones
}

func loadZone(areaName string) *world.Zone {
	data := loadFile(areaName)

	zone := &world.Zone{}
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
