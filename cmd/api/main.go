package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/app"
	"github.com/joho/godotenv"
)



func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — using system env")
	}

	application, err := app.New()

	if err != nil {
		log.Fatal(err)

	}

	log.Printf("✅ Listening on %s", application.Address())

	errCh := make(chan error, 1)
	go func() {
		errCh <- application.Start()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	select {
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	case sig := <-sigCh:
		log.Printf("shutting down after signal: %s", sig)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := application.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}
	}

}
