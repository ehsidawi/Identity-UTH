```markdown
# Identity-UTH: Universal Trust Hypervisor
## Global Access Kill Switch Orchestrator

**One trigger. Every identity silo. Instant revocation.**

[![GitHub Release](https://img.shields.io/github/v/release/ehsidawi/Identity-UTH)](https://github.com/ehsidawi/Identity-UTH/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue)](https://golang.org)

---

## 🚀 Quick Start (< 2 minutes)

```bash
git clone https://github.com/ehsidawi/Identity-UTH.git
cd Identity-UTH/control-plane
go build -o fabric-control-plane .
./fabric-control-plane --blast-radius global
```

The binary **finds `config.yaml` automatically** and executes a kill switch across **all configured identity providers** in parallel.

### Example Output (simplified)
```json
{
  "id": "kill-switch-...",
  "blast_radius": "global",
  "status": "completed",
  "results": [
    { "provider": "Okta", "category": "Workforce IAM", "status": "completed" },
    { "provider": "CyberArk", "category": "PAM", "status": "completed" },
    { "provider": "AWS IAM", "category": "Cloud IAM", "status": "completed" }
  ]
}
```

---

## 🎯 Coverage – 13 Identity Categories, 22+ Providers

| Category | Providers |
|----------|-----------|
| **Workforce IAM** | Okta, Azure AD, Ping Identity, OneLogin |
| **CIAM** | Auth0, AWS Cognito |
| **PAM** | CyberArk, Delinea, BeyondTrust |
| **PIM** | Azure AD PIM |
| **NHI / Secrets** | HashiCorp Vault, AWS Secrets Manager, Akeyless |
| **API Gateways** | Kong, AWS API Gateway |
| **Cloud IAM** | AWS IAM, GCP IAM |
| **Zero Trust** | Zscaler, Cloudflare Zero Trust |
| **Endpoint** | CrowdStrike |
| **Session Store** | Redis |
| **Service Mesh** | Istio |

All providers are **pluggable** via a simple Go interface.  
To add a new provider (any vendor), implement two functions and register it in the registry.

---

## 📋 How It Works

1. **Configuration** – `config.yaml` (in repo root) lists which providers to activate.  
2. **Parallel Revocation** – Each provider is called concurrently.  
3. **Audit Trail** – Every execution produces a machine‑readable JSON event.  
4. **Auto‑Discovery** – The binary finds `config.yaml` automatically (relative to its own location).

### Blast Radius Modes
- `global` – All users, all apps, all regions
- `region` – All workloads in a specific region
- `app` – Single application
- `tenant` – Single tenant (multi‑tenant systems)
- `user` – Individual user account

---

## 🔧 Configuration (`config.yaml`)

A fully populated example (included in the repo root) enables all providers.  
To disable a provider, simply remove its entry.  
For production, replace placeholder config values with real credentials (or use environment variables).

```yaml
providers:
  - type: okta
    config:
      org_url: "${OKTA_ORG_URL}"
      api_token: "${OKTA_API_TOKEN}"
  - type: azure_ad
    config:
      tenant_id: "${AZURE_TENANT_ID}"
      client_id: "${AZURE_CLIENT_ID}"
      client_secret: "${AZURE_CLIENT_SECRET}"
  # … enable/disable any of the 22+ providers
```

---

## 🧩 Adding a New Provider (Any Company)

1. Create a struct implementing the `RevocationProvider` interface:
   ```go
   type MyProvider struct{}
   func (p *MyProvider) Name() string { return "My Vendor" }
   func (p *MyProvider) Category() string { return "IAM" }
   func (p *MyProvider) Revoke(blastRadius string) error {
       // Real API call here
       return nil
   }
   ```
2. Add a constructor `NewMyProvider(cfg map[string]interface{}) (RevocationProvider, error)`.
3. Register it in the `registry` map inside `NewFabric()`.
4. Add its entry to `config.yaml`.

---

## 📁 Project Structure
```
Identity-UTH/
├── README.md                   ← you are here
├── QUICKSTART.md               ← 2‑minute setup guide
├── INSTALL.md                  ← installation methods
├── LICENSE                     ← MIT
├── config.yaml                 ← provider configuration (auto‑discovered)
├── control-plane/              ← Go binary source
│   ├── main.go                 ← orchestration logic
│   ├── providers.go            ← all provider stubs (22+)
│   ├── go.mod / go.sum
│   └── fabric-control-plane*   ← compiled binary (not in Git)
├── k8s/                        ← Kubernetes deployment
├── policies/                   ← OPA policies
├── orchestration/              ← SOAR workflow definition
├── terraform/                  ← AWS infrastructure (optional)
├── .github/workflows/          ← CI/CD pipeline
└── Dockerfile                  ← container image
```

---

## 🚢 Deployment Options

### Standalone CLI (instant – recommended for testing)
```bash
./fabric-control-plane --blast-radius global
```
- No external dependencies
- Runs on macOS, Linux, Windows (Go cross‑compile)

### Docker
```bash
docker build -t identity-uth .
docker run -it identity-uth
```

### Kubernetes (production)
```bash
kubectl apply -f k8s/fabric-deployment.yaml
```

### AWS Infrastructure (enterprise)
```bash
cd terraform
terraform init && terraform apply
```

---

## 🧪 Testing

```bash
cd control-plane
go test -v ./...
```

---

## 📖 Documentation

- **[QUICKSTART.md](QUICKSTART.md)** – 2‑minute setup
- **[INSTALL.md](INSTALL.md)** – All installation methods
- **[LICENSE](LICENSE)** – MIT open source

---

## 🛠️ Requirements

- **Go 1.21+** (to build from source)
- No other dependencies – the binary is self‑contained

---

## 🤝 Contributing

Contributions welcome!  
Open an issue or pull request at https://github.com/ehsidawi/Identity-UTH.

---

## 📜 License

MIT – see [LICENSE](LICENSE) for details.

---

**Built for enterprise security teams who need to stop everything, everywhere, right now.**
```
