package main

import (
	"main/model"
	v1 "main/v1"
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
	algo := v1.NewAlgo(&v1.Models{
		Edges:      Edges,
		Nodes:      Nodes,
		Streets:    Streets,
		UpStreams:  UpStreams,
		MidStreams: MidStreams})

	algo.Run()
}
