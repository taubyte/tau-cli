package functionTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/prompts"
)

func Query(function *structureSpec.Function) {
	prompts.RenderTable(getTableData(function, true))
}
