package prompts

import (
	"errors"
	"strings"

	"github.com/taubyte/tau/common"
	"github.com/taubyte/tau/flags"
	"github.com/taubyte/tau/validate"
	"github.com/urfave/cli/v2"
)

// TODO parse the source, and make this a selection based on exported functions.
func GetOrRequireACall(c *cli.Context, source common.Source, prev ...string) string {
	call := validateAndRequireString(c, validateRequiredStringHelper{
		field:  flags.Call.Name,
		prompt: CallPrompt,
		prev:   prev,
		validator: func(s string) error {
			err := validate.VariableDescriptionValidator(s)
			if err != nil {
				return err
			}

			// TODO REGEX
			if strings.HasSuffix(s, ".") == true {
				return errors.New("cannot end with `.`")
			}

			// TODO REGEX
			if strings.Contains(s, " ") == true {
				return errors.New("does not except spaces")
			}

			return nil
		},
	})

	return call
}
