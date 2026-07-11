# Workflow Playbooks Specification

## Purpose
Define the ready-to-customize subagent teams and workflows Pied Piper provides for recurring software-development activities.

## Requirements

### Requirement: Playbook team definition
The system SHALL provide playbooks as YAML team configurations containing named roles, role descriptions, nicknames, task-label workflows, and documentation workflows.

#### Scenario: Customize a playbook
- **WHEN** a user selects a bundled playbook as a starting point
- **THEN** the playbook SHALL supply a complete team definition that can be copied and customized for the user's project

### Requirement: Language migration playbook
The system SHALL provide a TypeScript-to-Python migration team that plans dependency order, migrates one file at a time, reviews migrated code and tests, and supports human approval.

#### Scenario: Migrate a file through review
- **WHEN** the migration planner schedules a TypeScript file for migration
- **THEN** a file planner SHALL produce a conversion plan, a Python programmer SHALL implement code and tests without modifying the TypeScript source, and a reviewer SHALL approve or reject the result

#### Scenario: Correct a rejected migration
- **WHEN** the Python reviewer rejects a migrated file
- **THEN** the Python programmer SHALL address the review comments and return the file for review

### Requirement: Test coverage playbook
The system SHALL provide a test-coverage team that plans unit-test work, writes tests, reviews them, summarizes the result, and closes approved tasks.

#### Scenario: Add tests using maker-checker review
- **WHEN** a task is labeled for unit-test work
- **THEN** the planner SHALL break down the work, the programmer SHALL implement tests, and the reviewer SHALL approve or reject them through task labels

### Requirement: Dream team feature playbook
The system SHALL provide a feature-development team with planning, coding, code review, performance and security review, and orchestration roles.

#### Scenario: Build and review a feature
- **WHEN** a dream-team feature enters the workflow
- **THEN** it SHALL progress through planning, coding, correctness review, performance and security review, human approval when configured, and closure

#### Scenario: Return rejected work to the coder
- **WHEN** either review role rejects the implementation
- **THEN** the coder SHALL address the feedback and resubmit the work

### Requirement: Dream team model strategies
The system SHALL provide both native Claude-model and Claude Code Router ensemble variants of the dream-team playbook.

#### Scenario: Use native Claude models
- **WHEN** the native Claude variant is selected
- **THEN** roles SHALL be assignable to `haiku`, `sonnet`, or `opus` according to their workload

#### Scenario: Use an ensemble of routed models
- **WHEN** the ensemble variant is selected
- **THEN** each role SHALL be assignable to an independent provider and model through `router_model`
