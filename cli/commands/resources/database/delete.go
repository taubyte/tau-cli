package database

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	databaseI18n "github.com/taubyte/tau-cli/i18n/database"
	databaseLib "github.com/taubyte/tau-cli/lib/database"
	databasePrompts "github.com/taubyte/tau-cli/prompts/database"
	databaseTable "github.com/taubyte/tau-cli/table/database"
)

func (link) Delete() common.Command {
	return (&resources.Delete[*structureSpec.Database]{
		PromptsGetOrSelect: databasePrompts.GetOrSelect,
		TableConfirm:       databaseTable.Confirm,
		PromptsDeleteThis:  databasePrompts.DeleteThis,
		LibDelete:          databaseLib.Delete,
		I18nDeleted:        databaseI18n.Deleted,
	}).Default()
}
