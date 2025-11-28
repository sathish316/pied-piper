# Pied Piper - Quick Start Guide

Get your team of AI SubAgents up and running in minutes in ClaudeCode or your favourite Coding CLI to work on a custom workflow.

## Custom Workflow for Unit test coverage

<placeholder: image for custom workflow>

As an example, let's use the following workflow:
* Roles & Responsibilities
    * unittest-orchestrator: start a unittest microsprint, end the microsprint, defines goal of microsprint
    * unittest-planner: The goal is to write unit tests for a file or a package. Unittest-planner breaks down the goal into smaller tasks, one per file.
    * unittest-programmer: Unittest-tester writes the unit tests for the files or packages. Unlike AI coding tools generating a mountain of unit tests and markdown files, unittest-programmer writes only 1 happy path test case and one or more edge cases per public function. This is enough for side-projects. Unittest-programmer also does not annoyingly modify/remove your code to make the tests pass.
    * unittest-reviewer: unittest-reviewer reviews the unit tests for the files or packages. If it finds issues, it rejects the test. If the test fails, it fixes the test. If both the code and tests are good, it approves the test.
    * unittest-summarizer: Generate a summary of tests implemented, concisely.
* Task management workflow
    * unittest-orchestrator
        * Incoming: #unittest
        * Outgoing: #unittest-plan
    * unittest-planner
        * Incoming: #unittest-plan
        * Outgoing: #unittest-ready-for-dev
    * unittest-programmer
        * Incoming: #unittest-ready-for-dev
        * Outgoing: #unittest-ready-for-review
    * unittest-reviewer
        * Incoming: #unittest-ready-for-review, #unittest-rejected
        * Outgoing: #unittest-approved, #unittest-rejected
    * unittest-summarizer
        * Incoming: #unittest-approved
        * Outgoing: #closed
* Wiki workflow
    * unittest-orchestrator
        * Incoming: -
        * Outgoing: UT_GOAL_<TASK_ID>.md
    * unittest-planner
        * Incoming: UT_GOAL_<TASK_ID>.md
        * Outgoing: UT_PLAN_<TASK_ID>.md
    * unittest-programmer
        * Incoming: UT_PLAN_<TASK_ID>.md
        * Outgoing: -
    * unittest-reviewer
        * Incoming: -
        * Outgoing: -
    * unittest-summarizer
        * Incoming: -
        * Outgoing: UT_SUMMARY_<TASK_ID>.md

## Installation

```bash
go install github.com/sathish316/pied-piper
pied-piper help
```

## Quick Start

### 1. Create Your Team and Add SubAgents

```bash
pied-piper team create --name "test-titans"
pied-piper subagent create --team-name "test-titans" --role "unittest-orchestrator" --nickname "Mike"
pied-piper subagent create --team-name "test-titans" --role "unittest-planner" --nickname "Peter"
pied-piper subagent create --team-name "test-titans" --role "unittest-programmer" --nickname "Tim"
pied-piper subagent create --team-name "test-titans" --role "unittest-reviewer" --nickname "Richard"
pied-piper subagent create --team-name "test-titans" --role "unittest-summarizer" --nickname "Sam"
```

### 2. View Your Team

```bash
# Show full team config
pied-piper team show --name "test-titans"

# List all SubAgents
pied-piper subagent list --team "test-titans"

# Show specific SubAgent
pied-piper subagent show --team "test-titans" --name "unittest-programmer"
```

### 3. Edit your team's workflow

```bash
vim ~/.pied-piper/test-titans/team-config.yml
``` 

```yml
name: "test-titans"
subagents:
  - role: "unittest-orchestrator"
    description: "Start a unittest microsprint, end the microsprint, defines goal of microsprint"
    nickname: "Mike"
    task_labels:
      incoming:
      - "#unittest"
      outgoing:
      - "#unittest-plan"
    wiki_labels:
      incoming: []
      outgoing:
      - "UT_GOAL_<TASK_ID>.md"
  - role: "unittest-planner"
    description: "The goal is to write unit tests for a file or a package. Unittest-planner breaks down the goal into smaller tasks, one per file."
    nickname: "Peter"
    task_labels:
      incoming:
      - "#unittest-plan"
      outgoing:
      - "#unittest-ready-for-dev"
    wiki_labels:
      incoming:
      - "UT_GOAL_<TASK_ID>.md"
      outgoing:
      - "UT_PLAN_<TASK_ID>.md"
  - role: "unittest-programmer"
    description: "Unittest-programmer writes the unit tests for the files or packages. Unittest-programmer writes only 1 happy path test case and one or more edge cases per public function. This is enough for side-projects. Unittest-programmer also does not modify/remove your code to make the tests pass."
    nickname: "Tim"
    task_labels:
      incoming:
      - "#unittest-ready-for-dev"
      outgoing:
      - "#unittest-ready-for-review"
    wiki_labels:
      incoming:
      - "UT_PLAN_<TASK_ID>.md"
      outgoing: []
  - role: "unittest-reviewer"
    description: "Unittest-reviewer reviews the unit tests for the files or packages. If it finds issues, it rejects the test. If the test fails, it fixes the test. If both the code and tests are good, it approves the test."
    nickname: "Richard"
    task_labels:
      incoming:
      - "#unittest-ready-for-review"
      - "#unittest-rejected"
      outgoing:
      - "#unittest-approved"
      - "#unittest-rejected"
      - "#unittest-ready-for-review"
    wiki_labels:
      incoming: []
      outgoing: []
  - role: "unittest-summarizer"
    nickname: "Sam"
    description: "Generate a summary of tests implemented, concisely. Keep it short, don't blabber"
    task_labels:
      incoming:
      - "#unittest-approved"
      outgoing:
      - "#closed"
    wiki_labels:
      incoming: []
      outgoing:
      - "UT_SUMMARY_<TASK_ID>.md"
```

List all SubAgents to verify the team config is correct.
```bash
pied-piper subagent list --team "test-titans"
```


### 3. Generate SubAgents for Claude Code

**Generate all SubAgents (global):**
```bash
pied-piper subagent generate-all --team "test-titans" --target claude-code
```

**Generate all SubAgents (project-specific):**
```bash
pied-piper subagent generate-all --team "test-titans" --target claude-code --target-dir /path/to/project
```

**Generate single SubAgent:**
```bash
pied-piper subagent generate --team "test-titans" --name "unittest-programmer" --target claude-code
```

### 4. Modify and Regenerate

**Edit team config:**
```bash
vim ~/.pied-piper/teams/test-titans/team-config.yml
```

**Regenerate SubAgent:**
```bash
pied-piper subagent generate --team "test-titans" --name "unittest-programmer" --target claude-code
```

### 5. Modify SubAgent's role description from Claude Code or your editor

```bash
vim /path/to/project/.subagents/unittest-programmer.yml
```

Modify the description to suit your project's needs.
```yml
----ROLE_DESCRIPTION STARTS----
This project uses python and the test framework is pytest.
We use underscore for test names. We use test setup and teardown methods if the file contains more than one test.
Keep the tests short and precise.
----ROLE_DESCRIPTION ENDS----
```

There is no need to regenerate the SubAgent, since you are directly editing in Claude folder.


### 6. Refine the prompt of subagents as required.

Generate metaprompt for each Subagent:
```bash
pied-piper subagent metaprompt
```

Go to cursor and use the prompt to update each agent:

### 7. Start improving Test coverage for your project in Claude Code using SubAgents

Once generated, use SubAgents in Claude Code with role or nickname or `@` mentions:
- `@unittest-orchestrator` or `@Mike`

$ bd create "test file foo.go"
$ bd create "test all go files in package foo"

> Ask unittest-orchestrator to start a microsprint
> Run "bd quickstart". List ready tasks
> Ask unittest-orchestrator to start task abc

### 8. Update memory of SubAgents

> Tell unittest-planner to keep the UT_PLAN_<TASK_ID>.md brief and concise with one line per file and public method to be tested. add a list of bullet points, one line per scenario. Update its memory in "MEMORY" section.

```bash
vim /path/to/project/.subagents/unittest-planner.yml
```
