package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"go.uber.org/fx"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}

func NewParser(logger *log.Logger, toParse string, splitter string) ([]string, error) {
	logger.Print("Executing NewHandler.")
	// TODO: add validators
	res := strings.Split(toParse, splitter)
	logger.Print(res)
	return res, nil
}

func Register() {
	NewParser(NewLogger(), "test string", " ")
}

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewParser,
		),
		fx.Invoke(Register),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
