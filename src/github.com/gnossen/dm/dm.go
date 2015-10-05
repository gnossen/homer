package dm

import (
	"errors"
	"fmt"
	"github.com/gnossen/mob"
	"math/rand"
	"strings"
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
    Alias   bool
}

type DM struct {
	Cmds      []Cmd
	Rand      *rand.Rand
	Active    []mob.ActiveMob
	Inactive  []mob.Mob
	Bestiary  []mob.MobClass
	directory string
}

func (dm *DM) SetDirectory(directory string) {
	dm.directory = directory
}

func (dm *DM) ParseCmd(cmdStr string) (string, error) {
	args := strings.Fields(cmdStr)
	if len(args) == 0 {
		return "", nil
	}
	for _, cmd := range dm.Cmds {
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

func (dm *DM) RegisterCmd(cmdName string, help string, alias bool, handler CmdHandler) {
	newCmd := Cmd{
		Name:    cmdName,
		HelpStr: help,
		Handler: handler,
        Alias: alias,
	}
	dm.Cmds = append(dm.Cmds, newCmd)
}
