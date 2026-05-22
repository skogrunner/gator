package main

import (
	"fmt"
	"github.com/skogrunner/gator/internal/config"
)

func main() {
	config := Read()
	config.SetUser("gary")
    fmt.Println(Read())
}