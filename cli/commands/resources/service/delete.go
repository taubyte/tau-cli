package service

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	serviceI18n "github.com/taubyte/tau-cli/i18n/service"
	serviceLib "github.com/taubyte/tau-cli/lib/service"
	servicePrompts "github.com/taubyte/tau-cli/prompts/service"
	serviceTable "github.com/taubyte/tau-cli/table/service"
)

func (link) Delete() common.Command {
	return (&resources.Delete[*structureSpec.Service]{
		PromptsGetOrSelect: servicePrompts.GetOrSelect,
		TableConfirm:       serviceTable.Confirm,
		PromptsDeleteThis:  servicePrompts.DeleteThis,
		LibDelete:          serviceLib.Delete,
		I18nDeleted:        serviceI18n.Deleted,
	}).Default()
}
