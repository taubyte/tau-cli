package applicationPrompts

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
	"github.com/urfave/cli/v2"
)

func Edit(ctx *cli.Context, prev *structureSpec.App) {
	prev.Description = prompts.GetOrAskForADescription(ctx, prev.Description)
	prev.Tags = prompts.GetOrAskForTags(ctx, prev.Tags)
}
