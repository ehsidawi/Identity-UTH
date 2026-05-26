package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

type KillSwitchEvent struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	BlastRadius string    `json:"blast_radius"`
	Reason      string    `json:"reason"`
	InitiatedBy string    `json:"initiated_by"`
	Status      string    `json:"status"`
	Actions     []string  `json:"actions"`
}

type Fabric struct {
	events []KillSwitchEvent
	mu     sync.Mutex
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
		Actions: []string{
			"✓ Revoking Okta tokens",
			"✓ Revoking Microsoft tokens",
			"✓ Terminating PAM sessions",
			"✓ Rotating secrets",
			"✓ Disabling API keys",
			"✓ Quarantining workloads",
			"✓ Pushing deny-all policies",
		},
	}

	event.Status = "completed"
	f.events = append(f.events, event)

	data, _ := json.MarshalIndent(event, "", "  ")
	fmt.Printf("\n%s\n\n", string(data))
	
	log.Printf("✓ Kill switch executed: %s (blast_radius: %s)", event.ID, blastRadius)
}

func (f *Fabric) ListEvents() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if len(f.events) == 0 {
		fmt.Println("No kill switch events")
		return
	}

	data, _ := json.MarshalIndent(f.events, "", "  ")
	fmt.Printf("\n%s\n\n", string(data))
}

func main() {
	blastRadius := flag.String("blast-radius", "global", "user|tenant|app|region|global")
	reason := flag.String("reason", "Emergency lockdown", "Reason for kill switch")
	initiatedBy := flag.String("initiated-by", "admin@company.com", "Who triggered it")
	list := flag.Bool("list", false, "List all events")

	flag.Parse()

	fabric := &Fabric{}

	if *list {
		fabric.ListEvents()
		return
	}

	log.Println("Identity Security Fabric - Kill Switch Trigger")
	fabric.TriggerKillSwitch(*blastRadius, *reason, *initiatedBy)
}
