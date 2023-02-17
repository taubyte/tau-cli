package websitePrompts

import (
	"errors"
	"fmt"
	"strings"

	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/flags"
	websiteI18n "github.com/taubyte/tau/i18n/website"
	websiteLib "github.com/taubyte/tau/lib/website"
	"github.com/taubyte/tau/prompts"

	"github.com/urfave/cli/v2"
)

/*
GetOrSelect will try to get the website from a name flag
if it is not set in the flag it will offer a selection menu
*/
func GetOrSelect(ctx *cli.Context) (*structureSpec.Website, error) {
	name := ctx.String(flags.Name.Name)

	resources, err := websiteLib.ListResources()
	if err != nil {
		return nil, err
	}

	// Try to select a website
	if len(name) == 0 && len(resources) > 0 {
		options := make([]string, len(resources))
		for idx, p := range resources {
			options[idx] = p.Name
		}

		name, err = prompts.SelectInterface(options, SelectPrompt, options[0])
		if err != nil {
			return nil, websiteI18n.SelectPromptFailed(err)
		}
	}

	if len(name) != 0 {
		LName := strings.ToLower(name)
		for _, website := range resources {
			if LName == strings.ToLower(website.Name) {
				return website, nil
			}
		}

		return nil, fmt.Errorf(NotFound, name)
	}

	return nil, errors.New(NoneFound)
}
