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

func (link) New() common.Command {
	return common.Create(
		&cli.Command{
			Flags: []cli.Flag{
				flags.Description,
				flags.Tags,
				flags.Yes,
			},
			Action: new,
		},
	)
}

func new(ctx *cli.Context) error {
	application, err := applicationPrompts.New(ctx)
	if err != nil {
		return err
	}

	name := application.Name

	confirm := applicationTable.Confirm(ctx, application, applicationPrompts.CreateThis)
	if confirm {
		err := applicationLib.New(application)
		if err != nil {
			return err
		}
		applicationI18n.Created(name)

		err = applicationLib.Select(ctx, name)
		if err != nil {
			return err
		}
		applicationI18n.Selected(name)

		return nil
	}

	return nil
}
