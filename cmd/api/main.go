package main

import (
	"flag"
	"github.com/mattcarabine/support-services/internal/server"
	"go.uber.org/zap"
	"log"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func initLogging() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	logger.Info("initialized zap logger")
	return logger
}

func main() {
	logger := initLogging()
	zap.ReplaceGlobals(logger)
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Fatalf("unable to flush logger: %s", err.Error())
		}
	}()

	flag.Parse()
	logger.Fatal("failed to start server", zap.String("address", *addr), zap.Error(server.SetupServer(*addr)))
}
