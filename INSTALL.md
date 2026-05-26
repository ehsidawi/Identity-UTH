# Installation Guide

## Requirements

- macOS 10.14+ OR Linux (Ubuntu 18.04+)
- **NO** additional dependencies required

## Method 1: Pre-built Binary (Recommended)

### Download
```bash
# macOS (ARM64/Apple Silicon)
curl -L https://github.com/ehsidawi/Identity-UTH/releases/download/v1.0.0/fabric-control-plane-macos -o fabric

# Linux (AMD64)
curl -L https://github.com/ehsidawi/Identity-UTH/releases/download/v1.0.0/fabric-control-plane-linux -o fabric
```

### Install
```bash
chmod +x fabric
sudo mv fabric /usr/local/bin/
```

### Verify
```bash
fabric --help
```

## Method 2: Clone & Build from Source

### Requirements
- Go 1.21+

### Build
```bash
git clone https://github.com/ehsidawi/Identity-UTH.git
cd Identity-UTH/control-plane
go mod tidy
go build -o fabric-control-plane main.go
./fabric-control-plane --help
```

### Install
```bash
sudo mv fabric-control-plane /usr/local/bin/fabric
```

## Method 3: Docker

### Build
```bash
git clone https://github.com/ehsidawi/Identity-UTH.git
cd Identity-UTH
docker build -t identity-uth:latest .
```

### Run
```bash
docker run -it identity-uth:latest --blast-radius global
```

## Verify Installation

```bash
fabric --version  # Shows help if not implemented
fabric --help     # Shows all options
```

## Uninstall

```bash
rm /usr/local/bin/fabric  # macOS/Linux
docker rmi identity-uth:latest  # Docker
```

---

**Installation complete. See [QUICKSTART.md](QUICKSTART.md) for next steps.**
