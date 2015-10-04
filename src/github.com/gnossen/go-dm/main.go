package main

import (
	"bufio"
	"fmt"
	"github.com/gnossen/dm"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Homer - Dungeon Master's Shell\n")
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
