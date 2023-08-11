package repositoryCommands

import (
	repositoryI18n "github.com/taubyte/tau-cli/i18n/repository"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
	"github.com/urfave/cli/v2"
)

func (lib *repositoryCommands) Import(ctx *cli.Context) error {
	profile, err := loginLib.GetSelectedProfile()
	if err != nil {
		return err
	}

	resource, err := lib.PromptsGetOrSelect(ctx)
	if err != nil {
		return err
	}

	auth, err := authClient.Load()
	if err != nil {
		return err
	}

	repoName := resource.Get().RepoName()
	if err = auth.RegisterRepository(resource.Get().RepoID()); err != nil {
		return repositoryI18n.RegisteringRepositoryFailed(repoName, err)
	}

	repositoryI18n.Imported(repoName, profile.Network)
	repositoryI18n.TriggerBuild()

	return nil
}
