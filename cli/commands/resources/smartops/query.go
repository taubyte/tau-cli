package smartops

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	smartopsLib "github.com/taubyte/tau/lib/smartops"
	smartopsPrompts "github.com/taubyte/tau/prompts/smartops"
	smartopsTable "github.com/taubyte/tau/table/smartops"
)

func (link) Query() common.Command {
	return (&resources.Query[*structureSpec.SmartOp]{
		LibListResources:   smartopsLib.ListResources,
		TableList:          smartopsTable.List,
		PromptsGetOrSelect: smartopsPrompts.GetOrSelect,
		TableQuery:         smartopsTable.Query,
	}).Default()
}

func (link) List() common.Command {
	return (&resources.List[*structureSpec.SmartOp]{
		LibListResources: smartopsLib.ListResources,
		TableList:        smartopsTable.List,
	}).Default()
}
