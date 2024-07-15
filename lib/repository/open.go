package repositoryLib

import (
	"fmt"

	libraryI18n "github.com/taubyte/tau-cli/i18n/library"
	websiteI18n "github.com/taubyte/tau-cli/i18n/website"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau-cli/states"
	"github.com/taubyte/tau/pkg/git"
)

func (info *Info) Open(project config.Project, url string) (*git.Repository, error) {
	profile, err := loginLib.GetSelectedProfile()
	if err != nil {
		return nil, err
	}

	repositoryPath, err := info.path(project)
	if err != nil {
		return nil, err
	}

	if !info.isCloned(repositoryPath) {
		switch info.Type {
		case "website":
			websiteI18n.Help().BeSureToCloneWebsite()
		case "library":
			libraryI18n.Help().BeSureToCloneLibrary()
		}
		return nil, fmt.Errorf("repository not cloned: `%s`", repositoryPath)
	}

	repo, err := git.New(states.Context,
		git.Root(repositoryPath),
		git.Author(profile.GitUsername, profile.GitEmail),
		git.URL(url),
		git.Token(profile.Token),

		// TODO branch, this breaks things
		// git.Branch(branch),
	)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
