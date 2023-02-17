package database

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	databaseLib "github.com/taubyte/tau/lib/database"
	databasePrompts "github.com/taubyte/tau/prompts/database"
	databaseTable "github.com/taubyte/tau/table/database"
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
