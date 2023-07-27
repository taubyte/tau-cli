package prompts

import (
	"github.com/taubyte/tau-cli/flags"
	"github.com/urfave/cli/v2"
)

func GetUseACodeTemplate(ctx *cli.Context) bool {
	return GetOrAskForBool(ctx, flags.UseCodeTemplate.Name, UseTemplatePrompt)
}
