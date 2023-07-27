package service

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	serviceFlags "github.com/taubyte/tau-cli/flags/service"
	serviceI18n "github.com/taubyte/tau-cli/i18n/service"
	serviceLib "github.com/taubyte/tau-cli/lib/service"
	servicePrompts "github.com/taubyte/tau-cli/prompts/service"
	serviceTable "github.com/taubyte/tau-cli/table/service"
	"github.com/urfave/cli/v2"
)

func (link) New() common.Command {
	return (&resources.New[*structureSpec.Service]{
		PromptsNew:        servicePrompts.New,
		TableConfirm:      serviceTable.Confirm,
		PromptsCreateThis: servicePrompts.CreateThis,
		LibNew:            serviceLib.New,
		I18nCreated:       serviceI18n.Created,

		UniqueFlags: []cli.Flag{
			serviceFlags.Protocol,
		},
	}).Default()
}
