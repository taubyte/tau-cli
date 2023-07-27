package website

import (
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/cli/common/options"
	"github.com/taubyte/tau-cli/i18n"
	"github.com/urfave/cli/v2"
)

// Base is the command that is proliferated to all sub commands
// Example options.NameFlagArg0() gives --name flag and args[0] name to
// new, edit, delete, etc...
func (link) Base() (*cli.Command, []common.Option) {
	return common.Base(
		&cli.Command{
			Name:      "website",
			ArgsUsage: i18n.ArgsUsageName,
		}, options.NameFlagArg0(),
	)
}
