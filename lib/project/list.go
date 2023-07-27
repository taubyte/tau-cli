package projectLib

import (
	client "github.com/taubyte/go-auth-http"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
)

func List() ([]string, error) {
	projects, err := ListResources()
	if err != nil {
		return nil, err
	}

	projectNames := make([]string, len(projects))
	for idx, project := range projects {
		projectNames[idx] = project.Name
	}

	return projectNames, nil
}

func ListResources() ([]*client.Project, error) {
	client, err := authClient.Load()
	if err != nil {
		return nil, err
	}

	return client.Projects()
}
