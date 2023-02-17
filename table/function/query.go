package functionTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
)

func Query(function *structureSpec.Function) {
	prompts.RenderTable(getTableData(function, true))
}
