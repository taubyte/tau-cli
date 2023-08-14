package project

import (
	"fmt"

	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/prompts/spinner"
	auth_client "github.com/taubyte/tau-cli/singletons/auth_client"
	projectTable "github.com/taubyte/tau-cli/table/project"
	httpClient "github.com/taubyte/tau/clients/http/auth"
	"github.com/urfave/cli/v2"
)

func list(ctx *cli.Context) error {
	client, err := auth_client.Load()
	if err != nil {
		return err
	}

	stopGlobe := spinner.Globe()
	projects, err := client.Projects()
	if err != nil {
		return fmt.Errorf("Query projects failed with %s", err.Error())
	}

	t := projectTable.ListNoRender(projects, func(project *httpClient.Project) string {
		return projectLib.Description(project)
	})

	stopGlobe()
	t.Render()

	return nil
}
