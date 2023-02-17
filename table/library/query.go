package libraryTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
)

func Query(library *structureSpec.Library) {
	prompts.RenderTable(getTableData(library, true))
}
