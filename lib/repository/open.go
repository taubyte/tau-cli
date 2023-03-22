package repositoryLib

import (
	"fmt"

	git "github.com/taubyte/go-simple-git"
	websiteI18n "github.com/taubyte/tau/i18n/website"
	loginLib "github.com/taubyte/tau/lib/login"
	"github.com/taubyte/tau/singletons/config"
	"github.com/taubyte/tau/states"
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

	if info.isCloned(repositoryPath) == false {
		websiteI18n.Help().BeSureToCloneWebsite()
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
