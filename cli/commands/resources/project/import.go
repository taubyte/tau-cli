package project

import (
	"fmt"
	"strings"

	httpClient "github.com/taubyte/go-auth-http"
	git "github.com/taubyte/go-simple-git"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	"github.com/taubyte/tau-cli/i18n"
	projectI18n "github.com/taubyte/tau-cli/i18n/project"
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
		return fmt.Errorf("getting selected profile failed with: %w", err)
	}

	repos, err := git.List(profile.GitUsername, profile.Token)
	if err != nil {
		return fmt.Errorf("listing `%s` repos failed with: %w", profile.GitUsername, err)
	}

	repoMap := make(map[string]*git.GitRepo, len(repos))
	configRepos := make([]string, 0, len(repos))
	codeRepos := make([]string, 0, len(repos))
	for _, repo := range repos {
		splitName := strings.SplitN(repo.Name, "_", 3)
		switch splitName[0] {
		case "tb":
			switch splitName[1] {
			case "library", "website":
				continue
			case "code":
				codeRepos = append(codeRepos, repo.FullName)
			default:
				configRepos = append(configRepos, repo.FullName)
			}
		default:
			continue
		}

		repoMap[repo.FullName] = repo
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
		return fmt.Errorf("selected code repo `%s` does not exist", codeRepoName)
	}

	configRepo, ok := repoMap[configRepoName]
	if !ok {
		return fmt.Errorf("selected config repo `%s` does not exist", configRepoName)
	}

	clientProject := &httpClient.Project{
		Name: configSplit[1],
	}

	auth, err := authClient.Load()
	if err != nil {
		return fmt.Errorf("loading auth client failed with: %w", err)
	}

	codeId := fmt.Sprintf("%d", codeRepo.ID)
	configId := fmt.Sprintf("%d", configRepo.ID)

	if err = auth.RegisterRepository(codeId); err != nil {
		return fmt.Errorf("registering code repo `%s` failed with: %w", codeRepo.FullName, err)
	}

	if err = auth.RegisterRepository(configId); err != nil {
		return fmt.Errorf("registering config repo `%s` failed with: %w", configRepo.FullName, err)
	}

	if err = clientProject.Create(auth, configId, codeId); err != nil {
		return fmt.Errorf("creating new project `%s` failed with: %w", clientProject.Name, err)
	}

	projectI18n.ImportedProject(clientProject.Name, profile.FQDN)

	if prompts.ConfirmPrompt(ctx, fmt.Sprintf("select `%s` as current project?", clientProject.Name)) {
		if err = session.Set().SelectedProject(clientProject.Name); err != nil {
			return fmt.Errorf("setting `%s` as current project failed with: %w", clientProject.Name, err)
		}

		projectI18n.SelectedProject(clientProject.Name)
		i18n.Help().BeSureToCloneProject()
	}

	return nil
}
