package config

import (
	"os"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Url string `json:"db_url"`
	Name string `json:"current_user_name"`
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

