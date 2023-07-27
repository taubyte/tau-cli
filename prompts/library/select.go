package libraryPrompts

import (
	"errors"
	"fmt"
	"strings"

	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/flags"
	libraryI18n "github.com/taubyte/tau-cli/i18n/library"
	libraryLib "github.com/taubyte/tau-cli/lib/library"
	"github.com/taubyte/tau-cli/prompts"

	"github.com/urfave/cli/v2"
)

/*
GetOrSelect will try to get the library from a name flag
if it is not set in the flag it will offer a selection menu
*/
func GetOrSelect(ctx *cli.Context) (*structureSpec.Library, error) {
	name := ctx.String(flags.Name.Name)

	resources, err := libraryLib.ListResources()
	if err != nil {
		return nil, err
	}

	// Try to select a library
	if len(name) == 0 && len(resources) > 0 {
		options := make([]string, len(resources))
		for idx, p := range resources {
			options[idx] = p.Name
		}

		name, err = prompts.SelectInterface(options, SelectPrompt, options[0])
		if err != nil {
			return nil, libraryI18n.SelectPromptFailed(err)
		}
	}

	if len(name) != 0 {
		LName := strings.ToLower(name)
		for _, library := range resources {
			if LName == strings.ToLower(library.Name) {
				return library, nil
			}
		}

		return nil, fmt.Errorf(NotFound, name)
	}

	return nil, errors.New(NoneFound)
}
