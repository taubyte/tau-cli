package service

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	serviceLib "github.com/taubyte/tau-cli/lib/service"
	servicePrompts "github.com/taubyte/tau-cli/prompts/service"
	serviceTable "github.com/taubyte/tau-cli/table/service"
)

func (link) Query() common.Command {
	return (&resources.Query[*structureSpec.Service]{
		LibListResources:   serviceLib.ListResources,
		TableList:          serviceTable.List,
		PromptsGetOrSelect: servicePrompts.GetOrSelect,
		TableQuery:         serviceTable.Query,
	}).Default()
}

func (link) List() common.Command {
	return (&resources.List[*structureSpec.Service]{
		LibListResources: serviceLib.ListResources,
		TableList:        serviceTable.List,
	}).Default()
}
