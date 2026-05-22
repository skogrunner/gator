package config

const configFileName = ".gatorconfig.json"

type Config struct {
	db_url string,
	current_user_name string
}

func Read() Config {
	return Config{}
}

func (c config) SetUser(name string) {
	pass
}

func getConfigFilePath() (string,error) {
	return "", nil
}

func write(cfg Config) error {
	return nil
}

