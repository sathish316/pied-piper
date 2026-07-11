# Subagent Generation Specification

## Purpose
Define how Pied Piper renders team subagent configurations into coding-agent-specific Markdown definitions and assigns models to them.

## Requirements

### Requirement: Generate a subagent specification
The system SHALL render a detailed Markdown subagent specification from the matching team entry and the template for the selected coding agent.

#### Scenario: Generate one subagent
- **WHEN** generation is requested for a configured role
- **THEN** the rendered Markdown SHALL include its role, nickname, description, task and documentation workflows, role description, and memory

### Requirement: Generate all configured subagents
The system SHALL support generating a coding-agent-specific specification for every subagent in a named team.

#### Scenario: Generate the complete default team
- **WHEN** the user runs `pied-piper subagent generate-all --target claude-code`
- **THEN** the system SHALL render one Claude Code specification for each subagent configured in the default team

### Requirement: Coding-agent target resolution
The system SHALL support `claude-code` and `rovodev` coding-agent targets and distinguish user-level from project-level target directories.

#### Scenario: Resolve a Claude Code user target
- **WHEN** `claude-code` is selected without a project directory
- **THEN** the target directory SHALL be `$HOME/.claude/agents`
- **AND** the target directory type SHALL be `user`

#### Scenario: Resolve a Claude Code project target
- **WHEN** `claude-code` is selected with a project directory
- **THEN** the target directory SHALL be `<project-dir>/.claude/agents`
- **AND** the target directory type SHALL be `project`

#### Scenario: Resolve a Rovo Dev user target
- **WHEN** `rovodev` is selected without a project directory
- **THEN** the target directory SHALL be `$HOME/.rovodev/subagents`
- **AND** generated files SHALL use the `.md` extension

#### Scenario: Resolve a Rovo Dev project target
- **WHEN** `rovodev` is selected with a project directory
- **THEN** the target directory SHALL be `<project-dir>/.rovodev/subagents`
- **AND** the target directory type SHALL be `project`

#### Scenario: Reject an unsupported target
- **WHEN** a coding-agent target other than `claude-code` or `rovodev` is selected
- **THEN** the system SHALL return an error stating that the target is unsupported

### Requirement: Subagent generation command inputs
The generation command SHALL require a subagent name and coding-agent target, default the team to `pied-piper`, and accept an optional project directory.

#### Scenario: Omit a required generation input
- **WHEN** the user invokes `pied-piper subagent generate` without a name or target
- **THEN** command validation SHALL report the missing required flag

### Requirement: Subagent generation metaprompt
The system SHALL expose the embedded metaprompt used to enrich generated subagent definitions.

#### Scenario: Show the metaprompt
- **WHEN** the user runs `pied-piper subagent metaprompt`
- **THEN** the system SHALL print the embedded subagent metaprompt

### Requirement: Per-subagent Claude model
The system SHALL allow each subagent to select its own Claude model in team configuration and SHALL default generated Claude Code definitions to `sonnet` when no model is set.

#### Scenario: Render explicitly selected models
- **WHEN** three subagents select `opus`, `sonnet`, and `haiku` respectively
- **THEN** each generated Claude Code definition SHALL contain the model selected for that role

#### Scenario: Default the Claude model
- **WHEN** a subagent does not configure a model
- **THEN** its generated Claude Code definition SHALL use `sonnet`

### Requirement: Claude Code Router model
The system SHALL allow a subagent to declare a Claude Code Router model independently of the native Claude model.

#### Scenario: Render a router model
- **WHEN** a subagent configures `router_model`
- **THEN** its Claude Code definition SHALL include that value in a `CCR-SUBAGENT-MODEL` marker

#### Scenario: Omit an unused router model
- **WHEN** a subagent does not configure `router_model`
- **THEN** the generated definition SHALL omit the router-model marker
