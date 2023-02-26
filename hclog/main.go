package main

import (
	"log"

	hclog "github.com/hashicorp/go-hclog"
)

func main() {
	stdLogging()
}

func stdLogging() {
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "my-app",
		Level: hclog.LevelFromString("DEBUG"),
	})
	log.SetOutput(appLogger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	log.SetPrefix("")
	log.SetFlags(0)

	log.Printf("[DEBUG] %d", 42)
}
