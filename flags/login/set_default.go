package loginFlags

import (
	"github.com/taubyte/tau/flags"
	"github.com/urfave/cli/v2"
)

var SetDefault = &flags.BoolWithInverseFlag{
	BoolFlag: &cli.BoolFlag{
		Name:    "set-default",
		Aliases: []string{"d"},
		Usage:   "Set the profile as the default profile",
	},
}
