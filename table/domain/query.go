package domainTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
)

func Query(domain *structureSpec.Domain) {
	prompts.RenderTable(getTableData(domain, true))
}
