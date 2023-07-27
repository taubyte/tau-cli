package function

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	functionLib "github.com/taubyte/tau-cli/lib/function"
	functionPrompts "github.com/taubyte/tau-cli/prompts/function"
	functionTable "github.com/taubyte/tau-cli/table/function"
)

func (link) Query() common.Command {
	return (&resources.Query[*structureSpec.Function]{
		LibListResources:   functionLib.ListResources,
		TableList:          functionTable.List,
		PromptsGetOrSelect: functionPrompts.GetOrSelect,
		TableQuery:         functionTable.Query,
	}).Default()
}

func (link) List() common.Command {
	return (&resources.List[*structureSpec.Function]{
		LibListResources: functionLib.ListResources,
		TableList:        functionTable.List,
	}).Default()
}
