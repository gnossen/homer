package dm

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"
    "io/ioutil"
    "encoding/json"
)

func (dm *DM) DefaultInit() {
	dm.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	dm.RegisterCmd("dice", "dice [num-dice] [num-sides]", DiceHandler)
	dm.RegisterCmd("d", "dice [num-dice] [num-sides]", DiceHandler)
	dm.RegisterCmd("seed", "seed [seed]", SeedHandler)
    dm.RegisterCmd("list", "list [active|inactive|class|all]", ListHandler)
    dm.RegisterCmd("ls", "list [active|inactive|class|all]", ListHandler)
    dm.RegisterCmd("l", "list [active|inactive|class|all]", ListHandler)
    dir, err := os.Getwd()
    if err != nil {
		fmt.Printf("Cannot read current working directory.\n")
		panic("Cannot read current working directory.\n")
	}
    dm.directory = dir
	dm.LoadActive()
	dm.LoadInactive()
	dm.LoadBestiary()
}

func (dm *DM) Dice(numDice int, numSides int) int {
	sum := 0
	for i := 0; i < numDice; i++ {
		sum += dm.Rand.Intn(numSides)
	}
	return sum
}

func (dm *DM) LoadActive() {
	filename := path.Join(dm.directory, "active.json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return
	}
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        panic("Could not read 'active.json'.")
    }
    err = json.Unmarshal(file, &dm.Active)
    if err != nil {
        fmt.Printf("Error parsing 'active.json': %s\n", err.Error())
        os.Exit(1)
    }
}

func (dm *DM) LoadInactive() {

}

func (dm *DM) LoadBestiary() {

}
