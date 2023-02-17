package prompts

import (
	"github.com/taubyte/tau/flags"
	"github.com/urfave/cli/v2"
)

func GetPrivate(ctx *cli.Context, prev ...bool) bool {
	return GetOrAskForBool(ctx, flags.Private.Name, PrivatePrompt, prev...)
}
