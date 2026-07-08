# 0001: Build the CLI in Go

Date: 2026-07-06
Status: accepted
Grandfathered: enacted before the evidence-ledger convention (ADR-0006); the CLI is built and verified in code.

## Context

The CLI is the third artifact — a read-only validator and context provider that sits alongside the loop (`loop.sh`) and the skills (`.agents/skills/`). The loop is bash; the CLI needs to parse markdown, check structural schemas, and embed/package assets. Doing that in bash would be painful.

Alternatives considered:
- **Bash** — already used for the loop, but markdown parsing and embedded assets are a mess in bash.
- **Python** — good for text processing, but adds a runtime dependency and packaging story. The loop already leans on Python for small inline scripts.
- **Rust** — fast single binary, but heavier toolchain and slower to iterate.
- **Go** — proven for this exact use case (litespec was Go). Single static binary, `embed` package for asset packaging, good stdlib for CLI tools, fast builds.

## Decision

We will build the CLI in Go. It will be a self-contained Go module under `cli/` with its own `go.mod`, producing a single static binary.

## Consequences

- The CLI can use `go:embed` to package the default skills into the binary itself — no external asset files needed at runtime.
- Single static binary distribution — no runtime dependency on the user's machine beyond the OS.
- Adds a Go toolchain requirement for contributors who want to build the CLI from source (pre-built binaries can be distributed).
- The loop (`loop.sh`) and the CLI remain independent artifacts in different languages, as the design intends.
