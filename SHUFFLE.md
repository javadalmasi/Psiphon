# Shuffle Mode Documentation for Psiphone

For more information about the project, visit the GitHub repository: https://github.com/javadalmasi/Psiphon

## Overview
Shuffle mode is a special operating mode that runs Psiphon proxies simultaneously across all available countries. It provides load balancing between different geographic locations and automatic refresh of connections to maintain stability.

## How It Works
1. **Multi-Country Operation**: Starts Psiphon proxies for all 31 supported countries simultaneously
2. **Port Assignment**: Each country gets a unique port starting from 11000 (e.g., AT=11000, AU=11001, etc.)
3. **Load Balancing**: Uses round-robin algorithm to distribute requests across all active proxies
4. **Auto-Refresh**: Automatically refreshes each country's connection every 15 minutes
5. **Staggered Refresh**: Refresh intervals are staggered to prevent all proxies from refreshing at once, ensuring continuous service availability

## Configuration
In shuffle mode:
- All 31 countries run simultaneously
- Proxies are accessible on ports 11000 to 11030
- Final load-balanced service is available on port 1080
- Each proxy refreshes every 15 minutes with different timing

## Command Line Usage
```bash
# Enable shuffle mode
./psiphon-proxy --shuffle

# Enable shuffle mode with verbose logging
./psiphon-proxy --shuffle -v

# Using Docker
docker run -p 1080:1080 psiphon-proxy --shuffle
```

## Docker Usage
```yaml
version: '3.8'
services:
  psiphon-proxy:
    image: ghcr.io/javadalmasi/psiphon-proxy:latest
    ports:
      - "1080:1080"
    command: ["--shuffle"]
    restart: unless-stopped
```

## Refresh Mechanism
- Initial refresh time varies by country with staggered delays
- After initial connection, each proxy refreshes every 15 minutes
- The staggered timing ensures that not all connections refresh simultaneously
- If one country's proxy becomes unavailable, traffic is automatically routed to other active proxies

## Performance Considerations
- Higher resource usage due to multiple concurrent Psiphon connections
- Potential latency variance depending on which country handles the request
- Increased connection establishment time during startup as all countries connect

## Troubleshooting
- Monitor logs for the status of individual country proxies
- Check individual proxy ports (11000-11030) if you need to test specific countries
- Ensure sufficient system resources for running multiple concurrent proxies