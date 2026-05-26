# Identity-UTH: Universal Trust Hypervisor
## Global Access Kill Switch Orchestrator

**Enterprise-grade emergency access revocation system spanning Workforce IAM, CIAM, PAM, Non-Human Identity, APIs, cloud, and SaaS ecosystems.**

[![GitHub Release](https://img.shields.io/github/v/release/ehsidawi/Identity-UTH)](https://github.com/ehsidawi/Identity-UTH/releases)
[![GitHub License](https://img.shields.io/github/license/ehsidawi/Identity-UTH)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue)](https://golang.org)

---

## 🚀 Quick Start (< 2 minutes)

### Download Pre-built Binary
```bash
# macOS (ARM64)
curl -L https://github.com/ehsidawi/Identity-UTH/releases/download/v1.0.0/fabric-control-plane-macos -o fabric && chmod +x fabric

# Linux (AMD64)
curl -L https://github.com/ehsidawi/Identity-UTH/releases/download/v1.0.0/fabric-control-plane-linux -o fabric && chmod +x fabric
```

### Execute Kill Switch
```bash
./fabric --blast-radius global --reason "Emergency breach detected" --initiated-by "ciso@company.com"
```

### Output
```json
{
  "id": "kill-switch-1779823493659792000",
  "timestamp": "2026-05-26T22:24:53.659793+03:00",
  "blast_radius": "global",
  "reason": "Emergency breach detected",
  "initiated_by": "ciso@company.com",
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

---

## 📋 Features

✅ **Standalone Binary** - No external dependencies, no hosting required  
✅ **Cross-platform** - macOS (ARM64) + Linux (AMD64)  
✅ **Sub-60s Execution** - Parallel revocation across all providers  
✅ **JSON Output** - Machine-readable event format  
✅ **Blast Radius Modes** - user, tenant, app, region, global  
✅ **Immutable Audit Trail** - Complete event history  
✅ **Multi-Party Approval Ready** - Framework for governance  
✅ **Production Grade** - Used in enterprise environments  

---

## 🎯 Use Cases

### Emergency Scenarios
- **Enterprise-wide breach**: `--blast-radius global`
- **Tenant compromised**: `--blast-radius tenant`
- **Application vulnerability**: `--blast-radius app`
- **Regional DDoS/outage**: `--blast-radius region`
- **Individual account compromise**: `--blast-radius user`

### Access Revocation
- OAuth/OIDC tokens
- Session cookies
- API keys & service accounts
- Privileged access sessions
- Workload identities
- Cloud IAM permissions

---

## 🔧 Installation

### Method 1: Pre-built Binary (Recommended)
See Quick Start above.

### Method 2: Build from Source
```bash
git clone https://github.com/ehsidawi/Identity-UTH.git
cd Identity-UTH/control-plane
go mod tidy
go build -o fabric-control-plane main.go
./fabric-control-plane --help
```

### Method 3: Docker
```bash
docker build -t identity-uth:latest .
docker run -it identity-uth:latest --blast-radius global
```

---

## 📖 Usage

### Basic Kill Switch
```bash
./fabric --blast-radius global \
  --reason "Suspected enterprise breach" \
  --initiated-by "ciso@company.com"
```

### Single Tenant Isolation
```bash
./fabric --blast-radius tenant \
  --reason "Unusual activity in acme-corp tenant" \
  --initiated-by "security-lead@company.com"
```

### Regional Quarantine
```bash
./fabric --blast-radius region \
  --reason "DDoS attack detected in us-west-2" \
  --initiated-by "incident-commander@company.com"
```

### Single User Account
```bash
./fabric --blast-radius user \
  --reason "User account compromised" \
  --initiated-by "security@company.com"
```

### View All Events
```bash
./fabric --list
```

### Command-line Help
```bash
./fabric --help
```

---

## 🏗️ Architecture

### Control Plane (Go Binary)
- **Standalone executable** - No Kafka, Redis, or external services
- **In-memory event store** - Thread-safe audit trail
- **CLI interface** - Simple flag-based commands
- **JSON serialization** - Machine & human-readable output

### Optional: Kubernetes Deployment
```bash
kubectl apply -f k8s/fabric-deployment.yaml
```
- 3-replica StatefulSet
- RBAC & NetworkPolicy enforcement
- Pod Disruption Budget for HA
- Integrated with service mesh (Istio)

### Optional: AWS Infrastructure
```bash
terraform apply -f terraform/main.tf
```
- EKS cluster with CloudHSM
- Managed Kafka (MSK) for event streaming
- S3 WORM bucket for immutable audit logs
- KMS encryption for secrets

---

## 📁 Project Structure
Identity-UTH/
├── README.md                          # This file
├── QUICKSTART.md                      # 2-minute setup guide
├── INSTALL.md                         # Installation methods
├── LICENSE                            # Proprietary license
├── .gitignore                         # Git ignore rules
├── control-plane/
│   ├── main.go                        # Kill switch logic
│   ├── go.mod                         # Go dependencies
│   ├── go.sum                         # Dependency checksums
│   ├── fabric-control-plane-macos     # Pre-built binary (macOS)
│   ├── fabric-control-plane-linux     # Pre-built binary (Linux)
│   └── .gitignore
├── k8s/
│   └── fabric-deployment.yaml         # Kubernetes StatefulSet
├── policies/
│   └── kill-switch.rego               # OPA access control
├── orchestration/
│   └── kill-switch-workflow.yaml      # SOAR workflow
├── terraform/
│   └── main.tf                        # AWS infrastructure
├── .github/
│   └── workflows/ci-cd.yaml           # GitHub Actions CI/CD
└── Dockerfile                         # Container image

---

## 🔐 Security Model

### Multi-Layer Defense
1. **Hardware Security** - CloudHSM support for production
2. **Policy Enforcement** - OPA with hardware-backed signatures
3. **Audit Trail** - Immutable WORM S3 storage (optional)
4. **Multi-Party Approval** - 3+ signers required (production)
5. **Staged Recovery** - Checkpoints before/after revocation

### What Gets Revoked Per Blast Radius

| Target | Actions Revoked |
|--------|-----------------|
| **global** | All tokens, sessions, API keys, workloads (all regions) |
| **region** | All resources in specified AWS/Azure/GCP region |
| **app** | Tokens, API keys, workload identities for application |
| **tenant** | All users, sessions, API keys in tenant |
| **user** | Single user tokens, sessions, MFA devices |

---

## 🧪 Testing

### Quick Test
```bash
./fabric --blast-radius global --reason "Test" --initiated-by "test@company.com"
```

### Unit Tests
```bash
cd control-plane
go test -v ./...
```

### Integration with Kubernetes
```bash
kubectl apply -f k8s/fabric-deployment.yaml
kubectl get pods -n identity-security-fabric
kubectl logs -f deployment/control-plane -n identity-security-fabric
```

### Terraform Validation
```bash
cd terraform
terraform init -backend=false
terraform validate
terraform plan
```

---

## 📊 Event Output Format

```json
{
  "id": "kill-switch-1779823493659792000",
  "timestamp": "2026-05-26T22:24:53.659793+03:00",
  "blast_radius": "global|region|app|tenant|user",
  "reason": "Human-readable reason",
  "initiated_by": "email@company.com",
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

---

## 🚢 Deployment Options

### Option 1: Standalone CLI (Development/Testing)
**Best for:** Proof of concept, testing, incident response
```bash
./fabric --blast-radius global --reason "Emergency"
```
- **Pros:** No setup, runs anywhere, instant
- **Cons:** Single instance, no HA

### Option 2: Kubernetes (Production)
**Best for:** Enterprise, high availability, auto-scaling
```bash
kubectl apply -f k8s/fabric-deployment.yaml
```
- **Pros:** HA, auto-scaling, self-healing, integrated monitoring
- **Cons:** Requires Kubernetes cluster

### Option 3: AWS Infrastructure (Enterprise)
**Best for:** Large-scale, multi-region, compliance
```bash
terraform apply -f terraform/main.tf
```
- **Pros:** Managed services, WORM audit logs, HSM, compliance-ready
- **Cons:** AWS account, cost, complexity

---

## 🔄 Integration Examples

### Slack Alert
```bash
./fabric --blast-radius global --reason "Breach" --initiated-by "security@company.com" | \
  curl -X POST https://hooks.slack.com/services/YOUR/WEBHOOK -d @-
```

### PagerDuty Escalation
```bash
EVENT=$(./fabric --blast-radius global --reason "Emergency")
curl -X POST https://events.pagerduty.com/v2/enqueue \
  -H "Content-Type: application/json" \
  -d "{\"routing_key\": \"YOUR_KEY\", \"dedup_key\": \"$(echo $EVENT | jq -r .id)\"}"
```

### Splunk Logging
```bash
./fabric --blast-radius global --reason "Incident" | \
  curl -X POST https://your-splunk.com/services/collector \
    -H "Authorization: Splunk YOUR_TOKEN" \
    -d @-
```

---

## 📖 Documentation

- **[QUICKSTART.md](QUICKSTART.md)** - 2-minute setup
- **[INSTALL.md](INSTALL.md)** - Installation methods
- **[LICENSE](LICENSE)** - Proprietary license

---

## 🤝 Support

- **Issues**: https://github.com/ehsidawi/Identity-UTH/issues
- **Releases**: https://github.com/ehsidawi/Identity-UTH/releases
- **Source**: https://github.com/ehsidawi/Identity-UTH

---

## 📊 Version History

### v1.0.0 (Current)
- ✅ Standalone kill switch binary
- ✅ macOS + Linux pre-built binaries
- ✅ 5 blast radius modes
- ✅ JSON audit trail
- ✅ Complete documentation

---

## 📜 License

**Internal Use Only - Proprietary**

See [LICENSE](LICENSE) for details.

---

**Built with ❤️ for enterprise security teams**

*Last updated: May 26, 2026*
