package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	var (
		showVersion = flag.Bool("version", false, "Show version information")
		port        = flag.String("port", "8080", "Server port")
		host        = flag.String("host", "localhost", "Server host")
		logLevel    = flag.String("log-level", "info", "Log level (debug, info, warn, error)")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("Example Server\n")
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Built: %s\n", date)
		return
	}

	// Setup logging based on level
	setupLogging(*logLevel)

	// Create HTTP server
	mux := http.NewServeMux()
	
	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "healthy",
			"version": "%s",
			"commit": "%s",
			"timestamp": "%s"
		}`, version, commit, time.Now().UTC().Format(time.RFC3339))
	})

	// API endpoint
	mux.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"service": "example-server",
			"version": "%s",
			"build_info": {
				"commit": "%s",
				"date": "%s"
			},
			"runtime": {
				"host": "%s",
				"port": "%s"
			}
		}`, version, commit, date, *host, *port)
	})

	// Root endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, `Example Server - Multi-binary demo
Version: %s
Endpoints:
- GET /health - Health check
- GET /api/info - Service information
`, version)
	})

	// Server configuration
	addr := fmt.Sprintf("%s:%s", *host, *port)
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}

func setupLogging(level string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	switch level {
	case "debug":
		log.SetOutput(os.Stdout)
		log.Println("Debug logging enabled")
	case "info":
		log.SetOutput(os.Stdout)
	case "warn", "error":
		log.SetOutput(os.Stderr)
	default:
		log.SetOutput(os.Stdout)
		log.Printf("Unknown log level '%s', defaulting to info", level)
	}
}
