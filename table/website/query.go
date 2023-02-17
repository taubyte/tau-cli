package websiteTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
)

func Query(website *structureSpec.Website) {
	prompts.RenderTable(getTableData(website, true))
}
