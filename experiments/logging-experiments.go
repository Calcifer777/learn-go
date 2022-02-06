package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
)

var sugarLogger *zap.SugaredLogger

func WithDevelopmentLogger() {
	jsonFile, _ := ioutil.ReadFile("log-config.json")
	var config zap.Config
	if err := json.Unmarshal(jsonFile, &config); err != nil {
		panic(err)
	}
	logger, _ := config.Build()
	sugarLogger = logger.Sugar()
}

func WithProductionLogger() {
	config := zap.NewProductionConfig()
	level := zap.NewAtomicLevelAt(zap.WarnLevel)
	config.Level = level
	logger, _ := config.Build()
	sugarLogger = logger.Sugar()
}

func main() {
	WithDevelopmentLogger()
	sugarLogger.Infof("Info message") // this will not be logged
	sugarLogger.Warnf("Warning message")
}
