package project

import (
	"github.com/taubyte/tau-cli/cli/common"
	projectFlags "github.com/taubyte/tau-cli/flags/project"
	projectI18n "github.com/taubyte/tau-cli/i18n/project"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	projectPrompts "github.com/taubyte/tau-cli/prompts/project"
	"github.com/taubyte/tau/pkg/git"
	"github.com/urfave/cli/v2"
)

func (link) Pull() common.Command {
	return common.Create(&cli.Command{
		Flags: []cli.Flag{
			projectFlags.ConfigOnly,
			projectFlags.CodeOnly,
		},
		Action: pull,
	})
}

func pull(ctx *cli.Context) error {
	project, err := projectPrompts.GetOrSelect(ctx, true)
	if err != nil {
		return err
	}

	repoHandler, err := projectLib.Repository(project.Name).Open()
	if err != nil {
		return err
	}

	err = (&dualRepoHandler{
		ctx:         ctx,
		repository:  repoHandler,
		projectName: project.Name,
		errorFormat: projectI18n.PullingProjectFailed,
		action: func(r *git.Repository) error {
			return r.Pull()
		},
	}).Run()
	if err != nil {
		return err
	}

	projectI18n.PulledProject(project.Name)
	return nil
}
