# SDLC Workflow Specification

## Purpose
Define how Pied Piper coordinates configurable role-based AI subagents through planning, implementation, layered review, and delivery while retaining optional human oversight.

## Requirements

### Requirement: Role-based SDLC team
The system SHALL allow a workflow to assign distinct planning, making, checking, and orchestration responsibilities to configurable subagent roles.

#### Scenario: Assign work by responsibility
- **WHEN** a microsprint advances through planning, implementation, correctness review, performance and security review, and delivery
- **THEN** the system SHALL route each stage to the role responsible for that stage

#### Scenario: Use custom roles
- **WHEN** a team defines roles or labels different from the bundled examples
- **THEN** the workflow SHALL use the roles and routing declared in that team's configuration

### Requirement: Microsprint issue decomposition
The system SHALL use tasks and task labels to represent feature, planning, implementation, review, and closure state within a microsprint.

#### Scenario: Start work on a feature
- **WHEN** a human or planner creates a feature task with the team's feature label
- **THEN** the configured planner and orchestrator roles SHALL coordinate its planning and implementation work

### Requirement: Automatic workflow progression
The system SHALL support progression through planning, making, checking, and shipping using the team's configured incoming and outgoing task labels.

#### Scenario: Complete the automatic happy path
- **WHEN** a feature is processed without rejection or validation failure
- **THEN** its work SHALL progress from planning through implementation, correctness review, performance and security review, approval, and closure

#### Scenario: Return rejected code for correction
- **WHEN** a correctness reviewer rejects an implementation
- **THEN** the configured coder SHALL receive the task and review feedback before resubmitting it

#### Scenario: Return performance or security findings
- **WHEN** the performance and security reviewer rejects an implementation
- **THEN** the configured coder SHALL address the findings and return the work through review

### Requirement: Human approval workflow
The system SHALL support a semi-autonomous mode with human review of plans and reviewed code, as well as an autonomous mode without those human approval pauses.

#### Scenario: Reject a plan
- **WHEN** the human engineer rejects a plan under review
- **THEN** the work SHALL return to the planning role for revision

#### Scenario: Reject an implementation
- **WHEN** the human engineer rejects implemented work under review
- **THEN** the work SHALL return to development

#### Scenario: Run autonomously
- **WHEN** the workflow is configured for autonomous mode
- **THEN** the orchestrator SHALL progress work without waiting for human plan or implementation approval

### Requirement: Role-specific knowledge artifacts
The system SHALL associate each role with configured incoming and outgoing local Markdown artifacts that provide durable context for long-running tasks.

#### Scenario: Produce planning artifacts
- **WHEN** a planner completes a feature plan
- **THEN** it SHALL write the configured plan artifact containing the feature overview, component breakdown, dependencies, and acceptance criteria

#### Scenario: Produce implementation artifacts
- **WHEN** a coder or reviewer starts work that declares an incoming plan artifact
- **THEN** that role SHALL read the configured artifact before implementing or reviewing the feature

### Requirement: Maker-checker review loop
The system SHALL support maker-checker and boomerang workflow patterns in which review roles can return work to the implementing role with feedback.

#### Scenario: Approve maker output
- **WHEN** a checker finds the implementation correct and compliant with its review responsibility
- **THEN** it SHALL emit the configured approval label for the next workflow stage

#### Scenario: Reject maker output
- **WHEN** a checker identifies issues
- **THEN** it SHALL record feedback, emit the configured rejection label, and route the task back to the maker
