package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	version = "dev"
	commit  = "xyz"
	date    = "Sun Jun 22 17:16:23 EDT 2025"
)

func main() {
	var (
		showVersion = flag.Bool("version", false, "Show version information")
		configFile  = flag.String("config", "", "Configuration file path")
		verbose     = flag.Bool("verbose", false, "Enable verbose output")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("Example CLI Tool\n")
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Built: %s\n", date)
		return
	}

	if *verbose {
		log.SetOutput(os.Stdout)
		log.Println("Starting CLI application...")
	}

	// Simulate CLI functionality
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Example CLI Tool - Multi-binary demo")
		fmt.Println("Usage: cli [options] <command>")
		fmt.Println("\nCommands:")
		fmt.Println("  process    Process data")
		fmt.Println("  status     Show status")
		fmt.Println("  help       Show this help")
		return
	}

	command := args[0]
	switch command {
	case "process":
		handleProcess(*configFile, *verbose)
	case "status":
		handleStatus(*verbose)
	case "help":
		fmt.Println("Available commands: process, status, help")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func handleProcess(configFile string, verbose bool) {
	if verbose {
		log.Printf("Processing with config: %s", configFile)
	}
	
	fmt.Println("Processing data...")
	time.Sleep(500 * time.Millisecond) // Simulate work
	fmt.Println("✅ Processing complete")
}

func handleStatus(verbose bool) {
	if verbose {
		log.Println("Checking system status...")
	}
	
	fmt.Printf("CLI Status: Running\n")
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Uptime: %v\n", time.Now().Format("2006-01-02 15:04:05"))
}
