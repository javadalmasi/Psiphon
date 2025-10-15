# Psiphone Usage Examples

## Basic Usage
```bash
# Run with default settings (Germany, port 1080)
./psiphon-proxy

# Run with verbose logging
./psiphon-proxy -v

# Run with a specific country
./psiphon-proxy -c US

# Run with custom bind address
./psiphon-proxy -b 0.0.0.0:1080 -c CA

# Run with custom port
./psiphon-proxy -b 127.0.0.1:8080 -c DE
```

## Shuffle Mode (All Countries)
```bash
# Run in shuffle mode (all countries with load balancing)
./psiphon-proxy --shuffle

# Run shuffle mode with verbose logging
./psiphon-proxy --shuffle -v
```

## Docker Usage

### Build the container
```bash
docker build -t psiphon-proxy .
```

### Run the container
```bash
# Default settings (Germany)
docker run -p 1080:1080 psiphon-proxy

# With specific country
docker run -p 1080:1080 psiphon-proxy -c US

# With verbose logging
docker run -p 1080:1080 psiphon-proxy -v -c NL
```

## Testing the Proxy

You can test the SOCKS5 proxy once it's running:

### Using curl with SOCKS5 proxy
```bash
curl --socks5-hostname 127.0.0.1:1080 https://ipinfo.io/country
```

### Using wget with SOCKS5 proxy
```bash
wget --bind-address=127.0.0.1 --execute="http_proxy=socks5://127.0.0.1:1080" -qO- https://ipinfo.io/country
```

## Environment Variables for Docker

You can also use environment variables if you modify your Docker setup:

```bash
docker run -p 1080:1080 -e PSIPHON_COUNTRY=JP psiphon-proxy
```