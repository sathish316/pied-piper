# Team Configuration Specification

## Purpose
Define how Pied Piper initializes, loads, displays, and interprets the YAML configuration for an SDLC subagent team.

## Requirements

### Requirement: Default team configuration initialization
The system SHALL initialize an embedded default team configuration at `$HOME/.pied-piper/config.yml` and ensure a sibling `subagents` directory exists.

#### Scenario: Create the default configuration
- **WHEN** the user runs `pied-piper team create`
- **THEN** the system SHALL create `$HOME/.pied-piper/config.yml` from the embedded sample configuration
- **AND** the system SHALL create `$HOME/.pied-piper/subagents`

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
- **WHEN** the user runs `pied-piper team show` and the configuration loads successfully
- **THEN** the system SHALL print the active configuration in YAML form

### Requirement: Role lookup
The system SHALL locate a configured subagent by its role.

#### Scenario: Find an existing role
- **WHEN** a caller requests the `software-engineer` role from the default team
- **THEN** the system SHALL return that role's configuration

#### Scenario: Find an unknown role
- **WHEN** a caller requests a role that is not configured
- **THEN** the system SHALL return a not-found error naming that role
