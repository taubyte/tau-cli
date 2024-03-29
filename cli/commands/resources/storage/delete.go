package storage

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	storageI18n "github.com/taubyte/tau-cli/i18n/storage"
	storageLib "github.com/taubyte/tau-cli/lib/storage"
	storagePrompts "github.com/taubyte/tau-cli/prompts/storage"
	storageTable "github.com/taubyte/tau-cli/table/storage"
)

func (link) Delete() common.Command {
	return (&resources.Delete[*structureSpec.Storage]{
		PromptsGetOrSelect: storagePrompts.GetOrSelect,
		TableConfirm:       storageTable.Confirm,
		PromptsDeleteThis:  storagePrompts.DeleteThis,
		LibDelete:          storageLib.Delete,
		I18nDeleted:        storageI18n.Deleted,
	}).Default()
}
