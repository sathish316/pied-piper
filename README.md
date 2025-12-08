# Overview

Pied Piper is a team of AI SubAgents that can autonomously or semi-autonomously work on long-running coding tasks with full End-to-end tracking and human-in-the-loop approvals. 

These SubAgents run on Coding Agents (like Claude Code), Docker, Cloud Desktop etc, so they can work even while you're AFK or sleeping or on vacation.

# Getting started

1. Install Pied-Piper
2. Create or configure a team
3. Generate SubAgents for your Coding Agent
4. Start assigning tasks to your SubAgents from Coding Agent (Claude Code)

Go to [QUICKSTART_CUSTOM_WORKFLOW.md](docs/QUICKSTART_CUSTOM_WORKFLOW.md) or [PLAYBOOK_TEST_COVERAGE.md](docs/playbook/PLAYBOOK_TEST_COVERAGE.md) to run a custom SDLC worfklow for Test coverage improvement using Pied-Piper.

Go to [PLAYBOOK_LANGUAGE_MIGRATION.md](docs/playbook/PLAYBOOK_LANGUAGE_MIGRATION.md) to run a custom SDLC worfklow for Language Migration from TypeScript to Python using Pied-Piper.

Go to Pied-Piper Commands section for docs on how to use Pied-Piper.

For an overview of Pied-Piper SDLC Workflow features like Subagents, Roles, Task workflows, Wiki workflows, Role Nicknames, go to [PIEDPIPER_SDLC_WORKFLOW_CONCEPTS.md](docs/PIEDPIPER_SDLC_WORKFLOW_CONCEPTS.md)

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

# Playbooks

Playbooks are repeatable workflows for different kinds of long-running or continuous-coding or boring tasks in software engineering, that can be executed by a team of Pied-Piper SubAgents:
1. Migration from library version x to version y - Rails 5 to Rails 8
2. Migration from language x to language y - Typescript to Python
4. Ensure Unit Test coverage is > 80%
5. Ensure Integration and Behavioural Test coverage is > 80%
6. Consolidate Microservices to Monolith
7. Change Tech Stack from x to y
8. Fix static code analysis violations in the codebase

You can find these playbooks in the [docs/PLAYBOOKS.MD](docs/PLAYBOOKS.MD)

### Language migration playbook

<img src="docs/assets/language_migration_playbook.png"/>

This is a sample Language migration playbook from Typescript to Python, using a team of Pied-Piper SubAgents.

Go to [PLAYBOOK_LANGUAGE_MIGRATION.md](docs/playbook/PLAYBOOK_LANGUAGE_MIGRATION.md) for detailed steps to run the language migration playbook using Pied-Piper.

Demo video (Claude Code):
TODO: LINK

### Unit test coverage playbook

<img src="docs/assets/unit_test_coverage_playbook.png"/>

This is a sample Unit test coverage improvement playbook, using a team of Pied-Piper SubAgents.

Go to [PLAYBOOK_TEST_COVERAGE.md](docs/playbook/PLAYBOOK_TEST_COVERAGE.md) for detailed steps to run the unit test coverage improvement playbook using Pied-Piper.

Demo video (Claude Code):
TODO: LINK

### Library version migration playbook

<TODO>

### Integration/Behavioural test coverage playbook

<TODO>

### Microservices to Monolith consolidation playbook

<TODO>

### Techstack migration playbook

<TODO>

### Static code analysis violation fix playbook

<TODO>

# Pied-Piper Commands

## Generate and Use SubAgents from Claude Code for SDLC Workflow

Pied-Piper generates SubAgents (*.md files) from simple specs that can be used from other Coding CLIs like Claude Code. 

SubAgents can be generated in User home directory or Project directory for each Coding Agent.

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

#### show-team
```bash
$ pied-piper team show --name "pied-piper"
```

File: **~/.pied-piper/pied-piper/config.yml**

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
```bash
$ pied-piper subagent show --team-name "pied-piper" --role "architect"
$ pied-piper subagent show --team-name "pied-piper" --role "architect" --nickname "Richard"
```

File: **~/.pied-piper/pied-piper/subagents/architect.yml
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

#### Generate SubAgents for a Coding CLI

Subagents can be generated in *.md format to target multiple Coding CLIs.

To generate all SubAgents for a team to target Claude Code:

```bash
$ pied-piper subagent generate --team-name "pied-piper" --all --target claude-code
```

#### Export SubAgents for a Coding CLI

Subagents can be exported in *.md format to the User directory (~/.claude) or Project directory (/path/to/project/.claude) for a target Coding CLI.

To export all SubAgents for a team to target Claude Code Project directory:

```bash
$ pied-piper subagent export all --team-name "pied-piper" --target claude-code --project-dir /path/to/project
```

To export an individual subagent to target Claude Code Project directory:

```bash
$ pied-piper subagent export subagent --team-name "pied-piper" --name "architect" --target claude-code --project-dir /path/to/project
```

### How to use SubAgents workflow from Claude Code?

Go to Claude Code:
1. > Onboard to beads task management using "bd quickstart"
2. > Create a new task with the label that starts your workflow
3. > Ask microsprint-orchestrator to work on beads open tasks

For more detailed steps, refer to the playbooks:
1. [PLAYBOOK_TEST_COVERAGE.md](docs/playbook/PLAYBOOK_TEST_COVERAGE.md)
2. [PLAYBOOK_LANGUAGE_MIGRATION.md](docs/playbook/PLAYBOOK_LANGUAGE_MIGRATION.md)

### How to use SubAgents from other Coding CLIs for SDLC Workflow?

Follow the same steps as above. While generating the subagents, change --target to your Coding CLI. 

Supported Coding CLIs are:
* Claude Code

# Release

TODO: Release to homebrew

# LICENSE

TODO: Add LICENSE

