package repositoryLib

import (
	git "github.com/taubyte/go-simple-git"
	"github.com/taubyte/tau/singletons/config"
)

func (info *Info) Pull(project config.Project, branch, url string) (*git.Repository, error) {
	repo, err := info.Open(project, branch, url)
	if err != nil {
		return nil, err
	}

	err = repo.Pull()
	if err != nil {
		return nil, err
	}

	return repo, nil
}
