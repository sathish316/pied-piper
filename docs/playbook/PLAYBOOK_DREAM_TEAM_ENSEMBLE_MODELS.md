# Pied Piper - Dream Team of Planner/Coder/Reviewer using Ensemble of the best Coding Models

Create a dream team of Planner/Coder/Reviewer using an ensemble of the best Coding Models GPT-5.1 or 5.2 Codex, Claude Opus 4.5, Gemini Pro 3.0:
* dream-team-planner: Plan to build X
* dream-team-coder: Build X
* dream-team-code-reviewer: Review code for X with the plan
* dream-team-performance-and-security-reviewer: Review code for X for performance and security issues

Use the Ensemble of the best Coding Models to build the feature X:
* GPT Codex 5.1 for Planning
* Claude Opus 4.5 for Coding
* Gemini Pro 3.0 for Code, Performance, Security Reviews

## Custom Workflow for Dream Team of Planner/Coder/Reviewers using an Ensemble of the best Coding Models

<placeholder-for-image>

This playbook demonstrates a systematic approach to creating a dream team of Planner/Coder/Reviewers to build any major/minor feature X using an Ensemble of the best Coding Models:
* Roles & Responsibilities
    * dream-team-planner: Plan to build X
    * dream-team-coder: Build X
    * dream-team-code-reviewer: Review code for X with the plan
    * dream-team-performance-and-security-reviewer: Review code for X for performance and security issues
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

## Prerequisites to make Claude Code work with OpenRouter

Claude Code can only use the following models as of now - Sonnet, Opus, Haiku.
To make Claude Code work with the best of OpenAI GPT-5.2 Codex and Gemini Pro 3.0, we will be using OpenRouter and make Claude Code use OpenRouter by using Claude-code-router: https://github.com/musistudio/claude-code-router

1. Install Claude-code-router

```bash
npm install -g @musistudio/claude-code-router
```

2. Add OPENROUTER_API_KEY to env
```bash
export OPENROUTER_API_KEY=<your-openrouter-api-key>
```

3. Configure Claude Code Router (CCR) with OpenRouter models:

Refer claude-code-router docs to add multiple OpenRouter models to Claude-Code. Sample config for using GPT-5.1 Codex or GPT-5.2, Opus 4.5, Gemini Pro 3.0 all from within Claude-Code is below:

File: ~/.config/claude-code-router/config.json
```json
{
  "HOST": "127.0.0.1",
  "PORT": 3456,
  "APIKEY": "",
  "API_TIMEOUT_MS": "600000",
  "Providers": [
    {
      "name": "openrouter",
      "api_base_url": "https://openrouter.ai/api/v1/chat/completions",
      "api_key": "<your-openrouter-api-key>",
      "models": [
        "openai/gpt-5.1-codex",
        "openai/gpt-5.1-codex-max",
        "openai/gpt-5.1-codex-mini",
        "openai/gpt-5.2",
        "openai/gpt-5.2-pro",
        "anthropic/claude-opus-4.5",
        "anthropic/claude-sonnet-4.5",
        "google/gemini-3-pro",
        "google/gemini-3-pro-preview"
      ]
    }
  ],
  "Router": {
    "default": "openrouter,anthropic/claude-opus-4.5",
    "think": "openrouter,openai/gpt-5.2",
    "longContext": "openrouter,google/gemini-3-pro-preview",
    "longContextThreshold": 60000,
  },
}
```

CCR allows using a different OpenRouter model per Subagent by specifying the following in Subagent.md file:
https://github.com/musistudio/claude-code-router?tab=readme-ov-file#subagent-routing

```md
<CCR-SUBAGENT-MODEL>openrouter,openai/gpt-5.2</CCR-SUBAGENT-MODEL>
```

## Pied-Piper Installation

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

Following is a trimmed down version of Pied-Piper Dream team workflow using Multiple Claude Models. For full config, go to [playbook/dream_team_ensemble.yml](../../playbook/dream_team/dream_team_ensemble.yml) and copy-paste it into your team config file.

```yml
name: "dream-team"
description: "A dream team of Planner/Coder/Reviewers to build major/minor features in your projects using Claude Models"
subagents:
  - role: "dream-team-planner"
    router_model: "openrouter,openai/gpt-5.2"
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

### 4. Select a different RouterModel per subagent

Let's use different models per subagent to finetune their behaviour:
* dream-team-planner: GPT-5.1 Codex or GPT-5.2 Codex
* dream-team-coder: Claude Opus 4.5
* dream-team-code-reviewer: Gemini Pro 3.0 or Gemini Pro 2.5
* dream-team-performance-and-security-reviewer: Gemini Pro 3.0 or Gemini Pro 2.5
* dream-team-orchestrator: Claude Haiku 4.5

Note: There are some integration issues between Gemini 3.0 Pro Preview and the tools used here. Due to this, rest of the doc will use Gemini 2.5 Pro wherever Gemini Pro 3.0 will be a better fit.

Change the config in ~/.pied-piper/dream-team/config.yml file to use the different models.
```yml
subagents:
  - name: dream-team-planner
    router_model: "openrouter,openai/gpt-5.1-codex"
  - name: dream-team-coder
    router_model: "openrouter,anthropic/claude-opus-4.5"
  - name: dream-team-code-reviewer
    router_model: "openrouter,google/gemini-3-pro-preview"
  - name: dream-team-performance-and-security-reviewer
    router_model: "openrouter,google/gemini-3-pro-preview"
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

Go to cursor or claudecode and use the above metaprompt to update each Subagent file Ex: **/path/to/project/.subagents/python-programmer.yml**

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
