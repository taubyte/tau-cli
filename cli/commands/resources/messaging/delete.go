package messaging

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	messagingI18n "github.com/taubyte/tau-cli/i18n/messaging"
	messagingLib "github.com/taubyte/tau-cli/lib/messaging"
	messagingPrompts "github.com/taubyte/tau-cli/prompts/messaging"
	messagingTable "github.com/taubyte/tau-cli/table/messaging"
)

func (link) Delete() common.Command {
	return (&resources.Delete[*structureSpec.Messaging]{
		PromptsGetOrSelect: messagingPrompts.GetOrSelect,
		TableConfirm:       messagingTable.Confirm,
		PromptsDeleteThis:  messagingPrompts.DeleteThis,
		LibDelete:          messagingLib.Delete,
		I18nDeleted:        messagingI18n.Deleted,
	}).Default()
}
