package config

import (
	"os"
	"strings"
)

const DefaultBaseURL = "https://api.openai.com/v1"

type Resolved struct {
	BaseURL string
	APIKey  string
	Model   string
}

type SafeView struct {
	BaseURL       string `json:"baseUrl"`
	APIKeyPresent bool   `json:"apiKeyPresent"`
	Model         string `json:"model"`
}

func Resolve(baseURLFlag, apiKeyFlag, modelFlag string) Resolved {
	baseURL := firstNonEmpty(strings.TrimSpace(baseURLFlag), strings.TrimSpace(os.Getenv("OPENAI_BASE_URL")), DefaultBaseURL)
	apiKey := firstNonEmpty(strings.TrimSpace(apiKeyFlag), strings.TrimSpace(os.Getenv("OPENAI_API_KEY")))
	model := firstNonEmpty(strings.TrimSpace(modelFlag), strings.TrimSpace(os.Getenv("OPENAI_MODEL")))

	return Resolved{
		BaseURL: strings.TrimRight(baseURL, "/"),
		APIKey:  apiKey,
		Model:   model,
	}
}

func (r Resolved) Safe() SafeView {
	return SafeView{
		BaseURL:       r.BaseURL,
		APIKeyPresent: r.APIKey != "",
		Model:         r.Model,
	}
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
