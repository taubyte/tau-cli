package service

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	serviceI18n "github.com/taubyte/tau/i18n/service"
	serviceLib "github.com/taubyte/tau/lib/service"
	servicePrompts "github.com/taubyte/tau/prompts/service"
	serviceTable "github.com/taubyte/tau/table/service"
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
