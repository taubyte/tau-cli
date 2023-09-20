//Comments add:

/*
Overall, this code defines a "login" CLI command 
that allows users to manage login profiles. It handles options, 
internationalization, profile creation, and profile selection based 
on user input and flags. The specific behavior of some functions (e.g., New and Select)
 is not provided in this code snippet, so their exact functionality would 
 depend on their implementations in other parts of the codebase.
package login
*/

/* IMPORT
The code starts by importing various packages and modules, 
including packages related to CLI handling, flags, internationalization 
(i18n), login functionality, prompts, and utility functions for string slices.
*/
import ( 
	"github.com/taubyte/tau-cli/cli/common/options"
	"github.com/taubyte/tau-cli/flags"
	loginFlags "github.com/taubyte/tau-cli/flags/login"
	"github.com/taubyte/tau-cli/i18n"
	loginI18n "github.com/taubyte/tau-cli/i18n/login"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	"github.com/taubyte/tau-cli/prompts"
	loginPrompts "github.com/taubyte/tau-cli/prompts/login"
	slices "github.com/taubyte/utils/slices/string"
	"github.com/urfave/cli/v2"
)

/* VAR
var Command Declaration: The code defines a variable named Command, 
which is a pointer to a cli.Command struct. This struct represents 
the definition of a CLI command.
*/ 
var Command = &cli.Command{
	Name: "login",
	Flags: flags.Combine(
		flags.Name,
		loginFlags.Token,
		loginFlags.Provider,
		loginFlags.New,
		loginFlags.SetDefault,
	),
	ArgsUsage: i18n.ArgsUsageName,
	Action:    Run,
	Before:    options.SetNameAsArgs0,
}

//Run Function: This function, Run(ctx *cli.Context) error, is executed when the "login" command is run.
func Run(ctx *cli.Context) error {
	_default, options, err := loginLib.GetProfiles()
	if err != nil {
		return loginI18n.GetProfilesFailed(err)
	}

	// New: if --new or no selectable profiles
	if ctx.Bool(loginFlags.New.Name) == true || len(options) == 0 {
		return New(ctx, options)
	}

	// Selection
	var name string
	if ctx.IsSet(flags.Name.Name) == true {
		name = ctx.String(flags.Name.Name)

		if slices.Contains(options, name) == false {
			return loginI18n.DoesNotExistIn(name, options)
		}
	} else {
		name, err = prompts.SelectInterface(options, loginPrompts.SelectAProfile, _default)
		if err != nil {
			return err
		}
	}

	return Select(ctx, name, ctx.Bool(loginFlags.SetDefault.Name))
}
