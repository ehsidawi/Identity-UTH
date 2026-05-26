# Quick Start - 2 Minutes

## Option 1: Download Pre-built Binary (Easiest)
```bash
curl -L https://github.com/ehsidawi/Identity-UTH/releases/download/v1.0.0/fabric-control-plane-macos -o fabric-control-plane
chmod +x fabric-control-plane
./fabric-control-plane --blast-radius global
```

## Option 2: Build from Source
```bash
git clone https://github.com/ehsidawi/Identity-UTH.git
cd Identity-UTH/control-plane
go build -o fabric-control-plane main.go
./fabric-control-plane --blast-radius global
```
