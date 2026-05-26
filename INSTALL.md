# Installation Guide

## Requirements
- macOS 10.14+ OR Linux (Ubuntu 18.04+)
- Go 1.21+ (for building from source)

## Method 1: Pre-built Binary (Recommended)
Download for your OS:
- macOS: `fabric-control-plane-macos`
- Linux: `fabric-control-plane-linux`

## Method 2: Homebrew
```bash
brew tap ehsidawi/tap
brew install identity-uth
```

## Method 3: From Source
```bash
git clone https://github.com/ehsidawi/Identity-UTH.git
cd Identity-UTH/control-plane
go mod tidy
go build -o fabric-control-plane main.go
sudo mv fabric-control-plane /usr/local/bin/
```

## Verify Installation
```bash
fabric-control-plane --help
```
