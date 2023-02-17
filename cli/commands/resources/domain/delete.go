package domain

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	domainI18n "github.com/taubyte/tau/i18n/domain"
	domainLib "github.com/taubyte/tau/lib/domain"
	domainPrompts "github.com/taubyte/tau/prompts/domain"
	domainTable "github.com/taubyte/tau/table/domain"
)

func (link) Delete() common.Command {
	return (&resources.Delete[*structureSpec.Domain]{
		PromptsGetOrSelect: domainPrompts.GetOrSelect,
		TableConfirm:       domainTable.Confirm,
		PromptsDeleteThis:  domainPrompts.DeleteThis,
		LibDelete:          domainLib.Delete,
		I18nDeleted:        domainI18n.Deleted,
	}).Default()
}
