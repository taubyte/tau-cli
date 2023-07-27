package projectLib

import (
	httpClient "github.com/taubyte/go-auth-http"
	projectI18n "github.com/taubyte/tau-cli/i18n/project"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
)

func projectByName(name string) (*httpClient.Project, error) {
	client, err := authClient.Load()
	if err != nil {
		return nil, err
	}

	projects, err := client.Projects()
	if err != nil {
		return nil, projectI18n.GettingProjectsFailed(err)
	}

	var project *httpClient.Project
	for _, _project := range projects {
		if _project.Name == name {
			project = _project
			break
		}
	}
	if project == nil {
		return nil, projectI18n.ProjectNotFound(name)
	}

	return project, nil
}
