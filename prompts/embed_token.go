package prompts

import (
	"github.com/taubyte/tau-cli/flags"
	"github.com/urfave/cli/v2"
)

func GetOrAskForEmbedToken(ctx *cli.Context, prev ...bool) bool {
	return GetOrAskForBool(ctx, flags.EmbedToken.Name, EmbedTokenPrompt, prev...)
}
