package prompts

import (
	"github.com/taubyte/tau/flags"
	"github.com/urfave/cli/v2"
)

func GetUseACodeTemplate(ctx *cli.Context) bool {
	return GetOrAskForBool(ctx, flags.UseCodeTemplate.Name, UseTemplatePrompt)
}
