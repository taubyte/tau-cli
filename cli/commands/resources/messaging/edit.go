package messaging

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	messagingFlags "github.com/taubyte/tau/flags/messaging"
	messagingI18n "github.com/taubyte/tau/i18n/messaging"
	messagingLib "github.com/taubyte/tau/lib/messaging"
	messagingPrompts "github.com/taubyte/tau/prompts/messaging"
	messagingTable "github.com/taubyte/tau/table/messaging"
)

func (link) Edit() common.Command {
	return (&resources.Edit[*structureSpec.Messaging]{
		PromptsGetOrSelect: messagingPrompts.GetOrSelect,
		PromptsEdit:        messagingPrompts.Edit,
		TableConfirm:       messagingTable.Confirm,
		PromptsEditThis:    messagingPrompts.EditThis,
		LibSet:             messagingLib.Set,
		I18nEdited:         messagingI18n.Edited,

		UniqueFlags: flags.Combine(
			flags.Local,
			flags.MatchRegex,
			flags.Match,
			messagingFlags.MQTT,
			messagingFlags.WebSocket,
		),
	}).Default()
}
