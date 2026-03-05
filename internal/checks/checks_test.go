package checks

import (
	"testing"

	"oaicheck/internal/config"
)

func TestBuildEnvelopeOK(t *testing.T) {
	cfg := config.Resolved{BaseURL: "https://api.openai.com/v1", APIKey: "x", Model: "gpt-4.1-mini"}
	results := []CheckResult{{Name: CheckPing, OK: true, Message: "reachable"}}
	env := BuildEnvelope("ping", cfg, results, PingData{Reachable: true, Status: 401})

	if !env.OK {
		t.Fatal("expected OK envelope")
	}
	if env.Error != nil {
		t.Fatal("expected nil error payload")
	}
	if !env.Config.APIKeyPresent {
		t.Fatal("expected redacted config to indicate apiKeyPresent")
	}
}

func TestBuildEnvelopeFailure(t *testing.T) {
	cfg := config.Resolved{BaseURL: "https://api.openai.com/v1"}
	results := []CheckResult{{Name: CheckModels, OK: false, Message: "missing API key"}}
	env := BuildEnvelope("models", cfg, results, ModelsData{})

	if env.OK {
		t.Fatal("expected failing envelope")
	}
	if env.Error == nil {
		t.Fatal("expected error payload")
	}
	if env.Error.Message != "missing API key" {
		t.Fatalf("unexpected error message: %q", env.Error.Message)
	}
}
