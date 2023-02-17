package project

import (
	"os"
	"path"

	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	projectFlags "github.com/taubyte/tau/flags/project"
	"github.com/taubyte/tau/i18n"
	projectI18n "github.com/taubyte/tau/i18n/project"
	projectLib "github.com/taubyte/tau/lib/project"
	"github.com/taubyte/tau/prompts"
	projectPrompts "github.com/taubyte/tau/prompts/project"
	"github.com/taubyte/tau/singletons/config"
	"github.com/urfave/cli/v2"
)

func (link) Clone() common.Command {
	return common.Create(
		&cli.Command{
			Flags: flags.Combine(
				flags.Yes,
				projectFlags.Loc,
				flags.Branch,
				flags.EmbedToken,
				flags.Select,
			),
			Action: clone,
		},
	)
}

func clone(c *cli.Context) error {
	checkEnv := c.Bool(flags.Select.Name) == false

	// TODO should select offer projects that are already cloned?
	project, err := projectPrompts.GetOrSelect(c, checkEnv)
	if err != nil {
		return err
	}

	configProject := config.Project{
		Name: project.Name,
	}

	// Check location flag, otherwise clone into cwd
	if c.IsSet(projectFlags.Loc.Name) == true {
		configProject.Location = c.String(projectFlags.Loc.Name)
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return i18n.GettingCwdFailed(err)
		}

		configProject.Location = path.Join(cwd, project.Name)
	}

	branch := prompts.GetOrRequireABranch(c)

	_, err = projectLib.Repository(project.Name, branch).Clone(configProject, prompts.GetOrAskForEmbedToken(c))
	if err != nil {
		return projectI18n.CloningProjectFailed(project.Name, err)
	}

	return nil
}
