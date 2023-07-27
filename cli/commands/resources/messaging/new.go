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

func (link) New() common.Command {
	return (&resources.New[*structureSpec.Messaging]{
		PromptsNew:        messagingPrompts.New,
		TableConfirm:      messagingTable.Confirm,
		PromptsCreateThis: messagingPrompts.CreateThis,
		LibNew:            messagingLib.New,
		I18nCreated:       messagingI18n.Created,

		UniqueFlags: flags.Combine(
			flags.Local,
			flags.MatchRegex,
			flags.Match,
			messagingFlags.MQTT,
			messagingFlags.WebSocket,
		),
	}).Default()
}
