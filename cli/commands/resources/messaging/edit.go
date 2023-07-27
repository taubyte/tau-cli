package messaging

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	messagingFlags "github.com/taubyte/tau-cli/flags/messaging"
	messagingI18n "github.com/taubyte/tau-cli/i18n/messaging"
	messagingLib "github.com/taubyte/tau-cli/lib/messaging"
	messagingPrompts "github.com/taubyte/tau-cli/prompts/messaging"
	messagingTable "github.com/taubyte/tau-cli/table/messaging"
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
