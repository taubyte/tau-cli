package repositoryCommands

import (
	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	"github.com/taubyte/tau/prompts"
	"github.com/urfave/cli/v2"
)

func (lib *repositoryCommands) PullCmd() common.Command {
	return common.Create(
		&cli.Command{
			Flags: []cli.Flag{
				flags.Branch,
			},
			Action: lib.Pull,
		},
	)
}

func (lib *repositoryCommands) Pull(ctx *cli.Context) error {
	project, resource, info, err := lib.selectResource(ctx)
	if err != nil {
		return err
	}

	branch := prompts.GetOrRequireABranch(ctx, resource.Get().Branch())

	_, err = info.Pull(project, branch, resource.Get().RepositoryURL())
	if err != nil {
		return err
	}
	lib.I18nPulled(resource.Get().RepositoryURL())

	return nil
}
