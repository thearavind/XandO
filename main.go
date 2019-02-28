package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/thearavind/XandO/game"
)

func SignalContext(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Printf("listening for shutdown signal")
		<-sigs
		fmt.Printf("shutdown signal received")
		signal.Stop(sigs)
		close(sigs)
		cancel()
	}()

	return ctx
}

func main() {

	ctx := SignalContext(context.Background())
	var err error

	err = game.Server().Run(ctx)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
