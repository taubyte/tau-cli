package project

import (
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/cli/common/options"
	"github.com/taubyte/tau-cli/env"
	"github.com/taubyte/tau-cli/i18n"
	"github.com/urfave/cli/v2"
)

func (link) Base() (*cli.Command, []common.Option) {
	selected, err := env.GetSelectedProject()
	if err != nil {
		selected = "selected"
	}

	return common.Base(&cli.Command{
		Name:      "project",
		ArgsUsage: i18n.ArgsUsageName,
	}, options.NameFlagSelectedArg0(selected))
}
