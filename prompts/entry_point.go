package prompts

import (
	"github.com/taubyte/tau-cli/flags"
	"github.com/taubyte/tau-cli/validate"
	"github.com/urfave/cli/v2"
)

func GetOrRequireAnEntryPoint(c *cli.Context, prev ...string) string {
	return validateAndRequireString(c, validateRequiredStringHelper{
		field:  flags.EntryPoint.Name,
		prompt: EntryPointPrompt,
		prev:   prev,

		// TODO better validator
		validator: validate.VariableDescriptionValidator,
	})
}
