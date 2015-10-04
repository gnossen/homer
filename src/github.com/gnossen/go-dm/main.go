package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "github.com/gnossen/dm"
)

func main() {
    fmt.Printf("go-dm Dungeon Master's Tool\n")
    dm := &dm.DM{}
    dm.Init()
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Printf("> ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Input error:\n%s\n", err)
            continue
        }
        input = strings.TrimRight(input, "\n")
        res, err := dm.ParseCmd(input)
        if err != nil {
            fmt.Printf("%s\n", err)
        }
        if res != "" {
            fmt.Printf("%s\n", res)
        }
    }
}
