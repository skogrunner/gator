package main

import (
	"fmt"
	"os"
	"github.com/skogrunner/gator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println("Error in Read", err)
		os.Exit(1)
	}
	st := State{State: &c}
    cm := Commands{CommandMap: make(map[string]func(*State, Command) error)}
    cm.Register("login", HandlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("no command specified")
		os.Exit(0)
	}
	err = cm.Run(st, Command{Name: os.Args[1], Args: os.Args[2:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}