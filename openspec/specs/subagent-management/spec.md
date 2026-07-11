# Subagent Management Specification

## Purpose
Define how users inspect configured subagents, maintain per-subagent specifications, and generate subagent YAML for supported coding-agent targets.

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
The system SHALL read and write each detailed subagent specification as `<team-config-dir>/subagents/<role>.yml`.

#### Scenario: Read a subagent specification
- **WHEN** the system requests the specification for `software-engineer`
- **THEN** it SHALL parse `subagents/software-engineer.yml` into role, nickname, task labels, wiki labels, workflow description, role description, and memory fields

#### Scenario: Update a subagent specification
- **WHEN** a generated or edited specification is saved for the `architect` role
- **THEN** the system SHALL write YAML to `subagents/architect.yml` and return that path

### Requirement: Generate a subagent specification
The system SHALL generate a detailed subagent specification from the matching team entry while preserving its role, description, nickname, task labels, and wiki labels.

#### Scenario: Generate the software engineer specification
- **WHEN** generation is requested for `software-engineer`
- **THEN** the generated YAML SHALL include role `software-engineer`, nickname `Gilfoyle`, its incoming task labels, and `@ready-for-code-review`

### Requirement: Claude Code target resolution
The system SHALL support `claude-code` as a coding-agent target and distinguish user-level from project-level target directories.

#### Scenario: Resolve a user target
- **WHEN** `claude-code` is selected without a project directory
- **THEN** the target directory SHALL be `$HOME/.claude/agents`
- **AND** the target directory type SHALL be `user`

#### Scenario: Resolve a project target
- **WHEN** `claude-code` is selected with a project directory
- **THEN** the target directory SHALL be `<project-dir>/.claude/agents`
- **AND** the target directory type SHALL be `project`

#### Scenario: Reject an unsupported target
- **WHEN** a coding-agent target other than `claude-code` is selected
- **THEN** the system SHALL return an error stating that the target is unsupported

### Requirement: Subagent generation command inputs
The generation command SHALL require a subagent name and coding-agent target, default the team to `pied-piper`, and accept an optional project directory.

#### Scenario: Omit a required generation input
- **WHEN** the user invokes `pied-piper subagent generate` without a name or target
- **THEN** command validation SHALL report the missing required flag
