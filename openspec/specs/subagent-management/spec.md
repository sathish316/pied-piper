# Subagent Management Specification

## Purpose
Define how users inspect, create, and maintain the subagents configured for a named team.

## Requirements

### Requirement: List configured subagents
The system SHALL list the roles configured for a team.

#### Scenario: List the default team
- **WHEN** the user runs `pied-piper subagent list` without overriding the team
- **THEN** the system SHALL list the roles in the default `pied-piper` team

### Requirement: Show a configured subagent
The system SHALL find and display a configured subagent by role or nickname.

#### Scenario: Show by role
- **WHEN** the user runs `pied-piper subagent show --name software-engineer`
- **THEN** the system SHALL display the software engineer configuration as YAML

#### Scenario: Show by nickname
- **WHEN** a caller requests a subagent using a configured nickname
- **THEN** the system SHALL return the matching subagent configuration

#### Scenario: Show an unknown subagent
- **WHEN** no configured role or nickname matches the request
- **THEN** the system SHALL return a not-found error naming the requested subagent

### Requirement: Per-subagent specification persistence
The system SHALL read and write each detailed subagent specification as `<team-config-dir>/subagents/<role>.md`.

#### Scenario: Read a subagent specification
- **WHEN** the system requests the specification for `software-engineer`
- **THEN** it SHALL parse `subagents/software-engineer.md` into role, nickname, task labels, wiki labels, generated task and wiki workflow descriptions, role description, and memory fields

#### Scenario: Update a subagent specification
- **WHEN** a generated or edited specification is saved for the `architect` role
- **THEN** the system SHALL write the rendered specification to `subagents/architect.md` and return that path

### Requirement: Create a configured subagent
The system SHALL append a new role to a named team's YAML configuration with an optional nickname and empty placeholders for descriptions and routing workflows.

#### Scenario: Create a new role
- **WHEN** the user runs `pied-piper subagent create --team <team-name> --role <role> --nickname <nickname>` for a role not already present
- **THEN** the system SHALL append the role and nickname to the team configuration
- **AND** it SHALL persist empty task-label, wiki-label, and workflow-description fields for later customization

#### Scenario: Reject a duplicate role
- **WHEN** the requested role already exists in the named team
- **THEN** the system SHALL report the duplicate and leave the configuration unchanged

#### Scenario: Omit required creation inputs
- **WHEN** the user omits the team or role from `pied-piper subagent create`
- **THEN** command validation SHALL report the missing required flag
