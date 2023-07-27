package repositoryCommands

import resources "github.com/taubyte/tau-cli/cli/commands/resources/common"

func InitCommand(new *LibCommands) Commands {
	resources.PanicIfMissingValue(new)

	return &repositoryCommands{new}
}
