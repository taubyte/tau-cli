package projectLib

import (
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau/pkg/git"
)

type Project struct {
	Id          string
	Name        string
	Description string
	Public      bool
}

type ProjectRepository interface {
	Config() (*git.Repository, error)
	Code() (*git.Repository, error)
	CurrentBranch() (string, error)
}

type RepositoryHandler interface {
	Open() (ProjectRepository, error)
	Clone(tauProject config.Project, embedToken bool) (ProjectRepository, error)
}
