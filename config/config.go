package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name     string `yaml:"name"`
	Server   Server
	Database Database
	Cache    Cache
}

type Server struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"portcognito"`
	Secret string `yaml:"string"`
}

type Database struct {
	User     string `yaml:"dbuser"`
	Port     string `yaml:"dbport"`
	Host     string `yaml:"dbhost"`
	Password string `yaml:"dbpassword"`
	Name     string `yaml:"dbname"`
}

type Cache struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

var config *Config

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	config = &conf
	return config, nil
}

func Get() *Config {
	return config
}
