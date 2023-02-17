package messagingTable

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
)

func Query(messaging *structureSpec.Messaging) {
	prompts.RenderTable(getTableData(messaging, true))
}
