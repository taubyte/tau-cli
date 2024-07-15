package repositoryLib

import (
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau/pkg/git"
)

func (info *Info) Pull(project config.Project, url string) (*git.Repository, error) {
	repo, err := info.Open(project, url)
	if err != nil {
		return nil, err
	}

	err = repo.Pull()
	if err != nil {
		return nil, err
	}

	return repo, nil
}
