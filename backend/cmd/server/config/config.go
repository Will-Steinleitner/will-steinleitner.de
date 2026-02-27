package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: "8090",
	}
}

func (c *Config) Load(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Println("Config file not found, using defaults:", err)
		return
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Println("Invalid config YAML, using defaults:", err)
		return
	}

	return
}
