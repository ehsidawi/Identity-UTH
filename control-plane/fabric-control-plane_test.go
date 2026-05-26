package main

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestKillSwitchOutput(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	fabric := &Fabric{}
	fabric.TriggerKillSwitch("global", "test", "tester")

	output := buf.String()
	if !strings.Contains(output, "kill-switch") {
		t.Error("Expected output to contain kill-switch event ID")
	}
}
