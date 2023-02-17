package smartops

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	smartopsI18n "github.com/taubyte/tau/i18n/smartops"
	smartopsLib "github.com/taubyte/tau/lib/smartops"
	smartopsPrompts "github.com/taubyte/tau/prompts/smartops"
	smartopsTable "github.com/taubyte/tau/table/smartops"
)

func (link) Delete() common.Command {
	return (&resources.Delete[*structureSpec.SmartOp]{
		PromptsGetOrSelect: smartopsPrompts.GetOrSelect,
		TableConfirm:       smartopsTable.Confirm,
		PromptsDeleteThis:  smartopsPrompts.DeleteThis,
		LibDelete:          smartopsLib.Delete,
		I18nDeleted:        smartopsI18n.Deleted,
	}).Default()
}
