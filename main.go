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
	err = c.SetUser("gary")
	if err != nil {
		fmt.Println("Error in SetUser", err)
		os.Exit(2)
	}
	c, err = config.Read()
	if err != nil {
		fmt.Println("Error in 2nd Read", err)
		os.Exit(3)
	}
    fmt.Println(c)
}