package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sw4mpd0nkey/smoothbrain/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func serve(server *http.Server, message string, serverErrors chan<- error) {
	go func() {
		fmt.Println(message)
		serverErrors <- server.ListenAndServe()
	}()
}

func shutdownServer(server *http.Server) error {
	shutdownContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(shutdownContext)
}

func run(ctx context.Context) error {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	appConfig, err := config.Load(workingDirectory)
	if err != nil {
		return fmt.Errorf("load configuration: %w", err)
	}
	fmt.Println(appConfig)

	return nil
}
