package output

import (
	"encoding/json"
	"fmt"
	"io"

	"oaicheck/internal/checks"
)

func RenderHuman(w io.Writer, env checks.Envelope) error {
	for _, check := range env.Checks {
		status := "PASS"
		icon := "✅"
		if !check.OK {
			status = "FAIL"
			icon = "❌"
		}

		if check.Details != "" {
			if _, err := fmt.Fprintf(w, "%s %s %s: %s (%s)\n", icon, status, check.Name, check.Message, check.Details); err != nil {
				return err
			}
			continue
		}

		if _, err := fmt.Fprintf(w, "%s %s %s: %s\n", icon, status, check.Name, check.Message); err != nil {
			return err
		}
	}

	if env.Command == "doctor" {
		if summary, ok := env.Data.(checks.DoctorData); ok {
			if _, err := fmt.Fprintf(w, "Summary: %d passed, %d failed\n", summary.Passed, summary.Failed); err != nil {
				return err
			}
		}
	}

	if env.Error != nil {
		if _, err := fmt.Fprintf(w, "Error: %s\n", env.Error.Message); err != nil {
			return err
		}
	}

	return nil
}

func RenderJSON(w io.Writer, env checks.Envelope) error {
	enc := json.NewEncoder(w)
	return enc.Encode(env)
}
