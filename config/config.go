package main

import "github.com/BurntSushi/toml"

type Path struct {
	EdgePath   string `toml:"edge_path"`
	NodePath   string `toml:"node_path"`
	StreetPath string `toml:"street_path"`

	MidStreamPath string `toml:"mid_stream_path"`
	UpStreamPath  string `toml:"up_stream_path"`
}

type Config struct {
	Path
	MinDis string
}

var config Config

func GetConfig() *Config {
	return &config
}

func InitConfig(configPath string) error {
	_, err := toml.DecodeFile(configPath, &config)
	return err
}
