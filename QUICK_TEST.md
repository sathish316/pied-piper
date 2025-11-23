## Quick Test with default team name

### Help commands
go run main.go --help

go run main.go team --help

go run main.go subagent --help

### Create a team
go run main.go team create

### Show a team
go run main.go team show --name pied-piper

### List subagents
go run main.go subagent list

### Show a subagent
go run main.go subagent show --name software-engineer

### Generate a subagent
go run main.go subagent generate --name software-engineer --target claude-code

### Generate all subagents for Claude Code
go run main.go subagent generate-all --target claude-code

### Generate all subagents for Claude Code in project directory
go run main.go subagent generate-all --target claude-code --projectDir ~/projects/foo