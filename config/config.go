package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Configuration struct {
	MaxRetries      int    `yaml:"maxRetries"`
	BaseDelay       int    `yaml:"baseDelay"`
	MaxDelay        int    `yaml:"maxDelay"`
	SecretKeyLength int    `yaml:"secretKeyLength"`
	SecretKey       string `yaml:"secretKey"`
}

var Config *Configuration

func Init() {
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic("Error reading YAML file")
	}

	// Parse the YAML content into a Configuration struct
	Config = new(Configuration)
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic("Error unmarshalling YAML")
	}

	secretKey, err := GenerateRandomSecretKey()
	if err != nil {
		panic("Error generating Secret Key")
	}
	Config.SecretKey = secretKey
}
