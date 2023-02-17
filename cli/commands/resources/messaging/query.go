package messaging

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	messagingLib "github.com/taubyte/tau/lib/messaging"
	messagingPrompts "github.com/taubyte/tau/prompts/messaging"
	messagingTable "github.com/taubyte/tau/table/messaging"
)

func (link) Query() common.Command {
	return (&resources.Query[*structureSpec.Messaging]{
		LibListResources:   messagingLib.ListResources,
		TableList:          messagingTable.List,
		PromptsGetOrSelect: messagingPrompts.GetOrSelect,
		TableQuery:         messagingTable.Query,
	}).Default()
}

func (link) List() common.Command {
	return (&resources.List[*structureSpec.Messaging]{
		LibListResources: messagingLib.ListResources,
		TableList:        messagingTable.List,
	}).Default()
}
