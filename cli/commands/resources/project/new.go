package project

import (
	"os"
	"path"

	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	projectFlags "github.com/taubyte/tau/flags/project"
	"github.com/taubyte/tau/i18n"
	projectI18n "github.com/taubyte/tau/i18n/project"
	projectLib "github.com/taubyte/tau/lib/project"
	projectPrompts "github.com/taubyte/tau/prompts/project"
	projectTable "github.com/taubyte/tau/table/project"
	"github.com/urfave/cli/v2"
)

func (link) New() common.Command {
	return common.Create(
		&cli.Command{
			Flags: []cli.Flag{
				flags.Description,
				projectFlags.Loc,
				flags.EmbedToken,
				projectFlags.Public,
				projectFlags.Private,
				flags.Yes,
			},
			Action: new,
		},
	)
}

func new(ctx *cli.Context) error {
	embedToken, project, err := projectPrompts.New(ctx)
	if err != nil {
		return err
	}

	name := project.Name

	// Check location flag, otherwise clone into cwd
	var location string
	if ctx.IsSet(projectFlags.Loc.Name) == true {
		location = ctx.String(projectFlags.Loc.Name)
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return i18n.GettingCwdFailed(err)
		}

		location = path.Join(cwd, project.Name)
	}

	if projectTable.Confirm(ctx, project, projectPrompts.CreateThisProject) == true {
		err = projectLib.New(project, location, embedToken)
		if err != nil {
			return err
		}
		projectI18n.CreatedProject(name)
		projectI18n.SelectedProject(name)

		return nil
	}

	return nil
}
