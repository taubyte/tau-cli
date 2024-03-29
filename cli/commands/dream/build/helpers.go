package build

import (
	"github.com/taubyte/go-project-schema/project"
	"github.com/taubyte/tau-cli/env"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/singletons/config"
)

type buildHelper struct {
	project       project.Project
	projectConfig config.Project
	currentBranch string
	selectedApp   string
}

func initBuild() (*buildHelper, error) {
	var err error
	helper := &buildHelper{}

	helper.project, err = projectLib.SelectedProjectInterface()
	if err != nil {
		return nil, err
	}

	helper.projectConfig, err = projectLib.SelectedProjectConfig()
	if err != nil {
		return nil, err
	}

	h := projectLib.Repository(helper.project.Get().Name())
	projectRepositories, err := h.Open()
	if err != nil {
		return nil, err
	}

	helper.currentBranch, err = projectRepositories.CurrentBranch()
	if err != nil {
		return nil, err
	}

	helper.selectedApp, _ = env.GetSelectedApplication()
	return helper, nil
}
