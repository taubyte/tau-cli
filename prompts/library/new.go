package libraryPrompts

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	libraryLib "github.com/taubyte/tau-cli/lib/library"
	"github.com/taubyte/tau-cli/prompts"
	loginPrompts "github.com/taubyte/tau-cli/prompts/login"
	"github.com/urfave/cli/v2"
)

func New(ctx *cli.Context) (interface{}, *structureSpec.Library, error) {
	library := &structureSpec.Library{}

	taken, err := libraryLib.List()
	if err != nil {
		return nil, nil, err
	}

	library.Name = prompts.GetOrRequireAUniqueName(ctx, NamePrompt, taken)
	library.Description = prompts.GetOrAskForADescription(ctx)
	library.Tags = prompts.GetOrAskForTags(ctx)

	library.Provider, err = loginPrompts.SelectAProvider(ctx)
	if err != nil {
		return nil, nil, err
	}

	info, err := RepositoryInfo(ctx, library, true)
	if err != nil {
		return nil, nil, err
	}

	library.Path = prompts.GetOrRequireAPath(ctx, "Path:")

	library.Branch = prompts.GetOrRequireABranch(ctx)

	return info, library, nil
}
