package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	backend "will-steinleitner.de"
	"will-steinleitner.de/internal/server"
)

func run(
	ctx context.Context,
	app *backend.Application,
) error {
	log.Println("Starting server at port", app.GetConfig().Port)
	handler := server.RegisterRoutes(
		app.GetHome(),
	)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(app.GetConfig().Host, app.GetConfig().Port),
		Handler: handler,
	}

	done := make(chan os.Signal, 1)
	go func() {
		<-ctx.Done()
		log.Println(ctx.Err())

		log.Println("Shutting down server...")

		// don't use ctx as parent contex here: it is already canceled after <-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		httpServer.SetKeepAlivesEnabled(false)
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("Could not gracefully shutdown the server: %v\n", err)
		}

		close(done)
	}()

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("ListenAndServe error: %v\n", err)
	}

	<-done
	log.Println("Server stopped")

	return nil
}

func main() {
	app := backend.NewApplication()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	if err := run(ctx, app); err != nil {
		os.Exit(1)
	}
	log.Println("Main Context", ctx.Err())
}
