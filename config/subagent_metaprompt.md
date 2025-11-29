Read <subagent.md> to understand the roles and responsibilities of <subagent>. Update the following sections of <subagent.md> file:
* ROLE_DESCRIPTION
* TASK_WORKFLOW_DESCRIPTION
* DOCUMENTATION_WORKFLOW_DESCRIPTION

Do not modify any text outside these three sections.

Update the ROLE_DESCRIPTION section to update the roles and responsibilities of the Agent in less than 300 words. Expand the high level roles and responsibilities of the Agent to update this section. Tell the agent how to identify themselves as <subagent-role> or their nickname.

In this project we use beads task management system. Update the TASK_WORKFLOW_DESCRIPTION section to tell the agent how to use beads task management system. Tell the agent to onboard to beads by running the bash command "bd quickstart". Tell the agent to only work on tasks assigned to them using Incoming labels of this Subagent. Tell the agent how to use bd to assign tasks to others by updating the labels with the Outgoing labels of this Subagent. Don't use more than 300-500 words for this section.

If the role is orchestrator or planner, tell the agent to run autonomously, poll for "bd ready" tasks with the orchestrator's incoming label and start working on them.

Give the agent instructions on what git tools they can use, only if applicable. In this project we use git-flow for feature branching. If the agent has a responsibility to write code, tell the agent to run the bash command "git-flow" to learn how to use git-flow, create feature branches, commit to feature branch. Tell the agent to create small commits to the feature branch. Tell the agent to create meaningful concise commit messages.

In this project, we use local wiki/markdown files for knowledge management. These are stored in wiki subdirectory. Update the DOCUMENTATION_WORKFLOW_DESCRIPTION section to tell the agent how to use wiki system. An agent can only create markdown files that are intended for it. An agent can only generate markdown files that it's allowed to generate according to the rules in High level Documentation workflow. Don't use more than 300-500 words for this section.

Give the agent instructions on what incoming wiki/markdown files to read before performing its task. Give this instruction only if applicable for the role of the agent.

Give the agent insructions on what outgoing wiki/markdown files to generate after performing its task. Give this instruction only if applicable for the role of the agent.

Tell the agent to update MEMORY section if the user gives new instructions to be remembered.