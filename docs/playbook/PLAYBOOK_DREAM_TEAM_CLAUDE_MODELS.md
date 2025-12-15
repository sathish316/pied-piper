# Pied Piper - Dream Team of Planner/Coder/Reviewer using Claude Models

Create a dream team of Planner/Coder/Reviewer using Claude Models:
* dream-team-planner: Plan to build X
* dream-team-coder: Build X
* dream-team-code-reviewer: Review code for X with the plan
* dream-team-performance-and-security-reviewer: Review code for X for performance and security issues
* dream-team-orchestrator: Orchestrate the dream team to build X

## Custom Workflow for Dream Team of Planner/Coder/Reviewers using Claude Models

<placeholder-for-image>

This playbook demonstrates a systematic approach to creating a dream team of Planner/Coder/Reviewers to build any major/minor feature X using Claude Models:
* Roles & Responsibilities
    * dream-team-planner: Plan to build X
    * dream-team-coder: Build X
    * dream-team-code-reviewer: Review code for X with the plan
    * dream-team-performance-and-security-reviewer: Review code for X for performance and security issues
    * dream-team-orchestrator: Orchestrate the dream team to build X
* Task management workflow
    * dream-team-planner
        * Incoming: #dream-team-feature
        * Outgoing: #dream-team-feature-plan-complete
    * dream-team-coder
        * Incoming: #dream-team-feature-plan-complete, #dream-team-feature-code-review-rejected, #dream-team-feature-performance-security-review-rejected
        * Outgoing: #dream-team-feature-code-complete
    * dream-team-code-reviewer
        * Incoming: #dream-team-feature-code-complete
        * Outgoing: #dream-team-feature-code-review-approved, #dream-team-feature-code-review-rejected
    * dream-team-performance-and-security-reviewer
        * Incoming: #dream-team-feature-code-review-approved
        * Outgoing: #dream-team-feature-performance-security-review-approved, #dream-team-feature-performance-security-review-rejected
    * dream-team-orchestrator
        * Incoming: #dream-team-feature
        * Outgoing: #closed

## Installation

```bash
go install github.com/sathish316/pied-piper
pied-piper help
```

If you're working from dev-env, you can replace "pied-piper" with "go run main.go" in below commands.

## Quick Start

### 1. Create Your Team and Add SubAgents

```bash
pied-piper team create --name "dream-team"

pied-piper subagent create --team "dream-team" --role "dream-team-planner" --nickname "Richard"

pied-piper subagent create --team "dream-team" --role "dream-team-coder" --nickname "Dinesh"

pied-piper subagent create --team "dream-team" --role "dream-team-code-reviewer" --nickname "Gavin"

pied-piper subagent create --team "dream-team" --role "dream-team-performance-and-security-reviewer" --nickname "Gilfoyle"

pied-piper subagent create --team "dream-team" --role "dream-team-orchestrator" --nickname "Erlich"

```

### 2. View Your Team members

```bash
# Show full team config
pied-piper team show --name dream-team

# List all SubAgents
pied-piper subagent list --team dream-team

# Show specific SubAgent
pied-piper subagent show --team dream-team --name dream-team-planner
```

### 3. Edit your team's workflow

```bash
vim ~/.pied-piper/dream-team/config.yml
``` 

Following is a trimmed down version of Pied-Piper Dream team workflow using Multiple Claude Models. For full config, go to [playbook/dream_team/dream_team_claude.yml](../../playbook/dream_team/dream_team_claude.yml) and copy-paste it into your team config file.

```yml
name: "dream-team"
description: "A dream team of Planner/Coder/Reviewers to build major/minor features in your projects using Claude Models"
subagents:
  - role: "dream-team-planner"
    model: "opus"
    description: |
      Plans the implementation of feature X. Analyzes requirements, breaks down the feature into implementable tasks, identifies dependencies, and creates a detailed implementation plan. Keep the plan concise with clear acceptance criteria per task. Do not use more than 5-6 sections and more than 500-600 words for the plan.
    nickname: "Richard"
    task_labels:
      incoming:
        - "#dream-team-feature"
      outgoing:
        - "#dream-team-plan-complete"
      task_workflow_description: |
        1. Dream-team-planner receives beads tasks with #dream-team-feature label containing the feature to build.
        It analyzes the feature requirements, identifies components to build, dependencies, and creates a detailed implementation plan.
        For complex features, it breaks down into smaller subtasks with clear acceptance criteria.
        After the plan is complete, it creates a wiki file with the implementation plan and updates the task label to #dream-team-plan-complete.
        The dream-team-coder will pick up tasks with #dream-team-plan-complete label.
    wiki_labels:
      incoming: []
      outgoing:
        - "DREAM_TEAM_PLAN_<TASK_ID>.md"
      wiki_workflow_description: |
        Wikis are created as local markdown files in "wiki" directory.
        1. Dream-team-planner receives beads tasks with #dream-team-feature label.
        Once it has analyzed the requirements and created an implementation plan, it creates a local wiki file called "DREAM_TEAM_PLAN_<TASK_ID>.md" with:
        - Feature overview and goals
        - Component breakdown with dependencies
        - Implementation steps
        - Acceptance criteria for each component
        - Technical considerations and edge cases
...
```

List all SubAgents to verify the team config is correct.
```bash
pied-piper subagent list --team dream-team
```

### 4. Select a different Model per subagent

Let's use different models per subagent to finetune their behaviour:
* dream-team-planner: opus
* dream-team-coder: sonnet
* dream-team-code-reviewer: sonnet
* dream-team-performance-and-security-reviewer: sonnet
* dream-team-orchestrator: haiku

Change the config in ~/.pied-piper/dream-team/config.yml file to use the different models.
```yml
subagents:
  - name: dream-team-planner
    model: opus
  - name: dream-team-coder
    model: sonnet
  - name: dream-team-code-reviewer
    model: sonnet
  - name: dream-team-performance-and-security-reviewer
    model: opus
  - name: dream-team-orchestrator
    model: haiku
```

### 5. Generate SubAgents for Claude Code and Export to Claude Directory

SubAgents are generated for a target CodingAgent in ~/.pied-piper/dream-team/subagents directory.

**Generate all SubAgents (global):**
```bash
pied-piper subagent generate-all --team dream-team --target claude-code
```

**Generate single SubAgent:**
```bash
pied-piper subagent generate --team dream-team --name dream-team-planner --target claude-code
```

**Export All SubAgents to Coding CLI to a project directory:**
```bash
pied-piper export all --team dream-team --target claude-code --project-dir /path/to/project
```

**Export single SubAgent to Coding CLI to a project directory:**
```bash
pied-piper export subagent --team "dream-team" --name "dream-team-planner" --target claude-code --project-dir /path/to/project
```

### 6. Enrich Subagent Description

Metaprompts are used to enrich the prompt of Subagent using AI tools like Cursor or Claude Code or by directly calling LLM APIs.

In order to make sure the Subagent honors the workflow, Enrich the prompt of Subagent using AI tools like Cursor or Claude Code or LLM APIs by following the steps below:

5.1 Generate or show metaprompt (applicable to all subagents):
```bash
pied-piper subagent metaprompt
```

Go to cursor or claudecode and use the above metaprompt to update each Subagent file Ex: **/path/to/project/.claude/agents/dream-team-code-reviewer.md**

### (Optional) 7. Modify Workflow and Regenerate Subagents

If you need to change the team workflow, Modify and Regenerate the Subagents

**Edit team config:**
```bash
vim ~/.pied-piper/dream-team/config.yml
```

**Regenerate SubAgent:**
```bash
pied-piper subagent generate --team "dream-team" --name "dream-team-code-reviewer" --target claude-code
```

**Export single SubAgent to Coding CLI to a target directory:**
```bash
pied-piper subagent export --team "dream-team" --name "dream-team-code-reviewer" --target claude-code --target-dir /path/to/project
```

### (Optional) 8. Modify SubAgent's role description from Claude Code or your editor

Once the SubAgent is generated and exported to Coding CLI, you can generate detailed workflow description and modify its behaviour using AI.

```bash
vim /path/to/project/.claude/agents/dream-team-code-reviewer.md
```

Modify the description to suit your project's needs.
```yml
----ROLE_DESCRIPTION STARTS----
<Generated Role Description>
<Project-specific custom Role description>
----ROLE_DESCRIPTION ENDS----
```

There is no need to regenerate the SubAgent, since you are directly editing in .claude/agents folder. 

### 9. Build feature X using the Dream Team of 3 best Coding Models

Initialize beads in project

```bash
projec-dir> bd init
```

If this is the first time you are using SubAgents and beads, run the following prompts in Claude-code:
> In this project, we use beads for task management. Run bash command "bd quickstart" to onboard to beads task management system

Ask Claude-code to build feature X using the Dream team subagents
> Use the dream-team subagents to implement Feature X

Watch the Pied-Piper Dream Team implement Feature X either semi-autonomously or autonomously, keeping you (human-reviewer) in the loop.
