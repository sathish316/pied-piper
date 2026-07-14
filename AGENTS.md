# AGENTS

See `README.md` for the product overview and the full CLI command reference.

## Cursor Cloud specific instructions

Pied Piper is a Go (Cobra) CLI that generates/manages a team of AI subagents.

- **Requires Go 1.24.2.** `go.mod` declares `go 1.24.2`; with an older Go on PATH the `go` toolchain
  auto-downloads 1.24.2 on first build. Go 1.24.2 is installed at `/usr/local/go`.
- Standard commands (from README): `go build ./...`, `go test ./...`, `go run main.go`.
- Runtime state (teams, subagents, generated specs) is written under `~/.pied-piper/<team>/` and
  generated agent files under `~/.claude/agents`, not in the repo.
- A typical smoke test: `go run main.go team create -n demo` then
  `go run main.go subagent create -t demo -r engineer` then `go run main.go subagent list -t demo`.
- No linter is configured (use `go vet` / `gofmt` if needed).
