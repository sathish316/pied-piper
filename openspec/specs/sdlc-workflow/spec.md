# SDLC Workflow Specification

## Purpose
Define how Pied Piper coordinates role-based AI subagents through planning, implementation, validation, and delivery while retaining human oversight.

## Requirements

### Requirement: Role-based SDLC team
The system SHALL define distinct subagent roles for microsprint orchestration, product management, architecture, software engineering, code review, code validation, and build engineering, with a human engineer acting as reviewer.

#### Scenario: Assign work by responsibility
- **WHEN** a microsprint advances through requirements, design, implementation, review, validation, and delivery
- **THEN** the system SHALL route each stage to the role responsible for that stage

### Requirement: Microsprint issue decomposition
The system SHALL represent a feature as a human-created feature issue plus planning and build issues coordinated within a microsprint.

#### Scenario: Start work on a feature
- **WHEN** a human engineer creates a feature issue with the `@open` label
- **THEN** the microsprint orchestrator SHALL coordinate a planning issue marked `@ready-for-plan` and a build issue marked `@ready-for-dev`

### Requirement: Automatic workflow progression
The system SHALL support automatic progression through planning, making, and shipping using the configured task labels.

#### Scenario: Complete the automatic happy path
- **WHEN** a feature is processed without rejection or validation failure
- **THEN** its work SHALL progress from requirement definition through high-level and low-level design, code review, code validation, merge readiness, and closure

#### Scenario: Return rejected code for correction
- **WHEN** code review emits `@code-review-rejected`
- **THEN** the software engineer SHALL receive the work for correction before another code review

#### Scenario: Retry failed validation
- **WHEN** validation emits `@code-validation-failed`
- **THEN** the work SHALL return to code validation until it succeeds

### Requirement: Human approval workflow
The system SHALL support a human-approval mode in which a human engineer can approve or reject planned and implemented work.

#### Scenario: Reject a plan
- **WHEN** the human engineer rejects a plan under review
- **THEN** the work SHALL return to requirement definition, high-level design, or low-level design as directed

#### Scenario: Reject an implementation
- **WHEN** the human engineer rejects implemented work under review
- **THEN** the work SHALL return to development

### Requirement: Role-specific knowledge artifacts
The system SHALL associate each SDLC role with the task, code, review, or wiki artifacts it consumes and produces.

#### Scenario: Produce planning artifacts
- **WHEN** a feature advances through product and architecture planning
- **THEN** the product manager SHALL produce a requirement artifact and the architect SHALL consume it to produce a high-level design artifact

#### Scenario: Produce implementation artifacts
- **WHEN** the software engineer receives design artifacts and a build task
- **THEN** the software engineer SHALL produce low-level design and code artifacts for review and validation
