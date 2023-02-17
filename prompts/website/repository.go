package websitePrompts

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
	httpClient "github.com/taubyte/go-auth-http"
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/common"
	"github.com/taubyte/tau/flags"
	projectLib "github.com/taubyte/tau/lib/project"
	repositoryLib "github.com/taubyte/tau/lib/repository"
	"github.com/taubyte/tau/prompts"
	authClient "github.com/taubyte/tau/singletons/auth_client"
	"github.com/taubyte/tau/singletons/templates"
	"github.com/urfave/cli/v2"
)

func RepositoryInfo(ctx *cli.Context, website *structureSpec.Website, new bool) (interface{}, error) {
	if new == true && prompts.GetGenerateRepository(ctx) {
		return repositoryInfoGenerate(ctx, website)
	}

	selectedRepository, err := prompts.SelectARepository(ctx, &repositoryLib.Info{
		Type:     repositoryLib.WebsiteRepositoryType,
		FullName: website.RepoName,
		ID:       website.RepoID,
	})
	if err != nil {
		return nil, err
	}

	website.RepoID = selectedRepository.ID
	website.RepoName = selectedRepository.FullName

	projectConfig, err := projectLib.SelectedProjectConfig()
	if err != nil {
		return nil, err
	}

	if selectedRepository.HasBeenCloned(projectConfig, website.Provider) == false {
		selectedRepository.DoClone = prompts.GetClone(ctx)
	}

	return selectedRepository, nil

}

func isRepositoryNameTaken(client *httpClient.Client, name string) (bool, error) {
	var fullName string
	if len(strings.Split(name, "/")) == 2 {
		fullName = name
	} else {
		userInfo, err := client.User().Get()
		if err != nil {
			return false, err
		}

		fullName = fmt.Sprintf("%s/%s", userInfo.Login, name)
	}

	// Considering name to be taken if err is nil
	_, err := client.GetRepositoryByName(fullName)
	if err == nil {
		return true, nil
	}

	return false, nil
}

// Only called by new
func repositoryInfoGenerate(ctx *cli.Context, website *structureSpec.Website) (*repositoryLib.InfoTemplate, error) {
	var repositoryName string
	if ctx.IsSet(flags.RepositoryName.Name) == true {
		repositoryName = ctx.String(flags.RepositoryName.Name)
	} else {
		repositoryName = fmt.Sprintf(common.WebsiteRepoPrefix, website.Name)
	}

	// Confirm name is valid
	client, err := authClient.Load()
	if err != nil {
		return nil, err
	}

	// Skipping prompt for repository name unless set, or generated name is already taken
	for taken, err := isRepositoryNameTaken(client, repositoryName); taken == true; {
		if err != nil {
			return nil, err
		}

		pterm.Warning.Printfln("Repository name %s is already taken", repositoryName)
		repositoryName = prompts.GetOrRequireARepositoryName(ctx)

		taken, err = isRepositoryNameTaken(client, repositoryName)
	}

	private := prompts.GetPrivate(ctx)

	templateMap, err := templates.Get().Websites()
	if err != nil {
		// TODO verbose
		return nil, err
	}

	templateUrl, err := prompts.SelectATemplate(ctx, templateMap)
	if err != nil {
		return nil, err
	}

	return &repositoryLib.InfoTemplate{
		RepositoryName: repositoryName,
		Info: templates.TemplateInfo{
			URL: templateUrl,
			// TODO Update website template description style
			// Description: website.Description,
		},
		Private: private,
	}, nil
}
