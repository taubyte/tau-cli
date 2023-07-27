package options

import (
	"fmt"

	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	"github.com/taubyte/tau-cli/i18n"
	"github.com/urfave/cli/v2"
)

func SetNameAsArgs0(ctx *cli.Context) error {
	first := ctx.Args().First()
	if len(first) == 0 {
		return nil
	}

	return ctx.Set(flags.Name.Name, first)
}

func NameFlagArg0() common.Option {
	return func(l common.Linker) {
		// Insert name flag into first position
		l.Flags().Shift(flags.Name)
		l.Before().Shift(SetNameAsArgs0)
	}
}

func NameFlagSelectedArg0(selected string) common.Option {
	return func(l common.Linker) {
		NameFlagArg0()(l)

		parentName := l.Parent().Name

		if parentName != "new" && parentName != "select" {
			l.Raw().ArgsUsage = fmt.Sprintf(i18n.ArgsUsageNameDefaultSelected, selected)
			l.Raw().Flags[0] = &cli.StringFlag{
				Name:        flags.Name.Name,
				Aliases:     flags.Name.Aliases,
				Usage:       "Will default to selected",
				DefaultText: selected,
			}
		}
	}
}
