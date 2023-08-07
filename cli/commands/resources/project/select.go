package project

import (
	"github.com/taubyte/tau-cli/cli/common"
	projectI18n "github.com/taubyte/tau-cli/i18n/project"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	projectPrompts "github.com/taubyte/tau-cli/prompts/project"
	"github.com/urfave/cli/v2"
)

func (link) Select() common.Command {
	return common.Create(
		&cli.Command{
			Action: _select,
		},
	)
}

func _select(ctx *cli.Context) error {
	project, deselect, err := projectPrompts.GetSelectOrDeselect(ctx)
	if err != nil {
		return err
	}

	if deselect {
		err = projectLib.Deselect(ctx, project.Name)
		if err != nil {
			return err
		}
		projectI18n.DeselectedProject(project.Name)
	} else {
		err = projectLib.Select(ctx, project.Name)
		if err != nil {
			return err
		}
		projectI18n.SelectedProject(project.Name)
		checkProjectClonedHelp(project.Name)
	}

	return nil
}
