package dm

import (
    "io"
    "os"
    "fmt"
    "strconv"
    "math/rand"
    "strings"
    "errors"
    "time"
)

const (
    PARSE_ERROR = iota
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

type DMController interface {
    io.ReadWriter
    ParseCmd(string) (string, error)
    RegisterCmd(string, string, CmdHandler)
}

func ToFile(dm DMController, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = io.Copy(dm, file)
    return err
}

func FromFile(dm DMController, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = io.Copy(file, dm)
    return err
}

type DM struct {
    cmds []Cmd
    Rand *rand.Rand
}

func DiceParser(args []string) (int, int, string) {
    if(len(args) != 2) {
        return 0, 0, "dice takes two arguments"
    }
    numDice, err1 := strconv.Atoi(args[0])
    diceSides, err2 := strconv.Atoi(args[1])
    if err1 != nil || err2 != nil || numDice < 1 || diceSides < 1{
        return 0, 0, "dice takes two positive integers."
    }
    return numDice, diceSides, ""
}

func DiceHandler(args []string, dm *DM) (string, *CmdError) {
    numDice, diceSides, errString := DiceParser(args)
    if errString != "" {
        return "", &CmdError {
            errType: PARSE_ERROR,
            errStr: errString,
        }
    }
    sum := 0
    for i := 0; i < numDice; i++ {
        sum += dm.Rand.Intn(diceSides)
    }
    return strconv.Itoa(sum), nil
}

func SeedParser(args []string) (int64, error) {
    if len(args) != 1 {
        return 0, errors.New("setseed takes one argument.")
    }
    seed, err := strconv.ParseInt(args[0], 0, 64)
    if err != nil {
        return 0, errors.New("setseed takes an integer.")
    }
    return seed, nil
}

func SeedHandler(args []string, dm *DM) (string, *CmdError) {
    seed, err := SeedParser(args)
    if err != nil {
        return "", &CmdError {
            errType: PARSE_ERROR,
            errStr: err.Error(),
        }
    }
    dm.Rand = rand.New(rand.NewSource(seed))
    return fmt.Sprintf("RNG set to seed %s.", args[0]), nil
}

func (dm *DM) Init() {
    dm.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    dm.RegisterCmd("dice", "dice [num-dice] [num-sides]", DiceHandler)
    dm.RegisterCmd("setseed", "dice [seed]", SeedHandler)
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
        Name: cmdName,
        HelpStr: help,
        Handler: handler,
    }
    dm.cmds = append(dm.cmds, newCmd)
}

func (dm *DM) Read(p []byte) (int, error) {
    return 0, nil
}

func (dm *DM) Write(p []byte) (int, error) {
    return 0, nil
}
