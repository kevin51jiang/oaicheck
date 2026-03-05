package output

import (
	"bytes"
	"encoding/json"
	"testing"

	"oaicheck/internal/checks"
	"oaicheck/internal/config"
)

func TestRenderJSON(t *testing.T) {
	env := checks.Envelope{
		OK:      true,
		Command: "ping",
		Config:  config.SafeView{BaseURL: "https://api.openai.com/v1", APIKeyPresent: true, Model: "gpt-4.1-mini"},
		Checks:  []checks.CheckResult{{Name: checks.CheckPing, OK: true, Message: "reachable"}},
		Data:    checks.PingData{Reachable: true, Status: 401},
		Error:   nil,
	}

	var buf bytes.Buffer
	if err := RenderJSON(&buf, env); err != nil {
		t.Fatalf("render json: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("unmarshal rendered json: %v", err)
	}
	if decoded["command"] != "ping" {
		t.Fatalf("expected command ping, got %#v", decoded["command"])
	}
}
