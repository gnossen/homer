package dm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"time"
)

func (dm *DM) DefaultInit() {
	dm.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	dm.RegisterCmd("dice", "dice [num-dice] [num-sides]", false, DiceHandler)
	dm.RegisterCmd("d", "dice [num-dice] [num-sides]", true, DiceHandler)
	dm.RegisterCmd("seed", "seed [seed]", false, SeedHandler)
	dm.RegisterCmd("list", "list [active|inactive|class|all]", false, ListHandler)
	dm.RegisterCmd("ls", "list [active|inactive|class|all]", true, ListHandler)
	dm.RegisterCmd("l", "list [active|inactive|class|all]", true, ListHandler)
	dm.RegisterCmd("help", "help [cmd]", false, HelpHandler)
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

func (dm *DM) Dice(numDice int, numSides int) ([]int, int) {
	rolls := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		rolls[i] = dm.Rand.Intn(numSides)
	}
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	return rolls, sum
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
	fmt.Printf("Loaded %d active mobs.\n", len(dm.Active))
}

func (dm *DM) LoadInactive() {
	filename := path.Join(dm.directory, "inactive.json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Could not read 'inactive.json'.")
	}
	err = json.Unmarshal(file, &dm.Inactive)
	if err != nil {
		fmt.Printf("Error parsing 'inactive.json': %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Loaded %d inactive mobs.\n", len(dm.Inactive))
}

func (dm *DM) LoadBestiary() {
	filename := path.Join(dm.directory, "class.json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Could not read 'class.json'.")
	}
	err = json.Unmarshal(file, &dm.Bestiary)
	if err != nil {
		fmt.Printf("Error parsing 'class.json': %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Loaded %d mob classes.\n", len(dm.Bestiary))
}
