# Pied Piper - Quick Start Guide

Get your team of AI SubAgents up and running in minutes in ClaudeCode or your favourite Coding CLI.

## Installation

```bash
go install github.com/sathish316/pied-piper
pied-piper help
```

## Quick Start

### 1. Create Your Team

```bash
pied-piper team create --name "pied-piper"
```

This creates a default SDLC team with 7 SubAgents: orchestrator, product-manager, architect, software-engineer, code-reviewer, code-validator, and build-engineer.

### 2. View Your Team

```bash
# Show full team config
pied-piper team show --name "pied-piper"

# List all SubAgents
pied-piper subagent list --team "pied-piper"

# Show specific SubAgent
pied-piper subagent show --team "pied-piper" --name "architect"
```

### 3. Generate SubAgents for Claude Code

**Generate all SubAgents (global):**
```bash
pied-piper subagent generate-all --team "pied-piper" --target claude-code
pied-piper subagent generate-all --team "pied-piper" --target claude-code --project-dir /path/to/project
```

**Generate single SubAgent:**
```bash
pied-piper subagent generate --team "pied-piper" --name "architect" --target claude-code
```

### 4. Modify and Regenerate

**Edit team config:**
```bash
vi ~/.pied-piper/teams/pied-piper/team-config.yml
```

**Regenerate SubAgent:**
```bash
pied-piper subagent generate --team "pied-piper" --name "architect" --target claude-code
```

### 5. Export SubAgents to Coding CLI

**Export all SubAgents to ClaudeCode User directory or Project directory:**
```bash
pied-piper subagent export --team "pied-piper" --target claude-code
pied-piper subagent export --team "pied-piper" --target claude-code --project-dir /path/to/project
```

**Export single SubAgent to ClaudeCode User directory or Project directory:**
```bash
pied-piper subagent export --team "pied-piper" --target claude-code --name "architect"
pied-piper subagent export --team "pied-piper" --target claude-code --name "architect" --project-dir /path/to/project
```

### 6. Start hacking in Claude Code using SubAgents

Once generated, use SubAgents in Claude Code with role or nickname or `@` mentions:
- `@architect` or `@Richard`
- `@software-engineer` or `@Gilfoyle`

$ bd create "implement shiny new feature x"
> Ask microsprint-orchestrator to start a microsprint
> Run "bd quickstart". List ready tasks
> Ask microsprint-orchestrator to start task abc

