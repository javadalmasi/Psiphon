package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/netip"
	"sync"
	"sync/atomic"
	"time"
)

// ShuffleProxyManager manages multiple Psiphon proxies with different countries
// and balances requests across them
type ShuffleProxyManager struct {
	ctx       context.Context
	logger    *slog.Logger
	countries []string
	ports     map[string]int // country -> port mapping
	proxyList []string       // list of proxy addresses in order
	current   uint64         // round-robin counter
	listener  net.Listener
	cancel    context.CancelFunc
}

// NewShuffleProxyManager creates a new shuffle proxy manager
func NewShuffleProxyManager(ctx context.Context, logger *slog.Logger) *ShuffleProxyManager {
	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(ctx)
	
	// Use all supported countries
	countries := getAllCountries()
	ports := make(map[string]int)
	var proxyList []string
	
	// Assign ports starting from 11000 to avoid conflicts
	for i, country := range countries {
		port := 11000 + i
		ports[country] = port
		proxyList = append(proxyList, fmt.Sprintf("127.0.0.1:%d", port))
	}
	
	return &ShuffleProxyManager{
		ctx:       ctx,
		logger:    logger,
		countries: countries,
		ports:     ports,
		proxyList: proxyList,
		cancel:    cancel,
	}
}

// getAllCountries returns all supported Psiphon countries
func getAllCountries() []string {
	return []string{
		"AT", "AU", "BE", "BG", "CA", "CH", "CZ", "DE", "DK", "EE",
		"ES", "FI", "FR", "GB", "HR", "HU", "IE", "IN", "IT", "JP",
		"LV", "NL", "NO", "PL", "PT", "RO", "RS", "SE", "SG", "SK", "US",
	}
}

// Start starts all Psiphon proxies and the load balancer server
func (spm *ShuffleProxyManager) Start() error {
	spm.logger.Info("🔄 Starting shuffle mode", "countries", len(spm.countries), "total_ports", len(spm.ports))

	// Start Psiphon proxies for each country with staggered refresh
	var wg sync.WaitGroup
	for i, country := range spm.countries {
		wg.Add(1)
		
		// Add delay for staggered refresh to prevent all proxies from refreshing at the same time
		refreshDelay := time.Duration(i) * (15 * time.Minute / time.Duration(len(spm.countries)))
		
		go func(country string, port int, delay time.Duration) {
			defer wg.Done()
			spm.startPsiphonProxy(country, port, delay)
		}(country, spm.ports[country], refreshDelay)
	}
	
	// Wait a bit for all proxies to start before starting the load balancer server
	time.Sleep(2 * time.Second)
	
	// Start the load balancer server on port 1080
	listener, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		return fmt.Errorf("failed to create listener on port 1080: %w", err)
	}
	spm.listener = listener
	
	spm.logger.Info("🎯 Shuffle mode load balancer started", "listen_port", 1080, "proxies_count", len(spm.countries))
	
	// Handle incoming connections
	go func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				select {
				case <-spm.ctx.Done():
					return
				default:
					spm.logger.Error("failed to accept connection", "error", err)
					continue
				}
			}
			
			go spm.handleConnection(conn)
		}
	}()
	
	// Wait for all Psiphon proxies to finish (they run indefinitely until context is cancelled)
	wg.Wait()
	
	return nil
}

// handleConnection handles a single client connection, forwarding it to one of the available proxies
func (spm *ShuffleProxyManager) handleConnection(clientConn net.Conn) {
	defer clientConn.Close()
	
	// Select a proxy in round-robin fashion
	proxyIndex := atomic.AddUint64(&spm.current, 1) % uint64(len(spm.proxyList))
	proxyAddr := spm.proxyList[proxyIndex]
	
	spm.logger.Debug("Forwarding connection", "client", clientConn.RemoteAddr(), "proxy_index", proxyIndex, "proxy_addr", proxyAddr)
	
	// Connect to selected proxy
	proxyConn, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		spm.logger.Error("failed to connect to proxy", "proxy", proxyAddr, "error", err)
		return
	}
	defer proxyConn.Close()
	
	// Start bidirectional copy between client and proxy
	go func() {
		_, _ = io.Copy(clientConn, proxyConn)
		clientConn.Close()
		proxyConn.Close()
	}()
	
	// Copy from client to proxy
	_, _ = io.Copy(proxyConn, clientConn)
}

// startPsiphonProxy starts a Psiphon proxy for a specific country with auto-refresh
func (spm *ShuffleProxyManager) startPsiphonProxy(country string, port int, refreshDelay time.Duration) {
	// Start the proxy with refresh mechanism
	go spm.runRefreshableProxy(country, port, refreshDelay)
}

// runRefreshableProxy runs a Psiphon proxy with auto-refresh every 15 minutes
func (spm *ShuffleProxyManager) runRefreshableProxy(country string, port int, refreshDelay time.Duration) {
	// Calculate the initial refresh time with the delay
	initialRefresh := 15*time.Minute + refreshDelay
	
	// Initial wait before starting
	time.Sleep(refreshDelay)
	
	// Start first proxy
	go spm.runSingleProxy(country, port)
	
	spm.logger.Info("🔄 Auto-refresh scheduler started", "country", country, "port", port, "refresh_delay", refreshDelay)
	
	// Set up ticker for refresh, starting after initial delay
	// Also start a regular 15-minute refresh cycle after the initial one
	go func() {
		// Wait for the initial period to complete
		time.Sleep(initialRefresh - refreshDelay)
		
		// Then start regular refresh cycle with the same staggered delay
		regularTicker := time.NewTicker(15 * time.Minute)
		defer regularTicker.Stop()
		
		for {
			select {
			case <-regularTicker.C:
				spm.logger.Info("🔄 Refreshing proxy", "country", country, "port", port)
				// Create new proxy to replace the old one
				go spm.runSingleProxy(country, port)
			case <-spm.ctx.Done():
				return
			}
		}
	}()
	
	// Wait for context to be cancelled
	<-spm.ctx.Done()
}

// runSingleProxy runs a single Psiphon proxy instance
func (spm *ShuffleProxyManager) runSingleProxy(country string, port int) {
	// Create a context for this specific proxy instance
	proxyCtx, cancel := context.WithCancel(spm.ctx)
	defer cancel()
	
	// Create bind address for this proxy
	bindAddrPort, err := netip.ParseAddrPort(fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		spm.logger.Error("failed to parse bind address", "port", port, "error", err)
		return
	}
	
	countryFlag := getCountryFlag(country)
	spm.logger.Info("🚀 Starting Psiphon proxy", "country", country, "flag", countryFlag, "port", port)
	
	// Run the Psiphon proxy
	if err := RunPsiphonProxy(proxyCtx, spm.logger.With("country", country, "port", port), bindAddrPort, country); err != nil {
		spm.logger.Error("error running Psiphon proxy", "country", country, "port", port, "error", err)
	} else {
		spm.logger.Info("🛑 Psiphon proxy stopped", "country", country, "port", port)
	}
}

// Stop stops all proxies and the load balancer
func (spm *ShuffleProxyManager) Stop() {
	spm.logger.Info("🛑 Shutting down shuffle mode")
	if spm.listener != nil {
		spm.listener.Close()
	}
	spm.cancel()
}

// getAllCountries returns all supported countries
func (spm *ShuffleProxyManager) getAllCountries() []string {
	return spm.countries
}