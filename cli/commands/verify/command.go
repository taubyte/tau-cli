package verify

import (
	"bitbucket.org/taubyte/config-compiler/compile"
	"bitbucket.org/taubyte/go-interfaces/services/patrick"
	loginLib "github.com/taubyte/tau/lib/login"
	projectLib "github.com/taubyte/tau/lib/project"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name: "verify",
	Subcommands: []*cli.Command{
		{
			Name:   "config",
			Action: Run,
		},
	},
}

func Run(ctx *cli.Context) error {
	project, err := projectLib.SelectedProjectInterface()
	if err != nil {
		return err
	}

	profile, err := loginLib.GetSelectedProfile()
	if err != nil {
		return err
	}

	rc, err := compile.CompilerConfig(project, patrick.Meta{
		HeadCommit: patrick.HeadCommit{
			ID: "1",
		},
		Repository: patrick.Repository{
			ID:         1,
			Provider:   profile.Provider,
			Branch:     "main",
			MainBranch: "main",
		},
	})
	if err != nil {
		return err
	}

	compiler, err := compile.New(rc, compile.Dev())
	if err != nil {
		return err
	}

	defer compiler.Close()
	err = compiler.Build()
	if err != nil {
		return err
	}

	return nil
}
