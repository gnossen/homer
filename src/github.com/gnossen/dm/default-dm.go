package dm

import (
	"math/rand"
	"time"
)

func (dm *DM) DefaultInit() {
	dm.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	dm.RegisterCmd("dice", "dice [num-dice] [num-sides]", DiceHandler)
	dm.RegisterCmd("d", "d [num-dice] [num-sides]", DiceHandler)
	dm.RegisterCmd("seed", "seed [seed]", SeedHandler)
}

func (dm *DM) Dice(numDice int, numSides int) int {
	sum := 0
	for i := 0; i < numDice; i++ {
		sum += dm.Rand.Intn(numSides)
	}
	return sum
}
