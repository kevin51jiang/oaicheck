package config

import "testing"

func TestResolvePrecedence(t *testing.T) {
	t.Setenv("OPENAI_BASE_URL", "https://env.example/v1")
	t.Setenv("OPENAI_API_KEY", "env-key")
	t.Setenv("OPENAI_MODEL", "env-model")

	cfg := Resolve("https://flag.example/v1", "flag-key", "flag-model")
	if cfg.BaseURL != "https://flag.example/v1" {
		t.Fatalf("expected flag base URL, got %q", cfg.BaseURL)
	}
	if cfg.APIKey != "flag-key" {
		t.Fatalf("expected flag API key, got %q", cfg.APIKey)
	}
	if cfg.Model != "flag-model" {
		t.Fatalf("expected flag model, got %q", cfg.Model)
	}
}

func TestResolveDefaults(t *testing.T) {
	t.Setenv("OPENAI_BASE_URL", "")
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("OPENAI_MODEL", "")

	cfg := Resolve("", "", "")
	if cfg.BaseURL != DefaultBaseURL {
		t.Fatalf("expected default base URL, got %q", cfg.BaseURL)
	}
	if cfg.APIKey != "" {
		t.Fatalf("expected empty API key, got %q", cfg.APIKey)
	}
	if cfg.Model != "" {
		t.Fatalf("expected empty model, got %q", cfg.Model)
	}
}
