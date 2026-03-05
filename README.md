# oaicheck

A tiny CLI for debugging OpenAI API config.

## Commands

- `oaicheck` (same as `oaicheck doctor`)
- `oaicheck doctor`
- `oaicheck ping`
- `oaicheck models`
- `oaicheck probe`

## Flags

- `--base-url` (or `OPENAI_BASE_URL`, default `https://api.openai.com/v1`)
- `--api-key` (or `OPENAI_API_KEY`)
- `--model` (or `OPENAI_MODEL`)
- `--json` (machine-readable output)

Flags override environment variables.

## What each check does

- `ping`: sends a request to `<base-url>/models` and treats any HTTP response as reachable.
- `models`: calls OpenAI models list to verify base URL + API key.
- `probe`: sends a tiny generation request to `/responses`; if that fails, retries with `/chat/completions`.
- `doctor`: runs all three checks and summarizes pass/fail.

## JSON mode

`--json` prints exactly one JSON object to stdout and nothing else.

Example:

```bash
oaicheck doctor --json
```

## Build

```bash
go build -o oaicheck .
```

## Quick run

```bash
OPENAI_API_KEY=... OPENAI_MODEL=gpt-4.1-mini oaicheck doctor
```
