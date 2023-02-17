package storage

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	storageI18n "github.com/taubyte/tau/i18n/storage"
	storageLib "github.com/taubyte/tau/lib/storage"
	storagePrompts "github.com/taubyte/tau/prompts/storage"
	storageTable "github.com/taubyte/tau/table/storage"
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
