package repositoryCommands

import (
	"github.com/taubyte/tau-cli/i18n"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	repositoryLib "github.com/taubyte/tau-cli/lib/repository"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/urfave/cli/v2"
)

func (lib *repositoryCommands) selectResource(ctx *cli.Context) (project config.Project, resource Resource, info *repositoryLib.Info, err error) {
	project, err = projectLib.SelectedProjectConfig()
	if err != nil {
		return
	}

	// Confirm project is cloned
	_, err = project.Interface()
	if err != nil {
		i18n.Help().BeSureToCloneProject()
		return
	}

	resource, err = lib.PromptsGetOrSelect(ctx)
	if err != nil {
		return
	}

	info = &repositoryLib.Info{
		FullName: resource.Get().RepoName(),
		ID:       resource.Get().RepoID(),
		Type:     lib.Type,
		DoClone:  true,
	}

	return
}
