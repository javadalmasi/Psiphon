// Package main provides functionality for running a Psiphon proxy server
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/netip"
	"path/filepath"

	"github.com/Psiphon-Labs/psiphon-tunnel-core/psiphon"
)

// RunPsiphonProxy starts the Psiphon proxy with the specified configuration
// and handles the lifecycle of the Psiphon connection
//
// Parameters:
// - ctx: Context for cancellation and timeout handling
// - l: Structured logger for application logging
// - bindAddrPort: Address and port where the SOCKS5 proxy will listen
// - country: Country code for the Psiphon egress location
//
// Returns:
// - error: An error if the proxy fails to start or run properly
func RunPsiphonProxy(ctx context.Context, l *slog.Logger, bindAddrPort netip.AddrPort, country string) error {
	// Get the flag emoji for the selected country
	countryFlag := getCountryFlag(country)
	
	// Log the start of the proxy with configuration details
	l.Info("🚀 Starting Psiphon proxy", "bind", bindAddrPort, "country", country, "flag", countryFlag)

	// Determine the listening interface based on the bind address
	// Use "any" if binding to a non-local address, otherwise use default
	host := ""
	if !netip.MustParsePrefix("127.0.0.0/8").Contains(bindAddrPort.Addr()) {
		// Binding to non-local address - allow connections from any interface
		host = "any"
	}

	// Set the connection timeout to 60 seconds
	timeout := 60
	
	// Define cache directory for storing Psiphon session data
	// In container environments, use temp directory
	cacheDir := "/tmp/psiphon-proxy"
	
	// Create Psiphon configuration with all required parameters
	config := &psiphon.Config{
		EgressRegion:                                 country,                 // Country to route traffic through
		ListenInterface:                              host,                    // Network interface to bind to
		LocalSocksProxyPort:                          int(bindAddrPort.Port()), // Port for SOCKS5 proxy
		DisableLocalHTTPProxy:                        true,                    // Disable HTTP proxy functionality
		PropagationChannelId:                         "FFFFFFFFFFFFFFFF",      // Channel ID for Psiphon
		RemoteServerListDownloadFilename:             "remote_server_list",    // Filename for server list
		RemoteServerListSignaturePublicKey:           "MIICIDANBgkqhkiG9w0BAQEFAAOCAg0AMIICCAKCAgEAt7Ls+/39r+T6zNW7GiVpJfzq/xvL9SBH5rIFnk0RXYEYavax3WS6HOD35eTAqn8AniOwiH+DOkvgSKF2caqk/y1dfq47Pdymtwzp9ikpB1C5OfAysXzBiwVJlCdajBKvBZDerV1cMvRzCKvKwRmvDmHgphQQ7WfXIGbRbmmk6opMBh3roE42KcotLFtqp0RRwLtcBRNtCdsrVsjiI1Lqz/lH+T61sGjSjQ3CHMuZYSQJZo/KrvzgQXpkaCTdbObxHqb6/+i1qaVOfEsvjoiyzTxJADvSytVtcTjijhPEV6XskJVHE1Zgl+7rATr/pDQkw6DPCNBS1+Y6fy7GstZALQXwEDN/qhQI9kWkHijT8ns+i1vGg00Mk/6J75arLhqcodWsdeG/M/moWgqQAnlZAGVtJI1OgeF5fsPpXu4kctOfuZlGjVZXQNW34aOzm8r8S0eVZitPlbhcPiR4gT/aSMz/wd8lZlzZYsje/Jr8u/YtlwjjreZrGRmG8KMOzukV3lLmMppXFMvl4bxv6YFEMiUTsOhbLTwFgh7KYNjodLj/LsqRVfwz31PgWQFTEPICV7GCvgVlPRxnofqKSjgTWI4mxDhBpVcATvaoBl1L/6WLbFvBsoAUBItWwctO2xalKxF5szhGm8lccoc5MZr8kfE0uxMgsxz4er68iCID+rsCAQM=", // Public key for server list verification
		RemoteServerListUrl:                          "https://s3.amazonaws.com//psiphon/web/mjr4-p23r-puwl/server_list_compressed", // URL for server list
		SponsorId:                                    "FFFFFFFFFFFFFFFF",      // Sponsor ID for Psiphon
		NetworkID:                                    "test",                  // Network identifier
		ClientPlatform:                               "Android_4.0.4_com.example.exampleClientLibraryApp", // Platform identifier
		AllowDefaultDNSResolverWithBindToDevice:      true,                    // Allow DNS resolution when binding to device
		EstablishTunnelTimeoutSeconds:                &timeout,                // Connection timeout in seconds
		DataRootDirectory:                            cacheDir,                // Directory for data storage
		MigrateDataStoreDirectory:                    cacheDir,                // Directory for data store migration
		MigrateObfuscatedServerListDownloadDirectory: cacheDir,                // Directory for server list migration
		MigrateRemoteServerListDownloadFilename:      filepath.Join(cacheDir, "server_list_compressed"), // Server list filename for migration
	}

	// Log attempt to establish Psiphon tunnel
	l.Info("🔄 Establishing Psiphon tunnel connection...", "country", country, "flag", countryFlag)
	
	// Commit the configuration to the Psiphon library
	if err := config.Commit(true); err != nil {
		return fmt.Errorf("config.Commit failed: %w", err)
	}

	// Open the Psiphon data store for session persistence
	if err := psiphon.OpenDataStore(config); err != nil {
		return fmt.Errorf("failed to open data store: %w", err)
	}

	// Import embedded server entries for initial connection
	if err := psiphon.ImportEmbeddedServerEntries(ctx, config, "", ""); err != nil {
		return fmt.Errorf("failed to import server entries: %w", err)
	}

	// Set up notice handling to provide status updates and logging
	psiphon.SetNoticeWriter(psiphon.NewNoticeReceiver(
		// Callback function to handle Psiphon notices and log them appropriately
		func(notice []byte) {
			var event map[string]interface{}
			if err := json.Unmarshal(notice, &event); err != nil {
				return
			}

			// Extract notice type from the received event
			noticeType := event["noticeType"].(string)
			switch noticeType {
			case "Connected":
				// Psiphon tunnel connection established successfully
				l.Info("✅ Psiphon tunnel connected successfully", "country", country, "flag", countryFlag)
			case "Connecting":
				// Psiphon is attempting to connect
				l.Info("⏳ Connecting to Psiphon network...", "country", country, "flag", countryFlag)
			case "Tunnels":
				// Check if any tunnels are available
				if count, ok := event["count"].(float64); ok && count > 0 {
					l.Info("🔗 Psiphon tunnel established", "country", country, "flag", countryFlag, "tunnel_count", int(count))
				}
			case "EstablishTunnelTimeout":
				// Connection attempt timed out
				l.Error("❌ Psiphon tunnel establishment timed out", "country", country, "flag", countryFlag)
			case "ActiveTunnel":
				// Active tunnel information
				l.Info("📡 Active Psiphon tunnel", "country", country, "flag", countryFlag)
			}
		}))

	// Create the Psiphon controller which manages the tunnel lifecycle
	controller, err := psiphon.NewController(config)
	if err != nil {
		return fmt.Errorf("psiphon.NewController failed: %w", err)
	}

	// Run the controller in a separate goroutine to allow non-blocking operation
	go func() {
		controller.Run(ctx)
	}()

	// Log successful startup
	l.Info("🎉 Psiphon proxy is running", "bind", bindAddrPort, "country", country, "flag", countryFlag)
	
	// Wait for context cancellation (signal) before shutting down
	<-ctx.Done()
	
	// Log shutdown and clean up resources
	l.Info("🛑 Psiphon proxy is shutting down...", "country", country, "flag", countryFlag)
	psiphon.CloseDataStore()
	
	return nil
}

// getCountryFlag returns the emoji flag for a given country code
func getCountryFlag(countryCode string) string {
	// Map country codes to their corresponding flag emojis
	flags := map[string]string{
		"AT": "🇦🇹", // Austria
		"AU": "🇦🇺", // Australia
		"BE": "🇧🇪", // Belgium
		"BG": "🇧🇬", // Bulgaria
		"CA": "🇨🇦", // Canada
		"CH": "🇨🇭", // Switzerland
		"CZ": "🇨🇿", // Czech Republic
		"DE": "🇩🇪", // Germany
		"DK": "🇩🇰", // Denmark
		"EE": "🇪🇪", // Estonia
		"ES": "🇪🇸", // Spain
		"FI": "🇫🇮", // Finland
		"FR": "🇫🇷", // France
		"GB": "🇬🇧", // United Kingdom
		"HR": "🇭🇷", // Croatia
		"HU": "🇭🇺", // Hungary
		"IE": "🇮🇪", // Ireland
		"IN": "🇮🇳", // India
		"IT": "🇮🇹", // Italy
		"JP": "🇯🇵", // Japan
		"LV": "🇱🇻", // Latvia
		"NL": "🇳🇱", // Netherlands
		"NO": "🇳🇴", // Norway
		"PL": "🇵🇱", // Poland
		"PT": "🇵🇹", // Portugal
		"RO": "🇷🇴", // Romania
		"RS": "🇷🇸", // Serbia
		"SE": "🇸🇪", // Sweden
		"SG": "🇸🇬", // Singapore
		"SK": "🇸🇰", // Slovakia
		"US": "🇺🇸", // United States
	}
	
	// Return the flag for the country, or a default globe emoji if not found
	if flag, exists := flags[countryCode]; exists {
		return flag
	}
	return "🌍" // Default globe emoji
}