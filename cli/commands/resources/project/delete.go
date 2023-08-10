package project

import (
	"errors"
	"fmt"
	"os"

	"github.com/pterm/pterm"
	client "github.com/taubyte/go-auth-http"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	"github.com/taubyte/tau-cli/i18n"
	projectI18n "github.com/taubyte/tau-cli/i18n/project"
	repositoryI18n "github.com/taubyte/tau-cli/i18n/repository"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/prompts"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/urfave/cli/v2"
)

func (link) Delete() common.Command {
	return common.Create(
		&cli.Command{
			Action: _delete,
			Flags: []cli.Flag{
				flags.Yes,
			},
		},
	)
}

func _delete(ctx *cli.Context) error {
	profile, err := loginLib.GetSelectedProfile()
	if err != nil {
		return err
	}

	projects, err := projectLib.ListResources()
	if err != nil {
		return err
	}

	projectMap := make(map[string]*client.Project, len(projects))
	projectList := make([]string, len(projects))
	for idx, project := range projects {
		projectList[idx] = project.Name
		projectMap[project.Name] = project
	}

	projectName := prompts.GetOrAskForSelection(ctx, "name", "Project:", projectList)
	project, ok := projectMap[projectName]
	if !ok {
		return i18n.ErrorDoesNotExist("project", projectName)
	}

	repoList, err := project.Repositories()
	if err != nil {
		return projectI18n.GettingRepositoriesFailed(projectName, err)
	}

	codeRepoFullName := repoList.Code.Fullname
	configRepoFullName := repoList.Configuration.Fullname

	printBullet := func(name string) string {
		return fmt.Sprint("  \u2022" + pterm.FgCyan.Sprint(name))
	}
	if prompts.ConfirmPrompt(
		ctx,
		fmt.Sprintf("Removing project `%s` will unregister the following repositories:\n%s\n%s\nProceed?",
			pterm.FgCyan.Sprint(projectName),
			printBullet(codeRepoFullName),
			printBullet(configRepoFullName),
		)) {

		prj, err := config.Projects().Get(projectName)
		if err != nil {
			os.RemoveAll(prj.Location)
		}

		if _, err = project.Delete(); err != nil {
			return projectI18n.ErrorDeleteProject(project.Name, err)
		}

		auth, err := authClient.Load()
		if err != nil {
			return err
		}

		codeRepo, err := auth.GetRepositoryByName(codeRepoFullName)
		if err != nil {
			return projectI18n.ErrorGettingRepositoryFailed(codeRepoFullName, err)
		}

		configRepo, err := auth.GetRepositoryByName(configRepoFullName)
		if err != nil {
			return projectI18n.ErrorGettingRepositoryFailed(configRepoFullName, err)
		}

		err = auth.UnregisterRepository(codeRepo.Get().ID())
		if _err := auth.UnregisterRepository(configRepo.Get().ID()); _err != nil {
			if err != nil {
				err = fmt.Errorf("%s:%w", err, _err)
			} else {
				err = _err
			}
		}
		if err != nil {
			return repositoryI18n.ErrorUnregisterRepositories(err)
		}

		projectI18n.RemovedProject(projectName, profile.Network)

		if prompts.ConfirmPrompt(ctx, "Remove from github?") {
			err = deleteRepo(ctx.Context, profile.Token, profile.GitUsername, repoList.Code.Name)
			if _err := deleteRepo(ctx.Context, profile.Token, profile.GitUsername, repoList.Configuration.Name); _err != nil {
				if err != nil {
					err = fmt.Errorf("%s:%w", err, _err)
				} else {
					err = _err
				}
			}
			if err != nil {
				if errors.Is(err, repositoryI18n.ErrorAdminRights) {
					pterm.Info.Println("Delete repositories manually")
					return nil
				}

				return err
			}
		}
	}

	return nil
}
