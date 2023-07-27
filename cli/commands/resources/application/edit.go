package application

import (
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	applicationI18n "github.com/taubyte/tau-cli/i18n/application"
	applicationLib "github.com/taubyte/tau-cli/lib/application"
	applicationPrompts "github.com/taubyte/tau-cli/prompts/application"
	applicationTable "github.com/taubyte/tau-cli/table/application"
	"github.com/urfave/cli/v2"
)

func (link) Edit() common.Command {
	return common.Create(
		&cli.Command{
			Flags: []cli.Flag{
				flags.Description,
				flags.Tags,
				flags.Select,
				flags.Yes,
			},
			Action: edit,
		},
	)
}

func edit(ctx *cli.Context) error {
	// If --select is set we should not check the user's currently selected application
	checkEnv := ctx.Bool(flags.Select.Name) == false

	application, err := applicationPrompts.GetOrSelect(ctx, checkEnv)
	if err != nil {
		return err
	}

	applicationPrompts.Edit(ctx, application)

	confirm := applicationTable.Confirm(ctx, application, applicationPrompts.EditThis)
	if confirm == true {

		err = applicationLib.Set(application)
		if err != nil {
			return err
		}
		applicationI18n.Edited(application.Name)

		return nil
	}

	return nil
}
