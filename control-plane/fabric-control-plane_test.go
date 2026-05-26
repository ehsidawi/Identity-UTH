package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestFabricInit(t *testing.T) {
	// Create a minimal config file
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("providers:\n  - type: okta\n    config: {}\n")
	tmpFile.Close()

	_, err = NewFabric(tmpFile.Name())
	if err != nil {
		t.Fatalf("Expected successful init, got %v", err)
	}
}

func TestKillSwitchExecution(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	f := &Fabric{
		providers: []RevocationProvider{&OktaProvider{}},
	}
	f.TriggerKillSwitch("global", "test", "tester")
	if len(f.events) != 1 {
		t.Error("Expected one event")
	}
}
