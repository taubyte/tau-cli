package database

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	databaseLib "github.com/taubyte/tau-cli/lib/database"
	databasePrompts "github.com/taubyte/tau-cli/prompts/database"
	databaseTable "github.com/taubyte/tau-cli/table/database"
)

func (link) Query() common.Command {
	return (&resources.Query[*structureSpec.Database]{
		LibListResources:   databaseLib.ListResources,
		TableList:          databaseTable.List,
		PromptsGetOrSelect: databasePrompts.GetOrSelect,
		TableQuery:         databaseTable.Query,
	}).Default()
}

func (link) List() common.Command {
	return (&resources.List[*structureSpec.Database]{
		LibListResources: databaseLib.ListResources,
		TableList:        databaseTable.List,
	}).Default()
}
