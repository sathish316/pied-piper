# Subagent Export Specification

## Purpose
Define how generated team subagent specifications are copied into user-level or project-level coding-agent configuration locations.

## Requirements

### Requirement: Supported export targets
The system SHALL export subagent specifications for Claude Code and Rovo Dev using each target's native configuration path.

#### Scenario: Select Claude Code
- **WHEN** the export target is `claude-code`
- **THEN** the system SHALL use `.claude/agents` as the coding-agent directory

#### Scenario: Select Rovo Dev
- **WHEN** the export target is `rovodev`
- **THEN** the system SHALL use `.rovodev/subagents` as the coding-agent directory

#### Scenario: Reject an unsupported export target
- **WHEN** the user selects an unrecognized coding-agent target
- **THEN** the export command SHALL report that the target is unsupported without copying files

### Requirement: Export one subagent
The system SHALL copy one generated `<role>.md` specification from the named team's `subagents` directory to the selected coding-agent location.

#### Scenario: Export to the user configuration
- **WHEN** the user exports one subagent without a project directory
- **THEN** the system SHALL copy that subagent to the selected coding agent's user-level configuration path

#### Scenario: Export to a project configuration
- **WHEN** the user exports one subagent with `--project-dir <project-dir>`
- **THEN** the system SHALL copy that subagent to the selected coding agent's configuration path beneath the project

### Requirement: Export all team subagents
The system SHALL support exporting every configured subagent in a named team.

#### Scenario: Export all to the user configuration
- **WHEN** the user runs `pied-piper export all --team <team-name>` without a project directory
- **THEN** the system SHALL copy every configured team subagent to the selected coding agent's user-level directory

#### Scenario: Export all to a project configuration
- **WHEN** the user runs `pied-piper export all --team <team-name> --project-dir <project-dir>`
- **THEN** the system SHALL copy every configured team subagent to the selected coding agent's project-level directory

### Requirement: Required export inputs
The export command SHALL require a team for bulk export and both a team and subagent name for single-subagent export.

#### Scenario: Omit a required export input
- **WHEN** the user invokes an export command without its required team or subagent flag
- **THEN** command validation SHALL report the missing required flag
