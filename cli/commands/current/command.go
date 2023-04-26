package current

import (
	"github.com/taubyte/tau/common"
	"github.com/taubyte/tau/env"
	"github.com/taubyte/tau/prompts"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:    "current",
	Usage:   "Display current selected values",
	Aliases: []string{"cur", "here", "this"},
	Action:  Run,
}

func parseIfEmpty(v string) string {
	if len(v) == 0 {
		return "(none)"
	}

	return v
}

func Run(c *cli.Context) error {
	selectedProfile, _ := env.GetSelectedUser()
	selectedProject, _ := env.GetSelectedProject()
	selectedApplication, _ := env.GetSelectedApplication()
	selectedNetwork, _ := env.GetSelectedNetwork()
	customNetworkUrl, _ := env.GetCustomNetworkUrl()

	defaultRender := [][]string{
		{"Profile", parseIfEmpty(selectedProfile)},
		{"Project", parseIfEmpty(selectedProject)},
		{"Application", parseIfEmpty(selectedApplication)},
		{"Network", parseIfEmpty(selectedNetwork)},
	}

	if selectedNetwork == common.CustomNetwork {
		defaultRender = append(defaultRender, []string{"FQDN", parseIfEmpty(customNetworkUrl)})
	}

	prompts.RenderTableWithMerge(defaultRender)
	return nil
}
