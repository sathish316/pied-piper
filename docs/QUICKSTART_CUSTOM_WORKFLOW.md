# Pied Piper - Quick Start Guide

Get your team of AI SubAgents up and running in minutes in ClaudeCode or your favourite Coding CLI to work on a custom workflow.

## Custom Workflow for Unit test coverage

<img src="assets/pied_piper_unittest_coverage_workflow.png" alt="Unit Test Coverage Playbook"/>

As an example, let's use the following workflow:
* Roles & Responsibilities
    * unittest-planner: The goal is to write unit tests for a file or a package. Unittest-planner breaks down the goal into smaller tasks, one per file.
    * unittest-programmer: Unittest-programmer writes the unit tests for the files or packages. Unlike AI coding tools generating a mountain of unit tests and markdown files, unittest-programmer writes only 1 happy path test case and one or more edge cases per public function. Unittest-programmer also does not annoyingly modify/remove your code to make the tests pass.
    * unittest-reviewer: unittest-reviewer reviews the unit tests for the files or packages. If it finds issues, it rejects the test. If the test fails, it fixes the test. If both the code and tests are good, it approves the test.
    * unittest-summarizer: Generate a summary of tests implemented, concisely.
* Task management workflow
    * unittest-planner
        * Incoming: #unittest
        * Outgoing: #unittest-plan-complete
    * unittest-programmer
        * Incoming: #unittest-ready-for-dev, #unittest-rejected
        * Outgoing: #unittest-ready-for-review
    * unittest-reviewer
        * Incoming: #unittest-ready-for-review
        * Outgoing: #unittest-approved, #unittest-rejected
    * unittest-summarizer
        * Incoming: #unittest-approved
        * Outgoing: closed
* Wiki workflow
    * unittest-planner
        * Incoming: []
        * Outgoing: UT_GOAL_<TASK_ID>.md
    * unittest-programmer
        * Incoming: UT_PLAN_<TASK_ID>.md
        * Outgoing: []
    * unittest-reviewer
        * Incoming: UT_PLAN_<TASK_ID>.md
        * Outgoing: []
    * unittest-summarizer
        * Incoming: []
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
#TODO: add subagents to team config
pied-piper subagent create --team "test-titans" --role "unittest-planner" --nickname "Peter"
pied-piper subagent create --team "test-titans" --role "unittest-programmer" --nickname "Tim"
pied-piper subagent create --team "test-titans" --role "unittest-reviewer" --nickname "Richard"
pied-piper subagent create --team "test-titans" --role "unittest-summarizer" --nickname "Sam"
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
  - role: "unittest-planner"
    description: |
        Start a unittest plan, end the unittest cycle. Unittest-planner breaks down the goal into smaller tasks, one per file or one per 5 public functions. Keep the goal short in less than 10 lines and 50 words.
    nickname: "Peter"
    task_labels:
      incoming:
      - "#unittest"
      - "#unittest-approved"
      outgoing:
      - "#unittest-plan-complete"
      - "#closed"
      task_workflow_description: |
        1. Unittest-planner receives beads tasks with #unittest label.
        It breaks down the goal into smaller tasks, one per file or one per 5 public functions.
        After the plan is complete, it comments the plan on beads task and updates the label of the task to #unittest-plan-complete.
        If the task is ready for implementation, update the label of the task to #unittest-ready-for-dev.
        If there are multiple complex testing tasks needed, create beads tasks that are linked to the original task with #unittest-ready-for-dev label.
        Tasks with $unittest-ready-for-dev label will be picked up by unittest-programmer.
        2. Unittest-planner receives beads tasks with #unittest-approved label.
        It's job is to close the beads task. It updates the label of the task to #closed.
        It takes the summary of the tests from unittest-summarizer and updates the beads task with the summary.
    wiki_labels:
      incoming: []
      outgoing:
      - "UT_GOAL_<TASK_ID>.md"
      wiki_workflow_description: |
        Wikis are created as local markdown files in "wiki" directory.
        1. Unittest-planner receives beads tasks with #unittest label.
        Once it has finalized a test plan, it creates a local wiki file called "UT_GOAL_<TASK_ID>.md" with the plan
  - role: "unittest-programmer"
    description: |
        Unittest-programmer writes the unit tests for the files or packages. Unittest-programmer writes only 1 happy path test case and one or more edge cases per public function. Unittest-programmer also does not modify/remove your code to make the tests pass.
    nickname: "Tim"
    task_labels:
      incoming:
      - "#unittest-ready-for-dev"
      - "#unittest-rejected"
      outgoing:
      - "#unittest-ready-for-review"
      task_workflow_description: |
        Unittest-programmer follows a boomerang workflow pattern with Unittest-reviewer. They can exchange tasks for a maximum of 3 times.
        This follows a Maker-Checker pattern where Unittest-programmer is the maker and Unittest-reviewer is the checker. Once all tests are written and the beads task is closed, Human reviewer is the final Checker.
        1. Unittest-programmer receives beads tasks with #unittest-ready-for-dev label. Once it has written the unit tests, verified those tests pass, it updates the label of the task to #unittest-ready-for-review.
        2. Unittest-programmer receives beads tasks with #unittest-rejected label. These tasks are rejected by unittest-reviewer with review comments as comments on the beads task. Unittest-programmer addresses the comments and modifies the unit tests accordingly, verified those tests pass, and then updates the label of the task to #unittest-ready-for-review.
    wiki_labels:
      incoming:
      - "UT_GOAL_<TASK_ID>.md"
      outgoing: []
      wiki_workflow_description: |
        Wikis are created as local markdown files in "wiki" directory.
        Unittest-programmer does not generate any wiki or markdown docs.
        1. Before starting on a task with #unittest-ready-for-dev label, Unittest-programmer reads the local wiki file "UT_GOAL_<TASK_ID>.md" to understand the goal of the test task and then proceeds to write the test.
  - role: "unittest-reviewer"
    description: |
        Unittest-reviewer reviews the unit tests for the files or packages. If it finds issues, it rejects the test. If the test fails, it fixes the test. If both the code and tests are good, it approves the test.
    nickname: "Richard"
    task_labels:
      incoming:
      - "#unittest-ready-for-review"
      outgoing:
      - "#unittest-approved"
      - "#unittest-rejected"
      task_workflow_description: |
        Unittest-reviewer follows a boomerang workflow pattern with Unittest-programmer. They can exchange tasks for a maximum of 3 times.
        This follows a Maker-Checker pattern where Unittest-programmer is the maker and Unittest-reviewer is the checker. Once all tests are written and the beads task is closed, Human reviewer is the final Checker.
        1. Unittest-reviewer receives beads tasks with #unittest-ready-for-review label. It reviews the unit tests for the files or packages. 
            a. If it finds issues, it rejects the test. It updates review comments in the beads task as comments. It rejects the test by updating the label of the task to #unittest-rejected.
            b. If both the code and tests are good, it approves the test. It approves the test by updating the label of the task to #unittest-approved.
    wiki_labels:
      incoming:
      - "UT_GOAL_<TASK_ID>.md"
      outgoing: []
      wiki_workflow_description: |
        Wikis are created as local markdown files in "wiki" directory.
        Unittest-reviewer does not generate any wiki or markdown docs.
        1. Before starting on a task with #unittest-ready-for-review label, Unittest-reviewer reads the local wiki file "UT_GOAL_<TASK_ID>.md" to understand the goal of the test task and then proceeds to review the test.
  - role: "unittest-summarizer"
    nickname: "Sam"
    description: "Generate a crisp summary of tests implemented, concisely. Keep the summary short in less than 10 lines and 500 words"
    task_labels:
      incoming:
      - "#unittest-approved"
      outgoing:
      - "closed"
      task_workflow_description: |
        1. Unittest-summarizer receives beads tasks with #unittest-approved label.
        It generates a summary of the tests implemented, concisely in a local wiki doc.
        Once the summary is generated, it updates the label of the task to #closed and hands-off control to unittest-orchestrator.
    wiki_labels:
      incoming: []
      outgoing:
      - "UT_SUMMARY_<TASK_ID>.md"
      wiki_workflow_description: |
        Wikis are created as local markdown files in "wiki" directory.
        unittest-summarizer looks at all comments in the beads task to generate the summary of tests implemented.
        The summary is concise and short in less than 10 lines and 500 words.
        Unittest-summarizer creates a local wiki file called "UT_SUMMARY_<TASK_ID>.md" with the summary of the tests implemented.
```

List all SubAgents to verify the team config is correct.
```bash
pied-piper subagent list --team "test-titans"
```


### 3. Generate SubAgents for Claude Code

SubAgents are generated for a target CodingAgent in ~/.pied-piper/test-titans/subagents directory.

**Generate all SubAgents (global):**
```bash
pied-piper subagent generate-all --team "test-titans" --target claude-code
```

**Generate single SubAgent:**
```bash
pied-piper subagent generate --team "test-titans" --name "unittest-programmer" --target claude-code
```

**Export SubAgents to Coding CLI:**
```bash
pied-piper subagent export --team "test-titans" --target claude-code
```

**Export single SubAgent to Coding CLI:**
```bash
pied-piper subagent export --team "test-titans" --name "unittest-programmer" --target claude-code
```

### 4. If you need to change the team workflow, Modify and Regenerate the Subagents

**Edit team config:**
```bash
vim ~/.pied-piper/teams/test-titans/team-config.yml
```

**Regenerate SubAgent:**
```bash
pied-piper subagent generate --team "test-titans" --name "unittest-programmer" --target claude-code
```

**Export single SubAgent to Coding CLI:**
```bash
pied-piper subagent export --team "test-titans" --name "unittest-programmer" --target claude-code
```

### 5. Modify SubAgent's role description from Claude Code or your editor

Once the SubAgent is generated and exported to Coding CLI, you can generate detailed workflow description and modify its behaviour using AI.

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

There is no need to regenerate the SubAgent, since you are directly editing in .claude folder. 

In order to make sure the Subagent honors the workflow, Refine the prompt of Subagent using AI tools like Cursor or Claude Code or by directly calling LLM APIs.

Generate metaprompt for each Subagent:
```bash
pied-piper subagent metaprompt
```

Go to cursor and use the prompt to update each Subagent file in **/path/to/project/.subagents/unittest-programmer.yml**

### 6. Start assigning work to your team of SubAgents

If this is the first time you are using SubAgents and beads, run the following prompts in Claude-code:
> Run bash command "bd quickstart" to onboard to beads task management system

```bash
$ bd create "add unit tests for file foo.go" --label "#unittest"
```

> Ask unittest-planner to plan and write unit tests for open tasks in beads

OR

> Use the subagents unittest-planner, unittest-programmer, unittest-reviewer, unittest-summarizer to run a complete unit testing workflow for the open task in beads

OR 

> Work on the open unit testing task in beads. Use the Subagents available for unit testing

Update Role description or Memory of each SubAgent to finetune its behaviour.