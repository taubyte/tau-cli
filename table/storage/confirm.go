package storageTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/urfave/cli/v2"
)

func Confirm(ctx *cli.Context, storage *structureSpec.Storage, prompt string) bool {
	return prompts.ConfirmData(ctx, prompt, getTableData(storage, false))
}
