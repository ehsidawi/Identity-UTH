package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type KillSwitchEvent struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	BlastRadius string    `json:"blast_radius"`
	Reason      string    `json:"reason"`
	InitiatedBy string    `json:"initiated_by"`
	Status      string    `json:"status"`
}

type IdentitySecurityFabric struct {
	kafkaWriter *kafka.Writer
	mu          sync.RWMutex
}

func NewIdentitySecurityFabric(brokers []string) (*IdentitySecurityFabric, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        "kill-switch-events",
		WriteTimeout: 10 * time.Second,
	})

	return &IdentitySecurityFabric{
		kafkaWriter: writer,
	}, nil
}

func (f *IdentitySecurityFabric) TriggerKillSwitch(ctx context.Context, blastRadius, reason, initiatedBy string) error {
	event := KillSwitchEvent{
		ID:          fmt.Sprintf("kill-switch-%d", time.Now().UnixNano()),
		Timestamp:   time.Now(),
		BlastRadius: blastRadius,
		Reason:      reason,
		InitiatedBy: initiatedBy,
		Status:      "executing",
	}
clear
	eventJSON, _ := json.Marshal(event)
	msg := kafka.Message{
		Key:   []byte(event.ID),
		Value: eventJSON,
	}

	if err := f.kafkaWriter.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("failed to publish kill switch event: %w", err)
	}

	event.Status = "completed"
	log.Printf("Kill switch executed: %s", event.ID)
	return nil
}

func main() {
	brokers := []string{"localhost:9092"}
	fabric, err := NewIdentitySecurityFabric(brokers)
	if err != nil {
		log.Fatalf("Failed to initialize fabric: %v", err)
	}

	ctx := context.Background()
	err = fabric.TriggerKillSwitch(ctx, "global", "Emergency lockdown", "admin@company.com")
	if err != nil {
		log.Fatalf("Kill switch failed: %v", err)
	}

	log.Println("Kill switch executed successfully")
}
