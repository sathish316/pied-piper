# Overview of Pied-Piper Subagents and SDLC Workflow Features

## SDLC Workflow 

Example of a typical SDLC Workflow using a Team of Pied-Piper SubAgents:
1. Receive feature request
2. Plan the implementation (architecture, problem breakdown, acceptance criteria)
3. Write code (in a feature branch)
4. Review code for correctness and adherence to plan
5. Review code for performance and security issues
6. Merge feature branch to main
7. Rinse and repeat for the next feature or bug fix

The creative work here occurs in both planning what to build and actually building it. Deciding what to build next requires a human. Verifying what is planned/built is what is needed requires a human to review the plan and code. Pied-Piper has human-in-the-loop approvals to involve humans at the right stage before approving the plan or before creating Pull requests.

The role of Pied-Piper SubAgents is to automate long-running or repetitive tasks like improving test coverage, performing repetitive migrations etc or sometimes even coding full features or bug fixes. Check the playbooks for examples of different types of workflows.

## SubAgent Roles & Responsibilities

Example of SubAgent Roles in Pied-Piper are:
* sdlc-orchestrator (Orchestrator): Orchestrates a team of subagents to build feature X, assigns tasks to subagents, coordinates the workflow
* sdlc-coder (Maker): Builds feature X based on the implementation plan, writes code, implements features
* sdlc-code-reviewer (Checker): Reviews code for X against the plan - checks correctness, adherence to requirements, and code quality
* sdlc-performance-and-security-reviewer (Checker): Reviews code for X for performance bottlenecks and security vulnerabilities
* human-engineer (Reviewer): Reviews the plan, code, and pull-request

Pied-Piper uses Maker-Checker and Boomerang workflow patterns for the SubAgents to review the code before handing it over to the human-engineer for approval.

Please note the roles are just examples and can be anything that suits a particular complex or long-running SDLC workflow.

## SubAgent Task Management

This project uses [beads](https://github.com/steveyegge/beads) as the Task management layer for both Agents and Humans.

To know all about beads run (Agents also run the same command to know all about beads):
```
$ bd quickstart
```

SubAgents use git-flow for every new feature
```
$ git-flow help
```

## SubAgent SDLC Workflow

SDLC Workflow is executed in every microsprint. Microsprint is like an epoch in the SDLC Workflow.

Beads tasks and task labels are used to manage the lifecycle of the tasks.
* feature-x: this task is created by human-engineer or planner with the label #feature
* plan-feature-x: this task moves through planning phase
* build-feature-x: this task moves through coding and review phases

Issues go through the following lifecycle:
* planning
    * #sdlc-feature -> #sdlc-plan-complete
* making
    * #sdlc-plan-complete -> #sdlc-code-complete
    * #sdlc-code-complete -> #sdlc-code-review-approved
    * #sdlc-code-complete -> #sdlc-code-review-rejected -> #sdlc-code-complete
    * #sdlc-code-review-approved -> #sdlc-performance-and-security-review-approved
    * #sdlc-code-review-approved -> #sdlc-performance-and-security-review-rejected -> #sdlc-code-complete
* shipping
    * #sdlc-performance-and-security-review-approved -> #closed

Please note the task labels and workflows are just examples and can be anything that suits a particular complex or long-running SDLC workflow.

## SubAgent Task Management/Workflow

Subagents need an Issue/Task management system that acts as both Agent layer and Human layer for managing long-running tasks.

Subagent are designed to respond to specific incoming and outgoing Task labels to participate in the SDLC Workflow.

SubAgents use [beads](https://github.com/steveyegge/beads) as the Task management layer for both Agents and Humans. Beads can sync to Jira or other Project management tools if needed.

Example of Incoming/Outgoing Task labels:
* sdlc-orchestrator
    * incoming: #sdlc-feature
    * outgoing: #closed
* sdlc-planner
    * incoming: #sdlc-feature
    * outgoing: #sdlc-plan-complete
* sdlc-coder
    * incoming: #sdlc-plan-complete, #sdlc-code-review-rejected, #sdlc-performance-and-security-review-rejected
    * outgoing: #sdlc-code-complete
* sdlc-code-reviewer
    * incoming: #sdlc-code-complete
    * outgoing: #sdlc-code-review-approved, #sdlc-code-review-rejected
* sdlc-performance-and-security-reviewer
    * incoming: #sdlc-code-review-approved
    * outgoing: #sdlc-performance-and-security-review-approved, #sdlc-performance-and-security-review-rejected

## SubAgent Knowledge Management

Subagents need a knowledge management system that acts as both Agent layer and Human layer for documenting and managing long-running tasks.

Subagents have access to local markdown docs, which act as wikis for knowledge management.

Each Subagent only reads and creates the files it is allowed to in each microsprint:
* sdlc-orchestrator:
    * incoming: []
    * outgoing: []
* sdlc-planner:
    * incoming: #sdlc-feature task
    * outgoing: SDLC_PLAN_<TASK_ID>.md (implementation plan with feature overview, component breakdown, acceptance criteria)
* sdlc-coder:
    * incoming: SDLC_PLAN_<TASK_ID>.md
    * outgoing: []
* sdlc-code-reviewer:
    * incoming: SDLC_PLAN_<TASK_ID>.md
    * outgoing: Comments in task, Status update in task
* sdlc-performance-and-security-reviewer:
    * incoming: SDLC_PLAN_<TASK_ID>.md
    * outgoing: Comments in task, Status update in task

## SubAgent Nicknames

SubAgents can have nicknames. You can change the nicknames after the SubAgents are generated. The SubAgents will respond to both roles and nicknames. In case you need to create multiple SubAgents of the same role, giving them different nicknames helps.

* sdlc-orchestrator: "Erlich"
* sdlc-planner: "Richard"
* sdlc-coder: "Dinesh"
* sdlc-code-reviewer: "Gavin"
* sdlc-performance-and-security-reviewer: "Gilfoyle"

## SubAgent Model Selection

This is an optional step using which you can select a different Claude model or a non-Anthropic model for each SubAgent. 

Each SubAgent can use a different Claude model optimized for its task:
* sdlc-planner: opus (best for complex planning and analysis)
* sdlc-orchestrator: haiku (fast coordination)
* sdlc-coder: sonnet (optimized for code generation)
* sdlc-code-reviewer: sonnet (efficient for code review)
* sdlc-performance-and-security-reviewer: opus (thorough analysis for security/performance)

See the playbook [PLAYBOOK_DREAM_TEAM_CLAUDE_MODELS.md](PLAYBOOK_DREAM_TEAM_CLAUDE_MODELS.md) or [PLAYBOOK_DREAM_TEAM_ENSEMBLE_MODELS.md](PLAYBOOK_DREAM_TEAM_ENSEMBLE_MODELS.md) for instrucitons on how to select different models per SubAgent.
