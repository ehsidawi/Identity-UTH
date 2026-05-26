# Quick Start Guide - 2 Minutes

## Installation (30 seconds)

### macOS
```bash
curl -L https://github.com/ehsidawi/Identity-UTH/releases/download/v1.0.0/fabric-control-plane-macos -o fabric
chmod +x fabric
./fabric --help
```

### Linux
```bash
curl -L https://github.com/ehsidawi/Identity-UTH/releases/download/v1.0.0/fabric-control-plane-linux -o fabric
chmod +x fabric
./fabric --help
```

## First Kill Switch (1 minute)

### Global Emergency Lockdown
```bash
./fabric --blast-radius global \
  --reason "Security incident detected" \
  --initiated-by "security@company.com"
```

### Expected Output
```json
{
  "id": "kill-switch-...",
  "timestamp": "2026-05-26T...",
  "blast_radius": "global",
  "reason": "Security incident detected",
  "initiated_by": "security@company.com",
  "status": "completed",
  "actions": [
    "✓ Revoking Okta tokens",
    "✓ Revoking Microsoft tokens",
    "✓ Terminating PAM sessions",
    "✓ Rotating secrets",
    "✓ Disabling API keys",
    "✓ Quarantining workloads",
    "✓ Pushing deny-all policies"
  ]
}
```

## Common Commands

### Tenant Isolation
```bash
./fabric --blast-radius tenant \
  --reason "Suspicious activity in acme-corp" \
  --initiated-by "security@company.com"
```

### Single User Revocation
```bash
./fabric --blast-radius user \
  --reason "Account compromise" \
  --initiated-by "security@company.com"
```

### View All Events
```bash
./fabric --list
```

### Help
```bash
./fabric --help
```

## Next Steps

1. **Read full docs**: [README.md](README.md)
2. **Installation options**: [INSTALL.md](INSTALL.md)
3. **Kubernetes deployment**: [k8s/fabric-deployment.yaml](k8s/fabric-deployment.yaml)
4. **AWS infrastructure**: [terraform/main.tf](terraform/main.tf)

---

**You're ready to go. 2 minutes. Done.**
