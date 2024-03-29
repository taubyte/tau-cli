package libraryPrompts

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/urfave/cli/v2"
)

func Edit(ctx *cli.Context, prev *structureSpec.Library) (interface{}, error) {
	prev.Description = prompts.GetOrAskForADescription(ctx, prev.Description)
	prev.Tags = prompts.GetOrAskForTags(ctx, prev.Tags)

	info, err := RepositoryInfo(ctx, prev, false)
	if err != nil {
		return nil, err
	}

	prev.Path = prompts.GetOrRequireAPath(ctx, "Path:", prev.Path)

	prev.Branch = prompts.GetOrRequireABranch(ctx, prev.Branch)

	return info, nil
}
