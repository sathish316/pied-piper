# Overview

Pied Piper is a team of AI SubAgents that can autonomously or semi-autonomously work on long-running coding tasks with full End-to-end tracking and human-in-the-loop approvals. 

These SubAgents run on Coding Agents (like Claude Code), Docker, Cloud Desktop etc, so they can work even while you're AFK or sleeping or on vacation.

# Getting started

Go to [QUICKSTART.md](QUICKSTART.md) for detailed steps to configure your team of SubAgents.

1. Install Pied-Piper
2. Create or configure a team
3. Generate SubAgents for your Coding Agent

# Development

Pied-piper is built using Go.

To get started with development, you can use the following commands:

1. Clone the repo

```bash
mkdir -p $GOPATH/src/github.com/
cd $GOPATH/src/github.com/sathish316
git clone https://github.com/sathish316/pied-piper.git
```

2. Run the project without building

```bash
go run main.go
```

3. Run tests

```bash
go test -v ./...
```

4. Build and Install pied-piper

4.1 Build
```bash
go build
```

4.2 Install
```bash
go install github.com/sathish316/pied-piper
```

4.3 Run Pied-Piper from anywhere
```bash
pied-piper
```

# SDLC Workflow 

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

Subagents have access to a wiki for markdown docs and knowledge management.
Subagents will create separate commits for code and docs.

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


## Generate and Use SubAgents from Claude Code for SDLC Workflow

Pied-Piper is not directly used by your Coding Agent. It gets out of the way after the SubAgents are created and configured in your coding agents. SubAgents can be generated in the home directory or project directory.

**Pied-piper CLI documentation**

#### help
```bash
$ pied-piper help
```

#### create-team
To create a default SDLC team with the name pied-piper

```bash
$ pied-piper team create --default
```

#### create-team
Create your custom team with the name pied-piper

```bash
$ pied-piper team create --name "pied-piper"
```

To create team for a given playbook, you can use the following command:
TODO: Implement
```bash
$ pied-piper team create --name "pied-piper" --playbook "microservice-to-monolith"
```

#### show-team
```bash
$ pied-piper team show --name "pied-piper"
```

teams/pied-piper/team-config.yml

```yml
name: "pied-piper"
subagents:
  - role: "architect"
    nickname: "Richard"
  - role: "software-engineer"
    nickname: "Gilfoyle"
  - role: "code-reviewer"
    nickname: "Dinesh"
  - role: "code-validator"
    nickname: "Erlich"
  - role: "build-engineer"
    nickname: "Jian Yang"
task_workflow:

```

#### create-subagent
If you've already updated subagents in team-config.yml, you can skip this step. Adding a subagent through CLI will update the config file.

```bash
$ pied-piper subagent create --team-name "pied-piper" --role "architect" --nickname "Richard"
```

#### show-subagent
TODO: Implement show by nickname
```bash
$ pied-piper subagent show --team-name "pied-piper" --role "architect"
$ pied-piper subagent show --team-name "pied-piper" --role "architect" --nickname "Richard"
```

teams/pied-piper/subagents/architect.yml
```yml
name: "architect"
role: "architect"
nickname: "Richard"
description: "..."
system_prompt: "..."
tools: default # configure in coding CLI
task_labels:
  incoming:
  - @ready-for-hld
  outgoing:
  - @ready-for-lld
wiki_labels:
  incoming:
  - GOAL_foo.md
  outgoing:
  - @ready-for-hld
  - @ready-for-lld
  - @plan-complete
  - @closed
```

#### customize an individual subagent

You can customize subagents either before they are generated into Coding CLI or after they are generated.

To change subagent before generation, edit **teams/<team-name>/subagents/<subagent-name>.yml** file ex: **teams/pied-piper/subagents/architect.yml** file

To change subagent after generation, directly update the Subagents in Claude or Coding CLI.

**.claude/subagents/<subagent-name>.yml** file

#### Generate SubAgents into Coding CLI

To generate all SubAgents for a team to Claude Code User directory:

```bash
$ pied-piper subagent generate --team-name "pied-piper" --all --target claude-code
```

To generate all SubAgents for a team to Claude Code Project directory:

```bash
$ pied-piper subagent generate --team-name "pied-piper" --all --target claude-code --target-dir /path/to/project
```

To generate or update an individual subagent into Claude Code User directory:

```bash
$ pied-piper subagent generate --team-name "pied-piper" --role "architect" --target claude-code
```

To generate or update an individual subagent into Claude Code Project directory:

```bash
$ pied-piper subagent generate --team-name "pied-piper" --role "architect" --target claude-code --target-dir /path/to/project
```

### How to use SubAgents from other Coding CLIs for SDLC Workflow?

Follow the same steps as above. While generating the subagents, change --target to your Coding CLI. 

Supported Coding CLIs are:
* Claude Code

## Building a feature with SubAgents

TODO: Add example of building a feature with SubAgents

# Playbooks

Playbooks are repeatable workflows for different kinds of long-running or continuous-coding tasks in software engineering, that can be executed by a team of Pied-Piper SubAgents:
1. Migration from library version x to version y - Rails 5 to Rails 8
2. Migration from language x to language y - Python to Typescript
4. Ensure Unit Test coverage is > 80%
5. Ensure Integration and Behavioural Test coverage is > 80%
6. Consolidate Microservices to Monolith
7. Change Techstack from x to y
8. Fix static code analysis violations in the codebase

You can find these playbooks in the [docs/PLAYBOOKS.MD](docs/PLAYBOOKS.MD)

# Release

TODO: Release to homebrew

# LICENSE

TODO: Add LICENSE

