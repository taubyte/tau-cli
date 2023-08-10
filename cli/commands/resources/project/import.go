package project

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v53/github"
	httpClient "github.com/taubyte/go-auth-http"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	"github.com/taubyte/tau-cli/i18n"
	projectI18n "github.com/taubyte/tau-cli/i18n/project"
	repositoryI18n "github.com/taubyte/tau-cli/i18n/repository"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	"github.com/taubyte/tau-cli/prompts"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
	"github.com/taubyte/tau-cli/singletons/session"
	slices "github.com/taubyte/utils/slices/string"
	"github.com/urfave/cli/v2"
)

func (link) Import() common.Command {
	return common.Create(
		&cli.Command{
			Action: _import,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "config",
				},
				&cli.StringFlag{
					Name: "code",
				},
				flags.Yes,
			},
		},
	)
}

func _import(ctx *cli.Context) error {
	profile, err := loginLib.GetSelectedProfile()
	if err != nil {
		return err
	}

	repos, err := ListRepos(ctx.Context, profile.GitUsername, profile.Token)
	if err != nil {
		return err
	}

	repoMap := make(map[string]*github.Repository, len(repos))
	configRepos := make([]string, 0, len(repos))
	codeRepos := make([]string, 0, len(repos))
	for _, repo := range repos {
		splitName := strings.SplitN(repo.GetName(), "_", 3)
		if len(splitName) < 2 || splitName[0] != "tb" || splitName[1] == "library" || splitName[1] == "website" {
			continue
		}

		fullName := repo.GetFullName()
		switch splitName[1] {
		case "code":
			codeRepos = append(codeRepos, fullName)
		default:
			configRepos = append(configRepos, fullName)
		}
		repoMap[fullName] = repo
	}

	configRepoName := prompts.GetOrAskForSelection(ctx, "config", "Config:", configRepos)
	configSplit := strings.SplitN(configRepoName, "_", 2)
	codeSplit := []string{configSplit[0], "code", configSplit[1]}
	codeRepoName := strings.Join(codeSplit, "_")

	var prev []string
	if slices.Contains(codeRepos, codeRepoName) {
		prev = append(prev, codeRepoName)
	}

	codeRepoName = prompts.GetOrAskForSelection(ctx, "code", "Code:", codeRepos, prev...)
	codeRepo, ok := repoMap[codeRepoName]
	if !ok {
		return i18n.ErrorDoesNotExist("code repo", codeRepoName)
	}

	configRepo, ok := repoMap[configRepoName]
	if !ok {
		return i18n.ErrorDoesNotExist("config repo", configRepoName)
	}

	clientProject := &httpClient.Project{
		Name: configSplit[1],
	}

	auth, err := authClient.Load()
	if err != nil {
		return err
	}

	codeId := fmt.Sprintf("%d", codeRepo.ID)
	configId := fmt.Sprintf("%d", configRepo.ID)

	if err = auth.RegisterRepository(codeId); err != nil {
		return repositoryI18n.RegisteringRepositoryFailed(codeRepoName, err)
	}

	if err = auth.RegisterRepository(configId); err != nil {
		return repositoryI18n.RegisteringRepositoryFailed(configRepoName, err)
	}

	if err = clientProject.Create(auth, configId, codeId); err != nil {
		return projectI18n.CreatingProjectFailed(err)
	}

	projectI18n.ImportedProject(clientProject.Name, profile.Network)

	if prompts.ConfirmPrompt(ctx, fmt.Sprintf("select `%s` as current project?", clientProject.Name)) {
		if err = session.Set().SelectedProject(clientProject.Name); err != nil {
			return projectI18n.SelectingAProjectPromptFailed(err)
		}

		projectI18n.SelectedProject(clientProject.Name)
		i18n.Help().BeSureToCloneProject()
	}

	return nil
}
