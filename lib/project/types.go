package projectLib

import (
	git "github.com/taubyte/go-simple-git"
	"github.com/taubyte/tau/singletons/config"
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
