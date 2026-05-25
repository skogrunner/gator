package config

import (
	"fmt"
	"os"
	"errors"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Url string `json:"db_url"`
	Name string `json:"current_user_name"`
}

type state struct {
	State *Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	CommandMap map[string]func(*state, command) error
}

func (c *commands) Run(s *state, cmd command) error {
	val, ok := c.CommandMap[cmd.Name]
	if !ok {
		return errors.New("Command does not exist")
	}
	err := val(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.CommandMap[name] = f
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	content, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}
	var c Config
	err = json.Unmarshal(content, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

func (c Config) SetUser(name string) error {
	c.Name = name
	return write(c)
}

func write(cfg Config) error {
	fileContent, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, fileContent, 0666)
}

func getConfigFilePath() (string,error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return hd + "/.gatorconfig.json", nil
}

func HandlerLogin(s *state, cmd command) error {
    if len(cmd.Args) != 1 {
		return errors.New("Invalid arguments. Usage: LOGIN <username>")
	}
	err := s.State.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set to", cmd.Args[0])
	return nil
}