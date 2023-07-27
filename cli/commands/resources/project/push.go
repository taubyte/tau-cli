package project

import (
	gosimplegit "github.com/taubyte/go-simple-git"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	projectFlags "github.com/taubyte/tau-cli/flags/project"
	projectI18n "github.com/taubyte/tau-cli/i18n/project"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/prompts"
	projectPrompts "github.com/taubyte/tau-cli/prompts/project"
	"github.com/urfave/cli/v2"
)

func (link) Push() common.Command {
	return common.Create(&cli.Command{
		Flags: []cli.Flag{
			flags.CommitMessage,
			projectFlags.ConfigOnly,
			projectFlags.CodeOnly,
		},
		Action: push,
	})
}

func push(ctx *cli.Context) error {
	project, err := projectPrompts.GetOrSelect(ctx, true)
	if err != nil {
		return err
	}

	repoHandler, err := projectLib.Repository(project.Name).Open()
	if err != nil {
		return err
	}

	commitMessage := prompts.GetOrRequireACommitMessage(ctx)

	err = (&dualRepoHandler{
		ctx:         ctx,
		repository:  repoHandler,
		projectName: project.Name,
		errorFormat: projectI18n.PullingProjectFailed,
		action: func(r *gosimplegit.Repository) error {
			err = r.Commit(commitMessage, ".")
			if err != nil {
				return err
			}

			return r.Push()
		},
	}).Run()
	if err != nil {
		return err
	}

	projectI18n.PushedProject(project.Name)
	return nil
}
