package cli

import (
	argsLib "github.com/taubyte/tau/cli/args"
	"github.com/taubyte/tau/i18n"
)

func Run(args ...string) error {
	app, err := New()
	if err != nil {
		return i18n.AppCreateFailed(err)
	}

	if len(args) == 1 {
		return app.Run(args)
	}

	args = argsLib.ParseArguments(app.Flags, app.Commands, args...)

	return app.Run(args)
}
