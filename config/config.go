package config

import "github.com/BurntSushi/toml"

type Path struct {
	EdgePath   string `toml:"edge_path"`
	NodePath   string `toml:"node_path"`
	StreetPath string `toml:"street_path"`

	MidStreamPath string `toml:"mid_stream_path"`
	UpStreamPath  string `toml:"up_stream_path"`

	MidToStreetPath string `toml:"mid_to_street_path"`
	UpToMidPath     string `toml:"up_to_mid_path"`
}

type Config struct {
	Path
	MinDis         float64 `toml:"min_dis"`
	NearCount      int     `toml:"near_count"`
	FoodsPerPerson string  `toml:"foods_per_person"`
	MaxDisMul      float64 `toml:"max_dis_mul"`

	LogLevel string `toml:"log_level"`
	LogPath  string `toml:"log_path"`
}

var config Config

func GetConfig() *Config {
	return &config
}

func InitConfig(configPath string) error {
	_, err := toml.DecodeFile(configPath, &config)
	return err
}
