package repositoryLib

import (
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau/pkg/git"
)

func (info *Info) Push(project config.Project, message, url string) (*git.Repository, error) {
	repo, err := info.Open(project, url)
	if err != nil {
		return nil, err
	}

	err = repo.Commit(message, ".")
	if err != nil {
		return nil, err
	}

	err = repo.Push()
	if err != nil {
		return nil, err
	}

	return repo, nil
}
