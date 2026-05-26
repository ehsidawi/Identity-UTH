package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

type RevocationProvider interface {
	Name() string
	Category() string   // IAM, CIAM, PAM, NHI, etc.
	Revoke(blastRadius string) error
}

type Config struct {
	Providers []ProviderConfig `yaml:"providers"`
}
type ProviderConfig struct {
	Type   string                 `yaml:"type"`
	Config map[string]interface{} `yaml:"config"`
}

type KillSwitchEvent struct {
	ID          string           `json:"id"`
	Timestamp   time.Time        `json:"timestamp"`
	BlastRadius string           `json:"blast_radius"`
	Reason      string           `json:"reason"`
	InitiatedBy string           `json:"initiated_by"`
	Status      string           `json:"status"`
	Results     []ProviderResult `json:"results"`
}
type ProviderResult struct {
	Provider string `json:"provider"`
	Category string `json:"category"`
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
}

type Fabric struct {
	providers []RevocationProvider
	events    []KillSwitchEvent
	mu        sync.Mutex
}

func NewFabric(configPath string) (*Fabric, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	// Universal provider registry – add any system here
	registry := map[string]func(map[string]interface{}) (RevocationProvider, error){
		// Workforce IAM
		"okta":            NewOktaProvider,
		"azure_ad":        NewAzureADProvider,
		"ping":            NewPingProvider,
		"onelogin":        NewOneLoginProvider,
		// CIAM
		"auth0":           NewAuth0Provider,
		"cognito":         NewCognitoProvider,
		// PAM
		"cyberark":        NewCyberArkProvider,
		"delinea":         NewDelineaProvider,
		"beyondtrust":     NewBeyondTrustProvider,
		// PIM
		"azure_pim":       NewAzurePIMProvider,
		// NHI / Secrets
		"vault":           NewVaultProvider,
		"aws_secrets":     NewAWSSecretsProvider,
		"akeyless":        NewAkeylessProvider,
		// API Gateways
		"kong":            NewKongProvider,
		"aws_api_gw":      NewAWSAPIGatewayProvider,
		// Cloud IAM
		"aws_iam":         NewAWSIAMProvider,
		"gcp_iam":         NewGCPIAMProvider,
		// Zero Trust / ZTNA
		"zscaler":         NewZscalerProvider,
		"cloudflare_zt":   NewCloudflareZTProvider,
		// Endpoint / EDR
		"crowdstrike":     NewCrowdStrikeProvider,
		// Session / Redis
		"redis":           NewRedisSessionProvider,
		// Service Mesh
		"istio":           NewIstioProvider,
	}

	var providers []RevocationProvider
	for _, pc := range cfg.Providers {
		factory, ok := registry[pc.Type]
		if !ok {
			log.Printf("WARNING: unknown provider type %q – skipping", pc.Type)
			continue
		}
		prov, err := factory(pc.Config)
		if err != nil {
			return nil, fmt.Errorf("init %s: %w", pc.Type, err)
		}
		providers = append(providers, prov)
	}
	if len(providers) == 0 {
		return nil, fmt.Errorf("no providers configured")
	}
	return &Fabric{providers: providers}, nil
}

func (f *Fabric) TriggerKillSwitch(blastRadius, reason, initiatedBy string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	event := KillSwitchEvent{
		ID:          fmt.Sprintf("kill-switch-%d", time.Now().UnixNano()),
		Timestamp:   time.Now(),
		BlastRadius: blastRadius,
		Reason:      reason,
		InitiatedBy: initiatedBy,
		Status:      "executing",
	}

	var wg sync.WaitGroup
	results := make([]ProviderResult, len(f.providers))
	for i, prov := range f.providers {
		wg.Add(1)
		go func(idx int, p RevocationProvider) {
			defer wg.Done()
			err := p.Revoke(blastRadius)
			res := ProviderResult{Provider: p.Name(), Category: p.Category()}
			if err != nil {
				res.Status = "failed"
				res.Message = err.Error()
				log.Printf("❌ [%s] %s: %v", p.Category(), p.Name(), err)
			} else {
				res.Status = "completed"
				log.Printf("✓ [%s] %s revoked", p.Category(), p.Name())
			}
			results[idx] = res
		}(i, prov)
	}
	wg.Wait()

	event.Results = results
	event.Status = "completed"
	f.events = append(f.events, event)

	data, _ := json.MarshalIndent(event, "", "  ")
	fmt.Println("\n" + string(data) + "\n")
	log.Printf("✓ Kill switch complete: %s\n", event.ID)
}

func (f *Fabric) ListEvents() {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.events) == 0 {
		fmt.Println("No events")
		return
	}
	data, _ := json.MarshalIndent(f.events, "", "  ")
	fmt.Println("\n" + string(data) + "\n")
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to provider config")
	blastRadius := flag.String("blast-radius", "global", "user|tenant|app|region|global")
	reason := flag.String("reason", "Emergency lockdown", "Reason")
	initiatedBy := flag.String("initiated-by", "admin@company.com", "Who triggered")
	list := flag.Bool("list", false, "List all events")
	flag.Parse()

	fabric, err := NewFabric(*configPath)
	if err != nil {
		log.Fatalf("Init failed: %v", err)
	}

	if *list {
		fabric.ListEvents()
		return
	}

	log.Println("Identity-UTH Universal Kill Switch")
	fabric.TriggerKillSwitch(*blastRadius, *reason, *initiatedBy)
}
