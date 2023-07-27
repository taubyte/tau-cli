package library

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	libraryI18n "github.com/taubyte/tau-cli/i18n/library"
	libraryLib "github.com/taubyte/tau-cli/lib/library"
	libraryPrompts "github.com/taubyte/tau-cli/prompts/library"
	libraryTable "github.com/taubyte/tau-cli/table/library"
)

func (link) Delete() common.Command {
	return (&resources.Delete[*structureSpec.Library]{
		PromptsGetOrSelect: libraryPrompts.GetOrSelect,
		TableConfirm:       libraryTable.Confirm,
		PromptsDeleteThis:  libraryPrompts.DeleteThis,
		LibDelete:          libraryLib.Delete,
		I18nDeleted:        libraryI18n.Deleted,
	}).Default()
}
