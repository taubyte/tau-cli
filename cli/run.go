package cli

import (
	argsLib "github.com/taubyte/tau-cli/cli/args"
	"github.com/taubyte/tau-cli/i18n"
)
// Run is a function that is responsible for running the CLI application
// It takes a variable number of string arguments
func Run(args ...string) error {
// Here it Creates a new instance of the CLI application
	app, err := New()
	if err != nil {
// If an error occurs during app creation, it returns an error message
		return i18n.AppCreateFailed(err)
	}
// Checks the number of arguments passed
	if len(args) == 1 {
// If there is only one argument, run the CLI application with that argument
		return app.Run(args)
	}
// Parse the arguments using the app's flags and commands
	args = argsLib.ParseArguments(app.Flags, app.Commands, args...)
// Run the CLI application with the parsed arguments
	return app.Run(args)
}
