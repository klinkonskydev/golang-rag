package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"rag-course/config"
	"rag-course/internal/app"
	"syscall"
)

func main() {
	// We need to:
	// - set up the app
	// - set up config
	// - set up an LLM client
	// - set up the Read-Eval-Print loop (REPL)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer stop()

	if err := app.Run(ctx, config.Load()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
