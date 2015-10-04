package dm

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	PARSE_ERROR     = iota
	EXECUTION_ERROR = iota
)

type CmdError struct {
	errType int
	errStr  string
}

func (err *CmdError) Error() string {
	return err.errStr
}

type CmdHandler func(args []string, dm *DM) (res string, err *CmdError)

type Cmd struct {
	Name    string
	HelpStr string
	Handler CmdHandler
}

type DM struct {
	cmds []Cmd
	Rand *rand.Rand
}

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

func (dm *DM) Init() {
	dm.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	dm.RegisterCmd("dice", "dice [num-dice] [num-sides]", DiceHandler)
	dm.RegisterCmd("d", "d [num-dice] [num-sides]", DiceHandler)
	dm.RegisterCmd("seed", "seed [seed]", SeedHandler)
}

func (dm *DM) ParseCmd(cmdStr string) (string, error) {
	args := strings.Fields(cmdStr)
	if len(args) == 0 {
		return "", nil
	}
	for _, cmd := range dm.cmds {
		if args[0] == cmd.Name {
			res, err := cmd.Handler(args[1:], dm)
			if err != nil && err.errType == PARSE_ERROR {
				return fmt.Sprintf("Usage: %s", cmd.HelpStr), err
			} else if err != nil {
				return "", err
			}
			return res, nil
		}
	}
	return "", errors.New(fmt.Sprintf("%s: Unrecognized command.", args[0]))
}

func (dm *DM) RegisterCmd(cmdName string, help string, handler CmdHandler) {
	newCmd := Cmd{
		Name:    cmdName,
		HelpStr: help,
		Handler: handler,
	}
	dm.cmds = append(dm.cmds, newCmd)
}

func (dm *DM) Dice(numDice int, numSides int) int {
	sum := 0
	for i := 0; i < numDice; i++ {
		sum += dm.Rand.Intn(numSides)
	}
	return sum
}
