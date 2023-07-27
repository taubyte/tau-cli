package domainPrompts

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/urfave/cli/v2"
)

func Edit(ctx *cli.Context, prev *structureSpec.Domain) error {
	prev.Description = prompts.GetOrAskForADescription(ctx, prev.Description)
	prev.Tags = prompts.GetOrAskForTags(ctx, prev.Tags)
	prev.Fqdn = GetOrRequireAnFQDN(ctx, prev.Fqdn)

	err := certificate(ctx, prev, false)
	if err != nil {
		return err
	}

	return nil
}
