package smartopsTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
)

func Query(smartops *structureSpec.SmartOp) {
	prompts.RenderTable(getTableData(smartops, true))
}
