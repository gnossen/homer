package dm

import (
	"errors"
	"fmt"
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
}

type DM struct {
	cmds []Cmd
	Rand *rand.Rand
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
