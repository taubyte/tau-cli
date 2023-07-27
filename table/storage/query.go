package storageTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/prompts"
)

func Query(storage *structureSpec.Storage) {
	prompts.RenderTable(getTableData(storage, true))
}
