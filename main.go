package main

import (
	"fmt"
	"main/config"
	"main/model"
)

var configPath = "config/config.toml"

var (
	Edges   []*model.Edge
	Nodes   []*model.Node
	Streets []*model.Street

	MidStreams []*model.MidStream
	UpStreams  []*model.UpStream
)

func main() {
	fmt.Println(config.GetConfig())
	fmt.Println(UpStreams)
}
