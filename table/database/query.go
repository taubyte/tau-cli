package databaseTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/prompts"
)

func Query(database *structureSpec.Database) {
	prompts.RenderTable(getTableData(database, true))
}
