package main

import (
	"action_for_reaction/internal/bot"
	"action_for_reaction/internal/processor"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("init: %v", err)
	}
}

func run() error {
	b, err := bot.New()
	if err != nil {
		return err
	}
	if err = processor.Process(b); err != nil {
		return err
	}
	return nil
}
