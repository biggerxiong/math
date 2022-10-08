package main

import (
	"fmt"
	"github.com/pkg/errors"
	"main/config"
	"main/reader"
)

func init() {
	InitConfig()
	InitModels()
}

func InitConfig() {
	err := config.InitConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("decode config file err:%v", err.Error()))
	}
}

func InitModels() {
	Edges, err = reader.ReadEdges(config.GetConfig().EdgePath)
	if err != nil {
		panic(errors.Wrap(err, "read Edges failed"))
	}

	Nodes, err = reader.ReadNodes(config.GetConfig().NodePath)
	if err != nil {
		panic(errors.Wrap(err, "read Nodes failed"))
	}

	Streets, err = reader.ReadStreets(config.GetConfig().StreetPath)
	if err != nil {
		panic(errors.Wrap(err, "read Streets failed"))
	}

	MidStreams, err = reader.ReadMidStreams(config.GetConfig().MidStreamPath)
	if err != nil {
		panic(errors.Wrap(err, "read MidStreams failed"))
	}

	UpStreams, err = reader.ReadUpStreams(config.GetConfig().UpStreamPath)
	if err != nil {
		panic(errors.Wrap(err, "read UpStreams failed"))
	}
}
