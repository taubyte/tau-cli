package application

import (
	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	applicationI18n "github.com/taubyte/tau/i18n/application"
	applicationLib "github.com/taubyte/tau/lib/application"
	applicationPrompts "github.com/taubyte/tau/prompts/application"
	applicationTable "github.com/taubyte/tau/table/application"
	"github.com/urfave/cli/v2"
)

func (link) Delete() common.Command {
	return common.Create(
		&cli.Command{
			Flags: []cli.Flag{
				flags.Select,
				flags.Yes,
			},
			Action: delete,
		},
	)
}

func delete(ctx *cli.Context) error {
	// If --select is set we should not check the user's currently selected application
	checkEnv := ctx.Bool(flags.Select.Name) == false

	application, err := applicationPrompts.GetOrSelect(ctx, checkEnv)
	if err != nil {
		return err
	}

	confirm := applicationTable.Confirm(ctx, application, applicationPrompts.DeleteThis)
	if confirm == true {
		err = applicationLib.Delete(application)
		if err != nil {
			return err
		}
		applicationI18n.Deleted(application.Name)

		return nil
	}

	return nil
}
