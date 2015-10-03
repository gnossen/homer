package main

import (
    "fmt"
    "github.com/gnossen/dm"
)

func main() {
    var dm dm.DM
    for {
        fmt.Printf("> ")
        var input string
        fmt.Scanln(&input)
        res, err := dm.ParseCmd(input)
        if err != nil {
            fmt.Printf("Parse error:\n%s\n", err)
        } else {
            fmt.Printf("%s\n", res)
        }
    }
}
