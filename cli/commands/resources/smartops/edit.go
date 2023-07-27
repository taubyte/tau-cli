package smartops

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	smartopsI18n "github.com/taubyte/tau-cli/i18n/smartops"
	smartopsLib "github.com/taubyte/tau-cli/lib/smartops"
	smartopsPrompts "github.com/taubyte/tau-cli/prompts/smartops"
	smartopsTable "github.com/taubyte/tau-cli/table/smartops"
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
