package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"main/config"
	"main/reader"

	nested "github.com/antonfisher/nested-logrus-formatter"
)

func init() {
	SetUpLog()
	InitConfig()
	InitModels()
}

func SetUpLog() {
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
}

func InitConfig() {
	err := config.InitConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("decode config file err:%v", err.Error()))
	}
	logrus.Infof("config: %v", config.GetConfig())
}

func InitModels() {
	var err error

	Edges, err = reader.ReadEdges(config.GetConfig().EdgePath)
	if err != nil {
		panic(errors.Wrap(err, "read Edges failed"))
	}
	logrus.Debugf("Edges: %v", Edges)

	Nodes, err = reader.ReadNodes(config.GetConfig().NodePath)
	if err != nil {
		panic(errors.Wrap(err, "read Nodes failed"))
	}
	logrus.Debugf("Nodes: %v", Nodes)

	Streets, err = reader.ReadStreets(config.GetConfig().StreetPath)
	if err != nil {
		panic(errors.Wrap(err, "read Streets failed"))
	}
	logrus.Debugf("Streets: %v", Streets)

	MidStreams, err = reader.ReadMidStreams(config.GetConfig().MidStreamPath)
	if err != nil {
		panic(errors.Wrap(err, "read MidStreams failed"))
	}
	logrus.Debugf("MidStreams: %v", MidStreams)

	UpStreams, err = reader.ReadUpStreams(config.GetConfig().UpStreamPath)
	if err != nil {
		panic(errors.Wrap(err, "read UpStreams failed"))
	}
	logrus.Debugf("UpStreams: %v", UpStreams)
}
