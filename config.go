package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	cfg *Config
)

type Config struct {
	Database struct {
		Host     string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
}

type Flag struct {
	configFile string
}

func parseFlags() *Flag {
	f := new(Flag)
	flag.StringVar(
		&f.configFile,
		"config",
		"alertman.yml",
		"Path to alertman configuration file",
	)
	flag.Parse()
	return f
}

func newConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	y := yaml.NewDecoder(file)
	config := &Config{}
	if err := y.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
