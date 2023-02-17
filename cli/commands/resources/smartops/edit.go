package smartops

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	smartopsI18n "github.com/taubyte/tau/i18n/smartops"
	smartopsLib "github.com/taubyte/tau/lib/smartops"
	smartopsPrompts "github.com/taubyte/tau/prompts/smartops"
	smartopsTable "github.com/taubyte/tau/table/smartops"
)

func (link) Edit() common.Command {
	return (&resources.Edit[*structureSpec.SmartOp]{
		PromptsGetOrSelect: smartopsPrompts.GetOrSelect,
		PromptsEdit:        smartopsPrompts.Edit,
		TableConfirm:       smartopsTable.Confirm,
		PromptsEditThis:    smartopsPrompts.EditThis,
		LibSet:             smartopsLib.Set,
		I18nEdited:         smartopsI18n.Edited,

		UniqueFlags: flags.Combine(
			flags.Timeout,
			flags.Memory,
			flags.MemoryUnit,
			flags.Source,
			flags.Call,
		),
	}).Default()
}
