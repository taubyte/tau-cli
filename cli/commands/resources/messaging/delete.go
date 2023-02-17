package messaging

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	messagingI18n "github.com/taubyte/tau/i18n/messaging"
	messagingLib "github.com/taubyte/tau/lib/messaging"
	messagingPrompts "github.com/taubyte/tau/prompts/messaging"
	messagingTable "github.com/taubyte/tau/table/messaging"
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
