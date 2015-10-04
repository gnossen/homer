package dm

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

func DiceParser(args []string) (int, int, string) {
	if len(args) != 2 {
		return 0, 0, "dice takes two arguments"
	}
	numDice, err1 := strconv.Atoi(args[0])
	diceSides, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil || numDice < 1 || diceSides < 1 {
		return 0, 0, "dice takes two positive integers."
	}
	return numDice, diceSides, ""
}

func DiceHandler(args []string, dm *DM) (string, *CmdError) {
	numDice, diceSides, errString := DiceParser(args)
	if errString != "" {
		return "", &CmdError{
			errType: PARSE_ERROR,
			errStr:  errString,
		}
	}
	res := dm.Dice(numDice, diceSides)
	return strconv.Itoa(res), nil
}

func SeedParser(args []string) (int64, error) {
	if len(args) != 1 {
		return 0, errors.New("seed takes one argument.")
	}
	seed, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return 0, errors.New("seed takes an integer.")
	}
	return seed, nil
}

func SeedHandler(args []string, dm *DM) (string, *CmdError) {
	seed, err := SeedParser(args)
	if err != nil {
		return "", &CmdError{
			errType: PARSE_ERROR,
			errStr:  err.Error(),
		}
	}
	dm.Rand = rand.New(rand.NewSource(seed))
	return fmt.Sprintf("RNG set to seed %s.", args[0]), nil
}
