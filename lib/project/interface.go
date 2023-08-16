package projectLib

import (
	"github.com/taubyte/go-project-schema/project"
	"github.com/taubyte/tau-cli/env"
	"github.com/taubyte/tau-cli/i18n"
	"github.com/taubyte/tau-cli/singletons/config"
)

func SelectedProjectInterface() (project.Project, error) {
	configProject, err := SelectedProjectConfig()
	if err != nil {
		return nil, err
	}

	project, err := configProject.Interface()
	if err != nil {
		i18n.Help().BeSureToCloneProject()
		return nil, err
	}

	return project, nil
}

func SelectedProjectConfig() (configProject config.Project, err error) {
	selectedProject, err := env.GetSelectedProject()
	if err != nil {
		i18n.Help().BeSureToSelectProject()
		return
	}

	configProject, err = config.Projects().Get(selectedProject)
	if err != nil {
		i18n.Help().BeSureToCloneProject()
	}

	return
}

func ConfirmSelectedProject() error {
	_, err := env.GetSelectedProject()
	if err != nil {
		i18n.Help().BeSureToSelectProject()
	}

	return err
}
