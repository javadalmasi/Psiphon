# Psiphone Wiki

## Table of Contents
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Supported Countries](#supported-countries)
- [Docker Deployment](#docker-deployment)
- [Troubleshooting](#troubleshooting)
- [FAQ](#faq)

## Getting Started

Psiphone is a lightweight containerized SOCKS5 proxy that connects to the Psiphon network. It allows you to route your traffic through different countries using the Psiphon censorship circumvention technology.

### Prerequisites
- Docker (for containerized deployment)
- Or Go 1.21+ (for building from source)

### Quick Start
```bash
# Using Docker (recommended)
docker run -p 1080:1080 ghcr.io/javadalmasi/psiphone:latest

# The proxy will be available at 127.0.0.1:1080
```

## Configuration

### Command Line Options
- `-v`: Enable verbose logging
- `-b string`: SOCKS bind address (default "127.0.0.1:1080")
- `-c string`: Psiphon country code (default "DE")
- `--shuffle`: Enable shuffle mode (run all countries with load balancing)
- `-config string`: Path to JSON config file
- `-h`: Show help

### Shuffle Mode
In shuffle mode, the application automatically runs Psiphon proxies on all available countries simultaneously. Each country gets its own port, and requests are load-balanced across all proxies. Each proxy is automatically refreshed every 15 minutes, but with staggered timing to ensure continuous service availability.

### Supported Countries
See [Countries](countries.md) for the full list of supported countries and their codes.

## Docker Deployment

### Basic Docker Run
```bash
docker run -p 1080:1080 psiphone
```

### With Specific Country
```bash
docker run -p 1080:1080 psiphone -c US
```

### In Docker Compose
```yaml
version: '3.8'
services:
  psiphone:
    image: ghcr.io/javadalmasi/psiphone:latest
    ports:
      - "1080:1080"
    command: ["-c", "NL"]
    restart: unless-stopped
```

## Troubleshooting

### Connection Issues
- Make sure your firewall allows outbound connections
- Some countries may block Psiphon connections
- Try different countries to see which works best for your location

### Logging
- Use the `-v` flag for verbose logging to diagnose issues
- Check container logs with `docker logs <container-name>`

## FAQ

### Q: How is this different from the original Warp Plus?
A: This version focuses exclusively on Psiphon functionality with a simplified configuration and container-first design. It removes all Cloudflare Warp dependencies.

### Q: Which countries are supported?
A: Over 25 countries are supported. See the countries list for full details.

### Q: Is this safe to use?
A: Yes, it uses the official Psiphon library for secure connections.

### Q: Why does connection establishment take time?
A: Psiphon needs to find available servers in the selected country and establish a secure connection, which can take up to a minute.

### Q: What is shuffle mode?
A: Shuffle mode automatically connects to all available countries simultaneously with load balancing. Each country runs on its own port, and requests are distributed across them. Each connection is refreshed every 15 minutes with staggered timing to prevent service downtime.