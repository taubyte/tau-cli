package dreamLib

import (
	"github.com/taubyte/go-project-schema/project"
	projectLib "github.com/taubyte/tau/lib/project"
	"github.com/taubyte/tau/singletons/config"
)

type ProdProject struct {
	Project project.Project
	Profile config.Profile
}

func (i *ProdProject) Import() error {
	h := projectLib.Repository(i.Project.Get().Name())
	projectRepositories, err := h.Open()
	if err != nil {
		return err
	}

	branch, err := projectRepositories.CurrentBranch()
	if err != nil {
		return err
	}

	return Execute("inject", "importProdProject",
		"--project-id", i.Project.Get().Id(),
		"--git-token", i.Profile.Token,
		"--branch", branch,
	)
}

func (i *ProdProject) Attach() error {
	return Execute("inject", "attachProdProject",
		"--project-id", i.Project.Get().Id(),
		"--git-token", i.Profile.Token,
	)
}
