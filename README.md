# Identity-UTH: Universal Trust Hypervisor
## Global Access Kill Switch Orchestrator

**Enterprise-grade emergency access revocation system for Workforce IAM, CIAM, PAM, Non-Human Identity, APIs, cloud, and SaaS.**

### ⚡ Quick Start (60 seconds)

```bash
git clone https://github.com/ehsidawi/Identity-UTH.git
cd Identity-UTH/control-plane
go build -o fabric-control-plane main.go
./fabric-control-plane --blast-radius global --reason "Emergency" --initiated-by "admin@company.com"
```

### 🎯 Blast Radius Modes
- **user**: Single user account
- **tenant**: All users in tenant  
- **app**: Single application
- **region**: All workloads in region
- **global**: All users, all apps, all regions

### 📋 Requirements
- Go 1.21+
- macOS/Linux
- No external dependencies (standalone)

### 🔧 Installation

```bash
git clone https://github.com/ehsidawi/Identity-UTH.git
cd Identity-UTH/control-plane
go mod tidy
go build -o fabric-control-plane main.go
```

### 🚀 Usage

```bash
# Global kill switch
./fabric-control-plane --blast-radius global --reason "Breach detected" --initiated-by "ciso@company.com"

# Single tenant isolation
./fabric-control-plane --blast-radius tenant --reason "Suspicious activity" --initiated-by "security@company.com"

# View all events
./fabric-control-plane --list
```

### 📁 Project Structure
Identity-UTH/
├── README.md                          # Documentation
├── control-plane/                     # Go binary
│   ├── main.go                        # Kill switch logic
│   ├── go.mod                         # Dependencies
│   ├── fabric-control-plane           # Compiled binary
│   └── .gitignore
├── k8s/
│   └── fabric-deployment.yaml         # Kubernetes deployment
├── policies/
│   └── kill-switch.rego               # OPA policies
├── orchestration/
│   └── kill-switch-workflow.yaml      # SOAR workflow
├── terraform/
│   └── main.tf                        # AWS infrastructure
├── .github/
│   └── workflows/ci-cd.yaml           # GitHub Actions
└── Dockerfile                         # Container image

### 🏗️ Deployment Options

**Option 1: Standalone CLI (Recommended)**
```bash
./fabric-control-plane --blast-radius global
```

**Option 2: Kubernetes**
```bash
kubectl apply -f k8s/fabric-deployment.yaml
```

**Option 3: AWS Infrastructure**
```bash
terraform apply -f terraform/main.tf
```

### 🔐 What Gets Revoked
- OAuth/OIDC tokens
- Session cookies
- API keys
- Service account credentials
- Privileged access sessions
- Workload identities
- Cloud IAM permissions

### 📊 Output Format
```json
{
  "id": "kill-switch-1779822119611966000",
  "timestamp": "2026-05-26T22:01:59.611969+03:00",
  "blast_radius": "global",
  "reason": "Emergency lockdown",
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

### 🧪 Testing
```bash
cd control-plane
go test -v ./...
```

### 📜 License
Internal Use Only - Towly.ai

### 📞 Support
- Issues: https://github.com/ehsidawi/Identity-UTH/issues
- Source: https://github.com/ehsidawi/Identity-UTH

**Built with ❤️ for enterprise security teams**
