package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := &http.Server {
		Addr: ":8080",
		Handler: routes(),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Channel to block until shutdown completes
    done := make(chan struct{})

    // Graceful shutdown goroutine
	go func() {
		<-quit // wait for SIGINT or SIGTERM
    log.Println("Shutting down...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    close(done) // signal to main that shutdown is done
	}()

	log.Println("Starting server on :8080")

	if err := server.ListenAndServe(); 
	   err != nil && 
	   err != http.ErrServerClosed {
		// If the error is not a server closed error, log it and exit
		log.Fatalf("Error starting server: %v", err)
		return
	}

	<-done // Wait until shutdown is complete
    log.Println("Server exited")
}