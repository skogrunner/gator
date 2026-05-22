package main

import (
	"fmt"
	"github.com/skogrunner/gator/internal/config"
)

func main() {
	c := config.Read()
	c.SetUser("gary")
    fmt.Println(config.Read())
}