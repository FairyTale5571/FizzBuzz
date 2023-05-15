package main

import (
	"github.com/fairytale5571/fizzBuzz/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	application := app.New()
	if err := application.Run(); err != nil {
		log.Fatalf("Error running app: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	_ = <-sigChan
	if err := application.Stop(); err != nil {
		log.Fatalf("Error stopping app: %v", err)
	}
	os.Exit(0)
}
