package project

import (
	"fmt"

	"github.com/pterm/pterm"
	client "github.com/taubyte/go-auth-http"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/prompts"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
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
	projects, err := projectLib.ListResources()
	if err != nil {
		return fmt.Errorf("listing projects failed with %w", err)
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
		return fmt.Errorf("project `%s` does not exist", projectName)
	}

	repoList, err := project.Repositories()
	if err != nil {
		return fmt.Errorf("getting project repos failed with: %w", err)
	}

	codeRepoName := repoList.Code.Fullname
	configRepoName := repoList.Configuration.Fullname
	if prompts.ConfirmPrompt(
		ctx,
		fmt.Sprintf("Continue with removing project %s, un-registering config: %s code: %s?",
			pterm.FgCyan.Sprint(projectName),
			pterm.FgCyan.Sprint(codeRepoName),
			pterm.FgCyan.Sprint(configRepoName),
		),
	) {
		if _, err = project.Delete(); err != nil {
			return fmt.Errorf("deleting project `%s` failed with: %w", project.Name, err)
		}

		auth, err := authClient.Load()
		if err != nil {
			return fmt.Errorf("loading auth client failed with: %w", err)
		}

		codeRepo, err := auth.GetRepositoryByName(codeRepoName)
		if err != nil {
			return fmt.Errorf("getting code repo `%s` from auth failed with: %w", codeRepoName, err)
		}

		configRepo, err := auth.GetRepositoryByName(configRepoName)
		if err != nil {
			return fmt.Errorf("getting config repo `%s` from auth failed with: %w", configRepoName, err)
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
			return fmt.Errorf("un-registering repos failed with: %w", err)
		}

		pterm.Success.Printfln("project: %s removed", projectName)
	}

	return nil
}
