package prompts

import (
	"github.com/taubyte/tau/flags"
	"github.com/taubyte/tau/validate"
	"github.com/urfave/cli/v2"
)

func GetOrAskForADescription(c *cli.Context, prev ...string) string {
	return validateString(c, validateStringHelper{
		field:     flags.Description.Name,
		prompt:    DescriptionPrompt,
		prev:      prev,
		validator: validate.VariableDescriptionValidator,
	})
}
