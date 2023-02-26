package main

import (
	"log"

	hclog "github.com/hashicorp/go-hclog"
)

func main() {
	defaultLogging()
	stdLogging()
	methodLogging()
}

func defaultLogging() {
	log.Printf("golang default logging")
}

func stdLogging() {
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "stdapp",
		Level: hclog.LevelFromString("DEBUG"),
	})
	log.SetOutput(appLogger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	log.SetPrefix("")
	log.SetFlags(0)

	log.Printf("[DEBUG] %d", 42)
}

func methodLogging() {
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "func",
		Level: hclog.LevelFromString("INFO"),
	})

	appLogger.Debug("this message will not print")
	appLogger.Info("this message will print")
	appLogger.Info("message with key-value", "key", 42)
}
