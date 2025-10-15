// Package main provides a Psiphon proxy implementation that creates a SOCKS5 proxy
// server connecting through the Psiphon network to bypass internet censorship.
// This application is designed to run in containerized environments.
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/netip"
	"os"
	"os/signal"
	"syscall"
)

// appName represents the name of this application for logging and display purposes
const appName = "Psiphon-Proxy"

// main is the entry point of the Psiphon proxy application
// It handles command-line argument parsing, configuration loading,
// logger initialization, and starts the proxy server
func main() {
	// Create a context that will be canceled when receiving termination signals (SIGINT, SIGTERM)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Define command-line flags
	verbose := flag.Bool("v", false, "enable verbose logging")
	bind := flag.String("b", "127.0.0.1:1080", "SOCKS bind address")
	country := flag.String("c", "DE", "psiphon country code")
	_ = flag.String("config", "", "path to JSON config file") // Config file not currently used
	help := flag.Bool("h", false, "show help")
	
	// Parse command-line flags
	flag.Parse()

	// Show help if requested
	if *help {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Initialize structured logger with appropriate log level based on verbose flag
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	if *verbose {
		// Enable debug logging if verbose flag is set
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	// Validate the bind address format (must be a valid IP:port combination)
	bindAddrPort, err := netip.ParseAddrPort(*bind)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: invalid bind address: %v\n", err)
		os.Exit(1)
	}

	// Start the Psiphon proxy server with the parsed configuration
	if err := RunPsiphonProxy(ctx, l, bindAddrPort, *country); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Wait for context cancellation (signal) before shutting down
	<-ctx.Done()
}