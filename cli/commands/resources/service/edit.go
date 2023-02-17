package service

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	serviceFlags "github.com/taubyte/tau/flags/service"
	serviceI18n "github.com/taubyte/tau/i18n/service"
	serviceLib "github.com/taubyte/tau/lib/service"
	servicePrompts "github.com/taubyte/tau/prompts/service"
	serviceTable "github.com/taubyte/tau/table/service"
	"github.com/urfave/cli/v2"
)

func (link) Edit() common.Command {
	return (&resources.Edit[*structureSpec.Service]{
		PromptsGetOrSelect: servicePrompts.GetOrSelect,
		PromptsEdit:        servicePrompts.Edit,
		TableConfirm:       serviceTable.Confirm,
		PromptsEditThis:    servicePrompts.EditThis,
		LibSet:             serviceLib.Set,
		I18nEdited:         serviceI18n.Edited,

		UniqueFlags: []cli.Flag{
			serviceFlags.Protocol,
		},
	}).Default()
}
