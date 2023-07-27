package storage

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	storageLib "github.com/taubyte/tau-cli/lib/storage"
	storagePrompts "github.com/taubyte/tau-cli/prompts/storage"
	storageTable "github.com/taubyte/tau-cli/table/storage"
)

func (link) Query() common.Command {
	return (&resources.Query[*structureSpec.Storage]{
		LibListResources:   storageLib.ListResources,
		TableList:          storageTable.List,
		PromptsGetOrSelect: storagePrompts.GetOrSelect,
		TableQuery:         storageTable.Query,
	}).Default()
}

func (link) List() common.Command {
	return (&resources.List[*structureSpec.Storage]{
		LibListResources: storageLib.ListResources,
		TableList:        storageTable.List,
	}).Default()
}
