# Team Configuration Specification

## Purpose
Define how Pied Piper initializes, loads, displays, and interprets the YAML configuration for an SDLC subagent team.

## Requirements

### Requirement: Named team configuration initialization
The system SHALL initialize each named team under `$HOME/.pied-piper/<team-name>/` with a configuration file, a `subagents` directory, and a `templates` directory.

#### Scenario: Create the default configuration
- **WHEN** the user runs `pied-piper team create`
- **THEN** the system SHALL create `$HOME/.pied-piper/pied-piper/config.yml` from the embedded default-team configuration
- **AND** the system SHALL create the team's `subagents` and `templates` directories

#### Scenario: Create a custom team
- **WHEN** the user runs `pied-piper team create --name test-titans`
- **THEN** the system SHALL initialize `$HOME/.pied-piper/test-titans/config.yml` from a blank team template

#### Scenario: Initialize coding-agent templates
- **WHEN** a team is initialized
- **THEN** the system SHALL create subagent templates for Claude Code and Rovo Dev in the team's `templates` directory

### Requirement: Configured team model
The system SHALL load team name, description, subagent roles, nicknames, task-label routing, and wiki-label routing from YAML.

#### Scenario: Load the embedded default team
- **WHEN** the initialized default configuration is loaded
- **THEN** the team SHALL be named `pied-piper`
- **AND** it SHALL contain seven configured subagents

#### Scenario: Reject unreadable configuration
- **WHEN** the configured YAML file cannot be read or parsed
- **THEN** the system SHALL return an error describing the read or parse failure

### Requirement: Default SDLC roles
The default team SHALL configure microsprint orchestrator, product manager, architect, software engineer, code reviewer, code validator, and build engineer roles with their incoming and outgoing routing labels.

#### Scenario: Load software engineer routing
- **WHEN** the default software engineer configuration is loaded
- **THEN** its nickname SHALL be `Gilfoyle`
- **AND** it SHALL accept development and review-rejection work and emit `@ready-for-code-review`

### Requirement: Team configuration display
The system SHALL serialize the loaded team configuration as YAML for display.

#### Scenario: Show the active team
- **WHEN** the user runs `pied-piper team show --name <team-name>` and the named configuration loads successfully
- **THEN** the system SHALL print the active configuration in YAML form

### Requirement: Role lookup
The system SHALL locate a configured subagent by its role.

#### Scenario: Find an existing role
- **WHEN** a caller requests the `software-engineer` role from the default team
- **THEN** the system SHALL return that role's configuration

#### Scenario: Find an unknown role
- **WHEN** a caller requests a role that is not configured
- **THEN** the system SHALL return a not-found error naming that role
