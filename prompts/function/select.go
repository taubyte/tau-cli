package functionPrompts

import (
	"errors"
	"fmt"
	"strings"

	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/flags"
	functionI18n "github.com/taubyte/tau-cli/i18n/function"
	functionLib "github.com/taubyte/tau-cli/lib/function"
	"github.com/taubyte/tau-cli/prompts"

	"github.com/urfave/cli/v2"
)

/*
GetOrSelect will try to get the function from a name flag
if it is not set in the flag it will offer a selection menu
*/
func GetOrSelect(ctx *cli.Context) (*structureSpec.Function, error) {
	name := ctx.String(flags.Name.Name)

	resources, err := functionLib.ListResources()
	if err != nil {
		return nil, err
	}

	// Try to select a function
	if len(name) == 0 && len(resources) > 0 {
		options := make([]string, len(resources))
		for idx, p := range resources {
			options[idx] = p.Name
		}

		name, err = prompts.SelectInterface(options, SelectPrompt, options[0])
		if err != nil {
			return nil, functionI18n.SelectPromptFailed(err)
		}
	}

	if len(name) != 0 {
		nameLC := strings.ToLower(name)
		for _, function := range resources {
			if nameLC == strings.ToLower(function.Name) {
				return function, nil
			}
		}

		return nil, fmt.Errorf(NotFound, name)
	}

	return nil, errors.New(NoneFound)
}
