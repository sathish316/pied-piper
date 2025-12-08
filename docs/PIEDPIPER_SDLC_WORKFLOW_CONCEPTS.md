# Overview of Pied-Piper SDLC Workflow features

## SDLC Workflow 

A typical SDLC workflow to implement a small feature or fix a bug consists of the following steps:
1. Requirement and Acceptance criteria
2. High level design or architecture
3. Low level design - Data model, API signatures etc
4. Write code (in a feature branch)
5. Write tests
6. Run unit tests to validate the implementation
7. Review code
8. Run unit and integration tests to validate the implementation
9. Create a Pull request for review
10. Merge feature branch to main
11. Rinse and repeat for the next feature or bug fix

The creative work here occurs in both planning what to build and actually building it. Deciding what to build next requires a human. Verifying what is planned/built is what is needed requires a human to review the plan and code.

The role of Pied-Piper SubAgents is to automate the repetitive tasks like writing unit tests, integration tests etc or sometimes even coding small features or bug fixes.

## SubAgent Roles & Responsibilities

<placeholder: image for roles>

The default Roles in Pied-Piper are:
* microsprint-orchestrator (Orchestrator): starts microsprint, assigns tasks to subagents, runs the microsprint either autonomously or semi-autonomously, ends microsprint
* product-manager (Planner): defines requirements and acceptance criteria
* architect (Planner): creates plan with high level design and architecture
* software-engineer (Maker): creates plan with low level design and data model, create feature branch,write code, write unit tests, run unit tests
* code-reviewer (Checker): review code in a git commmit, review code in a pull request
* code-validator (Checker): run unit and integration tests to validate the implementation, run quick-check or property based tests if required to validate the implementation
* build-engineer: creates pull request for review, merge feature branch to main if human-engineer approves
* human-engineer (Reviewer): reviews the plan, code, pull-request

## SubAgent Task Management

TODO: link to beads
This project uses beads for Task management by both Agents and Humans.

To know all about beads run:
```
$ bd quickstart
```

SubAgents use git-flow method for every new feature
```
$ git-flow help
```

## SubAgent SDLC Workflow

<placeholder: image for workflow>

SDLC Workflow is executed in every microsprint.
Microsprint is like an epoch in the SDLC Workflow.

2 Issues are created for a task.
* feature-x: this task is created by human-engineer with the label @open
* plan-feature-x: this task is created by microsprint-orchestrator with the label @ready-for-plan
* build-feature-x: this task is created by microsprint-orchestrator with the label @ready-for-dev

Issues go through the following lifecycle (in auto-approve mode):
* planning
    * @open / @ready-for-plan -> @define-requirement
    * @define-requirement -> @ready-for-hld
    * @ready-for-hld -> @ready-for-lld
    * @ready-for-lld -> @plan-complete -> @closed
* making
    * @open / @ready-for-dev -> @coding-done -> @ready-for-code-review
    * @ready-for-code-review -> @code-review-done -> @ready-for-code-validation
    * @ready-for-code-review -> @code-review-rejected -> @ready-for-code-review
    * @ready-for-code-validation -> @code-validation-done -> @ready-for-merge
    * @ready-for-code-validation -> @code-validation-failed -> @ready-for-code-validation
* shipping
    * @ready-for-merge -> @closed


Issues go through the following lifecycle (in human-approve mode):
* planning
    * @ready-for-lld -> @review-plan -> @approve-plan -> @closed
    * @ready-for-lld -> @review-plan -> @reject-plan -> (@define-requirement | @ready-for-hld | @ready-for-lld)
* making
    * @ready-for-build -> @review-make -> @approve-make -> @closed
    * @ready-for-build -> @review-make -> @reject-make -> @ready-for-dev
* shipping
    * @ready-for-merge -> @closed

## SubAgent Task management/workflow

Subagents need an Issue/Task management system that acts as both Agent layer and Human layer for managing long-running tasks.

Each Subagent responds to the following incoming and outgoing labels:
* microsprint-orchestrator
incoming: @open, @ready-for-plan, @approve-plan, @reject-plan, @approve-make, @reject-make
outgoing: @define-requirement
* product-manager
incoming: @define-requirement
outgoing: @ready-for-hld
* architect
incoming: @ready-for-hld
outgoing: @ready-for-lld
* software-engineer
incoming: @ready-for-lld, @ready-for-dev, @code-review-rejected
outgoing: @ready-for-code-review
* code-reviewer
incoming: @ready-for-code-review
outgoing: @code-review-done, @code-review-rejected
* code-validator
incoming: @ready-for-code-validation, @code-validation-failed
outgoing: @code-validation-done
* build-engineer
incoming: @ready-for-merge
outgoing: @closed
* human-engineer
incoming: @review-plan, @review-make
outgoing: @approve-plan, @approve-make, @reject-plan, @reject-make

## SubAgent Knowledge Management

Subagents need a knowledge management system that acts as both Agent layer and Human layer for managing long-running tasks.

Subagents have access to a wiki for markdown docs and knowledge management.

Each Subagent only reads and creates the following files in each microsprint. Other than docs, SubAgents also deal with a few other artifacts like Tasks, Code etc::
* microsprint-orchestrator:
incoming: plan-feature-x task
outgoing: GOAL_foo.md (where foo is the feature id or task id), build-feature-x task
* product-manager:
incoming: -
outgoing: REQUIREMENT_foo.md (where foo is the feature id or task id)
* architect:
incoming: REQUIREMENT_foo.md
outgoing: HLD_foo.md
* software-engineer:
incoming: HLD_foo.md, LLD_foo.md, build-feature-x task
outgoing: LLD_foo.md, Code
* code-reviewer:
incoming: git_sha..git_sha, Code,
outgoing: Comments in build-feature-x task, Future: Comments in github commit
* code-validator:
incoming: git_sha
outgoing: -
* build-engineer:
incoming: -
outgoing: -
* human-engineer:
incoming: -
outgoing: feature-x task for one or more tasks with specifications. Features can be configured to do plan + build, build only, plan only, bugfixes can be build + test only. By default do both plan + build.

## SubAgent Nicknames

SubAgents have nicknames from the fictional PiedPiper company. You can change the nicknames after the SubAgents are generated. The SubAgents will respond to both roles and nicknames. In case you need to create multiple SubAgents of the same role, giving them different nicknames helps.

* microsprint-orchestrator: "Pied-Piper"
* product-manager: "Jared"
* architect: "Richard"
* software-engineer: "Gilfoyle"
* code-reviewer: "Dinesh"
* code-validator: "Erlich"
* build-engineer: "Jian Yang"
* human-engineer: <Your name>
