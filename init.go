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
	InitConfig()
	SetUpLog()
	InitModels()
}

var strToLogLevel = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
	"trace": logrus.TraceLevel,
	"":      logrus.InfoLevel,
}

func SetUpLog() {
	logrus.SetLevel(strToLogLevel[config.GetConfig().LogLevel])
	if config.GetConfig().LogPath != "" {
		f, err := os.OpenFile(config.GetConfig().LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(f)
	} else {
		logrus.SetOutput(os.Stdout)
	}
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
	fmt.Printf("config: %v", config.GetConfig())
}

func InitModels() {
	var err error

	Edges, err = reader.ReadEdges(config.GetConfig().EdgePath)
	if err != nil {
		panic(errors.Wrap(err, "read Edges failed"))
	}
	logrus.Tracef("Edges: %v", Edges)

	Nodes, err = reader.ReadNodes(config.GetConfig().NodePath)
	if err != nil {
		panic(errors.Wrap(err, "read Nodes failed"))
	}
	logrus.Tracef("Nodes: %v", Nodes)

	Streets, err = reader.ReadStreets(config.GetConfig().StreetPath)
	if err != nil {
		panic(errors.Wrap(err, "read Streets failed"))
	}
	logrus.Tracef("Streets: %v", Streets)

	MidStreams, err = reader.ReadMidStreams(config.GetConfig().MidStreamPath)
	if err != nil {
		panic(errors.Wrap(err, "read MidStreams failed"))
	}
	logrus.Tracef("MidStreams: %v", MidStreams)

	UpStreams, err = reader.ReadUpStreams(config.GetConfig().UpStreamPath)
	if err != nil {
		panic(errors.Wrap(err, "read UpStreams failed"))
	}
	logrus.Tracef("UpStreams: %v", UpStreams)
}
