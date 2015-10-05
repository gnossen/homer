package dm

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	LIST_ALL      = iota
	LIST_ACTIVE   = iota
	LIST_INACTIVE = iota
	LIST_CLASS    = iota
)

func ListParser(args []string) (int, error) {
	if len(args) == 0 {
		return LIST_ALL, nil
	}
	errStr := "list takes one argument. One of 'active', 'inactive', 'class', and 'all"
	if len(args) != 1 {
		return 0, errors.New(errStr)
	}
	switch args[0] {
	case "active":
		return LIST_ACTIVE, nil
	case "inactive":
		return LIST_INACTIVE, nil
	case "class":
		return LIST_CLASS, nil
	case "all":
		return LIST_ALL, nil
	default:
		return 0, errors.New(errStr)
	}
}

func ListHandler(args []string, dm *DM) (string, *CmdError) {
	listType, err := ListParser(args)
	if err != nil {
		return "", &CmdError{
			errType: PARSE_ERROR,
			errStr:  err.Error(),
		}
	}
	var buf bytes.Buffer
	if listType == LIST_ACTIVE || listType == LIST_ALL {
		for _, mob := range dm.Active {
			buf.WriteString(fmt.Sprintf("%s (%s) (A)\n", mob.Name, mob.Nick))
		}
	}
	if listType == LIST_INACTIVE || listType == LIST_ALL {
		for _, mob := range dm.Inactive {
			buf.WriteString(fmt.Sprintf("%s (%s) (I)\n", mob.Name, mob.Nick))
		}
	}
	if listType == LIST_CLASS || listType == LIST_ALL {
		for _, mob := range dm.Bestiary {
			buf.WriteString(fmt.Sprintf("%s (%s) (C)\n", mob.Name, mob.Nick))
		}
	}
	return strings.TrimRight(buf.String(), "\n"), nil
}
