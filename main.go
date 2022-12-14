package main

import (
	"main/config"
	"main/model"
	v1 "main/v1"
	"main/writer"

	"github.com/sirupsen/logrus"
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

	// calc mid to street
	algo := v1.NewAlgo(&v1.Models{
		Edges:      Edges,
		Nodes:      Nodes,
		Streets:    Streets,
		UpStreams:  UpStreams,
		MidStreams: MidStreams})

	ans := algo.RunMidToStreet()

	logrus.Infof("write to file, path: %s", config.GetConfig().Path.MidToStreetPath)
	err := writer.WriteAnswerPath(config.GetConfig().MidToStreetPath, ans)
	if err != nil {
		logrus.Fatal(err)
	}

	ans2 := algo.RunUpToMid()
	logrus.Infof("write to file, path: %s", config.GetConfig().Path.UpToMidPath)
	err = writer.WriteAnswerPath(config.GetConfig().UpToMidPath, ans2)
	if err != nil {
		logrus.Fatal(err)
	}

	ans3 := algo.RunMidToStreetCars()
	logrus.Infof("write to file, path: %s", config.GetConfig().Path.MidToStreetCarPath)
	err = writer.WriteAnswerCar(config.GetConfig().MidToStreetCarPath, ans3)
	if err != nil {
		logrus.Fatal(err)
	}
}
