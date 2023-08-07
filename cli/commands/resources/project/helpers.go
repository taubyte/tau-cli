package project

import (
	"sync"

	"github.com/pterm/pterm"
	git "github.com/taubyte/go-simple-git"
	projectFlags "github.com/taubyte/tau-cli/flags/project"
	"github.com/taubyte/tau-cli/i18n"
	projectI18n "github.com/taubyte/tau-cli/i18n/project"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/urfave/cli/v2"
)

// See if they have cloned the project, if not show help
func checkProjectClonedHelp(name string) {
	project, err := config.Projects().Get(name)
	if err != nil || len(project.Location) == 0 {
		i18n.Help().BeSureToCloneProject()
	}
}

type dualRepoHandler struct {
	ctx         *cli.Context
	projectName string
	repository  projectLib.ProjectRepository
	action      func(*git.Repository) error
	errorFormat func(string) error
}

// Run will parse for config-only || code-only
// then Runs a go routine to commit the action on both
// config and code repositories asynchronously or run config/code only
func (h *dualRepoHandler) Run() error {
	config, code, err := projectFlags.ParseConfigCodeFlags(h.ctx)
	if err != nil {
		return err
	}

	var (
		configErr error
		codeErr   error
	)

	var wg sync.WaitGroup

	if config {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var configRepo *git.Repository

			configRepo, configErr = h.repository.Config()
			if configErr != nil {
				return
			}

			configErr = h.action(configRepo)
			if configErr != nil {
				pterm.Error.Printfln(projectI18n.ConfigRepo, configErr)
			}
		}()
	}

	if code {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var codeRepo *git.Repository

			codeRepo, codeErr = h.repository.Code()
			if codeErr != nil {
				return
			}

			codeErr = h.action(codeRepo)
			if codeErr != nil {
				pterm.Error.Printfln(projectI18n.CodeRepo, codeErr)
			}
		}()
	}

	wg.Wait()
	if configErr != nil || codeErr != nil {
		return h.errorFormat(h.projectName)
	}

	return nil
}
