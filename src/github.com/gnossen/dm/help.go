package dm

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func HelpParser(args []string, dm *DM) (int, error) {
	if len(args) == 0 {
		return -1, nil
	}
	if len(args) != 1 {
		err := errors.New("help takes a single argument - the name of a command.")
		return -1, err
	}
	for i, cmd := range dm.Cmds {
		if cmd.Name == args[0] {
			return i, nil
		}
	}
	err := errors.New(fmt.Sprintf("%s: Unrecognized command.", args[0]))
	return -1, err
}

func HelpHandler(args []string, dm *DM) (string, *CmdError) {
	cmdIndex, err := HelpParser(args, dm)
	if err != nil {
		return "", &CmdError{
			errType: PARSE_ERROR,
			errStr:  err.Error(),
		}
	}
	if cmdIndex == -1 {
		var buf bytes.Buffer
		for _, cmd := range dm.Cmds {
			if cmd.Alias {
				continue
			}
			buf.WriteString(cmd.Name)
			buf.WriteString("\n")
		}
		return strings.TrimRight(buf.String(), "\n"), nil
	}
	return fmt.Sprintf("Usage: %s", dm.Cmds[cmdIndex].HelpStr), nil
}
