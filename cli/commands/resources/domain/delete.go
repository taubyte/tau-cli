package domain

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	domainI18n "github.com/taubyte/tau-cli/i18n/domain"
	domainLib "github.com/taubyte/tau-cli/lib/domain"
	domainPrompts "github.com/taubyte/tau-cli/prompts/domain"
	domainTable "github.com/taubyte/tau-cli/table/domain"
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
