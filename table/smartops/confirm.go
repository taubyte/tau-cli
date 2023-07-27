package smartopsTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/urfave/cli/v2"
)

func Confirm(ctx *cli.Context, smartops *structureSpec.SmartOp, prompt string) bool {
	return prompts.ConfirmData(ctx, prompt, getTableData(smartops, false))
}
