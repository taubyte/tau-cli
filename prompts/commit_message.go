package prompts

import (
	"github.com/taubyte/tau/flags"
	"github.com/taubyte/tau/validate"
	"github.com/urfave/cli/v2"
)

func GetOrRequireACommitMessage(c *cli.Context) (ret string) {
	return validateAndRequireString(c, validateRequiredStringHelper{
		field:     flags.CommitMessage.Name,
		prompt:    CommitMessagePrompt,
		validator: validate.VariableDescriptionValidator,
	})
}
