package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Thresholds struct {
		CPU     float64 `yaml:"cpu"`
		Memory  float64 `yaml:"memory"`
		Disk    float64 `yaml:"disk"`
		Network float64 `yaml:"network"`
	} `yaml:"thresholds"`
	Alerts struct {
		Email struct {
			Enabled        bool   `yaml:"enabled"`
			SMTPServer     string `yaml:"smtp_server"`
			SMTPPort       int    `yaml:"smtp_port"`
			SenderEmail    string `yaml:"sender_email"`
			SenderPassword string `yaml:"sender_password"`
			RecipientEmail string `yaml:"recipient_email"`
		} `yaml:"email"`
	} `yaml:"alerts"`
}

func LoadConfig() (*Config, error) {
	configFile, err := os.Open("configs/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer configFile.Close()

	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
