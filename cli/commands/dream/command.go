package dream

import (
	"github.com/taubyte/tau/cli/commands/dream/build"
	dreamLib "github.com/taubyte/tau/lib/dream"
	projectLib "github.com/taubyte/tau/lib/project"
	"github.com/urfave/cli/v2"
)

const defaultBind = "node@1/verbose,seer@2/copies,node@2/copies"

var Command = &cli.Command{
	Name:  "dream",
	Usage: "Starts and interfaces with a local taubyte network.  All leading arguments to `tau dream ...` are passed to dreamland",
	Action: func(c *cli.Context) error {
		project, err := projectLib.SelectedProjectInterface()
		if err != nil {
			return err
		}

		h := projectLib.Repository(project.Get().Name())
		projectRepositories, err := h.Open()
		if err != nil {
			return err
		}

		branch, err := projectRepositories.CurrentBranch()
		if err != nil {
			return err
		}

		return dreamLib.Execute("new", "multiverse", "--bind", defaultBind, "--branch", branch)
	},

	Subcommands: []*cli.Command{
		injectCommand,
		attachCommand,
		build.Command,
	},
}
